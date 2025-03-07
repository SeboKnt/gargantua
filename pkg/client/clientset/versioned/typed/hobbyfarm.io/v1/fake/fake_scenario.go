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
	"context"

	hobbyfarmiov1 "github.com/hobbyfarm/gargantua/v3/pkg/apis/hobbyfarm.io/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeScenarios implements ScenarioInterface
type FakeScenarios struct {
	Fake *FakeHobbyfarmV1
	ns   string
}

var scenariosResource = schema.GroupVersionResource{Group: "hobbyfarm.io", Version: "v1", Resource: "scenarios"}

var scenariosKind = schema.GroupVersionKind{Group: "hobbyfarm.io", Version: "v1", Kind: "Scenario"}

// Get takes name of the scenario, and returns the corresponding scenario object, and an error if there is any.
func (c *FakeScenarios) Get(ctx context.Context, name string, options v1.GetOptions) (result *hobbyfarmiov1.Scenario, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(scenariosResource, c.ns, name), &hobbyfarmiov1.Scenario{})

	if obj == nil {
		return nil, err
	}
	return obj.(*hobbyfarmiov1.Scenario), err
}

// List takes label and field selectors, and returns the list of Scenarios that match those selectors.
func (c *FakeScenarios) List(ctx context.Context, opts v1.ListOptions) (result *hobbyfarmiov1.ScenarioList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(scenariosResource, scenariosKind, c.ns, opts), &hobbyfarmiov1.ScenarioList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &hobbyfarmiov1.ScenarioList{ListMeta: obj.(*hobbyfarmiov1.ScenarioList).ListMeta}
	for _, item := range obj.(*hobbyfarmiov1.ScenarioList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested scenarios.
func (c *FakeScenarios) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(scenariosResource, c.ns, opts))

}

// Create takes the representation of a scenario and creates it.  Returns the server's representation of the scenario, and an error, if there is any.
func (c *FakeScenarios) Create(ctx context.Context, scenario *hobbyfarmiov1.Scenario, opts v1.CreateOptions) (result *hobbyfarmiov1.Scenario, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(scenariosResource, c.ns, scenario), &hobbyfarmiov1.Scenario{})

	if obj == nil {
		return nil, err
	}
	return obj.(*hobbyfarmiov1.Scenario), err
}

// Update takes the representation of a scenario and updates it. Returns the server's representation of the scenario, and an error, if there is any.
func (c *FakeScenarios) Update(ctx context.Context, scenario *hobbyfarmiov1.Scenario, opts v1.UpdateOptions) (result *hobbyfarmiov1.Scenario, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(scenariosResource, c.ns, scenario), &hobbyfarmiov1.Scenario{})

	if obj == nil {
		return nil, err
	}
	return obj.(*hobbyfarmiov1.Scenario), err
}

// Delete takes name of the scenario and deletes it. Returns an error if one occurs.
func (c *FakeScenarios) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(scenariosResource, c.ns, name, opts), &hobbyfarmiov1.Scenario{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeScenarios) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(scenariosResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &hobbyfarmiov1.ScenarioList{})
	return err
}

// Patch applies the patch and returns the patched scenario.
func (c *FakeScenarios) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *hobbyfarmiov1.Scenario, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(scenariosResource, c.ns, name, pt, data, subresources...), &hobbyfarmiov1.Scenario{})

	if obj == nil {
		return nil, err
	}
	return obj.(*hobbyfarmiov1.Scenario), err
}
