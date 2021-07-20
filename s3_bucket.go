package main

type s3Bucket struct {
	name String
	corsRules s3BucketCorsRules
	lifecycleRules s3LifecycleRules
}

