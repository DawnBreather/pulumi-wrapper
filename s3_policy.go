package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type s3PolicySetup struct{

	CREATE policySetupCreateHelper
}

func (p *s3PolicySetup) Init(ctx *pulumi.Context) *s3PolicySetup {
	p.CREATE = policySetupCreateHelper{
		ctx:           ctx,
		s3Policies:    map[string]*s3.BucketPolicy{},
		conf:          &S3PolicyConf{},
		s3PolicySetup: p,
	}

	return p
}

//func (p *s3PolicySetup) SetRegion(region string) *s3PolicySetup {
//	p.CREATE.conf.Region = String(region)
//	return p
//}
func (p *s3PolicySetup) SetGroupName(groupName string) *s3PolicySetup {
	p.CREATE.conf.GroupName = String(groupName)
	return p
}

func (p *s3PolicySetup) SetEnvironmentName(environmentName string) *s3PolicySetup {
	p.CREATE.conf.EnvironmentName = String(environmentName)
	return p
}

func (p *s3PolicySetup) SetPulumiStackName(pulumiStackName string) *s3PolicySetup {
	p.CREATE.conf.PulumiStackName = String(pulumiStackName)
	return p
}

//func (p *s3PolicySetup) Export(){
//	for name, obj := range p.CREATE.s3Policies {
//		p.CREATE.ctx.Export(fmt.Sprintf("s3-%p-name", name), pulumi.String(name))
//		p.CREATE.ctx.Export(fmt.Sprintf("s3-%p-arn", name), obj.Arn)
//		p.CREATE.ctx.Export(fmt.Sprintf("s3-%p-domain", name), obj.BucketDomainName)
//	}
//}