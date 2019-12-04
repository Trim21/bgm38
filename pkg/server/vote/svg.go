package vote

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
)

type Person struct {
	ID string `uri:"id" binding:"required"`
	//Name string `uri:"name" binding:"required"`
}

const f = `<?xml version="1.0" standalone="yes"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">

<svg xmlns="http://www.w3.org/2000/svg" version="1.1">
  <circle cx="100" cy="50" r="40" stroke="black"
  stroke-width="2" fill="red" />
</svg>
`

func svg(ctx *gin.Context) {
	var person Person
	if err := ctx.ShouldBindUri(&person); err != nil {
		fmt.Println(err)
		ctx.JSON(400, gin.H{"msg": err})
		return
	}
	writeSvg(ctx, f)
}

func writeSvg(ctx *gin.Context, data string) {
	ctx.Header("Content-Type", "image/svg+xml; charset=utf-8")
	ctx.Status(200)
	_, _ = io.WriteString(ctx.Writer, data)
	return
}
