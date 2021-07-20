package main

import "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

type VpcConf struct {
	PulumiStackName String

	EnvironmentName String
	GroupName       String
	Name            String
	Region          String
}

func (v *VpcConf) GetTags() pulumi.StringMap{
	return pulumi.StringMap{
		"GroupName": v.GroupName.S(),
		"EnvironmentName": v.EnvironmentName.S(),
		"PulumiStack": v.PulumiStackName.S(),
	}
}