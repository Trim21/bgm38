package auth

import (
	"fmt"
	"reflect"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

// Part bind auth part to gin app
func Part(app *gin.Engine) {
	var router = app.Group("/auth")
	router.GET("/v1/bgm.tv/callback", callback)
	router.GET("/v1/bgm.tv/redirect", redirect)

	app.GET("/test/:a/:b/:c/:d", test)
}

type bind struct {
	A string `uri:"a" query:"a" binding:"required" json:"a" form:"a"`
	B string `uri:"b" binding:"required" form:"b"`
	C int    `uri:"c" binding:"required" form:"c"`
	D int    `uri:"d" binding:"required" form:"d"`
}

func test(ctx *gin.Context) {
	var v bind
	if err := ctx.ShouldBindQuery(&v); err != nil {
		fmt.Println(err)
		fmt.Println(reflect.TypeOf(err))
		fmt.Println(reflect.ValueOf(err))
		// for _, fieldErr := range err.(validator.ValidationErrors) {
		// 	fmt.Println(fieldErr)
		// 	return // exit on first error
		// }
		for _, fieldErr := range err.(validator.ValidationErrors) {
			fmt.Println(fieldErr.Field(),
				fieldErr.Kind(),
				fieldErr.Tag(),
				fieldErr.ActualTag(),
				fieldErr.Namespace(),
				fieldErr.StructNamespace())
		}
		ctx.String(422, err.Error())
		return
	}
	ctx.JSON(200, v)
}
