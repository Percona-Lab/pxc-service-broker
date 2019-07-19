package pxccontroller

import (
	"fmt"
	"log"

	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (p *Controller) DeployPXCOperator() error {
	deployClient := p.clientSet.AppsV1().Deployments("myproject")

	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: "percona-xtradb-cluster-operator",
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"name": "percona-xtradb-cluster-operator",
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"name": "percona-xtradb-cluster-operator",
					},
				},
				Spec: apiv1.PodSpec{
					ServiceAccountName: "percona-xtradb-cluster-operator",
					Containers: []apiv1.Container{
						apiv1.Container{
							Name:  "percona-xtradb-cluster-operator",
							Image: "percona/percona-xtradb-cluster-operator:1.1.0",
							Ports: []apiv1.ContainerPort{
								{
									Name:          "metrics",
									ContainerPort: 60000,
								},
							},
							Command:         []string{"percona-xtradb-cluster-operator"},
							ImagePullPolicy: apiv1.PullAlways,
							Env: []apiv1.EnvVar{
								apiv1.EnvVar{
									Name: "WATCH_NAMESPACE",
									ValueFrom: &apiv1.EnvVarSource{
										FieldRef: &apiv1.ObjectFieldSelector{
											FieldPath: "metadata.namespace",
										},
									},
								},
								apiv1.EnvVar{
									Name:  "OPERATOR_NAME",
									Value: "percona-xtradb-cluster-operator",
								},
							},
						},
					},
				},
			},
		},
	}

	// Create Deployment
	fmt.Println("Creating deployment...")
	err := deployClient.Delete("percona-xtradb-cluster-operator", nil)
	if err != nil {
		log.Println("Delete error:", err)
	}

	result, err := deployClient.Create(deployment)
	if err != nil {
		return err
	}
	fmt.Printf("Created deployment %q.\n", result.GetObjectMeta().GetName())

	/*fmt.Printf("Listing deployments in namespace %q:\n", apiv1.NamespaceDefault)
	list, err := deployClient.List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, d := range list.Items {
		fmt.Printf(" * %s (%d replicas)\n", d.Name, *d.Spec.Replicas)
	}*/
	return nil
}

func int32Ptr(i int32) *int32 { return &i }
