/*
Copyright The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	pxcv1 "github.com/Percona-Lab/pxc-service-broker/pkg/apis/pxc/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePerconaXtraDBClusterBackups implements PerconaXtraDBClusterBackupInterface
type FakePerconaXtraDBClusterBackups struct {
	Fake *FakePxcV1
	ns   string
}

var perconaxtradbclusterbackupsResource = schema.GroupVersionResource{Group: "pxc.percona.com", Version: "v1", Resource: "perconaxtradbclusterbackups"}

var perconaxtradbclusterbackupsKind = schema.GroupVersionKind{Group: "pxc.percona.com", Version: "v1", Kind: "PerconaXtraDBClusterBackup"}

// Get takes name of the perconaXtraDBClusterBackup, and returns the corresponding perconaXtraDBClusterBackup object, and an error if there is any.
func (c *FakePerconaXtraDBClusterBackups) Get(name string, options v1.GetOptions) (result *pxcv1.PerconaXtraDBClusterBackup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(perconaxtradbclusterbackupsResource, c.ns, name), &pxcv1.PerconaXtraDBClusterBackup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*pxcv1.PerconaXtraDBClusterBackup), err
}

// List takes label and field selectors, and returns the list of PerconaXtraDBClusterBackups that match those selectors.
func (c *FakePerconaXtraDBClusterBackups) List(opts v1.ListOptions) (result *pxcv1.PerconaXtraDBClusterBackupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(perconaxtradbclusterbackupsResource, perconaxtradbclusterbackupsKind, c.ns, opts), &pxcv1.PerconaXtraDBClusterBackupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &pxcv1.PerconaXtraDBClusterBackupList{ListMeta: obj.(*pxcv1.PerconaXtraDBClusterBackupList).ListMeta}
	for _, item := range obj.(*pxcv1.PerconaXtraDBClusterBackupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested perconaXtraDBClusterBackups.
func (c *FakePerconaXtraDBClusterBackups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(perconaxtradbclusterbackupsResource, c.ns, opts))

}

// Create takes the representation of a perconaXtraDBClusterBackup and creates it.  Returns the server's representation of the perconaXtraDBClusterBackup, and an error, if there is any.
func (c *FakePerconaXtraDBClusterBackups) Create(perconaXtraDBClusterBackup *pxcv1.PerconaXtraDBClusterBackup) (result *pxcv1.PerconaXtraDBClusterBackup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(perconaxtradbclusterbackupsResource, c.ns, perconaXtraDBClusterBackup), &pxcv1.PerconaXtraDBClusterBackup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*pxcv1.PerconaXtraDBClusterBackup), err
}

// Update takes the representation of a perconaXtraDBClusterBackup and updates it. Returns the server's representation of the perconaXtraDBClusterBackup, and an error, if there is any.
func (c *FakePerconaXtraDBClusterBackups) Update(perconaXtraDBClusterBackup *pxcv1.PerconaXtraDBClusterBackup) (result *pxcv1.PerconaXtraDBClusterBackup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(perconaxtradbclusterbackupsResource, c.ns, perconaXtraDBClusterBackup), &pxcv1.PerconaXtraDBClusterBackup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*pxcv1.PerconaXtraDBClusterBackup), err
}

// Delete takes name of the perconaXtraDBClusterBackup and deletes it. Returns an error if one occurs.
func (c *FakePerconaXtraDBClusterBackups) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(perconaxtradbclusterbackupsResource, c.ns, name), &pxcv1.PerconaXtraDBClusterBackup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePerconaXtraDBClusterBackups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(perconaxtradbclusterbackupsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &pxcv1.PerconaXtraDBClusterBackupList{})
	return err
}

// Patch applies the patch and returns the patched perconaXtraDBClusterBackup.
func (c *FakePerconaXtraDBClusterBackups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *pxcv1.PerconaXtraDBClusterBackup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(perconaxtradbclusterbackupsResource, c.ns, name, pt, data, subresources...), &pxcv1.PerconaXtraDBClusterBackup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*pxcv1.PerconaXtraDBClusterBackup), err
}