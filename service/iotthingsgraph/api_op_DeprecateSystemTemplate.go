// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotthingsgraph

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeprecateSystemTemplateInput struct {
	_ struct{} `type:"structure"`

	// The ID of the system to delete.
	//
	// The ID should be in the following format.
	//
	// urn:tdm:REGION/ACCOUNT ID/default:system:SYSTEMNAME
	//
	// Id is a required field
	Id *string `locationName:"id" type:"string" required:"true"`
}

// String returns the string representation
func (s DeprecateSystemTemplateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeprecateSystemTemplateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeprecateSystemTemplateInput"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeprecateSystemTemplateOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeprecateSystemTemplateOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeprecateSystemTemplate = "DeprecateSystemTemplate"

// DeprecateSystemTemplateRequest returns a request value for making API operation for
// AWS IoT Things Graph.
//
// Deprecates the specified system.
//
//    // Example sending a request using DeprecateSystemTemplateRequest.
//    req := client.DeprecateSystemTemplateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotthingsgraph-2018-09-06/DeprecateSystemTemplate
func (c *Client) DeprecateSystemTemplateRequest(input *DeprecateSystemTemplateInput) DeprecateSystemTemplateRequest {
	op := &aws.Operation{
		Name:       opDeprecateSystemTemplate,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeprecateSystemTemplateInput{}
	}

	req := c.newRequest(op, input, &DeprecateSystemTemplateOutput{})
	return DeprecateSystemTemplateRequest{Request: req, Input: input, Copy: c.DeprecateSystemTemplateRequest}
}

// DeprecateSystemTemplateRequest is the request type for the
// DeprecateSystemTemplate API operation.
type DeprecateSystemTemplateRequest struct {
	*aws.Request
	Input *DeprecateSystemTemplateInput
	Copy  func(*DeprecateSystemTemplateInput) DeprecateSystemTemplateRequest
}

// Send marshals and sends the DeprecateSystemTemplate API request.
func (r DeprecateSystemTemplateRequest) Send(ctx context.Context) (*DeprecateSystemTemplateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeprecateSystemTemplateResponse{
		DeprecateSystemTemplateOutput: r.Request.Data.(*DeprecateSystemTemplateOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeprecateSystemTemplateResponse is the response type for the
// DeprecateSystemTemplate API operation.
type DeprecateSystemTemplateResponse struct {
	*DeprecateSystemTemplateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeprecateSystemTemplate request.
func (r *DeprecateSystemTemplateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}