package vote

import (
	"strconv"

	"github.com/gofiber/fiber"
	"go.uber.org/zap"

	"bgm38/app/web/res"
	"bgm38/app/web/utils/handler"
	"bgm38/pkg/fetch"
	"bgm38/pkg/parser"
)

func Group(app *fiber.Group) {
	app.Get("/vote/:topic_id", handler.LogError(vote))
}

// @ID voteResult
// @Summary 解析帖子，生成投票结果
// @Description 解析帖子，生成投票结果
// @Produce  text/plain
// @Param topic_id path int true "user_id, 可以在个人主页的网址中找到"
// @Success 200 {string} string "text/calendar"
// @Failure 422 {object} res.Error
// @Failure 404 {object} res.Error
// @Failure 502 {object} res.Error
// @Router /bgm.tv/v1/vote/{topic_id} [get]
func vote(c *fiber.Ctx, logger *zap.Logger) error {
	var topicInput = c.Params("topic_id")
	topicID, err := strconv.Atoi(topicInput)
	if err != nil {
		return c.Status(401).JSON(res.Error{
			Status:  "error",
			Message: "`topic_id` should be int",
		})
	}

	doc, err := fetch.Topic(topicID)
	if err != nil {
		return err
	}

	t, err := parser.Topic(doc)
	if err != nil {
		return err
	}

	voteOptions, err := extraRawOption(t)
	if err != nil {
		return err
	}

	o, err := parseOption(voteOptions)
	if err != nil {
		// todo: show that options is not correct
		return err
	}

	return c.JSON(o)
}
