package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type S3PolicyConf struct {
	PulumiStackName String

	EnvironmentName String
	GroupName       String
	Region          String
}

func (s *S3PolicyConf) GetTags() pulumi.StringMap{
	return pulumi.StringMap{
		"GroupName":       s.GroupName.S(),
		"EnvironmentName": s.EnvironmentName.S(),
		"PulumiStack":     s.PulumiStackName.S(),
	}
}
//func (s *S3PolicyConf) GetBucketName(name string) String{
//	return String(fmt.Sprintf("s3-%s-%s-%s", name, s.EnvironmentName, s.GroupName))
//}