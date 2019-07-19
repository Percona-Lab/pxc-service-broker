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

package v1

import (
	"time"

	v1 "github.com/Percona-Lab/pxc-service-broker/pkg/apis/pxc/v1"
	scheme "github.com/Percona-Lab/pxc-service-broker/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// PerconaXtraDBClusterBackupsGetter has a method to return a PerconaXtraDBClusterBackupInterface.
// A group's client should implement this interface.
type PerconaXtraDBClusterBackupsGetter interface {
	PerconaXtraDBClusterBackups(namespace string) PerconaXtraDBClusterBackupInterface
}

// PerconaXtraDBClusterBackupInterface has methods to work with PerconaXtraDBClusterBackup resources.
type PerconaXtraDBClusterBackupInterface interface {
	Create(*v1.PerconaXtraDBClusterBackup) (*v1.PerconaXtraDBClusterBackup, error)
	Update(*v1.PerconaXtraDBClusterBackup) (*v1.PerconaXtraDBClusterBackup, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.PerconaXtraDBClusterBackup, error)
	List(opts metav1.ListOptions) (*v1.PerconaXtraDBClusterBackupList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.PerconaXtraDBClusterBackup, err error)
	PerconaXtraDBClusterBackupExpansion
}

// perconaXtraDBClusterBackups implements PerconaXtraDBClusterBackupInterface
type perconaXtraDBClusterBackups struct {
	client rest.Interface
	ns     string
}

// newPerconaXtraDBClusterBackups returns a PerconaXtraDBClusterBackups
func newPerconaXtraDBClusterBackups(c *PxcV1Client, namespace string) *perconaXtraDBClusterBackups {
	return &perconaXtraDBClusterBackups{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the perconaXtraDBClusterBackup, and returns the corresponding perconaXtraDBClusterBackup object, and an error if there is any.
func (c *perconaXtraDBClusterBackups) Get(name string, options metav1.GetOptions) (result *v1.PerconaXtraDBClusterBackup, err error) {
	result = &v1.PerconaXtraDBClusterBackup{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("perconaxtradbclusterbackups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of PerconaXtraDBClusterBackups that match those selectors.
func (c *perconaXtraDBClusterBackups) List(opts metav1.ListOptions) (result *v1.PerconaXtraDBClusterBackupList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.PerconaXtraDBClusterBackupList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("perconaxtradbclusterbackups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested perconaXtraDBClusterBackups.
func (c *perconaXtraDBClusterBackups) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("perconaxtradbclusterbackups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a perconaXtraDBClusterBackup and creates it.  Returns the server's representation of the perconaXtraDBClusterBackup, and an error, if there is any.
func (c *perconaXtraDBClusterBackups) Create(perconaXtraDBClusterBackup *v1.PerconaXtraDBClusterBackup) (result *v1.PerconaXtraDBClusterBackup, err error) {
	result = &v1.PerconaXtraDBClusterBackup{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("perconaxtradbclusterbackups").
		Body(perconaXtraDBClusterBackup).
		Do().
		Into(result)
	return
}

// Update takes the representation of a perconaXtraDBClusterBackup and updates it. Returns the server's representation of the perconaXtraDBClusterBackup, and an error, if there is any.
func (c *perconaXtraDBClusterBackups) Update(perconaXtraDBClusterBackup *v1.PerconaXtraDBClusterBackup) (result *v1.PerconaXtraDBClusterBackup, err error) {
	result = &v1.PerconaXtraDBClusterBackup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("perconaxtradbclusterbackups").
		Name(perconaXtraDBClusterBackup.Name).
		Body(perconaXtraDBClusterBackup).
		Do().
		Into(result)
	return
}

// Delete takes name of the perconaXtraDBClusterBackup and deletes it. Returns an error if one occurs.
func (c *perconaXtraDBClusterBackups) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("perconaxtradbclusterbackups").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *perconaXtraDBClusterBackups) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("perconaxtradbclusterbackups").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched perconaXtraDBClusterBackup.
func (c *perconaXtraDBClusterBackups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.PerconaXtraDBClusterBackup, err error) {
	result = &v1.PerconaXtraDBClusterBackup{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("perconaxtradbclusterbackups").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}