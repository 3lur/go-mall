package validator

import (
	"fmt"
	"net/http"

	"github.com/3lur/go-mall/internal/common/reason"
	"github.com/3lur/go-mall/pkg/console"
	"github.com/3lur/go-mall/pkg/e"
	"github.com/3lur/go-mall/pkg/response"
	"github.com/gin-gonic/gin"
)

// BindAndCheck
func BindAndCheck(ctx *gin.Context, data any) bool {
	// 支持 JSON、Form、URL Query
	if err := ctx.ShouldBind(data); err != nil {
		console.Warning(fmt.Sprintf("http_handle ShouldBind error: %s", err))
		response.Build(ctx, e.New(http.StatusBadRequest, reason.RequestBodyError), nil)
		return false
	}

	err := Validate.Struct(data)

	if err != nil {
		ctx.JSON(422, err.Error())
		return false
	}

	return true
}

// BindAndCheckReturnErr 返回错误字段信息
func BindAndCheckReturnErr(ctx *gin.Context, data any) (errFields []*FormErrorField) {
	if err := ctx.ShouldBind(data); err != nil {
		response.Build(ctx, e.New(http.StatusBadRequest, reason.RequestBodyError), nil)
		return nil
	}
	errFields, _ = Get().Check(data)

	return errFields
}
