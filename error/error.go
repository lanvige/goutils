package error

import "fmt"

// Error Error
type Error interface {
	Error() string
	GetCode() int
	GetMsg() string
}

// NewError New returns an error that formats as the given text.
func NewError(text string) Error {
	return &AppError{AppUnrecognized.ErrorCode, text}
}

// NewAppError New returns an error that formats as the given text.
func NewAppError(text string) Error {
	return &AppError{UnknownError.ErrorCode, text}
}

var (
	// AppUnrecognized APP
	// AppUnrecognized APP
	AppUnrecognized = AppError{4000, "未知错误"}
	ReqParamError   = AppError{4001, "请求参数错误"}
	ReqForbidden    = AppError{4003, "请求被拒绝"}
	ReqPathNotExist = AppError{4004, "请求路径不存在"}
	// UnknownError Server Error
	UnknownError      = AppError{5000, "服务端未知错误"}
	Timeout           = AppError{5001, "服务超时"}
	ServerUnavailable = AppError{5005, "服务当前不可用"}
	// AuthAccessForbid AuthAccessForbid
	AuthAccessForbid = AppError{4000, "没有访问权限"}
	AuthExpired      = AppError{4000, "证书过期"}
	AuthInvalid      = AppError{1003, "无效签名"}
	// UserAlreadyExist UserAlreadyExist
	UserAlreadyExist     = AppError{10001, "用户已存在"}
	UserNotExists        = AppError{10001, "用户不存在"}
	UserNamePwdIncorrect = AppError{10000, "用户名或密码错误"}
	UserPhoneIncorrect   = AppError{10002, "手机号码错误"}
	UserStatusError      = AppError{10002, "用户不可用"}
	WrongPassword        = AppError{10002, "密码错误"}
	// OAuthMixinAuthInvalid OAuth
	OAuthMixinAuthInvalid  = AppError{1004, "Mixin 无效签名1"}
	OAuthMixinAuth2Invalid = AppError{1004, "Mixin 无效签名2"}
	OAuthMixinAuth3Invalid = AppError{1004, "Mixin 无效签名3"}
	MixinCreateWalletError = AppError{1004, "Mixin 创建钱包出错"}
	FoxonePermissionDenied = AppError{1005, "权限不足"}
	FoxoneAuthInvalid      = AppError{1006, "Foxone鉴权失败"}
	FoxoneAuthRequired     = AppError{1007, "需Foxone重新授权"}
	DbQueryError           = AppError{5001, "数据库查询错误"}
	DBSessionError         = AppError{5001, "数据库事务错误"}
)

// AppError AppError
type AppError struct {
	ErrorCode int    `json:"errorCode"`
	ErrorMsg  string `json:"errorMsg"`
}

// GetCode GetCode
func (err AppError) GetCode() int {
	return err.ErrorCode
}

// GetMsg GetMsg
func (err AppError) GetMsg() string {
	return err.ErrorMsg
}

// Error Error
func (err AppError) Error() string {
	return fmt.Sprintf("errorCode: %d, errorMsg %s", err.ErrorCode, err.ErrorMsg)
}
