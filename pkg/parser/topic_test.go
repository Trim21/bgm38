package parser

import (
	"os"
	"testing"
	"time"

	"github.com/antchfx/htmlquery"
	"github.com/stretchr/testify/assert"
)

func Test_parseReply(t *testing.T) {
	f, _ := os.Open("../../tests/fixtures/topic.html")
	doc, _ := htmlquery.Parse(f)
	replyList, err := getReplyList(doc)
	assert.Nil(t, err)
	assert.Equal(t, 3, len(replyList))
	assert.Equal(t, 4, len(replyList[0].Replies))
	assert.Equal(t, 0, len(replyList[1].Replies))
	assert.Equal(t, 0, len(replyList[2].Replies))
	for _, reply := range replyList {
		assert.NotNil(t, reply.RawContent)
		for _, subReply := range reply.Replies {
			assert.NotNil(t, subReply.RawContent)
		}
	}
}

func Test_Topic(t *testing.T) {
	f, _ := os.Open("../../tests/fixtures/topic.html")
	doc, _ := htmlquery.Parse(f)
	topic, err := Topic(doc)
	assert.Nil(t, err)
	assert.NotNil(t, topic.RawContent)
	assert.Equal(t, "来推冻鳗群啦！", topic.Title)
	assert.Equal(t, "/user/455074", topic.Author)
	assert.Equal(t, 3, len(topic.Replies))
}

func Test_reInfo1(t *testing.T) {
	f, ti, err := reInfo("#1 - 2020-5-4 04:27")
	assert.Nil(t, err)
	assert.Equal(t, 1, f[0])
	assert.Equal(t, 0, f[1])
	assert.Equal(t, 2020, ti.Year(), ti)
	assert.Equal(t, time.Month(5), ti.Month(), ti)
	assert.Equal(t, 4, ti.Day(), ti)
	assert.Equal(t, 4, ti.Hour(), ti)
	assert.Equal(t, 27, ti.Minute(), ti)
	assert.Equal(t, 0, ti.Second(), ti)
	// assert.Equal(t, )
}
func Test_reInfo2(t *testing.T) {
	f, ti, err := reInfo("#2-1 - 2020-5-4 10:58")
	assert.Nil(t, err)
	assert.Equal(t, 2, f[0])
	assert.Equal(t, 1, f[1])
	assert.Equal(t, 2020, ti.Year(), ti)
	assert.Equal(t, time.Month(5), ti.Month(), ti)
	assert.Equal(t, 4, ti.Day(), ti)
	assert.Equal(t, 10, ti.Hour(), ti)
	assert.Equal(t, 58, ti.Minute(), ti)
	assert.Equal(t, 0, ti.Second(), ti)
	// assert.Equal(t, )
}
func Test_reInfo3(t *testing.T) {
	f, ti, err := reInfo("#5-9 - 2020-12-30 17:27")
	assert.Nil(t, err)
	assert.Equal(t, 5, f[0])
	assert.Equal(t, 9, f[1])
	assert.Equal(t, 2020, ti.Year(), ti)
	assert.Equal(t, time.Month(12), ti.Month(), ti)
	assert.Equal(t, 30, ti.Day(), ti)
	assert.Equal(t, 17, ti.Hour(), ti)
	assert.Equal(t, 27, ti.Minute(), ti)
	assert.Equal(t, 0, ti.Second(), ti)
	// assert.Equal(t, )
}
