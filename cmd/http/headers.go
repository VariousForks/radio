package http

// Standard S3 HTTP response constants
const (
	LastModified       = "Last-Modified"
	Date               = "Date"
	ETag               = "ETag"
	ContentType        = "Content-Type"
	ContentMD5         = "Content-Md5"
	ContentEncoding    = "Content-Encoding"
	Expires            = "Expires"
	ContentLength      = "Content-Length"
	ContentLanguage    = "Content-Language"
	ContentRange       = "Content-Range"
	Connection         = "Connection"
	AcceptRanges       = "Accept-Ranges"
	AmzBucketRegion    = "X-Amz-Bucket-Region"
	ServerInfo         = "Server"
	RetryAfter         = "Retry-After"
	Location           = "Location"
	CacheControl       = "Cache-Control"
	ContentDisposition = "Content-Disposition"
	Authorization      = "Authorization"
	Action             = "Action"
)

// Standard S3 HTTP request constants
const (
	IfModifiedSince   = "If-Modified-Since"
	IfUnmodifiedSince = "If-Unmodified-Since"
	IfMatch           = "If-Match"
	IfNoneMatch       = "If-None-Match"

	// S3 storage class
	AmzStorageClass = "x-amz-storage-class"

	// S3 extensions
	AmzCopySourceIfModifiedSince   = "x-amz-copy-source-if-modified-since"
	AmzCopySourceIfUnmodifiedSince = "x-amz-copy-source-if-unmodified-since"

	AmzCopySourceIfNoneMatch = "x-amz-copy-source-if-none-match"
	AmzCopySourceIfMatch     = "x-amz-copy-source-if-match"

	AmzCopySource                 = "X-Amz-Copy-Source"
	AmzCopySourceVersionID        = "X-Amz-Copy-Source-Version-Id"
	AmzCopySourceRange            = "X-Amz-Copy-Source-Range"
	AmzMetadataDirective          = "X-Amz-Metadata-Directive"
	AmzObjectLockMode             = "X-Amz-Object-Lock-Mode"
	AmzObjectLockRetainUntilDate  = "X-Amz-Object-Lock-Retain-Until-Date"
	AmzObjectLockLegalHold        = "X-Amz-Object-Lock-Legal-Hold"
	AmzObjectLockBypassGovernance = "X-Amz-Bypass-Governance-Retention"

	// Signature V4 related contants.
	AmzContentSha256        = "X-Amz-Content-Sha256"
	AmzDate                 = "X-Amz-Date"
	AmzAlgorithm            = "X-Amz-Algorithm"
	AmzExpires              = "X-Amz-Expires"
	AmzSignedHeaders        = "X-Amz-SignedHeaders"
	AmzSignature            = "X-Amz-Signature"
	AmzCredential           = "X-Amz-Credential"
	AmzSecurityToken        = "X-Amz-Security-Token"
	AmzDecodedContentLength = "X-Amz-Decoded-Content-Length"

	// Signature v2 related constants
	AmzSignatureV2 = "Signature"
	AmzAccessKeyID = "AWSAccessKeyId"

	// Response request id.
	AmzRequestID = "x-amz-request-id"

	// Deployment id.
	MinioDeploymentID = "x-minio-deployment-id"

	// Server-Status
	MinIOServerStatus = "x-minio-server-status"
)
