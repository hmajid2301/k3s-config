package main

import (
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	helmv3 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/helm/v3"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func addHA(ctx *pulumi.Context) error {
	namespace := "apps"

	_, err := corev1.NewNamespace(ctx, namespace, &corev1.NamespaceArgs{
		Metadata: &metav1.ObjectMetaArgs{
			Name: pulumi.String(namespace),
		},
	})
	if err != nil {
		return nil
	}

	_, err = helmv3.NewChart(ctx, "home-assistant", helmv3.ChartArgs{
		Chart:   pulumi.String("home-assistant"),
		Version: pulumi.String("0.2.59"),
		FetchArgs: helmv3.FetchArgs{
			Repo: pulumi.String(`http://pajikos.github.io/home-assistant-helm-chart/`),
		},
		Namespace: pulumi.String(namespace),
	})
	if err != nil {
		return err
	}

	return nil
}
