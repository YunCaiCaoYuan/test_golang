package enum

//go:generate stringer -type ErrCode -linecomment

type ErrCode int64 //错误码
const (
	ERR_CODE_OK             ErrCode = 0 // OK
	ERR_CODE_INVALID_PARAMS ErrCode = 1 // 无效参数
	ERR_CODE_TIMEOUT        ErrCode = 2 // 超时
)
