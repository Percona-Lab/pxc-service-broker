package pxccontroller

import (
	"k8s.io/api/rbac/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func ApplyRbac(clientSet *kubernetes.Clientset) error {
	role := v1beta1.ClusterRole{
		ObjectMeta: metav1.ObjectMeta{
			Name: "percona-xtradb-cluster-operator",
		},
		Rules: []v1beta1.PolicyRule{
			v1beta1.PolicyRule{
				APIGroups: []string{"pxc.percona.com"},
				Resources: []string{
					"perconaxtradbclusters",
					"perconaxtradbclusters/status",
					"perconaxtradbclusterbackups",
					"perconaxtradbclusterbackups/status",
					"perconaxtradbclusterrestores",
					"perconaxtradbclusterrestores/status",
				},
				Verbs: []string{
					"get",
					"list",
					"watch",
					"create",
					"update",
					"patch",
					"delete",
				},
			},
			v1beta1.PolicyRule{
				APIGroups: []string{""},
				Resources: []string{
					"pods",
					"pods/exec",
					"configmaps",
					"services",
					"persistentvolumeclaims",
					"secrets",
				},
				Verbs: []string{
					"get",
					"list",
					"watch",
					"create",
					"update",
					"patch",
					"delete",
				},
			},
			v1beta1.PolicyRule{
				APIGroups: []string{"apps"},
				Resources: []string{
					"deployments",
					"replicasets",
					"statefulsets",
				},
				Verbs: []string{
					"get",
					"list",
					"watch",
					"create",
					"update",
					"patch",
					"delete",
				},
			},
			v1beta1.PolicyRule{
				APIGroups: []string{"batch"},
				Resources: []string{
					"jobs",
					"cronjobs",
				},
				Verbs: []string{
					"get",
					"list",
					"watch",
					"create",
					"update",
					"patch",
					"delete",
				},
			},
			v1beta1.PolicyRule{
				APIGroups: []string{"policy"},
				Resources: []string{
					"poddisruptionbudgets",
				},
				Verbs: []string{
					"get",
					"list",
					"watch",
					"create",
					"update",
					"patch",
					"delete",
				},
			},
		},
	}
	rbac := clientSet.RbacV1beta1()
	rbac.ClusterRoles().Create(&role)

	bind := v1beta1.ClusterRoleBinding{
		ObjectMeta: metav1.ObjectMeta{
			Name: "service-account-percona-xtradb-cluster-operator",
		},
		Subjects: []v1beta1.Subject{
			v1beta1.Subject{
				Kind: "ServiceAccount",
				Name: "percona-xtradb-cluster-operator",
			},
		},
		RoleRef: v1beta1.RoleRef{
			Kind:     "ClusterRole",
			Name:     "percona-xtradb-cluster-operator",
			APIGroup: "rbac.authorization.k8s.io",
		},
	}
	roleBind := rbac.ClusterRoleBindings()
	roleBind.Create(&bind)
	return nil
}

/*
func ApplyCRDRbac(clientSet *kubernetes.Clientset) error {
	role := v1beta1.ClusterRole{
		ObjectMeta: metav1.ObjectMeta{
			Name: "pxc-service-broker",
		},
		Rules: []v1beta1.PolicyRule{
			v1beta1.PolicyRule{
				APIGroups: []string{"apiextensions.k8s.io"},
				Resources: []string{
					"customresourcedefinitions",
					"customresourcedefinition",
				},
				Verbs: []string{
					"get",
					"list",
					"watch",
					"create",
					"update",
					"patch",
					"delete",
				},
			},
		},
	}
	rbac := clientSet.RbacV1beta1()

	_, err := rbac.ClusterRoles().Create(&role)
	if err != nil {
		return err
	}

	bind := v1beta1.ClusterRoleBinding{
		ObjectMeta: metav1.ObjectMeta{
			Name: "pcx-dervice-broker-crd",
		},
		Subjects: []v1beta1.Subject{
			v1beta1.Subject{
				Kind:     "User",
				Name:     "system:serviceaccount:test:default",
				APIGroup: "rbac.authorization.k8s.io",
			},
		},
		RoleRef: v1beta1.RoleRef{
			Kind:     "ClusterRole",
			Name:     "pxc-service-broker",
			APIGroup: "rbac.authorization.k8s.io",
		},
	}
	roleBind := rbac.ClusterRoleBindings()
	_, err = roleBind.Create(&bind)
	if err != nil {
		return err
	}
	bind2 := v1beta1.ClusterRoleBinding{
		ObjectMeta: metav1.ObjectMeta{
			Name: "pxc-service-broker-admin",
		},
		Subjects: []v1beta1.Subject{
			v1beta1.Subject{
				Kind:     "User",
				Name:     "system:serviceaccount:test:default",
				APIGroup: "rbac.authorization.k8s.io",
			},
		},
		RoleRef: v1beta1.RoleRef{
			Kind:     "ClusterRole",
			Name:     "pxc-service-broker-admin",
			APIGroup: "rbac.authorization.k8s.io",
		},
	}

	_, err = roleBind.Create(&bind2)
	if err != nil {
		return err
	}
	return nil
}
*/
