package main

import (
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	helmv3 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/helm/v3"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func addGitlabRunner(ctx *pulumi.Context) error {
	namespace := "gitlab-runner"
	cfg := config.New(ctx, "")

	_, err := corev1.NewNamespace(ctx, namespace, &corev1.NamespaceArgs{
		Metadata: &metav1.ObjectMetaArgs{
			Name: pulumi.String(namespace),
		},
	})
	if err != nil {
		return nil
	}

	_, err = helmv3.NewChart(ctx, "gitlab-runner", helmv3.ChartArgs{
		Chart:   pulumi.String("gitlab-runner"),
		Version: pulumi.String("v0.64.1"),
		FetchArgs: helmv3.FetchArgs{
			Repo: pulumi.String(`https://charts.gitlab.io`),
		},
		Namespace: pulumi.String(namespace),
		Values: pulumi.Map{
			"gitlabUrl": pulumi.String("https://gitlab.com/"),
			"rbac": pulumi.Map{
				"create": pulumi.Bool(true),
			},
			"runnerToken":             cfg.GetSecret("gitlabRunnerToken"),
			"runnerRegistrationToken": pulumi.String(""),
		},
	})
	if err != nil {
		return err
	}
	return nil

}
