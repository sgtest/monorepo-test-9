/*
Copyright 2016 The Kubernetes Authors.

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

package util

import (
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/util/sets"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	"github.com/sourcegraph/monorepo-test-1/kubernetes-7/pkg/api/v1"
	"github.com/sourcegraph/monorepo-test-1/kubernetes-7/pkg/features"
)

// GetNamespacesFromPodAffinityTerm returns a set of names
// according to the namespaces indicated in podAffinityTerm.
// If namespaces is empty it considers the given pod's namespace.
func GetNamespacesFromPodAffinityTerm(pod *v1.Pod, podAffinityTerm *v1.PodAffinityTerm) sets.String {
	names := sets.String{}
	if len(podAffinityTerm.Namespaces) == 0 {
		names.Insert(pod.Namespace)
	} else {
		names.Insert(podAffinityTerm.Namespaces...)
	}
	return names
}

// PodMatchesTermsNamespaceAndSelector returns true if the given <pod>
// matches the namespace and selector defined by <affinityPod>`s <term>.
func PodMatchesTermsNamespaceAndSelector(pod *v1.Pod, namespaces sets.String, selector labels.Selector) bool {
	if !namespaces.Has(pod.Namespace) {
		return false
	}

	if !selector.Matches(labels.Set(pod.Labels)) {
		return false
	}
	return true
}

// NodesHaveSameTopologyKey checks if nodeA and nodeB have same label value with given topologyKey as label key.
// Returns false if topologyKey is empty.
func NodesHaveSameTopologyKey(nodeA, nodeB *v1.Node, topologyKey string) bool {
	if len(topologyKey) == 0 {
		return false
	}

	if nodeA.Labels == nil || nodeB.Labels == nil {
		return false
	}

	nodeALabel, okA := nodeA.Labels[topologyKey]
	nodeBLabel, okB := nodeB.Labels[topologyKey]

	// If found label in both nodes, check the label
	if okB && okA {
		return nodeALabel == nodeBLabel
	}

	return false
}

type Topologies struct {
	DefaultKeys []string
}

// NodesHaveSameTopologyKey checks if nodeA and nodeB have same label value with given topologyKey as label key.
// If the topologyKey is empty, check if the two nodes have any of the default topologyKeys, and have same corresponding label value.
func (tps *Topologies) NodesHaveSameTopologyKey(nodeA, nodeB *v1.Node, topologyKey string) bool {
	if utilfeature.DefaultFeatureGate.Enabled(features.AffinityInAnnotations) && len(topologyKey) == 0 {
		// assumes this is allowed only for PreferredDuringScheduling pod anti-affinity (ensured by api/validation)
		for _, defaultKey := range tps.DefaultKeys {
			if NodesHaveSameTopologyKey(nodeA, nodeB, defaultKey) {
				return true
			}
		}
		return false
	} else {
		return NodesHaveSameTopologyKey(nodeA, nodeB, topologyKey)
	}
}
