/*
Copyright 2023 The xxx Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/apiserver-runtime/pkg/builder/resource"
)

const ConfigSetPlural = "configsets"

var _ resource.Object = &ConfigSet{}
var _ resource.ObjectList = &ConfigSetList{}

func (ConfigSet) GetGroupVersionResource() schema.GroupVersionResource {
	return schema.GroupVersionResource{
		Group:    SchemeGroupVersion.Group,
		Version:  SchemeGroupVersion.Version,
		Resource: ConfigPlural,
	}
}

// IsStorageVersion returns true -- v1alpha1.Config is used as the internal version.
// IsStorageVersion implements resource.Object.
func (ConfigSet) IsStorageVersion() bool {
	return true
}

// GetObjectMeta implements resource.Object
func (r *ConfigSet) GetObjectMeta() *metav1.ObjectMeta {
	return &r.ObjectMeta
}

// NamespaceScoped returns true to indicate Fortune is a namespaced resource.
// NamespaceScoped implements resource.Object.
func (ConfigSet) NamespaceScoped() bool {
	return true
}

// New implements resource.Object
func (ConfigSet) New() runtime.Object {
	return &ConfigSet{}
}

// NewList implements resource.Object
func (ConfigSet) NewList() runtime.Object {
	return &ConfigSetList{}
}

// GetCondition returns the condition based on the condition kind
func (r *ConfigSet) GetCondition(t ConditionType) Condition {
	return r.Status.GetCondition(t)
}

// SetConditions sets the conditions on the resource. it allows for 0, 1 or more conditions
// to be set at once
func (r *ConfigSet) SetConditions(c ...Condition) {
	r.Status.SetConditions(c...)
}

func (r *ConfigSet) GetTarget() string {
	if len(r.Labels) == 0 {
		return ""
	}
	var sb strings.Builder
	targetNamespace, ok := r.Labels["targetNamespace"]
	if ok {
		sb.WriteString(targetNamespace)
		sb.WriteString("/")
	}
	targetName, ok := r.Labels["targetName"]
	if ok {
		sb.WriteString(targetName)
	}
	return sb.String()
}

// GetListMeta returns the ListMeta
func (r *ConfigSetList) GetListMeta() *metav1.ListMeta {
	return &r.ListMeta
}
