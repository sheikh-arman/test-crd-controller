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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/sheikh-arman/test-crd-controller/pkg/apis/arman.com/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// BlackmanLister helps list Blackmans.
// All objects returned here must be treated as read-only.
type BlackmanLister interface {
	// List lists all Blackmans in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Blackman, err error)
	// Blackmans returns an object that can list and get Blackmans.
	Blackmans(namespace string) BlackmanNamespaceLister
	BlackmanListerExpansion
}

// blackmanLister implements the BlackmanLister interface.
type blackmanLister struct {
	indexer cache.Indexer
}

// NewBlackmanLister returns a new BlackmanLister.
func NewBlackmanLister(indexer cache.Indexer) BlackmanLister {
	return &blackmanLister{indexer: indexer}
}

// List lists all Blackmans in the indexer.
func (s *blackmanLister) List(selector labels.Selector) (ret []*v1alpha1.Blackman, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Blackman))
	})
	return ret, err
}

// Blackmans returns an object that can list and get Blackmans.
func (s *blackmanLister) Blackmans(namespace string) BlackmanNamespaceLister {
	return blackmanNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// BlackmanNamespaceLister helps list and get Blackmans.
// All objects returned here must be treated as read-only.
type BlackmanNamespaceLister interface {
	// List lists all Blackmans in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Blackman, err error)
	// Get retrieves the Blackman from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Blackman, error)
	BlackmanNamespaceListerExpansion
}

// blackmanNamespaceLister implements the BlackmanNamespaceLister
// interface.
type blackmanNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Blackmans in the indexer for a given namespace.
func (s blackmanNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Blackman, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Blackman))
	})
	return ret, err
}

// Get retrieves the Blackman from the indexer for a given namespace and name.
func (s blackmanNamespaceLister) Get(name string) (*v1alpha1.Blackman, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("blackman"), name)
	}
	return obj.(*v1alpha1.Blackman), nil
}
