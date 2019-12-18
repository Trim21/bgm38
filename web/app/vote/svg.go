package vote

import (
	"bgm38/web/app/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"io"
	"text/template"
)

type person struct {
	ID uint `uri:"id" binding:"required"`
	//Name string `uri:"name" binding:"required"`
}

const f = `<?xml version="1.0" standalone="yes"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">

<svg xmlns="http://www.w3.org/2000/svg"  height="{{ step .Length 60 40 }}" version="1.1">
  <text x="20" y="30" font-size="30">
    {{ .Title }}
  </text>

{{ range $index, $element := .Options }} 

  <a href="./233">
	 <text x="25" y="{{ step $index 60 60 }}" font-size="20">
	    {{ $element.Text }}
     </text>
  </a>

 
  <rect id="option_{{ $index }}" 
	  height="30" 
	  width="{{ length $element.Count $.MaxCount $.MaxLength }}" 
	  y="{{ step $index 60 70 }}"
	  x="25"
      style="fill:rgb(0, 0, 0);stroke-width:1;stroke:rgb(0,0,0)"
  />

{{end}}

</svg>
`

var t *template.Template

func init() {
	var err error
	t, err = template.New("svg").Funcs(template.FuncMap{
		"step": func(sort int, step int, base int) int {
			return sort*step + base
		},
		"length": func(count int, maxCount, maxLength int) float32 {
			return float32(count) / float32(maxCount) * float32(maxLength)

		},
	}).Parse(f)
	if err != nil {
		panic(err)
	}

}

const notFound = `<?xml version="1.0" standalone="yes"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">

<svg xmlns="http://www.w3.org/2000/svg" version="1.1">
  <text x="250" y="150" 
        font-size="55">
    Not Found
  </text>
</svg>
  `

func svg(ctx *gin.Context) {
	var p person
	if err := ctx.ShouldBindUri(&p); err != nil {
		fmt.Println(err)
		ctx.JSON(400, gin.H{"msg": err})
		return
	}
	var vote, err = model.GetVoteFull(p.ID)

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			writeSvg(ctx, 404, notFound)
			return
		}
		ctx.String(502, "internal error")
		return
	}

	ctx.Header("Content-Type", "image/svg+xml; charset=utf-8")
	ctx.Status(200)
	err = t.Execute(ctx.Writer, gin.H{
		"Title":     vote.Title,
		"Options":   vote.Options,
		"MaxLength": 200,
		"MaxCount":  maxCount(vote.Options),
		"Length":    len(vote.Options),
	})
	if err != nil {
		fmt.Println(err)
	}

}

func writeSvg(ctx *gin.Context, status int, data string) {
	ctx.Header("Content-Type", "image/svg+xml; charset=utf-8")
	ctx.Status(status)
	_, _ = io.WriteString(ctx.Writer, data)
	return
}

func maxCount(m []model.VoteOption) int {
	var max int
	for _, v := range m {
		if v.Count > max {
			max = v.Count
		}
	}
	return max
}
