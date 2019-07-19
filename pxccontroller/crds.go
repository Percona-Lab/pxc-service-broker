package pxccontroller

import (
	apiextensionv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	apiextension "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func CreateClusterCRD(clientset apiextension.Interface) error {
	crd := &apiextensionv1beta1.CustomResourceDefinition{
		ObjectMeta: metav1.ObjectMeta{Name: "perconaxtradbclusters.pxc.percona.com"},
		Spec: apiextensionv1beta1.CustomResourceDefinitionSpec{
			Group: "pxc.percona.com",
			//Version: CRDVersion,
			Scope: apiextensionv1beta1.NamespaceScoped,
			Names: apiextensionv1beta1.CustomResourceDefinitionNames{
				Plural:   "perconaxtradbclusters",
				Singular: "perconaxtradbcluster",
				Kind:     "PerconaXtraDBCluster",
				ListKind: "PerconaXtraDBClusterList",
				ShortNames: []string{
					"pxc",
					"pxcs",
				},
			},
			Versions: []apiextensionv1beta1.CustomResourceDefinitionVersion{
				apiextensionv1beta1.CustomResourceDefinitionVersion{
					Name:    "v1",
					Storage: true,
					Served:  true,
				},
				apiextensionv1beta1.CustomResourceDefinitionVersion{
					Name:    "v1alpha1",
					Storage: false,
					Served:  true,
				},
			},
			AdditionalPrinterColumns: []apiextensionv1beta1.CustomResourceColumnDefinition{
				apiextensionv1beta1.CustomResourceColumnDefinition{
					Name:     "Endpoint",
					Type:     "string",
					JSONPath: ".status.host",
				},
				apiextensionv1beta1.CustomResourceColumnDefinition{
					Name:     "Status",
					Type:     "string",
					JSONPath: ".status.state",
				},
				apiextensionv1beta1.CustomResourceColumnDefinition{
					Name:        "PXC",
					Type:        "string",
					JSONPath:    ".status.pxc.ready",
					Description: "Ready pxc nodes",
				},
				apiextensionv1beta1.CustomResourceColumnDefinition{
					Name:        "proxysql",
					Type:        "string",
					JSONPath:    ".status.proxysql.ready",
					Description: "Ready pxc nodes",
				},
				apiextensionv1beta1.CustomResourceColumnDefinition{
					Name:     "Age",
					Type:     "date",
					JSONPath: ".metadata.creationTimestamp",
				},
			},
		},
	}

	_, err := clientset.ApiextensionsV1beta1().CustomResourceDefinitions().Create(crd)
	if err != nil && apierrors.IsAlreadyExists(err) {
		return nil
	}
	return err
}

func CreateBackupCRD(clientset apiextension.Interface) error {
	crd := &apiextensionv1beta1.CustomResourceDefinition{
		ObjectMeta: metav1.ObjectMeta{Name: "perconaxtradbclusterbackups.pxc.percona.com"},
		Spec: apiextensionv1beta1.CustomResourceDefinitionSpec{
			Group: "pxc.percona.com",
			//Version: CRDVersion,
			Scope: apiextensionv1beta1.NamespaceScoped,
			Names: apiextensionv1beta1.CustomResourceDefinitionNames{
				Plural:   "perconaxtradbclusterbackups",
				Singular: "perconaxtradbclusterbackup",
				Kind:     "PerconaXtraDBClusterBackup",
				ListKind: "PerconaXtraDBClusterBackupList",
				ShortNames: []string{
					"pxc-backup",
					"pxc-backups",
				},
			},
			Versions: []apiextensionv1beta1.CustomResourceDefinitionVersion{
				apiextensionv1beta1.CustomResourceDefinitionVersion{
					Name:    "v1",
					Storage: true,
					Served:  true,
				},
			},
			AdditionalPrinterColumns: []apiextensionv1beta1.CustomResourceColumnDefinition{
				apiextensionv1beta1.CustomResourceColumnDefinition{
					Name:        "Cluster",
					Type:        "string",
					Description: "Cluster name",
					JSONPath:    ".spec.pxcCluster",
				},
				apiextensionv1beta1.CustomResourceColumnDefinition{
					Name:        "Stotage",
					Type:        "string",
					Description: "Storage name from pxc spec",
					JSONPath:    ".status.storageName",
				},
				apiextensionv1beta1.CustomResourceColumnDefinition{
					Name:        "Destination",
					Type:        "string",
					Description: "Backup destination",
					JSONPath:    ".status.destination",
				},
				apiextensionv1beta1.CustomResourceColumnDefinition{
					Name:        "Status",
					Type:        "string",
					Description: "Job status",
					JSONPath:    ".status.state",
				},
				apiextensionv1beta1.CustomResourceColumnDefinition{
					Name:        "Completed",
					Description: "Completed time",
					Type:        "date",
					JSONPath:    ".status.completed",
				},
				apiextensionv1beta1.CustomResourceColumnDefinition{
					Name:     "Age",
					Type:     "date",
					JSONPath: ".metadata.creationTimestamp",
				},
			},
		},
	}

	_, err := clientset.ApiextensionsV1beta1().CustomResourceDefinitions().Create(crd)
	if err != nil && apierrors.IsAlreadyExists(err) {
		return nil
	}
	return err
}

func CreateRestoreCRD(clientset apiextension.Interface) error {
	crd := &apiextensionv1beta1.CustomResourceDefinition{
		ObjectMeta: metav1.ObjectMeta{Name: "perconaxtradbclusterrestores.pxc.percona.com"},
		Spec: apiextensionv1beta1.CustomResourceDefinitionSpec{
			Group: "pxc.percona.com",
			//Version: CRDVersion,
			Scope: apiextensionv1beta1.NamespaceScoped,
			Names: apiextensionv1beta1.CustomResourceDefinitionNames{
				Plural:   "perconaxtradbclusterrestores",
				Singular: "perconaxtradbclusterrestore",
				Kind:     "PerconaXtraDBClusterRestore",
				ListKind: "PerconaXtraDBClusterRestoreList",
				ShortNames: []string{
					"pxc-restore",
					"pxc-restores",
				},
			},
			Versions: []apiextensionv1beta1.CustomResourceDefinitionVersion{
				apiextensionv1beta1.CustomResourceDefinitionVersion{
					Name:    "v1",
					Storage: true,
					Served:  true,
				},
			},
			AdditionalPrinterColumns: []apiextensionv1beta1.CustomResourceColumnDefinition{
				apiextensionv1beta1.CustomResourceColumnDefinition{
					Name:        "Cluster",
					Type:        "string",
					Description: "Cluster name",
					JSONPath:    ".spec.pxcCluster",
				},
				apiextensionv1beta1.CustomResourceColumnDefinition{
					Name:        "Status",
					Type:        "string",
					Description: "Job status",
					JSONPath:    ".status.state",
				},
				apiextensionv1beta1.CustomResourceColumnDefinition{
					Name:        "Completed",
					Description: "Completed time",
					Type:        "date",
					JSONPath:    ".status.completed",
				},
				apiextensionv1beta1.CustomResourceColumnDefinition{
					Name:     "Age",
					Type:     "date",
					JSONPath: ".metadata.creationTimestamp",
				},
			},
		},
	}

	_, err := clientset.ApiextensionsV1beta1().CustomResourceDefinitions().Create(crd)
	if err != nil && apierrors.IsAlreadyExists(err) {
		return nil
	}
	return err
}
