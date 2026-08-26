package core_errors

import "errors"

type ErrorsMessage struct {
	Status     int    `json:"-"`
	StatusCode string `json:"status_code"`
	Details    string `json:"details"`
}

var (
	ErrBadRequest         = errors.New("BAD_REQUEST")
	ErrUserNotFound       = errors.New("USER_NOT_FOUND")
	ErrInvalidCredentials = errors.New("INVALID_CREDENTIALS")
	ErrUsernameExists     = errors.New("USERNAME_ALREADY_EXISTS")
	ErrValidation         = errors.New("VALIDATION_ERROR")
	ErrHashPassword       = errors.New("HASH_PASSWORD_ERROR")
	ErrCreateUser         = errors.New("CREATE_USER_ERROR")
	ErrGenerateJWT        = errors.New("GENERATE_JWT_ERROR")
	ErrCacheUnmarshal     = errors.New("CACHE_UNMARSHAL_ERROR")
	ErrTokenInvalid       = errors.New("TOKEN_INVALID")
)

var Errors = map[error]ErrorsMessage{
	ErrBadRequest: {
		Status:     400,
		StatusCode: ErrBadRequest.Error(),
		Details:    "bad request",
	},

	ErrUserNotFound: {
		Status:     404,
		StatusCode: ErrUserNotFound.Error(),
		Details:    "user not found",
	},

	ErrInvalidCredentials: {
		Status:     401,
		StatusCode: ErrInvalidCredentials.Error(),
		Details:    "invalid credentials",
	},

	ErrUsernameExists: {
		Status:     409,
		StatusCode: ErrUsernameExists.Error(),
		Details:    "username already exists",
	},

	ErrValidation: {
		Status:     400,
		StatusCode: ErrValidation.Error(),
		Details:    "validation error",
	},

	ErrHashPassword: {
		Status:     500,
		StatusCode: ErrHashPassword.Error(),
		Details:    "failed to hash password",
	},

	ErrCreateUser: {
		Status:     500,
		StatusCode: ErrCreateUser.Error(),
		Details:    "failed to create user",
	},

	ErrGenerateJWT: {
		Status:     500,
		StatusCode: ErrGenerateJWT.Error(),
		Details:    "failed to generate jwt",
	},

	ErrCacheUnmarshal: {
		Status:     500,
		StatusCode: ErrCacheUnmarshal.Error(),
		Details:    "failed to unmarshal cached user",
	},

	ErrTokenInvalid: {
		Status:     401,
		StatusCode: ErrTokenInvalid.Error(),
		Details:    "token invalid",
	},
}

func ErrorsResponse(err error) ErrorsMessage {
	errs, ok := Errors[err]
	if !ok {
		return Errors[ErrBadRequest]
	}

	return errs
}
