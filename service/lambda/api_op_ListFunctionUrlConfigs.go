// Code generated by smithy-go-codegen DO NOT EDIT.

package lambda

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

func (c *Client) ListFunctionUrlConfigs(ctx context.Context, params *ListFunctionUrlConfigsInput, optFns ...func(*Options)) (*ListFunctionUrlConfigsOutput, error) {
	if params == nil {
		params = &ListFunctionUrlConfigsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListFunctionUrlConfigs", params, optFns, c.addOperationListFunctionUrlConfigsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListFunctionUrlConfigsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListFunctionUrlConfigsInput struct {

	// This member is required.
	FunctionName *string

	Marker *string

	MaxItems *int32

	noSmithyDocumentSerde
}

type ListFunctionUrlConfigsOutput struct {

	// This member is required.
	FunctionUrlConfigs []types.FunctionUrlConfig

	NextMarker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListFunctionUrlConfigsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListFunctionUrlConfigs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListFunctionUrlConfigs{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpListFunctionUrlConfigsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListFunctionUrlConfigs(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

// ListFunctionUrlConfigsAPIClient is a client that implements the
// ListFunctionUrlConfigs operation.
type ListFunctionUrlConfigsAPIClient interface {
	ListFunctionUrlConfigs(context.Context, *ListFunctionUrlConfigsInput, ...func(*Options)) (*ListFunctionUrlConfigsOutput, error)
}

var _ ListFunctionUrlConfigsAPIClient = (*Client)(nil)

// ListFunctionUrlConfigsPaginatorOptions is the paginator options for
// ListFunctionUrlConfigs
type ListFunctionUrlConfigsPaginatorOptions struct {
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListFunctionUrlConfigsPaginator is a paginator for ListFunctionUrlConfigs
type ListFunctionUrlConfigsPaginator struct {
	options   ListFunctionUrlConfigsPaginatorOptions
	client    ListFunctionUrlConfigsAPIClient
	params    *ListFunctionUrlConfigsInput
	nextToken *string
	firstPage bool
}

// NewListFunctionUrlConfigsPaginator returns a new ListFunctionUrlConfigsPaginator
func NewListFunctionUrlConfigsPaginator(client ListFunctionUrlConfigsAPIClient, params *ListFunctionUrlConfigsInput, optFns ...func(*ListFunctionUrlConfigsPaginatorOptions)) *ListFunctionUrlConfigsPaginator {
	if params == nil {
		params = &ListFunctionUrlConfigsInput{}
	}

	options := ListFunctionUrlConfigsPaginatorOptions{}
	if params.MaxItems != nil {
		options.Limit = *params.MaxItems
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListFunctionUrlConfigsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListFunctionUrlConfigsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListFunctionUrlConfigs page.
func (p *ListFunctionUrlConfigsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListFunctionUrlConfigsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxItems = limit

	result, err := p.client.ListFunctionUrlConfigs(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextMarker

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListFunctionUrlConfigs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lambda",
		OperationName: "ListFunctionUrlConfigs",
	}
}