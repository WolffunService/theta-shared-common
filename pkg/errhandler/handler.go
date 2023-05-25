package errhandler

import (
	"errors"

	"github.com/WolffunService/thetan-shared-common/comm"
	"github.com/WolffunService/thetan-shared-common/pkg/validate"
	"github.com/WolffunService/thetan-shared-common/thetalog"
	"github.com/WolffunService/thetan-shared-common/thetanerr"
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

func HandleNotFound(ctx iris.Context) {
	resp := comm.ErrorResponse(thetanerr.ErrHTTPNotFound)
	_ = ctx.StopWithJSON(iris.StatusNotFound, resp)
}

func HandleInternalServerError(ctx iris.Context) {
	const op = "HandleInternalServerError"

	err := ctx.GetErr()
	if err != nil {
		thetalog.Err(err).Op(op).Send()
	} else {
		err = thetanerr.ErrHTTPInternalServerError
	}

	resp := comm.ErrorResponse(err)
	_ = ctx.StopWithJSON(iris.StatusInternalServerError, resp)
}

func Default(ctx iris.Context, err error) {
	if iris.IsErrPath(err) || errors.Is(err, context.ErrEmptyForm) {
		return // continue.
	}

	// Catch the first error from validator if exists
	if errs, ok := err.(validator.ValidationErrors); ok {
		for _, field := range errs {
			code := thetanerr.ErrWrongFormat
			message := field.Translate(validate.Trans)
			resp := comm.CodeResponse(code, message)
			_ = ctx.StopWithJSON(iris.StatusBadRequest, resp)
			return
		}
	}

	resp := comm.ErrorResponse(err)
	resp.Code = thetanerr.ErrWrongFormat

	if iris.IsErrEmptyJSON(err) {
		resp.Message = "Request body is empty"
	}

	_ = ctx.StopWithJSON(iris.StatusBadRequest, resp)
}
