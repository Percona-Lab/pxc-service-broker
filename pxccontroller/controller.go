package pxccontroller

import (
	"fmt"

	apiextension "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	pxcclientset "github.com/Percona-Lab/pxc-service-broker/pkg/client/clientset/versioned"
)

var scheme = runtime.NewScheme()

type Controller struct {
	clientSet    *kubernetes.Clientset
	pxcClientSet pxcclientset.Interface
}

func New() (Controller, error) {
	var pxc Controller

	config, err := rest.InClusterConfig()
	if err != nil {
		return pxc, err
	}

	crdClinetaSet, err := apiextension.NewForConfig(config)
	if err != nil {
		return pxc, err
	}

	err = CreateClusterCRD(crdClinetaSet)
	if err != nil {
		return pxc, err
	}

	err = CreateBackupCRD(crdClinetaSet)
	if err != nil {
		return pxc, err
	}
	err = CreateRestoreCRD(crdClinetaSet)
	if err != nil {
		return pxc, err
	}

	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		return pxc, err
	}
	err = ApplyRbac(clientSet)
	if err != nil {
		return pxc, err
	}

	pxcClient, err := pxcclientset.NewForConfig(config)
	if err != nil {
		return pxc, fmt.Errorf("getPXCClient: %v", err)
	}

	pxc.clientSet = clientSet
	pxc.pxcClientSet = pxcClient

	return pxc, nil
}
