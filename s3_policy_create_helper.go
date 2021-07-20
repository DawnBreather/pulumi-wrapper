package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

const (
	BUCKET_POLICY_DEFAULT_VERSION = "2012-10-17"
)

type policySetupCreateHelper struct{

	ctx *pulumi.Context

	s3Policies map[string]*s3.BucketPolicy

	conf          *S3PolicyConf
	s3PolicySetup *s3PolicySetup
}

type bucketPolicy struct {
	name string `json:"-"`
	s3Bucket string `json:"-"`

	version string `json:"Version"`
	statement []bucketPolicyStatement
}
func (bp *bucketPolicy) Statement() *bucketPolicyStatement{
	return &bucketPolicyStatement{}
}


type bucketPolicyStatement struct {
	name string `json:"Sid"`
	effect string `json:"Effect"`
	actions []bucketPolicyStatementAction `json:"Action"`
	resources []bucketPolicyStatementResource `json:"Resource"`
}

type bucketPolicyStatementAction string
type bucketPolicyStatementResource string


func (p *policySetupCreateHelper) Policy(name, s3Bucket string) *bucketPolicy{
	return &bucketPolicy{
		name:     name,
		s3Bucket: s3Bucket,
		version: BUCKET_POLICY_DEFAULT_VERSION,
	}
}