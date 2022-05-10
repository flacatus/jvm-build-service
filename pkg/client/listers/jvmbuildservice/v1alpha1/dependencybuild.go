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

// DependencyBuildLister helps list DependencyBuilds.
// All objects returned here must be treated as read-only.
type DependencyBuildLister interface {
	// List lists all DependencyBuilds in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DependencyBuild, err error)
	// DependencyBuilds returns an object that can list and get DependencyBuilds.
	DependencyBuilds(namespace string) DependencyBuildNamespaceLister
	DependencyBuildListerExpansion
}

// dependencyBuildLister implements the DependencyBuildLister interface.
type dependencyBuildLister struct {
	indexer cache.Indexer
}

// NewDependencyBuildLister returns a new DependencyBuildLister.
func NewDependencyBuildLister(indexer cache.Indexer) DependencyBuildLister {
	return &dependencyBuildLister{indexer: indexer}
}

// List lists all DependencyBuilds in the indexer.
func (s *dependencyBuildLister) List(selector labels.Selector) (ret []*v1alpha1.DependencyBuild, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DependencyBuild))
	})
	return ret, err
}

// DependencyBuilds returns an object that can list and get DependencyBuilds.
func (s *dependencyBuildLister) DependencyBuilds(namespace string) DependencyBuildNamespaceLister {
	return dependencyBuildNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DependencyBuildNamespaceLister helps list and get DependencyBuilds.
// All objects returned here must be treated as read-only.
type DependencyBuildNamespaceLister interface {
	// List lists all DependencyBuilds in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DependencyBuild, err error)
	// Get retrieves the DependencyBuild from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.DependencyBuild, error)
	DependencyBuildNamespaceListerExpansion
}

// dependencyBuildNamespaceLister implements the DependencyBuildNamespaceLister
// interface.
type dependencyBuildNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DependencyBuilds in the indexer for a given namespace.
func (s dependencyBuildNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DependencyBuild, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DependencyBuild))
	})
	return ret, err
}

// Get retrieves the DependencyBuild from the indexer for a given namespace and name.
func (s dependencyBuildNamespaceLister) Get(name string) (*v1alpha1.DependencyBuild, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("dependencybuild"), name)
	}
	return obj.(*v1alpha1.DependencyBuild), nil
}
