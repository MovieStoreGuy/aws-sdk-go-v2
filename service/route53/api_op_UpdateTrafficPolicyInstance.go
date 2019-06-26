// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// A complex type that contains information about the resource record sets that
// you want to update based on a specified traffic policy instance.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53-2013-04-01/UpdateTrafficPolicyInstanceRequest
type UpdateTrafficPolicyInstanceInput struct {
	_ struct{} `locationName:"UpdateTrafficPolicyInstanceRequest" type:"structure" xmlURI:"https://route53.amazonaws.com/doc/2013-04-01/"`

	// The ID of the traffic policy instance that you want to update.
	//
	// Id is a required field
	Id *string `location:"uri" locationName:"Id" min:"1" type:"string" required:"true"`

	// The TTL that you want Amazon Route 53 to assign to all of the updated resource
	// record sets.
	//
	// TTL is a required field
	TTL *int64 `type:"long" required:"true"`

	// The ID of the traffic policy that you want Amazon Route 53 to use to update
	// resource record sets for the specified traffic policy instance.
	//
	// TrafficPolicyId is a required field
	TrafficPolicyId *string `min:"1" type:"string" required:"true"`

	// The version of the traffic policy that you want Amazon Route 53 to use to
	// update resource record sets for the specified traffic policy instance.
	//
	// TrafficPolicyVersion is a required field
	TrafficPolicyVersion *int64 `min:"1" type:"integer" required:"true"`
}

// String returns the string representation
func (s UpdateTrafficPolicyInstanceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateTrafficPolicyInstanceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateTrafficPolicyInstanceInput"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}
	if s.Id != nil && len(*s.Id) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Id", 1))
	}

	if s.TTL == nil {
		invalidParams.Add(aws.NewErrParamRequired("TTL"))
	}

	if s.TrafficPolicyId == nil {
		invalidParams.Add(aws.NewErrParamRequired("TrafficPolicyId"))
	}
	if s.TrafficPolicyId != nil && len(*s.TrafficPolicyId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TrafficPolicyId", 1))
	}

	if s.TrafficPolicyVersion == nil {
		invalidParams.Add(aws.NewErrParamRequired("TrafficPolicyVersion"))
	}
	if s.TrafficPolicyVersion != nil && *s.TrafficPolicyVersion < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("TrafficPolicyVersion", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateTrafficPolicyInstanceInput) MarshalFields(e protocol.FieldEncoder) error {

	e.SetFields(protocol.BodyTarget, "UpdateTrafficPolicyInstanceRequest", protocol.FieldMarshalerFunc(func(e protocol.FieldEncoder) error {
		if s.TTL != nil {
			v := *s.TTL

			metadata := protocol.Metadata{}
			e.SetValue(protocol.BodyTarget, "TTL", protocol.Int64Value(v), metadata)
		}
		if s.TrafficPolicyId != nil {
			v := *s.TrafficPolicyId

			metadata := protocol.Metadata{}
			e.SetValue(protocol.BodyTarget, "TrafficPolicyId", protocol.StringValue(v), metadata)
		}
		if s.TrafficPolicyVersion != nil {
			v := *s.TrafficPolicyVersion

			metadata := protocol.Metadata{}
			e.SetValue(protocol.BodyTarget, "TrafficPolicyVersion", protocol.Int64Value(v), metadata)
		}
		return nil
	}), protocol.Metadata{XMLNamespaceURI: "https://route53.amazonaws.com/doc/2013-04-01/"})
	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Id", protocol.StringValue(v), metadata)
	}
	return nil
}

// A complex type that contains information about the resource record sets that
// Amazon Route 53 created based on a specified traffic policy.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53-2013-04-01/UpdateTrafficPolicyInstanceResponse
type UpdateTrafficPolicyInstanceOutput struct {
	_ struct{} `type:"structure"`

	// A complex type that contains settings for the updated traffic policy instance.
	//
	// TrafficPolicyInstance is a required field
	TrafficPolicyInstance *TrafficPolicyInstance `type:"structure" required:"true"`
}

// String returns the string representation
func (s UpdateTrafficPolicyInstanceOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateTrafficPolicyInstanceOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.TrafficPolicyInstance != nil {
		v := s.TrafficPolicyInstance

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "TrafficPolicyInstance", v, metadata)
	}
	return nil
}

const opUpdateTrafficPolicyInstance = "UpdateTrafficPolicyInstance"

// UpdateTrafficPolicyInstanceRequest returns a request value for making API operation for
// Amazon Route 53.
//
// Updates the resource record sets in a specified hosted zone that were created
// based on the settings in a specified traffic policy version.
//
// When you update a traffic policy instance, Amazon Route 53 continues to respond
// to DNS queries for the root resource record set name (such as example.com)
// while it replaces one group of resource record sets with another. Route 53
// performs the following operations:
//
// Route 53 creates a new group of resource record sets based on the specified
// traffic policy. This is true regardless of how significant the differences
// are between the existing resource record sets and the new resource record
// sets.
//
// When all of the new resource record sets have been created, Route 53 starts
// to respond to DNS queries for the root resource record set name (such as
// example.com) by using the new resource record sets.
//
// Route 53 deletes the old group of resource record sets that are associated
// with the root resource record set name.
//
//    // Example sending a request using UpdateTrafficPolicyInstanceRequest.
//    req := client.UpdateTrafficPolicyInstanceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53-2013-04-01/UpdateTrafficPolicyInstance
func (c *Client) UpdateTrafficPolicyInstanceRequest(input *UpdateTrafficPolicyInstanceInput) UpdateTrafficPolicyInstanceRequest {
	op := &aws.Operation{
		Name:       opUpdateTrafficPolicyInstance,
		HTTPMethod: "POST",
		HTTPPath:   "/2013-04-01/trafficpolicyinstance/{Id}",
	}

	if input == nil {
		input = &UpdateTrafficPolicyInstanceInput{}
	}

	req := c.newRequest(op, input, &UpdateTrafficPolicyInstanceOutput{})
	return UpdateTrafficPolicyInstanceRequest{Request: req, Input: input, Copy: c.UpdateTrafficPolicyInstanceRequest}
}

// UpdateTrafficPolicyInstanceRequest is the request type for the
// UpdateTrafficPolicyInstance API operation.
type UpdateTrafficPolicyInstanceRequest struct {
	*aws.Request
	Input *UpdateTrafficPolicyInstanceInput
	Copy  func(*UpdateTrafficPolicyInstanceInput) UpdateTrafficPolicyInstanceRequest
}

// Send marshals and sends the UpdateTrafficPolicyInstance API request.
func (r UpdateTrafficPolicyInstanceRequest) Send(ctx context.Context) (*UpdateTrafficPolicyInstanceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateTrafficPolicyInstanceResponse{
		UpdateTrafficPolicyInstanceOutput: r.Request.Data.(*UpdateTrafficPolicyInstanceOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateTrafficPolicyInstanceResponse is the response type for the
// UpdateTrafficPolicyInstance API operation.
type UpdateTrafficPolicyInstanceResponse struct {
	*UpdateTrafficPolicyInstanceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateTrafficPolicyInstance request.
func (r *UpdateTrafficPolicyInstanceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}