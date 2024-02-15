package zvirt

import (
	"log"
	"time"

	"k8s.io/client-go/informers"
	cloudprovider "k8s.io/cloud-provider"

	"k8s.io/client-go/tools/cache"
)

const (
	providerName = "zvirt"
)

type Cloud struct{}

func NewCloud() *Cloud {
	return &Cloud{}
}

func (zc *Cloud) Initialize(
	clientBuilder cloudprovider.ControllerClientBuilder,
	stop <-chan struct{},
) {
	clientset := clientBuilder.ClientOrDie("cloud-controller-manager")

	informerFactory := informers.NewSharedInformerFactory(clientset, time.Second*30)
	serviceInformer := informerFactory.Core().V1().Services()
	nodeInformer := informerFactory.Core().V1().Nodes()

	go serviceInformer.Informer().Run(stop)
	go nodeInformer.Informer().Run(stop)

	if !cache.WaitForCacheSync(stop, serviceInformer.Informer().HasSynced) {
		log.Fatal("Timed out waiting for caches to sync")
	}
	if !cache.WaitForCacheSync(stop, nodeInformer.Informer().HasSynced) {
		log.Fatal("Timed out waiting for caches to sync")
	}
}

// LoadBalancer returns a balancer interface if supported.
func (zc *Cloud) LoadBalancer() (cloudprovider.LoadBalancer, bool) {
	return nil, false
}

// Instances returns an instances interface if supported.
func (zc *Cloud) Instances() (cloudprovider.Instances, bool) {
	return nil, false
}

// Zones returns a zones interface if supported.
func (zc *Cloud) Zones() (cloudprovider.Zones, bool) {
	return nil, false
}

// Clusters returns a clusters interface if supported.
func (zc *Cloud) Clusters() (cloudprovider.Clusters, bool) {
	return nil, false
}

// Routes returns a routes interface if supported
func (zc *Cloud) Routes() (cloudprovider.Routes, bool) {
	return nil, false
}

// ProviderName returns the cloud provider ID.
func (zc *Cloud) ProviderName() string {
	return providerName
}

// HasClusterID returns true if the cluster has a clusterID
func (zc *Cloud) HasClusterID() bool {
	return true
}

func (zc *Cloud) InstancesV2() (cloudprovider.InstancesV2, bool) {
	return nil, false
}
