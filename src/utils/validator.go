package utils

import (
	"OpenSchedule/src/response"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"regexp"
	"sync"
)

var defaultValidator *validator.Validate
var validatorOnce sync.Once

func getValidator() *validator.Validate {
	validatorOnce.Do(func() {
		defaultValidator = validator.New()
		err := defaultValidator.RegisterValidation("hh:mm", ValidateDateTime)
		if err != nil {
			fmt.Println("validator RegisterValidation error: " + err.Error())
		}
	})
	return defaultValidator
}

func ValidateParam(ctx iris.Context, param interface{}) error {
	err := ctx.ReadJSON(param)
	if  err != nil {
		response.Fail(ctx, response.Error, err.Error(), nil)
		return  err
	}

	err = getValidator().Struct(param)
	if err != nil {
		response.Fail(ctx, response.Error, err.Error(), nil)
		return err
	}

	return nil
}

func ValidateDateTime(f validator.FieldLevel) bool {
	dateTime := f.Field().String()
	return CheckDateTime(dateTime)
}

func CheckDateTime(dateTime string) bool {
	regx, err := regexp.Compile("^([0-1]?[0-9]|2[0-3]):[0-5][0-9]$")
	if err != nil {
		fmt.Println(err.Error())
	}
	isValid := regx.MatchString(dateTime)
	return isValid
}