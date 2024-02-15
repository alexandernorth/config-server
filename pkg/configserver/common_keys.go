/*
Copyright 2024 Nokia.

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

package configserver

import (
	"context"
	"fmt"

	configv1alpha1 "github.com/sdcio/config-server/apis/config/v1alpha1"
	"github.com/henderiw/apiserver-store/pkg/storebackend"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	genericapirequest "k8s.io/apiserver/pkg/endpoints/request"
)

func (r *configCommon) getKey(ctx context.Context, name string) (storebackend.Key, error) {
	ns, namespaced := genericapirequest.NamespaceFrom(ctx)
	if namespaced != r.isNamespaced {
		return storebackend.Key{}, fmt.Errorf("namespace mismatch got %t, want %t", namespaced, r.isNamespaced)
	}
	return storebackend.Key{
		NamespacedName: types.NamespacedName{
			Name:      name,
			Namespace: ns,
		},
	}, nil
}

func (r *configCommon) getKeys(ctx context.Context, obj runtime.Object) (storebackend.Key, storebackend.Key, error) {
	accessor, err := meta.Accessor(obj)
	if err != nil {
		return storebackend.Key{}, storebackend.Key{}, err
	}
	// Get Key
	key, err := r.getKey(ctx, accessor.GetName())
	if err != nil {
		return storebackend.Key{}, storebackend.Key{}, err
	}
	// Get targetKey to intercat with the device
	targetKey, err := getTargetKey(accessor.GetLabels())
	if err != nil {
		return storebackend.Key{}, storebackend.Key{}, err
	}
	return key, targetKey, nil
}

func getTargetKey(labels map[string]string) (storebackend.Key, error) {
	var targetName, targetNamespace string
	if labels != nil {
		targetName = labels[configv1alpha1.TargetNameKey]
		targetNamespace = labels[configv1alpha1.TargetNamespaceKey]
	}
	if targetName == "" || targetNamespace == "" {
		return storebackend.Key{}, fmt.Errorf(" target namespace and name is required got %s.%s", targetNamespace, targetName)
	}
	return storebackend.Key{NamespacedName: types.NamespacedName{Namespace: targetNamespace, Name: targetName}}, nil
}
