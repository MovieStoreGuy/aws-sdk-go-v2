// Code generated by smithy-go-codegen DO NOT EDIT.

package emr

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/emr/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Provides summaries of all notebook executions. You can filter the list based on
// multiple criteria such as status, time range, and editor id. Returns a maximum
// of 50 notebook executions and a marker to track the paging of a longer notebook
// execution list across multiple ListNotebookExecution calls.
func (c *Client) ListNotebookExecutions(ctx context.Context, params *ListNotebookExecutionsInput, optFns ...func(*Options)) (*ListNotebookExecutionsOutput, error) {
	if params == nil {
		params = &ListNotebookExecutionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListNotebookExecutions", params, optFns, addOperationListNotebookExecutionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListNotebookExecutionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListNotebookExecutionsInput struct {

	// The unique ID of the editor associated with the notebook execution.
	EditorId *string

	// The beginning of time range filter for listing notebook executions. The default
	// is the timestamp of 30 days ago.
	From *time.Time

	// The pagination token, returned by a previous ListNotebookExecutions call, that
	// indicates the start of the list for this ListNotebookExecutions call.
	Marker *string

	// The status filter for listing notebook executions.
	//
	// * START_PENDING indicates
	// that the cluster has received the execution request but execution has not
	// begun.
	//
	// * STARTING indicates that the execution is starting on the cluster.
	//
	// *
	// RUNNING indicates that the execution is being processed by the cluster.
	//
	// *
	// FINISHING indicates that execution processing is in the final stages.
	//
	// *
	// FINISHED indicates that the execution has completed without error.
	//
	// * FAILING
	// indicates that the execution is failing and will not finish successfully.
	//
	// *
	// FAILED indicates that the execution failed.
	//
	// * STOP_PENDING indicates that the
	// cluster has received a StopNotebookExecution request and the stop is pending.
	//
	// *
	// STOPPING indicates that the cluster is in the process of stopping the execution
	// as a result of a StopNotebookExecution request.
	//
	// * STOPPED indicates that the
	// execution stopped because of a StopNotebookExecution request.
	Status types.NotebookExecutionStatus

	// The end of time range filter for listing notebook executions. The default is the
	// current timestamp.
	To *time.Time
}

type ListNotebookExecutionsOutput struct {

	// A pagination token that a subsequent ListNotebookExecutions can use to determine
	// the next set of results to retrieve.
	Marker *string

	// A list of notebook executions.
	NotebookExecutions []types.NotebookExecutionSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListNotebookExecutionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListNotebookExecutions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListNotebookExecutions{}, middleware.After)
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
	if err = awsmiddleware.AddAttemptClockSkewMiddleware(stack); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListNotebookExecutions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListNotebookExecutions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticmapreduce",
		OperationName: "ListNotebookExecutions",
	}
}