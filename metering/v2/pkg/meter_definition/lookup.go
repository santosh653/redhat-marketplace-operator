// Copyright 2020 IBM Corp.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package meter_definition

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"emperror.dev/errors"
	"github.com/cespare/xxhash"
	"github.com/go-logr/logr"
	olmv1alpha1 "github.com/operator-framework/api/pkg/operators/v1alpha1"
	"github.com/redhat-marketplace/redhat-marketplace-operator/v2/apis/marketplace/v1beta1"
	rhmclient "github.com/redhat-marketplace/redhat-marketplace-operator/v2/pkg/client"
	. "github.com/redhat-marketplace/redhat-marketplace-operator/v2/pkg/utils/reconcileutils"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
)

type MeterDefWorkload = types.NamespacedName

type MeterDefinitionLookupFilter struct {
	MeterDefName MeterDefWorkload
	workloads    []v1beta1.ResourceFilter
	filters      [][]FilterRuntimeObject
	cc           ClientCommandRunner
	log          logr.Logger
	findOwner    *rhmclient.FindOwnerHelper
}

var (
	log = logf.Log.WithName("meterDefLookupFilter")
)

func NewMeterDefinitionLookupFilter(
	cc ClientCommandRunner,
	meterdef *v1beta1.MeterDefinition,
	findOwner *rhmclient.FindOwnerHelper,
) (*MeterDefinitionLookupFilter, error) {
	log.V(0).Info("building filters", "meterdef", meterdef)

	s := &MeterDefinitionLookupFilter{
		MeterDefName: types.NamespacedName{Name: meterdef.Name, Namespace: meterdef.Namespace},
		findOwner:    findOwner,
		cc:           cc,
		log:          log.WithValues("meterdefName", meterdef.Name, "meterdefNamespace", meterdef.Namespace),
	}

	ns, err := s.findNamespaces(meterdef)
	if err != nil {
		log.Error(err, "error creating find namespaces")
		return nil, err
	}
	filters, err := s.createFilters(meterdef, ns)
	if err != nil {
		s.log.Error(err, "error creating filters")
		return nil, err
	}

	s.workloads = meterdef.Spec.ResourceFilters
	s.filters = filters

	return s, nil
}

func (s *MeterDefinitionLookupFilter) Hash() string {
	h := xxhash.New()

	h.Write([]byte(fmt.Sprintf("%v", s.MeterDefName)))
	for k, v := range s.workloads {
		h.Write([]byte(fmt.Sprintf("%v", k)))
		h.Write([]byte(fmt.Sprintf("%v", v)))
	}

	return fmt.Sprintf("%x", h.Sum(nil))
}

func (s *MeterDefinitionLookupFilter) String() string {
	return fmt.Sprintf("MeterDef{workloads=%v, filters=%v}", len(s.workloads), len(s.filters))
}

func (s *MeterDefinitionLookupFilter) Matches(obj interface{}) (bool, error) {
	o, ok := obj.(metav1.Object)

	if !ok {
		err := errors.New("type is not a metav1 Object")
		s.log.Error(err, "failed to find workload due to error")
		return false, err
	}

	filterLogger := s.log.WithValues("obj", o.GetName()+"/"+o.GetNamespace(), "type", fmt.Sprintf("%T", obj), "filterLen", len(s.filters))
	debugFilterLogger := filterLogger

	for key, workloadFilters := range s.filters {
		debugFilterLogger.Info("testing", "key", key, "filters", printFilterList(workloadFilters))
		results := []bool{}
		for i, filter := range workloadFilters {
			ans, err := filter.Filter(obj)

			if err != nil {
				filterLogger.Error(err, "workload failed due to error", "workloadStatus", "fail", "filters", printFilterList(workloadFilters), "i", i, "filter", filter)
				return false, err
			}

			if !ans {
				break
			}

			results = append(results, ans)
		}

		if len(results) == 0 || len(results) != len(workloadFilters) {
			debugFilterLogger.Info("workload did not pass all filters", "workloadStatus", "fail", "filters", printFilterList(workloadFilters))
			continue
		}

		debugFilterLogger.Info("workload passed all filters", "workloadStatus", "pass", "filters", printFilterList(workloadFilters))
		return true, nil
	}

	return false, nil
}

func (s *MeterDefinitionLookupFilter) findNamespaces(
	instance *v1beta1.MeterDefinition,
) (namespaces []string, err error) {
	cc := s.cc
	functionError := errors.NewWithDetails("error with findNamespaces", "meterdef", instance.Name+"/"+instance.Namespace)
	reqLogger := s.log.WithValues("func", "findNamespaces", "meterdef", instance.Name+"/"+instance.Namespace)

	for _, resourceFilter := range instance.Spec.ResourceFilters {
		if resourceFilter.Namespace == nil {
			reqLogger.Info("operatorGroup is for all namespaces")
			namespaces = []string{corev1.NamespaceAll}
			return
		}

		if resourceFilter.Namespace.UseOperatorGroup {
			reqLogger.Info("operatorGroup vertex")
			csv := &olmv1alpha1.ClusterServiceVersion{}

			if instance.Spec.InstalledBy == nil {
				reqLogger.Info("installedBy not provided, falling back to namespace")

				return []string{instance.GetNamespace()}, nil
			}

			reqLogger.Info("installedBy provided, looking for operatorgroup")

			result, _ := cc.Do(context.TODO(),
				GetAction(instance.Spec.InstalledBy.ToTypes(), csv),
			)

			if result.Is(NotFound) {
				reqLogger.Info("installedBy not found, falling back to namespace")

				return []string{instance.GetNamespace()}, nil
			}

			if !result.Is(Continue) {
				err = errors.Wrap(functionError, "csv not found due to error")
				reqLogger.Error(err, "installed by is not found")

				return
			}

			olmNamespacesStr, ok := csv.GetAnnotations()["olm.targetNamespaces"]

			if !ok {
				err = errors.Wrap(functionError, "olmNamspaces on CSV not found")
				// set condition and requeue for later
				reqLogger.Error(err, "")
				return
			}

			if olmNamespacesStr == "" {
				reqLogger.Info("operatorGroup is for all namespaces")
				namespaces = []string{corev1.NamespaceAll}
				return
			}

			namespaces = strings.Split(olmNamespacesStr, ",")
			return
		}

		if resourceFilter.Namespace.LabelSelector != nil {
			reqLogger.Info("namespace vertex with filter")

			if resourceFilter.Namespace.LabelSelector == nil {
				reqLogger.Info("namespace vertex is for all namespaces")
				break
			}

			namespaceList := &corev1.NamespaceList{}

			var selector labels.Selector
			selector, err = metav1.LabelSelectorAsSelector(resourceFilter.Namespace.LabelSelector)

			if err != nil {
				return
			}

			result, _ := cc.Do(
				context.TODO(),
				ListAction(namespaceList, client.MatchingLabelsSelector{Selector: selector}),
			)

			if !result.Is(Continue) {
				err = errors.Wrap(functionError, "csv not found")
				reqLogger.Info("csv not found", "csv", instance.Spec.InstalledBy)

				return
			}

			for _, ns := range namespaceList.Items {
				namespaces = append(namespaces, ns.GetName())
			}
		}
	}
	return
}

func (s *MeterDefinitionLookupFilter) createFilters(
	instance *v1beta1.MeterDefinition,
	namespaces []string,
) ([][]FilterRuntimeObject, error) {

	// Bottom Up
	// Start with pods, filter, go to owner. If owner not provided, stop.
	filters := [][]FilterRuntimeObject{}

	for _, filter := range instance.Spec.ResourceFilters {
		runtimeFilters := []FilterRuntimeObject{&WorkloadNamespaceFilter{namespaces: namespaces}}

		var err error
		typeFilter := &WorkloadTypeFilter{}
		switch filter.WorkloadType {
		case v1beta1.WorkloadTypePod:
			gvk := reflect.TypeOf(&corev1.Pod{})
			typeFilter.gvks = []reflect.Type{gvk}
		case v1beta1.WorkloadTypePVC:
			gvk := reflect.TypeOf(&corev1.PersistentVolumeClaim{})
			typeFilter.gvks = []reflect.Type{gvk}
		case v1beta1.WorkloadTypeService:
			gvk1 := reflect.TypeOf(&corev1.Service{})
			typeFilter.gvks = []reflect.Type{gvk1}
		default:
			s.log.Error(err, "unknown type filter", "type", filter.WorkloadType)
			err = errors.NewWithDetails("unknown type filter", "type", filter.WorkloadType)
			return nil, err
		}

		runtimeFilters = append(runtimeFilters, typeFilter)

		if filter.Label != nil && filter.Label.LabelSelector != nil {
			selector, err := metav1.LabelSelectorAsSelector(filter.Label.LabelSelector)

			if err != nil {
				return nil, err
			}

			runtimeFilters = append(runtimeFilters, &WorkloadLabelFilter{
				labelSelector: selector,
			})
		}

		if filter.Annotation != nil && filter.Annotation.AnnotationSelector != nil {
			selector, err := metav1.LabelSelectorAsSelector(filter.Annotation.AnnotationSelector)

			if err != nil {
				return nil, err
			}

			runtimeFilters = append(runtimeFilters, &WorkloadAnnotationFilter{
				annotationSelector: selector,
			})
		}

		if filter.OwnerCRD != nil {
			runtimeFilters = append(runtimeFilters, NewWorkloadFilterForOwner(*filter.OwnerCRD, s.findOwner))
		}

		filters = append(filters, runtimeFilters)
	}
	return filters, nil
}
