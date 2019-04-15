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
	clientset "github.com/kubernetes-sigs/volcano/pkg/client/clientset/versioned"
	batchv1alpha1 "github.com/kubernetes-sigs/volcano/pkg/client/clientset/versioned/typed/batch/v1alpha1"
	fakebatchv1alpha1 "github.com/kubernetes-sigs/volcano/pkg/client/clientset/versioned/typed/batch/v1alpha1/fake"
	busv1alpha1 "github.com/kubernetes-sigs/volcano/pkg/client/clientset/versioned/typed/bus/v1alpha1"
	fakebusv1alpha1 "github.com/kubernetes-sigs/volcano/pkg/client/clientset/versioned/typed/bus/v1alpha1/fake"
	schedulingv1alpha1 "github.com/kubernetes-sigs/volcano/pkg/client/clientset/versioned/typed/scheduling/v1alpha1"
	fakeschedulingv1alpha1 "github.com/kubernetes-sigs/volcano/pkg/client/clientset/versioned/typed/scheduling/v1alpha1/fake"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/discovery"
	fakediscovery "k8s.io/client-go/discovery/fake"
	"k8s.io/client-go/testing"
)

// NewSimpleClientset returns a clientset that will respond with the provided objects.
// It's backed by a very simple object tracker that processes creates, updates and deletions as-is,
// without applying any validations and/or defaults. It shouldn't be considered a replacement
// for a real clientset and is mostly useful in simple unit tests.
func NewSimpleClientset(objects ...runtime.Object) *Clientset {
	o := testing.NewObjectTracker(scheme, codecs.UniversalDecoder())
	for _, obj := range objects {
		if err := o.Add(obj); err != nil {
			panic(err)
		}
	}

	cs := &Clientset{}
	cs.discovery = &fakediscovery.FakeDiscovery{Fake: &cs.Fake}
	cs.AddReactor("*", "*", testing.ObjectReaction(o))
	cs.AddWatchReactor("*", func(action testing.Action) (handled bool, ret watch.Interface, err error) {
		gvr := action.GetResource()
		ns := action.GetNamespace()
		watch, err := o.Watch(gvr, ns)
		if err != nil {
			return false, nil, err
		}
		return true, watch, nil
	})

	return cs
}

// Clientset implements clientset.Interface. Meant to be embedded into a
// struct to get a default implementation. This makes faking out just the method
// you want to test easier.
type Clientset struct {
	testing.Fake
	discovery *fakediscovery.FakeDiscovery
}

func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	return c.discovery
}

var _ clientset.Interface = &Clientset{}

// BatchV1alpha1 retrieves the BatchV1alpha1Client
func (c *Clientset) BatchV1alpha1() batchv1alpha1.BatchV1alpha1Interface {
	return &fakebatchv1alpha1.FakeBatchV1alpha1{Fake: &c.Fake}
}

// Batch retrieves the BatchV1alpha1Client
func (c *Clientset) Batch() batchv1alpha1.BatchV1alpha1Interface {
	return &fakebatchv1alpha1.FakeBatchV1alpha1{Fake: &c.Fake}
}

// BusV1alpha1 retrieves the BusV1alpha1Client
func (c *Clientset) BusV1alpha1() busv1alpha1.BusV1alpha1Interface {
	return &fakebusv1alpha1.FakeBusV1alpha1{Fake: &c.Fake}
}

// Bus retrieves the BusV1alpha1Client
func (c *Clientset) Bus() busv1alpha1.BusV1alpha1Interface {
	return &fakebusv1alpha1.FakeBusV1alpha1{Fake: &c.Fake}
}

// SchedulingV1alpha1 retrieves the SchedulingV1alpha1Client
func (c *Clientset) SchedulingV1alpha1() schedulingv1alpha1.SchedulingV1alpha1Interface {
	return &fakeschedulingv1alpha1.FakeSchedulingV1alpha1{Fake: &c.Fake}
}

// Scheduling retrieves the SchedulingV1alpha1Client
func (c *Clientset) Scheduling() schedulingv1alpha1.SchedulingV1alpha1Interface {
	return &fakeschedulingv1alpha1.FakeSchedulingV1alpha1{Fake: &c.Fake}
}
