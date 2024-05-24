package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		err := addDashboard(ctx)
		if err != nil {
			return err
		}

		err = addMonitoring(ctx)
		if err != nil {
			return err
		}

		err = addHA(ctx)
		if err != nil {
			return err
		}

		err = addImmich(ctx)
		if err != nil {
			return err
		}

		err = addGitlabRunner(ctx)
		if err != nil {
			return err
		}

		return nil
	})
}
