package constants

const (
	ERROR_SIGNATURE_REQUIRED   = "Signature is required"
	ERROR_UNAUTHORIZED         = "Unauthorized"
	ERROR_INVALID_TOKEN        = ERROR_UNAUTHORIZED + ", token not valid"
	ERROR_USER_TOKEN_NOT_FOUND = ERROR_UNAUTHORIZED + ", user not found"

	FAILED_BUILD_QUERY           = "failed to build SQL query"
	FAILED_EXECUTE_QUERY         = "failed to execute SQL query"
	FAILED_SCANNING_QUERY_PARAMS = "failed to scan SQL query parameters"
)
