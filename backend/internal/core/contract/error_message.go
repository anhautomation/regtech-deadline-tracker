package contract

var ErrorMessage = map[string]string{
	SUCCESS:        "success",
	INVALID:        "invalid request",
	NOT_FOUND:      "resource not found",
	ALREADY_EXISTS: "resource already exists",
	UNAUTHORIZED:   "unauthorized",
	FORBIDDEN:      "forbidden",
	UPSTREAM_ERROR: "upstream provider error",
	TIMEOUT:        "request timeout",
	CONFLICT:       "conflict",
	INTERNAL:       "internal server error",
}
