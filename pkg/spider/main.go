package spider

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/antchfx/htmlquery"
	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
	"golang.org/x/net/html"

	"bgm38/pkg/db"
	"bgm38/pkg/utils/log"
)

var logger *zap.Logger

func Start() error {
	fmt.Println("spider start")
	logger = log.BindMeta(log.CreateLogger("bgm38-spider-v1"))
	db.InitDB()
	prepareStmt()
	var urlToFetch = make(chan string)
	var resQueue = make(chan response)
	var workerCount = 20
	var parserCount = 3

	for i := 0; i < parserCount; i++ {
		go parser(resQueue)
	}
	for i := 0; i < workerCount; i++ {
		go downloader(urlToFetch, resQueue)
	}
	go dispatcher(urlToFetch)
	// urlToFetch <- fmt.Sprintf("https://mirror.bgm.rin.cat/subject/%d", 889)
	//
	// for i := 1; i < 296800; i++ {
	// 	urlToFetch <- fmt.Sprintf("https://mirror.bgm.rin.cat/subject/%d", i)
	// }
	ch := make(chan int)
	<-ch
	return nil
}

func getSubjectID(url string) (int, error) {
	sp := strings.Split(url, "/")
	return strconv.Atoi(sp[len(sp)-1])

}

type response struct {
	url string
	res *resty.Response
}

func getImageURL(docs *html.Node) string {
	cover := htmlquery.SelectAttr(htmlquery.FindOne(docs, `//*[@id="bangumiInfo"]/div/div/a/img`), "src")
	if cover == "" {
		return "lain.bgm.tv/img/no_icon_subject.png"
	}
	return strings.Replace(cover, "//lain.bgm.tv/pic/cover/c/", "lain.bgm.tv/pic/cover/g/", -1)
}

func getScore(doc *html.Node) string {
	v := htmlquery.FindOne(doc,
		`//*[@id="panelInterestWrapper"]//div[@class="global_score"]/span[1]`)
	if v == nil {
		return ""
	}
	return htmlquery.InnerText(v)
}
