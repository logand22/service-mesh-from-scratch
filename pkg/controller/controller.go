package controller

import (
	"time"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/informers"
	informerv1 "k8s.io/client-go/informers/core/v1"
	"k8s.io/client-go/kubernetes"
)

const DefaultResyncTimeout = time.Minute * 10

type Controller struct {
	si informerv1.ServiceInformer
	ei informerv1.EndpointsInformer

	sf informers.SharedInformerFactory
}

// NewController returns a controller that knows how to return services using k8s shared informers
// Must call <foo>.Informer() to register it in map of informers.
// This makes call to start work.
func NewController(kubeClient *kubernetes.Clientset) *Controller {
	control := &Controller{}
	control.sf = informers.NewSharedInformerFactory(kubeClient, DefaultResyncTimeout)

	control.si = control.sf.Core().V1().Services()
	control.si.Informer()

	return control
}

// Start the registered shared informers
func (c *Controller) Start(stopCh <-chan struct{}) {
	c.sf.Start(stopCh)
}

// GetServices returns services from the controller. If namespace is empty returns all services
// If name is empty, returns all services in a namespace
// If both are provided gets a specific service from the controller.
func (c *Controller) GetServices(namespace, name string, selector labels.Selector) ([]*v1.Service, error) {
	var services []*v1.Service
	var err error

	if namespace == "" {
		services, err = c.si.Lister().List(selector)
	} else if name == "" {
		services, err = c.si.Lister().Services(namespace).List(selector)
	} else {
		var service *v1.Service
		service, err = c.si.Lister().Services(namespace).Get(name)
		services = []*v1.Service{service}
	}

	return services, err
}

// GetEndpoints returns endpoints from the controller. If namespace is empty returns all endpoints
// If name is empty, returns all endpoints in a namespace
// If both are provided gets a specific endpoint from the controller.
func (c *Controller) GetEndpoints(namespace, name string, selector labels.Selector) ([]*v1.Endpoints, error) {
	var endpoints []*v1.Endpoints
	var err error

	if namespace == "" {
		endpoints, err = c.ei.Lister().List(selector)
	} else if name == "" {
		endpoints, err = c.ei.Lister().Endpoints(namespace).List(selector)
	} else {
		var endpoint *v1.Endpoints
		endpoint, err = c.ei.Lister().Endpoints(namespace).Get(name)
		endpoints = []*v1.Endpoints{endpoint}
	}

	return endpoints, err
}
