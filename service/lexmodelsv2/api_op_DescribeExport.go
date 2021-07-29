// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelsv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelsv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets information about a specific export.
func (c *Client) DescribeExport(ctx context.Context, params *DescribeExportInput, optFns ...func(*Options)) (*DescribeExportOutput, error) {
	if params == nil {
		params = &DescribeExportInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeExport", params, optFns, c.addOperationDescribeExportMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeExportOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeExportInput struct {

	// The unique identifier of the export to describe.
	//
	// This member is required.
	ExportId *string

	noSmithyDocumentSerde
}

type DescribeExportOutput struct {

	// The date and time that the export was created.
	CreationDateTime *time.Time

	// A pre-signed S3 URL that points to the bot or bot locale archive. The URL is
	// only available for 5 minutes after calling the DescribeExport operation.
	DownloadUrl *string

	// The unique identifier of the described export.
	ExportId *string

	// The status of the export. When the status is Complete the export archive file is
	// available for download.
	ExportStatus types.ExportStatus

	// If the exportStatus is failed, contains one or more reasons why the export could
	// not be completed.
	FailureReasons []string

	// The file format used in the files that describe the bot or bot locale.
	FileFormat types.ImportExportFileFormat

	// The last date and time that the export was updated.
	LastUpdatedDateTime *time.Time

	// The bot, bot ID, and optional locale ID of the exported bot or bot locale.
	ResourceSpecification *types.ExportResourceSpecification

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeExportMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeExport{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeExport{}, middleware.After)
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
	if err = addOpDescribeExportValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeExport(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeExport(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "DescribeExport",
	}
}