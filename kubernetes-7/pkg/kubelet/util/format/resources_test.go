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

package format

import (
	"testing"

	"k8s.io/apimachinery/pkg/api/resource"
	"github.com/sourcegraph/monorepo-test-1/kubernetes-7/pkg/api/v1"
)

func TestResourceList(t *testing.T) {
	resourceList := v1.ResourceList{}
	resourceList[v1.ResourceCPU] = resource.MustParse("100m")
	resourceList[v1.ResourceMemory] = resource.MustParse("5Gi")
	actual := ResourceList(resourceList)
	expected := "cpu=100m,memory=5Gi"
	if actual != expected {
		t.Errorf("Unexpected result, actual: %v, expected: %v", actual, expected)
	}
}
