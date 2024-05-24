package main

import (
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	helmv3 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/helm/v3"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func addImmich(ctx *pulumi.Context) error {
	namespace := "apps"

	_, err := corev1.NewPersistentVolume(ctx, "pv", &corev1.PersistentVolumeArgs{
		Spec: &corev1.PersistentVolumeSpecArgs{
			Capacity: pulumi.StringMap{
				"storage": pulumi.String("200Gi"),
			},
			VolumeMode:       pulumi.StringPtr("Filesystem"),
			AccessModes:      pulumi.StringArray{pulumi.String("ReadWriteOnce")},
			StorageClassName: pulumi.String("local-storage"),
			Local: &corev1.LocalVolumeSourceArgs{
				Path: pulumi.String("/data"),
			},
			PersistentVolumeReclaimPolicy: pulumi.String("Retain"),
			NodeAffinity: &corev1.VolumeNodeAffinityArgs{
				Required: &corev1.NodeSelectorArgs{
					NodeSelectorTerms: corev1.NodeSelectorTermArray{
						&corev1.NodeSelectorTermArgs{
							MatchExpressions: corev1.NodeSelectorRequirementArray{
								&corev1.NodeSelectorRequirementArgs{
									Key:      pulumi.String("kubernetes.io/hostname"),
									Operator: pulumi.String("In"),
									Values:   pulumi.StringArray{pulumi.String("frameworkedup")},
								},
							},
						},
					},
				},
			},
		},
	})
	if err != nil {
		return err
	}

	pvc, err := corev1.NewPersistentVolumeClaim(ctx, "immich-pvc", &corev1.PersistentVolumeClaimArgs{
		Spec: corev1.PersistentVolumeClaimSpecArgs{
			AccessModes: pulumi.StringArray{
				pulumi.String("ReadWriteOnce"),
			},
			Resources: corev1.VolumeResourceRequirementsArgs{
				Requests: pulumi.StringMap{
					"storage": pulumi.String("100Gi"),
				},
			},
			StorageClassName: pulumi.String("local-storage"),
		},
		Metadata: metav1.ObjectMetaArgs{
			Namespace: pulumi.String(namespace),
		},
	})
	if err != nil {
		return err
	}

	_, err = helmv3.NewChart(ctx, "immich", helmv3.ChartArgs{
		Chart:   pulumi.String("immich"),
		Version: pulumi.String("0.6.0"),
		FetchArgs: helmv3.FetchArgs{
			Repo: pulumi.String(`https://immich-app.github.io/immich-charts`),
		},
		Namespace: pulumi.String(namespace),
		Values: pulumi.Map{
			"immich": pulumi.Map{
				"metrics": pulumi.Map{
					"enabled": pulumi.Bool(true),
				},
				"persistence": pulumi.Map{
					"library": pulumi.Map{
						"existingClaim": pvc.Metadata.Name().Elem(),
					},
				},
			},
			"image": pulumi.Map{
				"tag": pulumi.String("v1.105.1"),
			},
			"postgresql": pulumi.Map{
				"enabled": pulumi.Bool(true),
			},
			"redis": pulumi.Map{
				"enabled": pulumi.Bool(true),
			},
		},
	})
	if err != nil {
		return err
	}

	return nil
}
