// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-2019 Datadog, Inc.

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"time"

	v1alpha1 "github.com/DataDog/datadog-operator/pkg/apis/datadoghq/v1alpha1"
	scheme "github.com/DataDog/datadog-operator/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DatadogMetricsGetter has a method to return a DatadogMetricInterface.
// A group's client should implement this interface.
type DatadogMetricsGetter interface {
	DatadogMetrics(namespace string) DatadogMetricInterface
}

// DatadogMetricInterface has methods to work with DatadogMetric resources.
type DatadogMetricInterface interface {
	Create(*v1alpha1.DatadogMetric) (*v1alpha1.DatadogMetric, error)
	Update(*v1alpha1.DatadogMetric) (*v1alpha1.DatadogMetric, error)
	UpdateStatus(*v1alpha1.DatadogMetric) (*v1alpha1.DatadogMetric, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.DatadogMetric, error)
	List(opts v1.ListOptions) (*v1alpha1.DatadogMetricList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DatadogMetric, err error)
	DatadogMetricExpansion
}

// datadogMetrics implements DatadogMetricInterface
type datadogMetrics struct {
	client rest.Interface
	ns     string
}

// newDatadogMetrics returns a DatadogMetrics
func newDatadogMetrics(c *DatadoghqV1alpha1Client, namespace string) *datadogMetrics {
	return &datadogMetrics{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the datadogMetric, and returns the corresponding datadogMetric object, and an error if there is any.
func (c *datadogMetrics) Get(name string, options v1.GetOptions) (result *v1alpha1.DatadogMetric, err error) {
	result = &v1alpha1.DatadogMetric{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("datadogmetrics").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DatadogMetrics that match those selectors.
func (c *datadogMetrics) List(opts v1.ListOptions) (result *v1alpha1.DatadogMetricList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.DatadogMetricList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("datadogmetrics").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested datadogMetrics.
func (c *datadogMetrics) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("datadogmetrics").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a datadogMetric and creates it.  Returns the server's representation of the datadogMetric, and an error, if there is any.
func (c *datadogMetrics) Create(datadogMetric *v1alpha1.DatadogMetric) (result *v1alpha1.DatadogMetric, err error) {
	result = &v1alpha1.DatadogMetric{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("datadogmetrics").
		Body(datadogMetric).
		Do().
		Into(result)
	return
}

// Update takes the representation of a datadogMetric and updates it. Returns the server's representation of the datadogMetric, and an error, if there is any.
func (c *datadogMetrics) Update(datadogMetric *v1alpha1.DatadogMetric) (result *v1alpha1.DatadogMetric, err error) {
	result = &v1alpha1.DatadogMetric{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("datadogmetrics").
		Name(datadogMetric.Name).
		Body(datadogMetric).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *datadogMetrics) UpdateStatus(datadogMetric *v1alpha1.DatadogMetric) (result *v1alpha1.DatadogMetric, err error) {
	result = &v1alpha1.DatadogMetric{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("datadogmetrics").
		Name(datadogMetric.Name).
		SubResource("status").
		Body(datadogMetric).
		Do().
		Into(result)
	return
}

// Delete takes name of the datadogMetric and deletes it. Returns an error if one occurs.
func (c *datadogMetrics) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("datadogmetrics").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *datadogMetrics) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("datadogmetrics").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched datadogMetric.
func (c *datadogMetrics) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DatadogMetric, err error) {
	result = &v1alpha1.DatadogMetric{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("datadogmetrics").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
