package main

import (
	"fmt"
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type s3Setup struct{

	CREATE s3SetupCreateHelper
}

func (s *s3Setup) Init(ctx *pulumi.Context) *s3Setup{
	s.CREATE = s3SetupCreateHelper{
		ctx:       ctx,
		s3Buckets: map[string]*s3.Bucket{},
		conf:      &S3Conf{},
		s3Setup:   s,
	}

	return s
}

func (s *s3Setup) SetRegion(region string) *s3Setup {
	s.CREATE.conf.Region = String(region)
	return s
}
func (s *s3Setup) SetGroupName(groupName string) *s3Setup{
	s.CREATE.conf.GroupName = String(groupName)
	return s
}

func (s *s3Setup) SetEnvironmentName(environmentName string) *s3Setup{
	s.CREATE.conf.EnvironmentName = String(environmentName)
	return s
}

func (s *s3Setup) SetPulumiStackName(pulumiStackName string) *s3Setup{
	s.CREATE.conf.PulumiStackName = String(pulumiStackName)
	return s
}

func (s *s3Setup) Export(){
	for name, obj := range s.CREATE.s3Buckets {
		s.CREATE.ctx.Export(fmt.Sprintf("s3-%s-name", name), pulumi.String(name))
		s.CREATE.ctx.Export(fmt.Sprintf("s3-%s-arn", name), obj.Arn)
		s.CREATE.ctx.Export(fmt.Sprintf("s3-%s-domain", name), obj.BucketDomainName)
	}
}