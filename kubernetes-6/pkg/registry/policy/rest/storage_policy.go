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

package rest

import (
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/registry/rest"
	genericapiserver "k8s.io/apiserver/pkg/server"
	serverstorage "k8s.io/apiserver/pkg/server/storage"
	"github.com/sourcegraph/monorepo-test-1/kubernetes-6/pkg/api"
	"github.com/sourcegraph/monorepo-test-1/kubernetes-6/pkg/apis/policy"
	policyapiv1beta1 "github.com/sourcegraph/monorepo-test-1/kubernetes-6/pkg/apis/policy/v1beta1"
	poddisruptionbudgetstore "github.com/sourcegraph/monorepo-test-1/kubernetes-6/pkg/registry/policy/poddisruptionbudget/storage"
)

type RESTStorageProvider struct{}

func (p RESTStorageProvider) NewRESTStorage(apiResourceConfigSource serverstorage.APIResourceConfigSource, restOptionsGetter generic.RESTOptionsGetter) (genericapiserver.APIGroupInfo, bool) {
	apiGroupInfo := genericapiserver.NewDefaultAPIGroupInfo(policy.GroupName, api.Registry, api.Scheme, api.ParameterCodec, api.Codecs)

	if apiResourceConfigSource.AnyResourcesForVersionEnabled(policyapiv1beta1.SchemeGroupVersion) {
		apiGroupInfo.VersionedResourcesStorageMap[policyapiv1beta1.SchemeGroupVersion.Version] = p.v1beta1Storage(apiResourceConfigSource, restOptionsGetter)
		apiGroupInfo.GroupMeta.GroupVersion = policyapiv1beta1.SchemeGroupVersion
	}
	return apiGroupInfo, true
}

func (p RESTStorageProvider) v1beta1Storage(apiResourceConfigSource serverstorage.APIResourceConfigSource, restOptionsGetter generic.RESTOptionsGetter) map[string]rest.Storage {
	version := policyapiv1beta1.SchemeGroupVersion
	storage := map[string]rest.Storage{}
	if apiResourceConfigSource.ResourceEnabled(version.WithResource("poddisruptionbudgets")) {
		poddisruptionbudgetStorage, poddisruptionbudgetStatusStorage := poddisruptionbudgetstore.NewREST(restOptionsGetter)
		storage["poddisruptionbudgets"] = poddisruptionbudgetStorage
		storage["poddisruptionbudgets/status"] = poddisruptionbudgetStatusStorage
	}
	return storage
}

func (p RESTStorageProvider) GroupName() string {
	return policy.GroupName
}
