/*
Copyright 2019 The OpenEBS Authors

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
	v1alpha1 "github.com/openebs/maya/pkg/apis/openebs.io/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// RunTaskLister helps list RunTasks.
// All objects returned here must be treated as read-only.
type RunTaskLister interface {
	// List lists all RunTasks in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.RunTask, err error)
	// RunTasks returns an object that can list and get RunTasks.
	RunTasks(namespace string) RunTaskNamespaceLister
	RunTaskListerExpansion
}

// runTaskLister implements the RunTaskLister interface.
type runTaskLister struct {
	indexer cache.Indexer
}

// NewRunTaskLister returns a new RunTaskLister.
func NewRunTaskLister(indexer cache.Indexer) RunTaskLister {
	return &runTaskLister{indexer: indexer}
}

// List lists all RunTasks in the indexer.
func (s *runTaskLister) List(selector labels.Selector) (ret []*v1alpha1.RunTask, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RunTask))
	})
	return ret, err
}

// RunTasks returns an object that can list and get RunTasks.
func (s *runTaskLister) RunTasks(namespace string) RunTaskNamespaceLister {
	return runTaskNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// RunTaskNamespaceLister helps list and get RunTasks.
// All objects returned here must be treated as read-only.
type RunTaskNamespaceLister interface {
	// List lists all RunTasks in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.RunTask, err error)
	// Get retrieves the RunTask from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.RunTask, error)
	RunTaskNamespaceListerExpansion
}

// runTaskNamespaceLister implements the RunTaskNamespaceLister
// interface.
type runTaskNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all RunTasks in the indexer for a given namespace.
func (s runTaskNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.RunTask, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RunTask))
	})
	return ret, err
}

// Get retrieves the RunTask from the indexer for a given namespace and name.
func (s runTaskNamespaceLister) Get(name string) (*v1alpha1.RunTask, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("runtask"), name)
	}
	return obj.(*v1alpha1.RunTask), nil
}
