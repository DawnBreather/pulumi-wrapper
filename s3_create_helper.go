package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type s3SetupCreateHelper struct{

	ctx *pulumi.Context

	s3Buckets map[string]*s3.Bucket

	conf *S3Conf
	s3Setup *s3Setup
}

func (s *s3SetupCreateHelper) DONE() *s3Setup{
	return s.s3Setup
}

func (s *s3SetupCreateHelper) S3Media(name string, allowedOriginDomainsForCors []string) *s3SetupCreateHelper{

	s3.NewBucketPolicy(s.ctx, s.conf.GetBucketName(name).R(), &s3.BucketPolicyArgs{
		Bucket: String("").S(),
		Policy: String("").S(),
	})

	s.s3Buckets[name], _ = s3.NewBucket(s.ctx, s.conf.GetBucketName(name).R(), &s3.BucketArgs{
		AccelerationStatus:                nil,
		Acl:                               nil,
		Arn:                               nil,
		Bucket:                            nil,
		BucketPrefix:                      nil,
		CorsRules:                         s3.BucketCorsRuleArray{
			s3.BucketCorsRuleArgs{},
		},
		ForceDestroy:                      nil,
		Grants:                            nil,
		HostedZoneId:                      nil,
		LifecycleRules:                    s3.BucketLifecycleRuleArray{
			s3.BucketLifecycleRuleArgs{
				AbortIncompleteMultipartUploadDays: nil,
				Enabled:                            nil,
				Expiration:                         nil,
				Id:                                 nil,
				NoncurrentVersionExpiration:        nil,
				NoncurrentVersionTransitions:       nil,
				Prefix:                             nil,
				Tags:                               nil,
				Transitions:                        nil,
			},
		},
		Loggings:                          nil,
		ObjectLockConfiguration:           nil,
		Policy:                            pulumi.String(""),
		ReplicationConfiguration:          nil,
		RequestPayer:                      nil,
		ServerSideEncryptionConfiguration: nil,
		Tags:                              nil,
		TagsAll:                           nil,
		Versioning:                        nil,
		Website:                           nil,
		WebsiteDomain:                     nil,
		WebsiteEndpoint:                   nil,
	})

	return s
}