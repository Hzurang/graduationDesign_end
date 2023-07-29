package service

import (
	"fmt"
	"ginStudy/config"
	"ginStudy/dao"
	"ginStudy/model"
	"github.com/PuerkitoBio/goquery"
	"github.com/axgle/mahonia"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func EssayPageNumSpiderService(ctx *colly.Context, essayType string) {
	decoder := mahonia.NewDecoder("gbk")

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

	ct.OnHTML("div[class='page'] > li > a", func(element *colly.HTMLElement) {
		str := element.Text
		str = decoder.ConvertString(str)
		if str == "末页" {
			href := element.Attr("href")
			doc := strings.Index(href, ".")
			href = href[5:doc]
			ctx.Put("totalPageNum", href)
		}
	})
	if essayType == "love" || essayType == "shuangyu" {
		ct.Visit(fmt.Sprintf("https://www.enread.com/story/%s/index.html", essayType))
	} else {
		ct.Visit(fmt.Sprintf("https://www.enread.com/%s/index.html", essayType))
	}
}

func EssayFirstSpiderService(ctx *colly.Context, essayType string) {
	//var hr string
	//var end string
	var list = make([]string, 0, 3000)

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

	ct.OnHTML("div[class='list']", func(element *colly.HTMLElement) {
		var pageList = make([]string, 0, 35)
		element.ForEach("div[class='node_list']", func(i int, ee *colly.HTMLElement) {
			str, _ := ee.DOM.Find("h2 > a").Attr("href")
			str = fmt.Sprintf("http://www.enread.com%s", str)
			pageList = append(pageList, str)
		})
		for i := len(pageList) - 1; i >= 0; i-- {
			list = append(list, pageList[i])
		}
	})

	strGet := ctx.Get("totalPageNum")
	num, _ := strconv.Atoi(strGet)
	fmt.Println(num)
	if essayType == "love" || essayType == "shuangyu" {
		for i := num; i >= 1; i-- {
			ct.Visit(fmt.Sprintf("https://www.enread.com/story/%s/list_%d.html", essayType, i))
		}
	} else {
		for i := num; i >= 1; i-- {
			ct.Visit(fmt.Sprintf("https://www.enread.com/%s/list_%d.html", essayType, i))
		}
	}
	ctx.Put("Urlhref", list)
}

func EssaySpiderService(ctx *colly.Context, essayType string) {
	// 创建一个 GB18030 编码器
	decoder := mahonia.NewDecoder("gbk")
	re := regexp.MustCompile(`文章作者：(\w+) 发布时间：([\d-]+\s[\d:]+)`)
	re1 := regexp.MustCompile(`发布时间：([\d-]+\s[\d:]+)`)

	pattern := regexp.MustCompile(`([\p{Han}]+)。([^ ])`)
	dateformat := "2006-01-02 15:04"

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
		html := string(r.Body)
		html = strings.ReplaceAll(html, "&ldquo;", `"`)
		html = strings.ReplaceAll(html, "&rdquo;", `"`)
		html = strings.ReplaceAll(html, "&mdash;", "--")
		html = strings.ReplaceAll(html, "&quot;", "\"")
		html = strings.ReplaceAll(html, "&#39;", "'")
		html = strings.ReplaceAll(html, "&lsquo;", "'")
		html = strings.ReplaceAll(html, "&rsquo;", "'")
		html = strings.ReplaceAll(html, "&cent;", "￠")
		html = strings.ReplaceAll(html, "&pound;", "£")
		html = strings.ReplaceAll(html, "&yen;", "¥")
		html = strings.ReplaceAll(html, "&nbsp;", " ")
		html = strings.ReplaceAll(html, "&hellip;", "...")
		r.Body = []byte(html)
		fmt.Println("收到响应后调用:", r.Request.URL)
	})

	ct.OnHTML("div[class='index']", func(el *colly.HTMLElement) {
		var txt string
		var str string
		var essayId uint64
		essay := new(model.Essay)
		essayId, _ = config.GenID()
		essay.EssayId = essayId
		// 文章标题
		title := el.ChildText("div[align='center'] > font[style='font-size: 18px; font-weight: bold;']")
		essay.EssayTitle = decoder.ConvertString(title)
		//fmt.Println(title)
		// 文章作者 match[1] 和发布时间 match[2]
		m := el.ChildText("td[bgcolor='#E7F1FA'] > div[align='center']")
		m = decoder.ConvertString(m)
		match := re.FindStringSubmatch(m)
		if len(match) != 0 {
			essay.EssayAuthor = match[1]
			essay.PublishAt, _ = time.Parse(dateformat, match[2])
		} else {
			match1 := re1.FindStringSubmatch(m)
			essay.EssayAuthor = "乐学"
			essay.PublishAt, _ = time.Parse(dateformat, match1[1])
		}
		// 文章内容txt
		el.ForEach("p", func(i int, element *colly.HTMLElement) {
			str = element.Text
			element.DOM.Find("sup").Each(func(i int, s *goquery.Selection) {
				str = strings.ReplaceAll(str, s.Text(), "")
			})
			utf8Str := decoder.ConvertString(str)
			utf8Str = strings.TrimSpace(utf8Str)
			txt = txt + utf8Str + "\n"
		})
		if txt == "" {
			el.ForEach("#dede_content > div", func(i int, ee *colly.HTMLElement) {
				index := ee.Index
				total := ee.DOM.Parent().Children().Length()
				if total-index == 2 {
					ee.ForEach("div", func(i int, eet *colly.HTMLElement) {
						str = eet.Text
						eet.DOM.Find("sup").Each(func(i int, s *goquery.Selection) {
							str = strings.ReplaceAll(str, s.Text(), "")
						})
						utf8Str := decoder.ConvertString(str)
						utf8Str = strings.TrimSpace(utf8Str)
						txt = txt + utf8Str + "\n"
					})
					return
				}
				str = ee.Text
				ee.DOM.Find("sup").Each(func(i int, s *goquery.Selection) {
					str = strings.ReplaceAll(str, s.Text(), "")
				})
				utf8Str := decoder.ConvertString(str)
				utf8Str = strings.TrimSpace(utf8Str)
				txt = txt + utf8Str + "\n"
			})
		}
		essay.EssayContent = txt
		switch essayType {
		case "novel":
			essay.EssayType = 1 //小说
		case "love":
			essay.EssayType = 2 //情感故事
		case "essays":
			essay.EssayType = 3 //英语美文
		case "shuangyu":
			essay.EssayType = 4 //双语故事
		}
		essay.EssayCollectNum = 0
		if essay.EssayContent == "" {
			el.Request.Abort()
			return
		}
		dao.CreateEssay(essay)
		var essayWordList = make([]*model.EssayWord, 0, 120)
		// 文章单词内容
		el.ForEach("div[class='zc_boxl']", func(i int, et *colly.HTMLElement) {
			word := new(model.EssayWord)
			word.WordNum, _ = strconv.Atoi(et.ChildText("span[class='circle']"))
			word.Word = et.ChildText("span[class='danci'] > a[target='_blank']")
			b := et.DOM.Find("img[src='/images/play.gif']")
			mp3Tag, _ := b.Attr("alt")
			mTag := string(mp3Tag[0])
			word.WordMusic = fmt.Sprintf("http://sound.yywz123.com/qsbdcword/%s/%s.mp3", mTag, mp3Tag)
			word.WordId, _ = config.GenID()
			word.EssayId = essayId
			strMean := et.ChildText("td[class='zhushi']")
			word.WordMeaning = decoder.ConvertString(strMean)
			sentence := et.ChildText("div[class='liju'] > ul > li")
			sentence = decoder.ConvertString(sentence)
			word.WordSentence = pattern.ReplaceAllString(sentence, "${1}。\n$2")
			essayWordList = append(essayWordList, word)
			dao.CreateEssayWord(word)
		})
	})
	id := ctx.GetAny("Urlhref")
	for _, v := range id.([]string) {
		if v != "" {
			ct.Visit(v)
		}
	}
}
