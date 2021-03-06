/*
Copyright 2018 The Kubernetes Authors.

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

package processors

import (
	"k8s.io/autoscaler/cluster-autoscaler/processors/nodegroups"
	"k8s.io/autoscaler/cluster-autoscaler/processors/pods"
)

// AutoscalingProcessors are a set of customizable processors used for encapsulating
// various heuristics used in different parts of Cluster Autoscaler code.
type AutoscalingProcessors struct {
	// PodListProcessor is used to process list of unschedulable pods before autoscaling.
	PodListProcessor pods.PodListProcessor
	// NodeGroupListProcessor is used to process list of NodeGroups that can be used in scale-up.
	NodeGroupListProcessor nodegroups.NodeGroupListProcessor
}

// DefaultProcessors returns default set of processors.
func DefaultProcessors() *AutoscalingProcessors {
	return &AutoscalingProcessors{
		PodListProcessor:       pods.NewDefaultPodListProcessor(),
		NodeGroupListProcessor: nodegroups.NewDefaultNodeGroupListProcessor()}
}

// TestProcessors returns a set of simple processors for use in tests.
func TestProcessors() *AutoscalingProcessors {
	return &AutoscalingProcessors{
		PodListProcessor:       &pods.NoOpPodListProcessor{},
		NodeGroupListProcessor: &nodegroups.NoOpNodeGroupListProcessor{}}
}
