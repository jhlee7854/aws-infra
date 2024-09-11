package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func NewVpc(ctx *pulumi.Context, name string, cidr string) (*ec2.Vpc, error) {
	return ec2.NewVpc(ctx, name, &ec2.VpcArgs{
		CidrBlock:          pulumi.String(cidr),
		EnableDnsHostnames: pulumi.Bool(true),
		EnableDnsSupport:   pulumi.Bool(true),
		InstanceTenancy:    pulumi.String("default"),
		Tags: pulumi.StringMap{
			"Name": pulumi.String(name),
		},
	})
}
