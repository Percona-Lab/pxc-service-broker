package pxccontroller

import (
	"log"

	v1 "github.com/Percona-Lab/pxc-service-broker/pkg/apis/pxc/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (p *Controller) DeployPXCCluster() error {
	q, err := resource.ParseQuantity("2Gi")
	if err != nil {
		log.Println(err)
	}

	none := "none"
	body := v1.PerconaXtraDBCluster{
		TypeMeta: metav1.TypeMeta{
			Kind:       "PerconaXtraDBCluster",
			APIVersion: "pxc.percona.com/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "some-name",
		},
		Spec: v1.PerconaXtraDBClusterSpec{
			SecretsName: "my-cluster-secrets",
			PXC: &v1.PodSpec{
				Size:  3,
				Image: "perconalab/percona-xtradb-cluster-operator:master-pxc",
				VolumeSpec: &v1.VolumeSpec{
					PersistentVolumeClaim: &corev1.PersistentVolumeClaimSpec{
						Resources: corev1.ResourceRequirements{
							Requests: corev1.ResourceList{corev1.ResourceStorage: q},
						},
					},
				},
				Resources: nil,
				Affinity: &v1.PodAffinity{
					TopologyKey: &none,
				},
			},
			ProxySQL: &v1.PodSpec{
				Size:  1,
				Image: "perconalab/percona-xtradb-cluster-operator:master-proxysql",
				VolumeSpec: &v1.VolumeSpec{
					PersistentVolumeClaim: &corev1.PersistentVolumeClaimSpec{
						Resources: corev1.ResourceRequirements{
							Requests: corev1.ResourceList{corev1.ResourceStorage: q},
						},
					},
				},
				Resources: nil,
				Affinity: &v1.PodAffinity{
					TopologyKey: &none,
				},
			},
			PMM: &v1.PMMSpec{
				Enabled:    false,
				Image:      "perconalab/pmm-client:1.17.1",
				ServerHost: "monitoring-service",
				ServerUser: "pmm",
			},
			Backup: &v1.PXCScheduledBackup{
				Image:              "perconalab/percona-xtradb-cluster-operator:master-backup",
				ServiceAccountName: "percona-xtradb-cluster-operator",
				Storages: map[string]*v1.BackupStorageSpec{
					"pvc": &v1.BackupStorageSpec{
						Type: "filesystem",
						Volume: &v1.VolumeSpec{
							PersistentVolumeClaim: &corev1.PersistentVolumeClaimSpec{
								AccessModes: []corev1.PersistentVolumeAccessMode{
									corev1.ReadWriteOnce,
								},
								Resources: corev1.ResourceRequirements{
									Requests: corev1.ResourceList{corev1.ResourceStorage: q},
								},
							},
						},
					},
				},
			},
		},
	}

	log.Println("Apply PXC")
	err = p.pxcClientSet.PxcV1().PerconaXtraDBClusters("myproject").Delete("some-name", nil)
	if err != nil {
		log.Printf("get pxc client: %v", err)
	}

	_, err = p.pxcClientSet.PxcV1().PerconaXtraDBClusters("myproject").Create(&body)
	if err != nil {
		return err
	}

	return nil
}
