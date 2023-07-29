package service

import (
	"fmt"
	"ginStudy/config"
	"ginStudy/dao"
	"ginStudy/model"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
	"html"
	"regexp"
	"strconv"
	"strings"
	"time"
)

/*
ListenPageNumSpiderService
@author: LJR
@Description: 获取总页数
@param ctx
@param listenType
@return totalPageNum
*/
func ListenPageNumSpiderService(listenType string) (totalPageNum int) {
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

	ct.OnHTML("div[class='page th'] > strong > span[id='total']", func(element *colly.HTMLElement) {
		total := element.Text
		totalPageNum, _ = strconv.Atoi(total)
	})

	ct.Visit(fmt.Sprintf("http://www.kekenet.com/Article/%s/", listenType))
	return totalPageNum
}

/*
ListenUrlSpiderService
@author: LJR
@Description: 爬取每页的子项链接
@param ctx
@param listenType
@param totalPageNum
*/
func ListenUrlSpiderService(ctx *colly.Context, listenType string, totalPageNum int, isCron bool) {
	var list = make([]string, 0, 600)

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

	ct.OnHTML("ul[id='menu-list']", func(element *colly.HTMLElement) {
		var urlList = make([]string, 0, 20)
		if listenType == "media" {
			element.ForEach("li", func(i int, ee *colly.HTMLElement) {
				url, _ := ee.DOM.Find("a").Attr("href")
				urlWhole := fmt.Sprintf("http://www.kekenet.com%s", url)
				urlList = append(urlList, urlWhole)
			})
		} else if listenType == "17698" {
			element.ForEach("li", func(i int, ee *colly.HTMLElement) {
				url, _ := ee.DOM.Find("h2 > a").Attr("href")
				urlList = append(urlList, url)
			})
		} else if listenType == "chuji" {
			element.ForEach("li", func(i int, ee *colly.HTMLElement) {
				url := ee.ChildAttr("h2 a:nth-of-type(2)", "href")
				urlList = append(urlList, url)
			})
		}
		for i := len(urlList) - 1; i >= 0; i-- {
			list = append(list, urlList[i])
		}
	})
	if isCron == true {
		ct.Visit(fmt.Sprintf("http://www.kekenet.com/Article/%s/", listenType))
		ctx.Put("urlList", list)
	}
	if isCron == false {
		// 改1
		for i := 1; i <= totalPageNum; i++ {
			time.Sleep(120 * time.Millisecond)
			if i == totalPageNum {
				ct.Visit(fmt.Sprintf("http://www.kekenet.com/Article/%s/", listenType))
				break
			}
			ct.Visit(fmt.Sprintf("http://www.kekenet.com/Article/%s/List_%d.shtml", listenType, i))
		}
		ctx.Put("urlList", list)
	}
}

/*
ListenSpiderService
@author: LJR
@Description: 爬取每一项里面的具体内容写入数据库
@param ctx
@param listenType
@return cnt1
@return cnt2
*/
func ListenSpiderService(ctx *colly.Context, listenType string) (cnt1 int, cnt2 int) {
	layout := "2006-01-02 15:04:05"
	re1 := regexp.MustCompile(`编辑:(\w+)`)
	re2 := regexp.MustCompile(`thunder_url\s*=\s*"([^"]+)"`)
	re3 := regexp.MustCompile(`var thunder_url ="(.*)";`)
	cnt1 = 0
	cnt2 = 0

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
		//ht = strings.ReplaceAll(ht, "<div style=\"width:360px;min-height:44px;position:absolute;left:310px;z-index:999;border:1px solid #A9D4DB;font-size:12px;background-color:#FFF;display:none;\" class=\"lx_box\">", "")
		r.Body = []byte(ht)
		fmt.Println("收到响应后调用:", r.Request.URL)
	})

	ct.OnHTML("div[class='lastPage_left']", func(element *colly.HTMLElement) {
		cnt := 1
		flag := false
		var contentStr string
		listen := new(model.Listen)
		listen.ListenSource = "可可英语"
		listen.ListenId, _ = config.GenID()
		listen.DeleteIsOk = 0
		listen.ListenCollectNum = 0
		listen.ListenTitle = element.DOM.Find("div[class='e_title'] > h1[id='nrtitle']").Text()

		timeStr := element.DOM.Find("div[class='e_title'] > cite > time").Text()
		timeStr = strings.Replace(timeStr, "时间:", "", 1)
		listen.PublishAt, _ = time.Parse(layout, timeStr)

		editor := element.DOM.Find("div[class='e_title'] > cite").Text()
		match1 := re1.FindStringSubmatch(editor)
		listen.ListenEditor = match1[1]

		element.ForEach("div[class='info-qh'] > div", func(i int, ee *colly.HTMLElement) {
			contentStr += ee.Text + "\n"
		})
		// 去掉最后一个换行符
		contentStr = strings.TrimRight(contentStr, "\n")
		listen.ListenContent = contentStr

		switch listenType {
		case "brand":
			listen.ListenMp3Path, _ = element.DOM.Find("div[id='show_box'] > div[id='player_box'] > audio[id='myaudio']").Attr("src")
			listen.ListenSecondType = element.DOM.Find("p[class='infoNav'] > a").Last().Text()
			listen.ListenMediaPath = ""
			listen.ListenType = 5
		case "jiaoxue":
			listen.ListenMp3Path, _ = element.DOM.Find("div[id='show_box'] > div[id='player_box'] > audio[id='myaudio']").Attr("src")
			listen.ListenSecondType = element.DOM.Find("p[class='infoNav'] > a").Last().Text()
			listen.ListenMediaPath = ""
			listen.ListenType = 4
		case "chuji":
			listen.ListenMp3Path, _ = element.DOM.Find("div[id='show_box'] > div[id='player_box'] > audio[id='myaudio']").Attr("src")
			listen.ListenSecondType = element.DOM.Find("p[class='infoNav'] > a").Last().Text()
			listen.ListenMediaPath = ""
			listen.ListenType = 3
		case "media":
			listen.ListenMp3Path, _ = element.DOM.Find("div[id='show_box'] > div[id='player_box'] > audio[id='myaudio']").Attr("src")
			listen.ListenSecondType = element.DOM.Find("p[class='infoNav'] > a").Last().Text()
			listen.ListenMediaPath = ""
			listen.ListenType = 2
		case "17698":
			thunderUrl := element.DOM.Find("div[style='margin-bottom:4px;'] > script[type='text/javascript']").Text()
			/*  表达式 thunder_url\s*=\s*"([^"]+)"
			\s* 表示 0 个或多个空格
			[^"]+ 表示至少一个非双引号字符
			() 表示捕获组，[^"]+ 匹配到的内容会被保存到捕获组中
			*/
			if thunderUrl == "" {
				thunderUrl = element.ChildText("div[style='margin-bottom:4px;'] > script")
				match := re3.FindStringSubmatch(thunderUrl)
				listen.ListenMediaPath = "http://k6.kekenet.com/" + match[1]
			} else {
				match2 := re2.FindStringSubmatch(thunderUrl)
				listen.ListenMediaPath = "http://k6.kekenet.com/" + match2[1]
			}
			listen.ListenMp3Path = ""
			listen.ListenSecondType = ""
			listen.ListenType = 1
		}
		err := dao.CreateListen(listen)
		if err == nil {
			cnt1++
		}
		element.ForEach("table[class='wordList'] > tbody > tr", func(i int, et *colly.HTMLElement) {
			listenWord := new(model.ListenWord)
			if flag == true {
				listenWord.ListenId = listen.ListenId
				listenWord.WordId, _ = config.GenID()
				listenWord.DeleteIsOk = 0
				listenWord.WordNum = cnt
				listenWord.WordMusic, _ = et.DOM.Find("td > a[class='play']").Attr("data-url")
				listenWord.WordPhonetic = et.DOM.Find("td > span[class='py']").Text()
				listenWord.Word = et.DOM.Find("td > span > a[target='_blank']").Text()
				meaning := et.DOM.Find("td > div[class='explain'] > div[class='content'] > p").Text()
				listenWord.WordMeaning = strings.ReplaceAll(meaning, " ", "")
				err := dao.CreateListenWord(listenWord)
				if err == nil {
					cnt2++
				}
				cnt++
			}
			flag = true
		})
	})

	urlList := ctx.GetAny("urlList")

	for _, url := range urlList.([]string) {
		if url != "" {
			time.Sleep(120 * time.Millisecond)
			ct.Visit(url)
		}
	}
	return cnt1, cnt2
}

/*
ListenCronSpiderService
@author: LJR
@Description: 定时任务更新数据业务逻辑
@param ctx
@param listenType
@return cnt1
@return cnt2
*/
func ListenCronSpiderService(ctx *colly.Context, listenType string) (cnt1 int, cnt2 int) {
	layout := "2006-01-02 15:04:05"
	re1 := regexp.MustCompile(`编辑:(\w+)`)
	re2 := regexp.MustCompile(`thunder_url\s*=\s*"([^"]+)"`)
	cnt1 = 0
	cnt2 = 0

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

	ct.OnHTML("div[class='lastPage_left']", func(element *colly.HTMLElement) {
		cnt := 1
		flag := false
		var contentStr string
		listen := new(model.Listen)
		listen.ListenSource = "可可英语"
		listen.ListenId, _ = config.GenID()
		listen.ListenTitle = element.DOM.Find("div[class='e_title'] > h1[id='nrtitle']").Text()
		_, err := dao.GetListenByTitle(listen.ListenTitle)
		if err != nil {
			return
		}
		listen.DeleteIsOk = 0
		listen.ListenCollectNum = 0
		timeStr := element.DOM.Find("div[class='e_title'] > cite > time").Text()
		timeStr = strings.Replace(timeStr, "时间:", "", 1)
		listen.PublishAt, _ = time.Parse(layout, timeStr)

		editor := element.DOM.Find("div[class='e_title'] > cite").Text()
		match1 := re1.FindStringSubmatch(editor)
		listen.ListenEditor = match1[1]

		element.ForEach("div[class='info-qh'] > div", func(i int, ee *colly.HTMLElement) {
			contentStr += ee.Text + "\n"
		})
		// 去掉最后一个换行符
		contentStr = strings.TrimRight(contentStr, "\n")
		listen.ListenContent = contentStr

		switch listenType {
		case "brand":
			listen.ListenMp3Path, _ = element.DOM.Find("div[id='show_box'] > div[id='player_box'] > audio[id='myaudio']").Attr("src")
			listen.ListenSecondType = element.DOM.Find("p[class='infoNav'] > a").Last().Text()
			listen.ListenMediaPath = ""
			listen.ListenType = 4
		case "jiaoxue":
			listen.ListenMp3Path, _ = element.DOM.Find("div[id='show_box'] > div[id='player_box'] > audio[id='myaudio']").Attr("src")
			listen.ListenSecondType = element.DOM.Find("p[class='infoNav'] > a").Last().Text()
			listen.ListenMediaPath = ""
			listen.ListenType = 3
		case "chuji":
			listen.ListenMp3Path, _ = element.DOM.Find("div[id='show_box'] > div[id='player_box'] > audio[id='myaudio']").Attr("src")
			listen.ListenSecondType = element.DOM.Find("p[class='infoNav'] > a").Last().Text()
			listen.ListenMediaPath = ""
			listen.ListenType = 2
		case "media":
			listen.ListenMp3Path, _ = element.DOM.Find("div[id='show_box'] > div[id='player_box'] > audio[id='myaudio']").Attr("src")
			listen.ListenSecondType = element.DOM.Find("p[class='infoNav'] > a").Last().Text()
			listen.ListenMediaPath = ""
			listen.ListenType = 1
		case "17698":
			thunderUrl := element.DOM.Find("div[style='margin-bottom:4px;'] > script[type='text/javascript']").Text()
			/*  表达式 thunder_url\s*=\s*"([^"]+)"
			\s* 表示 0 个或多个空格
			[^"]+ 表示至少一个非双引号字符
			() 表示捕获组，[^"]+ 匹配到的内容会被保存到捕获组中
			*/
			match2 := re2.FindStringSubmatch(thunderUrl)
			listen.ListenMediaPath = "http://k6.kekenet.com/" + match2[1]
			listen.ListenMp3Path = ""
			listen.ListenSecondType = ""
			listen.ListenType = 0
		}
		err = dao.CreateListen(listen)
		if err == nil {
			cnt1++
		}
		element.ForEach("table[class='wordList'] > tbody > tr", func(i int, et *colly.HTMLElement) {
			listenWord := new(model.ListenWord)
			if flag == true {
				listenWord.ListenId = listen.ListenId
				listenWord.WordId, _ = config.GenID()
				listenWord.DeleteIsOk = 0
				listenWord.WordNum = cnt
				listenWord.WordMusic, _ = et.DOM.Find("td > a[class='play']").Attr("data-url")
				listenWord.WordPhonetic = et.DOM.Find("td > span[class='py']").Text()
				listenWord.Word = et.DOM.Find("td > span > a[target='_blank']").Text()
				meaning := et.DOM.Find("td > div[class='explain'] > div[class='content'] > p").Text()
				listenWord.WordMeaning = strings.ReplaceAll(meaning, " ", "")
				err := dao.CreateListenWord(listenWord)
				if err == nil {
					cnt2++
				}
				cnt++
			}
			flag = true
		})
	})

	// 改1
	urlList := ctx.GetAny("urlList")
	for _, url := range urlList.([]string) {
		if url != "" {
			time.Sleep(120 * time.Millisecond)
			ct.Visit(url)
		}
	}
	return cnt1, cnt2
}
