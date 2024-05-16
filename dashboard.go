package main

import (
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	helmv3 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/helm/v3"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	rbacv1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/rbac/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func addDashboard(ctx *pulumi.Context) error {
	namespace := "monitoring"

	_, err := helmv3.NewChart(ctx, "kubernetes-dashboard", helmv3.ChartArgs{
		Chart: pulumi.String("kubernetes-dashboard"),
		FetchArgs: helmv3.FetchArgs{
			Repo: pulumi.String(`https://kubernetes.github.io/dashboard/`),
		},
		Namespace: pulumi.String(namespace),
	})
	if err != nil {
		return err
	}

	_, err = corev1.NewNamespace(ctx, namespace, &corev1.NamespaceArgs{
		Metadata: &metav1.ObjectMetaArgs{
			Name: pulumi.String(namespace),
		},
	})
	if err != nil {
		return nil
	}

	svc, err := corev1.NewServiceAccount(ctx, "kubernetes-dashboard", &corev1.ServiceAccountArgs{
		Metadata: metav1.ObjectMetaArgs{
			Name:      pulumi.String("admin-user"),
			Namespace: pulumi.String(namespace),
		},
	})
	if err != nil {
		return err
	}

	_, err = rbacv1.NewClusterRoleBinding(ctx, "cluster-admin-binding", &rbacv1.ClusterRoleBindingArgs{
		Metadata: metav1.ObjectMetaArgs{
			Name: pulumi.String("admin-user"),
		},
		RoleRef: rbacv1.RoleRefArgs{
			Name:     pulumi.String("cluster-admin"),
			ApiGroup: pulumi.String("rbac.authorization.k8s.io"),
			Kind:     pulumi.String("ClusterRole"),
		},
		Subjects: rbacv1.SubjectArray{
			rbacv1.SubjectArgs{
				Kind:      pulumi.String("ServiceAccount"),
				Name:      svc.Metadata.Name().Elem(),
				Namespace: pulumi.String(namespace),
			},
		},
	})
	if err != nil {
		return err
	}

	return nil
}
