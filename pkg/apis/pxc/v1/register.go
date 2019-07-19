package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/Percona-Lab/pxc-service-broker/pkg/apis/pxc"
)

// GroupVersion is the identifier for the API which includes
// the name of the group and the version of the API
var SchemeGroupVersion = schema.GroupVersion{
	Group:   pxc.GroupName,
	Version: "v1",
}

/*
var (
	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	SchemeBuilder = &scheme.Builder{GroupVersion: SchemeGroupVersion}
)

func init() {
	SchemeBuilder.Register(
		&PerconaXtraDBCluster{}, &PerconaXtraDBClusterList{},
		&PerconaXtraDBClusterBackup{}, &PerconaXtraDBClusterBackupList{},
	)
}*/

// create a SchemeBuilder which uses functions to add types to
// the scheme
var AddPXCToScheme = runtime.NewSchemeBuilder(addKnownTypes)
var AddToScheme = AddPXCToScheme.AddToScheme

func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

// addKnownTypes adds our types to the API scheme by registering
// MyResource and MyResourceList
func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(
		SchemeGroupVersion,
		&PerconaXtraDBCluster{}, &PerconaXtraDBClusterList{},
		&PerconaXtraDBClusterBackup{}, &PerconaXtraDBClusterBackupList{},
		&PerconaXtraDBClusterRestore{}, &PerconaXtraDBClusterRestoreList{},
	)

	// register the type in the scheme
	meta_v1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
