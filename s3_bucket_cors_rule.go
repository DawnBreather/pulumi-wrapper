package main

import (
	"fmt"
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type HttpScheme string
type HttpMethod string

const (
	HTTP 	HttpScheme 	= "http"
	HTTPS 	HttpScheme 	= "https"

	GET     HttpMethod = "GET"
	POST    HttpMethod = "POST"
	OPTIONS HttpMethod = "OPTIONS"
	PUT     HttpMethod = "PUT"
	DELETE  HttpMethod = "DELETE"
	TRACE   HttpMethod = "TRACE"
	CONNECT HttpMethod = "CONNECT"
	HEAD    HttpMethod = "HEAD"
	PATCH   HttpMethod = "PATCH"
)

func (s *s3Bucket) S3BucketCorsRule() *s3BucketCorsRule{

	if s.corsRules == nil {
		s.corsRules = s3BucketCorsRules{}
	}

	corsRule := s3BucketCorsRule{}
	s.corsRules.Append(corsRule)
	return &corsRule
}

type s3BucketCorsRules []s3BucketCorsRule
func (s *s3BucketCorsRules) Append(corsRule s3BucketCorsRule) *s3BucketCorsRules{
	*s = append(*s, corsRule)
	return s
}
func (s s3BucketCorsRules) ToPulumiCorsRules() (res s3.BucketCorsRuleArray){
	for _, i := range s{
		res = append(res, i.ToPulumiCorsRuleArgs())
	}

	return res
}

type s3BucketCorsRule struct{
	allowedOrigins []string
	allowedMethods []HttpMethod
	allowedHeaders []string

	s3bucket *s3Bucket
}

func (s *s3BucketCorsRule) Done() *s3Bucket{
	return s.s3bucket
}

func (s *s3BucketCorsRule) AllowedOrigin(scheme HttpScheme, domainName string) *s3BucketCorsRule{

	if s.allowedOrigins == nil {
		s.allowedOrigins = []string{}
	}

	origin := fmt.Sprintf("%s://%s", scheme, domainName)
	s.allowedOrigins = append(s.allowedOrigins, origin)

	return s
}

func (s *s3BucketCorsRule) AllowedMethods(methods ...HttpMethod) *s3BucketCorsRule{

	if s.allowedMethods == nil {
		s.allowedMethods = []HttpMethod{}
	}

	s.allowedMethods = append(s.allowedMethods, methods...)

	return s
}

func (s *s3BucketCorsRule) AllowedHeaders(headers ...string) *s3BucketCorsRule{

	if s.allowedHeaders == nil {
		s.allowedHeaders = []string{}
	}

	s.allowedHeaders = append(s.allowedHeaders, headers...)

	return s
}

func (s *s3BucketCorsRule) ToPulumiCorsRuleArgs() s3.BucketCorsRuleArgs{

	pAllowedHeaders := pulumi.StringArray{}
	pAllowedMethods := pulumi.StringArray{}
	pAllowedOrigins := pulumi.StringArray{}

	for _, h := range s.allowedHeaders{
		pAllowedHeaders = append(pAllowedHeaders, pulumi.String(h))
	}
	for _, m := range s.allowedMethods{
		pAllowedMethods = append(pAllowedMethods, pulumi.String(m))
	}
	for _, o := range s.allowedOrigins{
		pAllowedOrigins = append(pAllowedOrigins, pulumi.String(o))
	}

	return s3.BucketCorsRuleArgs{
		AllowedHeaders: pAllowedHeaders,
		AllowedMethods: pAllowedMethods,
		AllowedOrigins: pAllowedOrigins,
		//ExposeHeaders:  nil,
		//MaxAgeSeconds:  nil,
	}
}










