/*
Copyright 2021-2022 Red Hat, Inc.

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
// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/redhat-appstudio/jvm-build-service/pkg/apis/jvmbuildservice/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ArtifactBuildRequestLister helps list ArtifactBuildRequests.
// All objects returned here must be treated as read-only.
type ArtifactBuildRequestLister interface {
	// List lists all ArtifactBuildRequests in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ArtifactBuildRequest, err error)
	// ArtifactBuildRequests returns an object that can list and get ArtifactBuildRequests.
	ArtifactBuildRequests(namespace string) ArtifactBuildRequestNamespaceLister
	ArtifactBuildRequestListerExpansion
}

// artifactBuildRequestLister implements the ArtifactBuildRequestLister interface.
type artifactBuildRequestLister struct {
	indexer cache.Indexer
}

// NewArtifactBuildRequestLister returns a new ArtifactBuildRequestLister.
func NewArtifactBuildRequestLister(indexer cache.Indexer) ArtifactBuildRequestLister {
	return &artifactBuildRequestLister{indexer: indexer}
}

// List lists all ArtifactBuildRequests in the indexer.
func (s *artifactBuildRequestLister) List(selector labels.Selector) (ret []*v1alpha1.ArtifactBuildRequest, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ArtifactBuildRequest))
	})
	return ret, err
}

// ArtifactBuildRequests returns an object that can list and get ArtifactBuildRequests.
func (s *artifactBuildRequestLister) ArtifactBuildRequests(namespace string) ArtifactBuildRequestNamespaceLister {
	return artifactBuildRequestNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ArtifactBuildRequestNamespaceLister helps list and get ArtifactBuildRequests.
// All objects returned here must be treated as read-only.
type ArtifactBuildRequestNamespaceLister interface {
	// List lists all ArtifactBuildRequests in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ArtifactBuildRequest, err error)
	// Get retrieves the ArtifactBuildRequest from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ArtifactBuildRequest, error)
	ArtifactBuildRequestNamespaceListerExpansion
}

// artifactBuildRequestNamespaceLister implements the ArtifactBuildRequestNamespaceLister
// interface.
type artifactBuildRequestNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ArtifactBuildRequests in the indexer for a given namespace.
func (s artifactBuildRequestNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ArtifactBuildRequest, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ArtifactBuildRequest))
	})
	return ret, err
}

// Get retrieves the ArtifactBuildRequest from the indexer for a given namespace and name.
func (s artifactBuildRequestNamespaceLister) Get(name string) (*v1alpha1.ArtifactBuildRequest, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("artifactbuildrequest"), name)
	}
	return obj.(*v1alpha1.ArtifactBuildRequest), nil
}
