package main

import (
	helmv3 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/helm/v3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func addMonitoring(ctx *pulumi.Context) error {
	namespace := "monitoring"

	_, err := helmv3.NewRelease(ctx, "prometheus", &helmv3.ReleaseArgs{
		Chart:   pulumi.String("kube-prometheus-stack"),
		Version: pulumi.String("58.6.0"),
		RepositoryOpts: helmv3.RepositoryOptsArgs{
			Repo: pulumi.String("https://prometheus-community.github.io/helm-charts"),
		},
		Namespace: pulumi.String(namespace),
	})
	if err != nil {
		return err
	}
	return nil

}
