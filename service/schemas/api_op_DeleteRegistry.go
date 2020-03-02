// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package schemas

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

type DeleteRegistryInput struct {
	_ struct{} `type:"structure"`

	// RegistryName is a required field
	RegistryName *string `location:"uri" locationName:"registryName" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteRegistryInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteRegistryInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteRegistryInput"}

	if s.RegistryName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RegistryName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteRegistryInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.RegistryName != nil {
		v := *s.RegistryName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "registryName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteRegistryOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteRegistryOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteRegistryOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteRegistry = "DeleteRegistry"

// DeleteRegistryRequest returns a request value for making API operation for
// Schemas.
//
// Deletes a Registry.
//
//    // Example sending a request using DeleteRegistryRequest.
//    req := client.DeleteRegistryRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/schemas-2019-12-02/DeleteRegistry
func (c *Client) DeleteRegistryRequest(input *DeleteRegistryInput) DeleteRegistryRequest {
	op := &aws.Operation{
		Name:       opDeleteRegistry,
		HTTPMethod: "DELETE",
		HTTPPath:   "/v1/registries/name/{registryName}",
	}

	if input == nil {
		input = &DeleteRegistryInput{}
	}

	req := c.newRequest(op, input, &DeleteRegistryOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteRegistryRequest{Request: req, Input: input, Copy: c.DeleteRegistryRequest}
}

// DeleteRegistryRequest is the request type for the
// DeleteRegistry API operation.
type DeleteRegistryRequest struct {
	*aws.Request
	Input *DeleteRegistryInput
	Copy  func(*DeleteRegistryInput) DeleteRegistryRequest
}

// Send marshals and sends the DeleteRegistry API request.
func (r DeleteRegistryRequest) Send(ctx context.Context) (*DeleteRegistryResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteRegistryResponse{
		DeleteRegistryOutput: r.Request.Data.(*DeleteRegistryOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteRegistryResponse is the response type for the
// DeleteRegistry API operation.
type DeleteRegistryResponse struct {
	*DeleteRegistryOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteRegistry request.
func (r *DeleteRegistryResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}