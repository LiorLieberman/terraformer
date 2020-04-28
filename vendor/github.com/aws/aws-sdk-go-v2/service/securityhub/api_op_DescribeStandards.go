// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package securityhub

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DescribeStandardsInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of standards to return.
	MaxResults *int64 `location:"querystring" locationName:"MaxResults" min:"1" type:"integer"`

	// The token that is required for pagination. On your first call to the DescribeStandards
	// operation, set the value of this parameter to NULL.
	//
	// For subsequent calls to the operation, to continue listing data, set the
	// value of this parameter to the value returned from the previous response.
	NextToken *string `location:"querystring" locationName:"NextToken" type:"string"`
}

// String returns the string representation
func (s DescribeStandardsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeStandardsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeStandardsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeStandardsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "MaxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DescribeStandardsOutput struct {
	_ struct{} `type:"structure"`

	// The pagination token to use to request the next page of results.
	NextToken *string `type:"string"`

	// A list of available standards.
	Standards []Standard `type:"list"`
}

// String returns the string representation
func (s DescribeStandardsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeStandardsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Standards != nil {
		v := s.Standards

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Standards", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opDescribeStandards = "DescribeStandards"

// DescribeStandardsRequest returns a request value for making API operation for
// AWS SecurityHub.
//
// Returns a list of the available standards in Security Hub.
//
// For each standard, the results include the standard ARN, the name, and a
// description.
//
//    // Example sending a request using DescribeStandardsRequest.
//    req := client.DescribeStandardsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/securityhub-2018-10-26/DescribeStandards
func (c *Client) DescribeStandardsRequest(input *DescribeStandardsInput) DescribeStandardsRequest {
	op := &aws.Operation{
		Name:       opDescribeStandards,
		HTTPMethod: "GET",
		HTTPPath:   "/standards",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeStandardsInput{}
	}

	req := c.newRequest(op, input, &DescribeStandardsOutput{})
	return DescribeStandardsRequest{Request: req, Input: input, Copy: c.DescribeStandardsRequest}
}

// DescribeStandardsRequest is the request type for the
// DescribeStandards API operation.
type DescribeStandardsRequest struct {
	*aws.Request
	Input *DescribeStandardsInput
	Copy  func(*DescribeStandardsInput) DescribeStandardsRequest
}

// Send marshals and sends the DescribeStandards API request.
func (r DescribeStandardsRequest) Send(ctx context.Context) (*DescribeStandardsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeStandardsResponse{
		DescribeStandardsOutput: r.Request.Data.(*DescribeStandardsOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeStandardsRequestPaginator returns a paginator for DescribeStandards.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeStandardsRequest(input)
//   p := securityhub.NewDescribeStandardsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeStandardsPaginator(req DescribeStandardsRequest) DescribeStandardsPaginator {
	return DescribeStandardsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeStandardsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// DescribeStandardsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeStandardsPaginator struct {
	aws.Pager
}

func (p *DescribeStandardsPaginator) CurrentPage() *DescribeStandardsOutput {
	return p.Pager.CurrentPage().(*DescribeStandardsOutput)
}

// DescribeStandardsResponse is the response type for the
// DescribeStandards API operation.
type DescribeStandardsResponse struct {
	*DescribeStandardsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeStandards request.
func (r *DescribeStandardsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
