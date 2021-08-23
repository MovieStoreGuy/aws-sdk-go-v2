// Code generated by smithy-go-codegen DO NOT EDIT.

package endpoints

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/endpoints"
	"regexp"
)

// Options is the endpoint resolver configuration options
type Options struct {
	DisableHTTPS bool
}

// Resolver S3 Control endpoint resolver
type Resolver struct {
	partitions endpoints.Partitions
}

// ResolveEndpoint resolves the service endpoint for the given region and options
func (r *Resolver) ResolveEndpoint(region string, options Options) (endpoint aws.Endpoint, err error) {
	if len(region) == 0 {
		return endpoint, &aws.MissingRegionError{}
	}

	opt := endpoints.Options{
		DisableHTTPS: options.DisableHTTPS,
	}
	return r.partitions.ResolveEndpoint(region, opt)
}

// New returns a new Resolver
func New() *Resolver {
	return &Resolver{
		partitions: defaultPartitions,
	}
}

var partitionRegexp = struct {
	Aws      *regexp.Regexp
	AwsCn    *regexp.Regexp
	AwsIso   *regexp.Regexp
	AwsIsoB  *regexp.Regexp
	AwsUsGov *regexp.Regexp
}{

	Aws:      regexp.MustCompile("^(us|eu|ap|sa|ca|me|af)\\-\\w+\\-\\d+$"),
	AwsCn:    regexp.MustCompile("^cn\\-\\w+\\-\\d+$"),
	AwsIso:   regexp.MustCompile("^us\\-iso\\-\\w+\\-\\d+$"),
	AwsIsoB:  regexp.MustCompile("^us\\-isob\\-\\w+\\-\\d+$"),
	AwsUsGov: regexp.MustCompile("^us\\-gov\\-\\w+\\-\\d+$"),
}

var defaultPartitions = endpoints.Partitions{
	{
		ID: "aws",
		Defaults: endpoints.Endpoint{
			Hostname:          "s3-control.{region}.amazonaws.com",
			Protocols:         []string{"https"},
			SignatureVersions: []string{"s3v4"},
		},
		RegionRegex:    partitionRegexp.Aws,
		IsRegionalized: true,
		Endpoints: endpoints.Endpoints{
			"ap-northeast-1": endpoints.Endpoint{
				Hostname:          "s3-control.ap-northeast-1.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
				CredentialScope: endpoints.CredentialScope{
					Region: "ap-northeast-1",
				},
			},
			"ap-northeast-2": endpoints.Endpoint{
				Hostname:          "s3-control.ap-northeast-2.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
				CredentialScope: endpoints.CredentialScope{
					Region: "ap-northeast-2",
				},
			},
			"ap-northeast-3": endpoints.Endpoint{
				Hostname:          "s3-control.ap-northeast-3.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
				CredentialScope: endpoints.CredentialScope{
					Region: "ap-northeast-3",
				},
			},
			"ap-south-1": endpoints.Endpoint{
				Hostname:          "s3-control.ap-south-1.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
				CredentialScope: endpoints.CredentialScope{
					Region: "ap-south-1",
				},
			},
			"ap-southeast-1": endpoints.Endpoint{
				Hostname:          "s3-control.ap-southeast-1.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
				CredentialScope: endpoints.CredentialScope{
					Region: "ap-southeast-1",
				},
			},
			"ap-southeast-2": endpoints.Endpoint{
				Hostname:          "s3-control.ap-southeast-2.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
				CredentialScope: endpoints.CredentialScope{
					Region: "ap-southeast-2",
				},
			},
			"ca-central-1": endpoints.Endpoint{
				Hostname:          "s3-control.ca-central-1.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
				CredentialScope: endpoints.CredentialScope{
					Region: "ca-central-1",
				},
			},
			"ca-central-1-fips": endpoints.Endpoint{
				Hostname:          "s3-control-fips.ca-central-1.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
				CredentialScope: endpoints.CredentialScope{
					Region: "ca-central-1",
				},
			},
			"eu-central-1": endpoints.Endpoint{
				Hostname:          "s3-control.eu-central-1.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
				CredentialScope: endpoints.CredentialScope{
					Region: "eu-central-1",
				},
			},
			"eu-north-1": endpoints.Endpoint{
				Hostname:          "s3-control.eu-north-1.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
				CredentialScope: endpoints.CredentialScope{
					Region: "eu-north-1",
				},
			},
			"eu-west-1": endpoints.Endpoint{
				Hostname:          "s3-control.eu-west-1.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
				CredentialScope: endpoints.CredentialScope{
					Region: "eu-west-1",
				},
			},
			"eu-west-2": endpoints.Endpoint{
				Hostname:          "s3-control.eu-west-2.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
				CredentialScope: endpoints.CredentialScope{
					Region: "eu-west-2",
				},
			},
			"eu-west-3": endpoints.Endpoint{
				Hostname:          "s3-control.eu-west-3.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
				CredentialScope: endpoints.CredentialScope{
					Region: "eu-west-3",
				},
			},
			"sa-east-1": endpoints.Endpoint{
				Hostname:          "s3-control.sa-east-1.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
				CredentialScope: endpoints.CredentialScope{
					Region: "sa-east-1",
				},
			},
			"us-east-1": endpoints.Endpoint{
				Hostname:          "s3-control.us-east-1.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
				CredentialScope: endpoints.CredentialScope{
					Region: "us-east-1",
				},
			},
			"us-east-1-fips": endpoints.Endpoint{
				Hostname:          "s3-control-fips.us-east-1.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
				CredentialScope: endpoints.CredentialScope{
					Region: "us-east-1",
				},
			},
			"us-east-2": endpoints.Endpoint{
				Hostname:          "s3-control.us-east-2.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
				CredentialScope: endpoints.CredentialScope{
					Region: "us-east-2",
				},
			},
			"us-east-2-fips": endpoints.Endpoint{
				Hostname:          "s3-control-fips.us-east-2.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
				CredentialScope: endpoints.CredentialScope{
					Region: "us-east-2",
				},
			},
			"us-west-1": endpoints.Endpoint{
				Hostname:          "s3-control.us-west-1.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
				CredentialScope: endpoints.CredentialScope{
					Region: "us-west-1",
				},
			},
			"us-west-1-fips": endpoints.Endpoint{
				Hostname:          "s3-control-fips.us-west-1.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
				CredentialScope: endpoints.CredentialScope{
					Region: "us-west-1",
				},
			},
			"us-west-2": endpoints.Endpoint{
				Hostname:          "s3-control.us-west-2.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
				CredentialScope: endpoints.CredentialScope{
					Region: "us-west-2",
				},
			},
			"us-west-2-fips": endpoints.Endpoint{
				Hostname:          "s3-control-fips.us-west-2.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
				CredentialScope: endpoints.CredentialScope{
					Region: "us-west-2",
				},
			},
		},
	},
	{
		ID: "aws-cn",
		Defaults: endpoints.Endpoint{
			Hostname:          "s3-control.{region}.amazonaws.com.cn",
			Protocols:         []string{"https"},
			SignatureVersions: []string{"s3v4"},
		},
		RegionRegex:    partitionRegexp.AwsCn,
		IsRegionalized: true,
		Endpoints: endpoints.Endpoints{
			"cn-north-1": endpoints.Endpoint{
				Hostname:          "s3-control.cn-north-1.amazonaws.com.cn",
				SignatureVersions: []string{"s3v4"},
				CredentialScope: endpoints.CredentialScope{
					Region: "cn-north-1",
				},
			},
			"cn-northwest-1": endpoints.Endpoint{
				Hostname:          "s3-control.cn-northwest-1.amazonaws.com.cn",
				SignatureVersions: []string{"s3v4"},
				CredentialScope: endpoints.CredentialScope{
					Region: "cn-northwest-1",
				},
			},
		},
	},
	{
		ID: "aws-iso",
		Defaults: endpoints.Endpoint{
			Hostname:          "s3-control.{region}.c2s.ic.gov",
			Protocols:         []string{"https"},
			SignatureVersions: []string{"v4"},
		},
		RegionRegex:    partitionRegexp.AwsIso,
		IsRegionalized: true,
	},
	{
		ID: "aws-iso-b",
		Defaults: endpoints.Endpoint{
			Hostname:          "s3-control.{region}.sc2s.sgov.gov",
			Protocols:         []string{"https"},
			SignatureVersions: []string{"v4"},
		},
		RegionRegex:    partitionRegexp.AwsIsoB,
		IsRegionalized: true,
	},
	{
		ID: "aws-us-gov",
		Defaults: endpoints.Endpoint{
			Hostname:          "s3-control.{region}.amazonaws.com",
			Protocols:         []string{"https"},
			SignatureVersions: []string{"s3v4"},
		},
		RegionRegex:    partitionRegexp.AwsUsGov,
		IsRegionalized: true,
		Endpoints: endpoints.Endpoints{
			"us-gov-east-1": endpoints.Endpoint{
				Hostname:          "s3-control.us-gov-east-1.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
				CredentialScope: endpoints.CredentialScope{
					Region: "us-gov-east-1",
				},
			},
			"us-gov-east-1-fips": endpoints.Endpoint{
				Hostname:          "s3-control-fips.us-gov-east-1.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
				CredentialScope: endpoints.CredentialScope{
					Region: "us-gov-east-1",
				},
			},
			"us-gov-west-1": endpoints.Endpoint{
				Hostname:          "s3-control.us-gov-west-1.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
				CredentialScope: endpoints.CredentialScope{
					Region: "us-gov-west-1",
				},
			},
			"us-gov-west-1-fips": endpoints.Endpoint{
				Hostname:          "s3-control-fips.us-gov-west-1.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
				CredentialScope: endpoints.CredentialScope{
					Region: "us-gov-west-1",
				},
			},
		},
	},
}
