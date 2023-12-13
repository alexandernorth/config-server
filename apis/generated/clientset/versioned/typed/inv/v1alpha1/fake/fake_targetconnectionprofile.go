// Copyright 2023 The xxx Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/iptecharch/config-server/apis/inv/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeTargetConnectionProfiles implements TargetConnectionProfileInterface
type FakeTargetConnectionProfiles struct {
	Fake *FakeInvV1alpha1
	ns   string
}

var targetconnectionprofilesResource = schema.GroupVersionResource{Group: "inv.nephio.org", Version: "v1alpha1", Resource: "targetconnectionprofiles"}

var targetconnectionprofilesKind = schema.GroupVersionKind{Group: "inv.nephio.org", Version: "v1alpha1", Kind: "TargetConnectionProfile"}

// Get takes name of the targetConnectionProfile, and returns the corresponding targetConnectionProfile object, and an error if there is any.
func (c *FakeTargetConnectionProfiles) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.TargetConnectionProfile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(targetconnectionprofilesResource, c.ns, name), &v1alpha1.TargetConnectionProfile{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TargetConnectionProfile), err
}

// List takes label and field selectors, and returns the list of TargetConnectionProfiles that match those selectors.
func (c *FakeTargetConnectionProfiles) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.TargetConnectionProfileList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(targetconnectionprofilesResource, targetconnectionprofilesKind, c.ns, opts), &v1alpha1.TargetConnectionProfileList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.TargetConnectionProfileList{ListMeta: obj.(*v1alpha1.TargetConnectionProfileList).ListMeta}
	for _, item := range obj.(*v1alpha1.TargetConnectionProfileList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested targetConnectionProfiles.
func (c *FakeTargetConnectionProfiles) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(targetconnectionprofilesResource, c.ns, opts))

}

// Create takes the representation of a targetConnectionProfile and creates it.  Returns the server's representation of the targetConnectionProfile, and an error, if there is any.
func (c *FakeTargetConnectionProfiles) Create(ctx context.Context, targetConnectionProfile *v1alpha1.TargetConnectionProfile, opts v1.CreateOptions) (result *v1alpha1.TargetConnectionProfile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(targetconnectionprofilesResource, c.ns, targetConnectionProfile), &v1alpha1.TargetConnectionProfile{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TargetConnectionProfile), err
}

// Update takes the representation of a targetConnectionProfile and updates it. Returns the server's representation of the targetConnectionProfile, and an error, if there is any.
func (c *FakeTargetConnectionProfiles) Update(ctx context.Context, targetConnectionProfile *v1alpha1.TargetConnectionProfile, opts v1.UpdateOptions) (result *v1alpha1.TargetConnectionProfile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(targetconnectionprofilesResource, c.ns, targetConnectionProfile), &v1alpha1.TargetConnectionProfile{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TargetConnectionProfile), err
}

// Delete takes name of the targetConnectionProfile and deletes it. Returns an error if one occurs.
func (c *FakeTargetConnectionProfiles) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(targetconnectionprofilesResource, c.ns, name, opts), &v1alpha1.TargetConnectionProfile{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTargetConnectionProfiles) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(targetconnectionprofilesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.TargetConnectionProfileList{})
	return err
}

// Patch applies the patch and returns the patched targetConnectionProfile.
func (c *FakeTargetConnectionProfiles) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.TargetConnectionProfile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(targetconnectionprofilesResource, c.ns, name, pt, data, subresources...), &v1alpha1.TargetConnectionProfile{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TargetConnectionProfile), err
}
