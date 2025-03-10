//
// Copyright 2019 HAProxy Technologies LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by lister-gen. DO NOT EDIT.

package v3

import (
	v3 "github.com/haproxytech/kubernetes-ingress/crs/api/ingress/v3"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// BackendLister helps list Backends.
// All objects returned here must be treated as read-only.
type BackendLister interface {
	// List lists all Backends in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v3.Backend, err error)
	// Backends returns an object that can list and get Backends.
	Backends(namespace string) BackendNamespaceLister
	BackendListerExpansion
}

// backendLister implements the BackendLister interface.
type backendLister struct {
	indexer cache.Indexer
}

// NewBackendLister returns a new BackendLister.
func NewBackendLister(indexer cache.Indexer) BackendLister {
	return &backendLister{indexer: indexer}
}

// List lists all Backends in the indexer.
func (s *backendLister) List(selector labels.Selector) (ret []*v3.Backend, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v3.Backend))
	})
	return ret, err
}

// Backends returns an object that can list and get Backends.
func (s *backendLister) Backends(namespace string) BackendNamespaceLister {
	return backendNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// BackendNamespaceLister helps list and get Backends.
// All objects returned here must be treated as read-only.
type BackendNamespaceLister interface {
	// List lists all Backends in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v3.Backend, err error)
	// Get retrieves the Backend from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v3.Backend, error)
	BackendNamespaceListerExpansion
}

// backendNamespaceLister implements the BackendNamespaceLister
// interface.
type backendNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Backends in the indexer for a given namespace.
func (s backendNamespaceLister) List(selector labels.Selector) (ret []*v3.Backend, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v3.Backend))
	})
	return ret, err
}

// Get retrieves the Backend from the indexer for a given namespace and name.
func (s backendNamespaceLister) Get(name string) (*v3.Backend, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v3.Resource("backend"), name)
	}
	return obj.(*v3.Backend), nil
}
