package oss

import (
	"context"

	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss/signer"
)

type MetaQueryAggregation struct {
	// The field name.
	Field *string `xml:"Field"`

	// The operator for aggregate operations.*   min*   max*   average*   sum*   count*   distinct*   group
	Operation *string `xml:"Operation"`

	// The result of the aggregate operation.
	Value *float64 `xml:"Value"`
}

type MetaQueryAggregations struct {
	// The container that stores the information about a single aggregate operation.
	Aggregations []MetaQueryAggregation `xml:"Aggregation"`
}

type MetaQueryUserMeta struct {
	// The key of the user metadata item.
	Key *string `xml:"Key"`

	// The value of the user metadata item.
	Value *string `xml:"Value"`
}

type MetaQueryFile struct {
	// The time when the object was last modified.
	FileModifiedTime *string `xml:"FileModifiedTime"`

	// The type of the object.Valid values:*   Multipart        :        The object is uploaded by using multipart upload        .*   Symlink        :        The object is a symbolic link that was created by calling the PutSymlink operation.    *   Appendable        :        The object is uploaded by using AppendObject        .*   Normal        :        The object is uploaded by using PutObject.
	OSSObjectType *string `xml:"OSSObjectType"`

	// The ETag of the object.
	ETag *string `xml:"ETag"`

	// The server-side encryption algorithm used when the object was created.
	ServerSideEncryptionCustomerAlgorithm *string `xml:"ServerSideEncryptionCustomerAlgorithm"`

	// The number of the tags of the object.
	OSSTaggingCount *int64 `xml:"OSSTaggingCount"`

	// The tags.
	OSSTagging []MetaQueryTagging `xml:"OSSTagging>Tagging"`

	// The user metadata items.
	OSSUserMeta []MetaQueryUserMeta `xml:"OSSUserMeta>UserMeta"`

	// The full path of the object.
	Filename *string `xml:"Filename"`

	// The storage class of the object.Valid values:*   Archive        :        the Archive storage class        .*   ColdArchive        :        the Cold Archive storage class        .*   IA        :        the Infrequent Access (IA) storage class        .*   Standard        :        The Standard storage class        .
	OSSStorageClass *string `xml:"OSSStorageClass"`

	// The access control list (ACL) of the object.Valid values:*   default        :        the ACL of the bucket        .*   private        :        private        .*   public-read        :        public-read        .*   public-read-write        :        public-read-write        .
	ObjectACL *string `xml:"ObjectACL"`

	// The CRC-64 value of the object.
	OSSCRC64 *string `xml:"OSSCRC64"`

	// The server-side encryption of the object.
	ServerSideEncryption *string `xml:"ServerSideEncryption"`

	// The object size.
	Size *int64 `xml:"Size"`
}

type MetaQuery struct {
	// The maximum number of objects to return. Valid values: 0 to 100. If this parameter is not set or is set to 0, up to 100 objects are returned.
	MaxResults *int64 `xml:"MaxResults"`

	// The query conditions. A query condition includes the following elements:*   Operation: the operator. Valid values: eq (equal to), gt (greater than), gte (greater than or equal to), lt (less than), lte (less than or equal to), match (fuzzy query), prefix (prefix query), and (AND), or (OR), and not (NOT).*   Field: the field name.*   Value: the field value.*   SubQueries: the subquery conditions. Options that are included in this element are the same as those of simple query. You need to set subquery conditions only when Operation is set to and, or, or not.
	Query *string `xml:"Query"`

	// The field based on which the results are sorted.
	Sort *string `xml:"Sort"`

	// The sort order.
	Order MetaQueryOrderType `xml:"Order"`

	// The container that stores the information about aggregate operations.
	Aggregations *MetaQueryAggregations `xml:"Aggregations"`

	// The pagination token used to obtain information in the next request. The object information is returned in alphabetical order starting from the value of NextToken.
	NextToken *string `xml:"NextToken"`
}

type MetaQueryStatus struct {
	// The time when the metadata index library was created. The value follows the RFC 3339 standard in the YYYY-MM-DDTHH:mm:ss+TIMEZONE format. YYYY-MM-DD indicates the year, month, and day. T indicates the beginning of the time element. HH:mm:ss indicates the hour, minute, and second. TIMEZONE indicates the time zone.
	CreateTime *string `xml:"CreateTime"`

	// The time when the metadata index library was updated. The value follows the RFC 3339 standard in the YYYY-MM-DDTHH:mm:ss+TIMEZONE format. YYYY-MM-DD indicates the year, month, and day. T indicates the beginning of the time element. HH:mm:ss indicates the hour, minute, and second. TIMEZONE indicates the time zone.
	UpdateTime *string `xml:"UpdateTime"`

	// The status of the metadata index library. Valid values:- Ready: The metadata index library is being prepared after it is created.In this case, the metadata index library cannot be used to query data.- Stop: The metadata index library is paused.- Running: The metadata index library is running.- Retrying: The metadata index library failed to be created and is being created again.- Failed: The metadata index library failed to be created.- Deleted: The metadata index library is deleted.
	State *string `xml:"State"`

	// The scan type. Valid values:- FullScanning: Full scanning is in progress.- IncrementalScanning: Incremental scanning is in progress.
	Phase *string `xml:"Phase"`
}

type MetaQueryTagging struct {
	// The tag key.
	Key *string `xml:"Key"`

	// The tag value.
	Value *string `xml:"Value"`
}

type GetMetaQueryStatusRequest struct {
	// The name of the bucket.
	Bucket *string `input:"host,bucket,required"`

	RequestCommon
}

type GetMetaQueryStatusResult struct {
	// The container that stores the metadata information.
	MetaQueryStatus *MetaQueryStatus `output:"body,MetaQueryStatus,xml"`

	ResultCommon
}

// GetMetaQueryStatus Queries the information about the metadata index library of a bucket.
func (c *Client) GetMetaQueryStatus(ctx context.Context, request *GetMetaQueryStatusRequest, optFns ...func(*Options)) (*GetMetaQueryStatusResult, error) {
	var err error
	if request == nil {
		request = &GetMetaQueryStatusRequest{}
	}
	input := &OperationInput{
		OpName: "GetMetaQueryStatus",
		Method: "GET",
		Headers: map[string]string{
			HTTPHeaderContentType: contentTypeXML,
		},
		Parameters: map[string]string{
			"metaQuery": "",
		},
		Bucket: request.Bucket,
	}
	input.OpMetadata.Set(signer.SubResource, []string{"metaQuery"})

	if err = c.marshalInput(request, input, updateContentMd5); err != nil {
		return nil, err
	}
	output, err := c.invokeOperation(ctx, input, optFns)
	if err != nil {
		return nil, err
	}

	result := &GetMetaQueryStatusResult{}

	if err = c.unmarshalOutput(result, output, unmarshalBodyXmlMix); err != nil {
		return nil, c.toClientError(err, "UnmarshalOutputFail", output)
	}

	return result, err
}

type CloseMetaQueryRequest struct {
	// The name of the bucket.
	Bucket *string `input:"host,bucket,required"`

	RequestCommon
}

type CloseMetaQueryResult struct {
	ResultCommon
}

// CloseMetaQuery Disables the metadata management feature for an Object Storage Service (OSS) bucket. After the metadata management feature is disabled for a bucket, OSS automatically deletes the metadata index library of the bucket and you cannot perform metadata indexing.
func (c *Client) CloseMetaQuery(ctx context.Context, request *CloseMetaQueryRequest, optFns ...func(*Options)) (*CloseMetaQueryResult, error) {
	var err error
	if request == nil {
		request = &CloseMetaQueryRequest{}
	}
	input := &OperationInput{
		OpName: "CloseMetaQuery",
		Method: "POST",
		Headers: map[string]string{
			HTTPHeaderContentType: contentTypeXML,
		},
		Parameters: map[string]string{
			"comp":      "delete",
			"metaQuery": "",
		},
		Bucket: request.Bucket,
	}
	input.OpMetadata.Set(signer.SubResource, []string{"metaQuery", "comp"})

	if err = c.marshalInput(request, input, updateContentMd5); err != nil {
		return nil, err
	}
	output, err := c.invokeOperation(ctx, input, optFns)
	if err != nil {
		return nil, err
	}

	result := &CloseMetaQueryResult{}

	if err = c.unmarshalOutput(result, output, unmarshalBodyXmlMix); err != nil {
		return nil, c.toClientError(err, "UnmarshalOutputFail", output)
	}

	return result, err
}

type DoMetaQueryRequest struct {
	// The name of the bucket.
	Bucket *string `input:"host,bucket,required"`

	// The request body schema.
	MetaQuery *MetaQuery `input:"body,MetaQuery,xml,required"`

	RequestCommon
}

type DoMetaQueryResult struct {
	// The token that is used for the next query when the total number of objects exceeds the value of MaxResults.The value of NextToken is used to return the unreturned results in the next query.This parameter has a value only when not all objects are returned.
	NextToken *string `xml:"NextToken"`

	// The list of file information.
	Files []MetaQueryFile `xml:"Files>File"`

	// The list of file information.
	Aggregations []MetaQueryAggregation `xml:"Aggregations>Aggregation"`

	ResultCommon
}

// DoMetaQuery Queries the objects in a bucket that meet the specified conditions by using the data indexing feature. The information about the objects is listed based on the specified fields and sorting methods.
func (c *Client) DoMetaQuery(ctx context.Context, request *DoMetaQueryRequest, optFns ...func(*Options)) (*DoMetaQueryResult, error) {
	var err error
	if request == nil {
		request = &DoMetaQueryRequest{}
	}
	input := &OperationInput{
		OpName: "DoMetaQuery",
		Method: "POST",
		Headers: map[string]string{
			HTTPHeaderContentType: contentTypeXML,
		},
		Parameters: map[string]string{
			"comp":      "query",
			"metaQuery": "",
		},
		Bucket: request.Bucket,
	}
	input.OpMetadata.Set(signer.SubResource, []string{"metaQuery", "comp"})

	if err = c.marshalInput(request, input, updateContentMd5); err != nil {
		return nil, err
	}
	output, err := c.invokeOperation(ctx, input, optFns)
	if err != nil {
		return nil, err
	}

	result := &DoMetaQueryResult{}

	if err = c.unmarshalOutput(result, output, unmarshalBodyXmlMix); err != nil {
		return nil, c.toClientError(err, "UnmarshalOutputFail", output)
	}

	return result, err
}

type OpenMetaQueryRequest struct {
	// The name of the bucket.
	Bucket *string `input:"host,bucket,required"`

	RequestCommon
}

type OpenMetaQueryResult struct {
	ResultCommon
}

// OpenMetaQuery Enables metadata management for a bucket. After you enable the metadata management feature for a bucket, Object Storage Service (OSS) creates a metadata index library for the bucket and creates metadata indexes for all objects in the bucket. After the metadata index library is created, OSS continues to perform quasi-real-time scans on incremental objects in the bucket and creates metadata indexes for the incremental objects.
func (c *Client) OpenMetaQuery(ctx context.Context, request *OpenMetaQueryRequest, optFns ...func(*Options)) (*OpenMetaQueryResult, error) {
	var err error
	if request == nil {
		request = &OpenMetaQueryRequest{}
	}
	input := &OperationInput{
		OpName: "OpenMetaQuery",
		Method: "POST",
		Headers: map[string]string{
			HTTPHeaderContentType: contentTypeXML,
		},
		Parameters: map[string]string{
			"comp":      "add",
			"metaQuery": "",
		},
		Bucket: request.Bucket,
	}
	input.OpMetadata.Set(signer.SubResource, []string{"metaQuery", "comp"})

	if err = c.marshalInput(request, input, updateContentMd5); err != nil {
		return nil, err
	}
	output, err := c.invokeOperation(ctx, input, optFns)
	if err != nil {
		return nil, err
	}

	result := &OpenMetaQueryResult{}

	if err = c.unmarshalOutput(result, output, unmarshalBodyXmlMix); err != nil {
		return nil, c.toClientError(err, "UnmarshalOutputFail", output)
	}

	return result, err
}
