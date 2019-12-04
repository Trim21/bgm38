package vote

import (
	"bgm38/pkg/server/model"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"strconv"
)

func index(ctx *gin.Context) {
	//ctx.String(200, "hello world")
	ctx.HTML(200, "vote/index.tmpl", map[string]string{
		"Title":   "hello Vote Function",
		"Message": "hello world",
	})
}

func createUI(ctx *gin.Context) {
	ctx.HTML(200, "vote/create.tmpl", map[string]string{
		"Title":   "hello Vote Function",
		"Message": "hello world",
	})
}

type SubmitVote struct {
	Title   string   `form:"title" binding:"required"`
	Options []string `form:"options" binding:"required"`
}

func create(ctx *gin.Context) {
	var form = &SubmitVote{}
	if ctx.ShouldBind(form) != nil {
		ctx.String(401, "submit data error")
		return
	}

	var vote = model.Vote{Title: form.Title, Creator: 280000}

	model.DB.Create(&vote)

	for _, text := range form.Options {
		model.DB.Create(&model.VoteOption{Text: text, VoteID: vote.ID})
	}

	ctx.JSON(201, form)
}

func json(ctx *gin.Context) {
	var id, err = strconv.Atoi(ctx.Query("id"))
	if err != nil {
		ctx.JSON(401, gin.H{"error": "id not int"})
		return
	}

	var voteFull model.VoteFull

	voteFull, err = model.GetVoteFull(uint(id))

	if gorm.IsRecordNotFoundError(err) {
		ctx.JSON(404, gin.H{"error": "vote not found"})
		return
	}

	ctx.JSON(200, voteFull)
}
