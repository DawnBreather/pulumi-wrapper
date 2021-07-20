package main

import (
	"fmt"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type S3Conf struct {
	PulumiStackName String

	EnvironmentName String
	GroupName       String
	Region          String
}

func (s *S3Conf) GetTags() pulumi.StringMap{
	return pulumi.StringMap{
		"GroupName":       s.GroupName.S(),
		"EnvironmentName": s.EnvironmentName.S(),
		"PulumiStack":     s.PulumiStackName.S(),
	}
}
func (s *S3Conf) GetBucketName(name string) String{
	return String(fmt.Sprintf("s3-%s-%s-%s", name, s.EnvironmentName, s.GroupName))
}