package apierr

type ApiError struct {
	Error string `json:"error"`
}

func NewApiError(error string) ApiError {
	return ApiError{Error: error}
}

var InternalServerError = ApiError{
	Error: "INTERNAL_SERVER",
}

var InvalidJsonError = ApiError{
	Error: "INVALID_JSON",
}

var NotFoundError = ApiError{
	Error: "NOT_FOUND",
}

var WrongCredentials = ApiError{
	Error: "WRONG_CREDENTIALS",
}

var CookieNotExists = ApiError{
	Error: "COOKIE_NOT_EXISTS",
}

var SessionNotFound = ApiError{
	Error: "SESSION_NOT_FOUND",
}

var UserNotFound = ApiError{
	Error: "USER_NOT_FOUND",
}

var Unauthorized = ApiError{
	Error: "UNAUTHORIZED",
}

var UserWithThisUsernameExists = ApiError{
	Error: "USER_WITH_THIS_USERNAME_EXISTS",
}

var EventNotFound = ApiError{
	Error: "EVENT_NOT_FOUND",
}

var InvalidQueryParams = ApiError{
	Error: "INVALID_QUERY_PARAMS",
}
