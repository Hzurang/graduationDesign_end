package service

import (
	"fmt"
	"ginStudy/config"
	"ginStudy/dao"
	"ginStudy/model"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
	"go.uber.org/zap"
	"html"
	"regexp"
	"strconv"
	"strings"
	"time"
)

/*
WordPageNumSpiderService
@author: LJR
@Description: 获取总页数
@param ctx
@param wordType CET4 CET6 TEM4 TEM8 KAOYAN GRE TOEFL IELTS
@return totalPageNum
*/
func WordPageNumSpiderService(wordType string) (totalPageNum int) {
	ct := colly.NewCollector(
		colly.Debugger(&debug.LogDebugger{}), // 开启debug
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36"),
		colly.AllowURLRevisit(),
	)
	ct.OnRequest(func(r *colly.Request) {
	})

	ct.OnError(func(_ *colly.Response, err error) {
		fmt.Println("请求期间发生错误,则调用:", err)
	})

	ct.OnResponse(func(r *colly.Response) {
		fmt.Println("收到响应后调用:", r.Request.URL)
	})

	ct.OnHTML("div[class='container yd-tags'] > a", func(element *colly.HTMLElement) {
		num := element.DOM.Last().Text()
		totalPageNum, _ = strconv.Atoi(num)
	})

	ct.Visit(fmt.Sprintf("https://www.quword.com/tags/%s/0", wordType))
	return totalPageNum
}

/*
WordUrlSpiderService
@author: LJR
@Description: 爬取每页的单词子链接
@param ctx
@param listenType
@param totalPageNum
*/
func WordUrlSpiderService(ctx *colly.Context, wordType string, totalPageNum int) {
	var list = make([]string, 0, 1500)

	ct := colly.NewCollector(
		colly.Debugger(&debug.LogDebugger{}), // 开启debug
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36"),
		colly.AllowURLRevisit(),
	)
	ct.OnRequest(func(r *colly.Request) {

	})

	ct.OnError(func(_ *colly.Response, err error) {
		fmt.Println("请求期间发生错误,则调用:", err)
	})

	ct.OnResponse(func(r *colly.Response) {
		fmt.Println("收到响应后调用:", r.Request.URL)
	})

	ct.OnHTML("div[class='container'] > div[class='panel panel-default'] > div[class='panel-body'] > div[class='row']", func(element *colly.HTMLElement) {
		element.ForEach("div[class='col-sm-6 col-md-3']", func(i int, ee *colly.HTMLElement) {
			wordUrl, _ := ee.DOM.Find("div[class='thumbnail'] > div[class='caption'] > h3 > a").Attr("href")
			urlWord := fmt.Sprintf("https://www.quword.com%s", wordUrl)
			list = append(list, urlWord)
		})
	})
	// 记得改回来total
	for i := 0; i <= totalPageNum-1; i++ {
		time.Sleep(100 * time.Millisecond)
		ct.Visit(fmt.Sprintf("https://www.quword.com/tags/%s/%d", wordType, i))
	}
	for _, v := range list {
		fmt.Println(v)
	}
	ctx.Put("urlList", list)
}

/*
WordSpiderService
@author: LJR
@Description: 爬取每一项里面的具体内容写入数据库
@param ctx
@param wordType
*/
func WordSpiderService(ctx *colly.Context, wordType string) (wordList []*model.Word) {
	wordList = make([]*model.Word, 0, 1000)
	re := regexp.MustCompile(`\[(.+)\]`)

	ct := colly.NewCollector(
		colly.Debugger(&debug.LogDebugger{}), // 开启debug
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36"),
		colly.AllowURLRevisit(),
	)

	ct.OnRequest(func(r *colly.Request) {

	})

	ct.OnError(func(_ *colly.Response, err error) {
		fmt.Println("请求期间发生错误,则调用:", err)
	})

	ct.OnResponse(func(r *colly.Response) {
		ht := string(r.Body)
		ht = html.UnescapeString(ht)
		r.Body = []byte(ht)
		fmt.Println("收到响应后调用:", r.Request.URL)
	})

	ct.OnHTML("div[class='container'] > div[class='row'] > div[id='yd-content']", func(element *colly.HTMLElement) {
		word := new(model.Word)
		word.WordId, _ = config.GenID()
		selection := element.DOM
		word.Word = selection.Find("div[class='row'] > div[id='yd-word-info'] > div[class='row'] > h3[id='yd-word']").Text()
		phoneticStr := selection.Find("div[class='row'] > div[id='yd-word-info'] > div[id='yd-word-pron']").Text()
		phoneticResult := strings.Split(phoneticStr, "\n")
		if len(phoneticResult) == 1 {
			matches := re.FindStringSubmatch(phoneticResult[0])
			if len(matches) == 0 {
				w, _ := GetWordMeaningService(word.Word)
				word.PhoneticTransEng = w.PhoneticTrans1
				word.PhoneticTransAme = w.PhoneticTrans2
			} else {
				word.PhoneticTransEng = matches[0]
				word.PhoneticTransAme = matches[0]
			}
		} else {
			word.PhoneticTransEng = phoneticResult[0]
			word.PhoneticTransAme = phoneticResult[1]
		}
		word.WordMeaning = selection.Find("div[class='row'] > div[id='yd-word-info'] > div[id='yd-word-meaning']").Text()
		word.WordMeaning = strings.Trim(word.WordMeaning, "\n")
		word.MnemonicAid = selection.Find("div[style='font-family:SimSun,serif;']").Text()
		word.ChiEtymology = selection.Find("div[id='yd-ciyuan']").Text()
		element.ForEach("div[id='yd-liju'] > dl", func(i int, ee *colly.HTMLElement) {
			if i == 0 {
				word.SentenceEng1, _ = ee.DOM.Find("dt").Html()
				word.SentenceEng1 = word.SentenceEng1[3:]
				word.SentenceChi1 = ee.DOM.Find("dd").Text()
			}
			if i == 1 {
				word.SentenceEng2, _ = ee.DOM.Find("dt").Html()
				word.SentenceEng2 = word.SentenceEng1[3:]
				word.SentenceChi2 = ee.DOM.Find("dd").Text()
			}
			if i == 2 {
				word.SentenceEng3, _ = ee.DOM.Find("dt").Html()
				word.SentenceEng3 = word.SentenceEng1[3:]
				word.SentenceChi3 = ee.DOM.Find("dd").Text()
			}
		})
		switch wordType {
		case "CET4":
			word.WordType = 1
		case "CET6":
			word.WordType = 2
		case "TEM4":
			word.WordType = 3
		case "TEM8":
			word.WordType = 4
		case "KAOYAN":
			word.WordType = 5
		case "GRE":
			word.WordType = 6
		case "TOEFL":
			word.WordType = 7
		case "IELTS":
			word.WordType = 8
		}
		err := dao.CreateWord(word)
		wordList = append(wordList, word)
		if err != nil {
			zap.L().Error(word.Word, zap.Error(err))
		}
	})
	//ct.Visit("https://www.quword.com/w/allowed")
	//return nil

	urlList := ctx.GetAny("urlList")
	for _, url := range urlList.([]string) {
		if url != "" {
			time.Sleep(100 * time.Millisecond)
			ct.Visit(url)
		}
	}
	return wordList
}
