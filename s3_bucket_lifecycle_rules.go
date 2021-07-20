package main

import "github.com/pulumi/pulumi-aws/sdk/v4/go/aws/s3"

type s3LifecycleRules []s3LifecycleRule
func (s s3LifecycleRules) ToPulumiLifecycleRuleArray() (res s3.BucketLifecycleRuleArray){
	for _, i := range s{
		res = append(res, i.ToPulumiLifecycleRuleArgs())
	}

	return res
}

type s3LifecycleRule struct{
	Id String
	Enabled Bool
	AbortIncompleteMultipartUploadDays Int
}

func (s *s3LifecycleRule) ToPulumiLifecycleRuleArgs() s3.BucketLifecycleRuleArgs{
	return s3.BucketLifecycleRuleArgs{
		Id:                                 s.Id.S(),
		Enabled:                            s.Enabled.B(),
		AbortIncompleteMultipartUploadDays: s.AbortIncompleteMultipartUploadDays.I(),
		//Expiration:                         nil,
		//NoncurrentVersionExpiration:        nil,
		//NoncurrentVersionTransitions:       nil,
	}
}