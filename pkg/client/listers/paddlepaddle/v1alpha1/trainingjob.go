/*
Copyright (c) 2016 PaddlePaddle Authors All Rights Reserved.

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
// This file was automatically generated by lister-gen

package v1alpha1

import (
	v1alpha1 "github.com/paddleflow/paddle-operator/pkg/apis/paddlepaddle/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// TrainingJobLister helps list TrainingJobs.
type TrainingJobLister interface {
	// List lists all TrainingJobs in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.TrainingJob, err error)
	// TrainingJobs returns an object that can list and get TrainingJobs.
	TrainingJobs(namespace string) TrainingJobNamespaceLister
	TrainingJobListerExpansion
}

// trainingJobLister implements the TrainingJobLister interface.
type trainingJobLister struct {
	indexer cache.Indexer
}

// NewTrainingJobLister returns a new TrainingJobLister.
func NewTrainingJobLister(indexer cache.Indexer) TrainingJobLister {
	return &trainingJobLister{indexer: indexer}
}

// List lists all TrainingJobs in the indexer.
func (s *trainingJobLister) List(selector labels.Selector) (ret []*v1alpha1.TrainingJob, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.TrainingJob))
	})
	return ret, err
}

// TrainingJobs returns an object that can list and get TrainingJobs.
func (s *trainingJobLister) TrainingJobs(namespace string) TrainingJobNamespaceLister {
	return trainingJobNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// TrainingJobNamespaceLister helps list and get TrainingJobs.
type TrainingJobNamespaceLister interface {
	// List lists all TrainingJobs in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.TrainingJob, err error)
	// Get retrieves the TrainingJob from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.TrainingJob, error)
	TrainingJobNamespaceListerExpansion
}

// trainingJobNamespaceLister implements the TrainingJobNamespaceLister
// interface.
type trainingJobNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all TrainingJobs in the indexer for a given namespace.
func (s trainingJobNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.TrainingJob, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.TrainingJob))
	})
	return ret, err
}

// Get retrieves the TrainingJob from the indexer for a given namespace and name.
func (s trainingJobNamespaceLister) Get(name string) (*v1alpha1.TrainingJob, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("trainingjob"), name)
	}
	return obj.(*v1alpha1.TrainingJob), nil
}