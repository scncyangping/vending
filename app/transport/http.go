package transport

import (
	"github.com/gin-gonic/gin"
	"vending/config/log"
)

const (
	DefaultStatus             = -1
	StatusOK                  = 200
	StatusMovedPermanently    = 301
	StatusFound               = 302
	StatusBadRequest          = 400
	StatusNotFound            = 404
	StatusInternalServerError = 500

	// Request Error
	RequestParameterError    = 1001
	RequestCheckTokenError   = 1002
	RequestCheckTokenTimeOut = 1003
	RequestTokenNotFound     = 1004
	CreateTokenError         = 1005
	// System  Error
	DataConvertError      = 2001
	ParameterConvertError = 2002
	// DataBase Error
	InitDataBaseError = 3001
	QueryDBError      = 3002
	UserNotFound      = 3003
	AddUserError      = 3004
	DataNotFound      = 3005
	//
	RateLimit = 4001
)

const (
	MethodGet     = "GET"
	MethodHead    = "HEAD"
	MethodPost    = "POST"
	MethodPut     = "PUT"
	MethodPatch   = "PATCH"
	MethodDelete  = "DELETE"
	MethodConnect = "CONNECT"
	MethodOptions = "OPTIONS"
	MethodTrace   = "TRACE"

	ConnectTypeJson = "application/json"
	ConnectTypeWww  = "application/x-www-form-urlencoded"
)

var statusText = map[int]string{
	DefaultStatus:             "",
	StatusOK:                  "OK",
	StatusBadRequest:          "Bad Request",
	StatusMovedPermanently:    "Moved Permanently",
	StatusFound:               "Found",
	StatusNotFound:            "Not Found",
	StatusInternalServerError: "Internal Server Error",
	RequestParameterError:     "Request Parameter Error",
	DataConvertError:          "Data Convert Error",
	RequestCheckTokenError:    "Token Is Not Exists, Please Login",
	ParameterConvertError:     "Parameter Error, Please Check Parameter",
	InitDataBaseError:         "Init DataBase Error",
	QueryDBError:              "Query DataBase Error",
	RequestCheckTokenTimeOut:  "request check token time out",
	RequestTokenNotFound:      "request token not found, please login first",
	UserNotFound:              "user not found",
	CreateTokenError:          "create token error",
	AddUserError:              "add user error",
	DataNotFound:              "data not found",
	RateLimit:                 "rate limit",
}

func StatusText(code int) string {
	return statusText[code]
}

type ListQuery struct {
	Skip  int    `json:"skip"`
	Limit int    `json:"limit"`
	Sort  string `json:"sort"`
}

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// ctx data code msg
func SendSuccess(ctx *gin.Context, arg ...interface{}) {
	var (
		result   *Result
		actionId = ctx.GetString("ActionId")
	)

	switch len(arg) {
	case 0:
		result = &Result{StatusOK, StatusText(StatusOK), nil}
	case 1:
		result = &Result{StatusOK, StatusText(StatusOK), arg[0]}
	case 2:
		result = &Result{arg[1].(int), StatusText(StatusOK), arg[0]}
	case 3:
		result = &Result{arg[1].(int), arg[2].(string), arg[0]}
	default:
		panic("parameter error")
	}

	log.Logger.Infof("Response [%s] To Client: %v", actionId, result)
	ctx.JSON(StatusOK, result)
}

func SendFailure(ctx *gin.Context, arg ...interface{}) {
	var (
		result   *Result
		actionId = ctx.GetString("ActionId")
	)

	switch len(arg) {
	case 0:
		result = &Result{StatusBadRequest, StatusText(StatusBadRequest), nil}
	case 1:
		result = &Result{StatusBadRequest, StatusText(StatusBadRequest), arg[0]}
	case 2:
		result = &Result{arg[1].(int), StatusText(StatusBadRequest), arg[0]}
	case 3:
		result = &Result{arg[1].(int), arg[2].(string), arg[0]}
	default:
		panic("parameter error")
	}

	log.Logger.Infof("Response [%s] To Client: %v", actionId, result)
	ctx.JSON(StatusOK, result)
}
