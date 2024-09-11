package main

import (
	"fmt"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		conf := config.New(ctx, "")
		roundName := conf.Require("round-name")
		vpcCidr := conf.Require("vpc-cidr")

		vpcName := fmt.Sprintf("eks-vpc-%s", roundName)
		vpc, err := NewVpc(ctx, vpcName, vpcCidr)
		if err != nil {
			return err
		}

		ctx.Export("vpc.ID", vpc.ID())
		return nil
	})
}
