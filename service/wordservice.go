package service

import (
	"fmt"
	"ginStudy/config"
	"ginStudy/dao"
	"ginStudy/model"
	"ginStudy/utils"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
	"html"
	"regexp"
	"strconv"
	"strings"
	"time"
)

/*
GetWordMeaningService
@author: LJR
@Description: 用户获取单词信息API业务逻辑
@param word
@return wordTranslationInfo
@return err
*/
func GetWordMeaningService(word string) (wordTranslationInfo *model.WordTranslationInfo, err error) {
	var meaning string
	wordTranslationInfo = new(model.WordTranslationInfo)
	wordTranslationInfo.NetworkDefinition = make([]model.NetworkDefinition, 0, 6)
	wordTranslationInfo.WordPhrase = make([]model.WordPhrase, 0, 25)
	wordTranslationInfo.WordNearSynonym = make([]model.WordNearSynonym, 0, 5)
	wordTranslationInfo.WordSentence = make([]model.WordSentence, 0, 5)
	re1 := regexp.MustCompile(`\s+`)
	re2 := regexp.MustCompile(`,{1,}`)
	re3 := regexp.MustCompile(`([\p{Han}]+)|([\p{Latin}\s]+)`)

	ct := colly.NewCollector(
		colly.Debugger(&debug.LogDebugger{}), // 开启debug
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36"),
		colly.AllowURLRevisit(),
	)
	ct.OnRequest(func(r *colly.Request) {

	})

	ct.OnError(func(_ *colly.Response, errResponse error) {
		err = errResponse
		fmt.Println("请求期间发生错误,则调用:", errResponse)
	})

	ct.OnResponse(func(r *colly.Response) {
		ht := string(r.Body)
		ht = html.UnescapeString(ht)
		r.Body = []byte(ht)
		fmt.Println("收到响应后调用:", r.Request.URL)
	})

	ct.OnHTML("div[id='container']", func(element *colly.HTMLElement) {
		cnt := 0
		wordTranslationInfo.Word = word
		element.ForEach("h2[class='wordbook-js'] > div[class='baav'] > span[class='pronounce']", func(i int, ee *colly.HTMLElement) {
			if cnt == 0 {
				wordTranslationInfo.PhoneticTrans1 = ee.DOM.Find("span[class='phonetic']").Text()
				wordTranslationInfo.WordMp3Path1 = fmt.Sprintf("https://dict.youdao.com/dictvoice?audio=%s&type=1", word)
			}
			if cnt == 1 {
				wordTranslationInfo.PhoneticTrans2 = ee.DOM.Find("span[class='phonetic']").Text()
				wordTranslationInfo.WordMp3Path2 = fmt.Sprintf("https://dict.youdao.com/dictvoice?audio=%s&type=2", word)
			}
			cnt++
		})

		element.ForEach("div[id='phrsListTab'] > div[class='trans-container'] > ul > li", func(i int, ee1 *colly.HTMLElement) {
			meaning = meaning + ee1.Text + "\n"
		})
		wordTranslationInfo.WordMeaning = strings.TrimSuffix(meaning, "\n")
		wordTranslationInfo.WordForm = element.DOM.Find("div[id='phrsListTab'] > div[class='trans-container'] > p[class='additional']").Text()
		wordTranslationInfo.WordForm = re1.ReplaceAllString(wordTranslationInfo.WordForm, " ")

		element.ForEach("div[id='webTransToggle'] > div[id='tWebTrans'] > div", func(i int, ee2 *colly.HTMLElement) {
			if ee2.Attr("class") == "wt-container" || ee2.Attr("class") == "wt-container wt-collapse" {
				var netWorkDefinition model.NetworkDefinition
				netWorkDefinition.Meaning = ee2.DOM.Find("div[class='title'] > span").Text()
				netWorkDefinition.Meaning = strings.TrimSpace(netWorkDefinition.Meaning)

				str, _ := ee2.DOM.Find("p[class='collapse-content']").Html()
				str = strings.TrimSpace(str)
				if strings.Contains(str, "[gap=496]") {
					// 去掉[gap=496]
					str = strings.ReplaceAll(str, "[gap=496]", "")
					// 将关键词后的】改成【
					str = strings.ReplaceAll(str, "关键词】", "关键词【")
					str = strings.ReplaceAll(str, "Key words】", "】Key words:")
				}
				netWorkDefinition.Sentence = str
				wordTranslationInfo.NetworkDefinition = append(wordTranslationInfo.NetworkDefinition, netWorkDefinition)
			}
		})

		element.ForEach("div[id='transformToggle'] > div[id='wordGroup'] > p", func(i int, ee *colly.HTMLElement) {
			var wordPhrase model.WordPhrase
			wordPhrase.Phrase = ee.DOM.Find("span[class='contentTitle'] > a[class='search-js']").Text()
			wordPhrase.Meaning = ee.Text
			wordPhrase.Meaning = strings.ReplaceAll(wordPhrase.Meaning, wordPhrase.Phrase, "")
			wordPhrase.Meaning = strings.TrimSpace(wordPhrase.Meaning)
			wordTranslationInfo.WordPhrase = append(wordTranslationInfo.WordPhrase, wordPhrase)
		})

		element.ForEach("div[id='transformToggle'] > div[id='synonyms'] > ul > li", func(i int, ee *colly.HTMLElement) {
			var wordNearSynonym model.WordNearSynonym
			wordNearSynonym.Meaning = ee.DOM.Text()
			wordNearSynonym.English = re2.ReplaceAllString(ee.DOM.Parent().Find("p").Eq(i).Text(), ", ")
			wordNearSynonym.English = re1.ReplaceAllString(wordNearSynonym.English, "")
			wordTranslationInfo.WordNearSynonym = append(wordTranslationInfo.WordNearSynonym, wordNearSynonym)
		})
		element.ForEach("div[id='examplesToggle'] > div[id='bilingual'] > ul[class='ol'] > li", func(i int, ee *colly.HTMLElement) {
			var wordSentence model.WordSentence
			str := ee.DOM.Find("p[class!='example-via']").Text()
			str = re1.ReplaceAllString(str, " ")
			matches := re3.FindAllStringSubmatch(str, -1)
			var chinese, english string
			for _, match := range matches {
				if match[1] != "" {
					chinese += match[1]
				} else if match[2] != "" {
					english += match[2]
				}
			}
			wordSentence.Sentence = english
			wordSentence.Meaning = chinese
			mp3Url, _ := ee.DOM.Find("p > a[title='点击发音']").Attr("data-rel")
			wordSentence.Mp3Path = fmt.Sprintf("https://dict.youdao.com/dictvoice?audio=%s", mp3Url)
			wordTranslationInfo.WordSentence = append(wordTranslationInfo.WordSentence, wordSentence)
		})
	})

	err = ct.Visit(fmt.Sprintf("https://www.youdao.com/w/%s/", word))
	if err != nil {
		return nil, utils.GetError(40004)
	}
	return wordTranslationInfo, nil
}

/*
GetWordDetailService
@author: LJR
@Description: 用户记单词时获取单词详情业务逻辑
@param wordId
@return wordDetail
@return err
*/
func GetWordDetailService(wordId uint64) (wordDetail *model.WordDetail, err error) {
	word, err := dao.GetWordByWordId(wordId)
	if err != nil {
		return nil, err
	}
	wordDetail = new(model.WordDetail)
	wordDetail.WordId = word.WordId
	wordDetail.Word = word.Word
	wordDetail.PhoneticTransEng = word.PhoneticTransEng
	wordDetail.PhoneticTransAme = word.PhoneticTransAme
	wordDetail.WordMeaning = word.WordMeaning
	wordDetail.MnemonicAid = word.MnemonicAid
	wordDetail.ChiEtymology = word.ChiEtymology
	wordDetail.SentenceEng1 = word.SentenceEng1
	wordDetail.SentenceChi1 = word.SentenceChi1
	wordDetail.SentenceEng2 = word.SentenceEng2
	wordDetail.SentenceChi2 = word.SentenceChi2
	wordDetail.SentenceEng3 = word.SentenceEng3
	wordDetail.SentenceChi3 = word.SentenceChi3
	wordDetail.WordType = word.WordType
	wordDetail.WordMp3Path1 = fmt.Sprintf("https://dict.youdao.com/dictvoice?audio=%s&type=1", wordDetail.Word)
	wordDetail.WordMp3Path2 = fmt.Sprintf("https://dict.youdao.com/dictvoice?audio=%s&type=2", wordDetail.Word)
	wordDetail.SentenceMp3Path1 = strings.ReplaceAll(wordDetail.SentenceEng1, " ", "+")
	wordDetail.SentenceMp3Path2 = strings.ReplaceAll(wordDetail.SentenceEng2, " ", "+")
	wordDetail.SentenceMp3Path3 = strings.ReplaceAll(wordDetail.SentenceEng3, " ", "+")
	return wordDetail, nil
}

/*
CollectWordService
@author: LJR
@Description: 用户收藏单词业务逻辑
@param wordId
@param userId
@return err
*/
func CollectWordService(wordId uint64, userId uint64) (err error) {
	_, err = dao.GetWordByWordId(wordId)
	if err != nil {
		return utils.GetError(40010)
	}
	wordCollect, err := dao.GetSoftDeleteCollectWordByWordIdAndUserId(wordId, userId)
	if err == nil {
		// 软删除更新
		if wordCollect.DeleteIsOk == 1 {
			err = dao.UpdateSoftDeleteCollectWord(wordCollect)
			if err != nil {
				return err
			}
			return nil
		} else if wordCollect.DeleteIsOk == 0 {
			return utils.GetError(40012)
		}
	}
	wordCollect = &model.WordCollect{
		WordId: wordId,
		UserId: userId,
	}
	err = dao.CreateCollectWord(wordCollect)
	if err != nil {
		return err
	}
	return nil
}

/*
CancelCollectWordService
@author: LJR
@Description: 用户取消单词收藏业务逻辑
@param wordId
@param userId
@return err
*/
func CancelCollectWordService(wordId uint64, userId uint64) (err error) {
	_, err = dao.GetWordByWordId(wordId)
	if err != nil {
		return utils.GetError(40010)
	}
	err = dao.DeleteCollectWordByWordIdAndUserId(wordId, userId)
	if err != nil {
		return err
	}
	return nil
}

/*
GetCollectWordService
@author: LJR
@Description: 用户获取单词收藏列表业务逻辑
@param userId
@return wordCollectInfoList
@return err
*/
func GetCollectWordService(userId uint64) (wordCollectInfoList []model.WordCollectInfo, err error) {
	wordList, err := dao.GetWordCollectListByUserId(userId)
	if err != nil {
		return nil, err
	}
	wordCollectInfoList = make([]model.WordCollectInfo, 0, len(wordList))
	for _, word := range wordList {
		wordCollectInfo := model.WordCollectInfo{
			WordId:      word.WordId,
			Word:        word.Word,
			WordMeaning: word.WordMeaning,
			WordType:    word.WordType,
		}
		wordCollectInfoList = append(wordCollectInfoList, wordCollectInfo)
	}
	return wordCollectInfoList, nil
}

/*
InsertWordService
@author: LJR
@Description: 手动添加单词业务逻辑
@param p
@return err
*/
func InsertWordService(p *model.ParamInsertWord) (err error) {
	_, err = dao.GetWordByWordAndWordType(p.Word, p.WordType)
	if err != nil {
		return err
	}
	wordTranslationInfo, err := GetWordMeaningService(p.Word)
	wordId, _ := config.GenID()
	//wordType, _ := strconv.ParseInt(p.WordType, 10, 64)
	word := &model.Word{
		CreatedAt:        time.Now(),
		WordId:           wordId,
		Word:             p.Word,
		PhoneticTransEng: wordTranslationInfo.PhoneticTrans1,
		PhoneticTransAme: wordTranslationInfo.PhoneticTrans2,
		WordMeaning:      wordTranslationInfo.WordMeaning,
		MnemonicAid:      p.MnemonicAid,
		ChiEtymology:     p.ChiEtymology,
		SentenceEng1:     wordTranslationInfo.WordSentence[0].Sentence,
		SentenceChi1:     wordTranslationInfo.WordSentence[0].Meaning,
		SentenceEng2:     wordTranslationInfo.WordSentence[1].Sentence,
		SentenceChi2:     wordTranslationInfo.WordSentence[1].Meaning,
		SentenceEng3:     wordTranslationInfo.WordSentence[2].Sentence,
		SentenceChi3:     wordTranslationInfo.WordSentence[2].Meaning,
		WordType:         p.WordType,
		DeleteIsOk:       0,
	}
	err = dao.CreateWord(word)
	if err != nil {
		return err
	}
	return nil
}

/*
UpdateWordService
@author: LJR
@Description: 修改具体单词信息业务逻辑
@param p
@return err
*/
func UpdateWordService(p *model.ParamWordInfo) (err error) {
	_, err = dao.GetWordByWordId(p.WordId)
	if err != nil {
		return err
	}
	word := &model.Word{
		WordId:           p.WordId,
		Word:             p.Word,
		PhoneticTransEng: p.PhoneticTransEng,
		PhoneticTransAme: p.PhoneticTransAme,
		WordMeaning:      p.WordMeaning,
		MnemonicAid:      p.MnemonicAid,
		ChiEtymology:     p.ChiEtymology,
		SentenceEng1:     p.SentenceEng1,
		SentenceChi1:     p.SentenceChi1,
		SentenceEng2:     p.SentenceEng2,
		SentenceChi2:     p.SentenceChi2,
		SentenceEng3:     p.SentenceEng3,
		SentenceChi3:     p.SentenceChi3,
		WordType:         p.WordType,
	}
	if err = dao.UpdateWordByWordId(word); err != nil {
		return err
	}
	return nil
}

/*
GetWordService
@author: LJR
@Description: 获取具体单词信息业务逻辑
@param id
@return wordInfo, err
*/
func GetWordService(id string) (wordInfo *model.WordInfo, err error) {
	num, _ := strconv.Atoi(id)
	wordId := uint64(num)
	word, err := dao.GetWordByWordId(wordId)
	if err != nil {
		return nil, err
	}
	wordInfo = new(model.WordInfo)
	wordInfo.WordId = word.WordId
	wordInfo.Word = word.Word
	wordInfo.PhoneticTransEng = word.PhoneticTransEng
	wordInfo.PhoneticTransAme = word.PhoneticTransAme
	wordInfo.WordMeaning = word.WordMeaning
	wordInfo.MnemonicAid = word.MnemonicAid
	wordInfo.ChiEtymology = word.ChiEtymology
	wordInfo.SentenceEng1 = word.SentenceEng1
	wordInfo.SentenceChi1 = word.SentenceChi1
	wordInfo.SentenceEng2 = word.SentenceEng2
	wordInfo.SentenceChi2 = word.SentenceChi2
	wordInfo.SentenceEng3 = word.SentenceEng3
	wordInfo.SentenceChi3 = word.SentenceChi3
	wordInfo.WordType = word.WordType
	return wordInfo, nil
}

/*
DeleteWordService
@author: LJR
@Description: 删除单词业务逻辑
@param id
@return err
*/
func DeleteWordService(id string) (err error) {
	num, _ := strconv.Atoi(id)
	wordId := uint64(num)
	err = dao.DeleteWordByWordId(wordId)
	if err != nil {
		return err
	}
	// 补充用户收藏单词取消，因为单词软删除了
	_ = dao.DeleteCollectWordByWordId(wordId)
	return nil
}

/*
DetermineVocabularyService
@author: LJR
@Description: 判断用户是否背完这本书 传 word_type 背完返回1，没有背完返回0
@param userId
@param word_type
@return isEnd, err
*/
func DetermineVocabularyService(userId uint64, word_type string) (isEnd int8, err error) {
	_, err = dao.GetUserByUserId(userId)
	if err != nil {
		return 2, err
	}
	var wordType int8
	switch word_type {
	case "CET4":
		wordType = 1
	case "CET6":
		wordType = 2
	case "TEM4":
		wordType = 3
	case "TEM8":
		wordType = 4
	case "KAOYAN":
		wordType = 5
	case "GRE":
		wordType = 6
	case "TOEFL":
		wordType = 7
	case "IELTS":
		wordType = 8
	}
	count, err := dao.GetWordLearnCountByUserIdAndWordType(userId, wordType)
	if err != nil {
		return 2, err
	}
	switch word_type {
	case "CET4":
		if count == 1 {
			return 1, nil
		}
	case "CET6":
		if count == 1 {
			return 1, nil
		}
	case "TEM4":
		if count == 1 {
			return 1, nil
		}
	case "TEM8":
		if count == 1 {
			return 1, nil
		}
	case "KAOYAN":
		if count == 1 {
			return 1, nil
		}
	case "GRE":
		if count == 1 {
			return 1, nil
		}
	case "TOEFL":
		if count == 1 {
			return 1, nil
		}
	case "IELTS":
		if count == 1 {
			return 1, nil
		}
	}
	return 0, nil
}

/*
GetWordAllService
@author: LJR
@Description: 根据单词类型和页数等加载单词列表
@param word_type
@param page_num
@param page_size
@return WordList, total, err
*/
func GetWordAllService(word_type string, page_num string, page_size string) (wordList []model.Word, total int64, err error) {
	pageNum, _ := strconv.Atoi(page_num)
	pageSize, _ := strconv.Atoi(page_size)
	var wordType int8
	switch word_type {
	case "CET4":
		wordType = 1
	case "CET6":
		wordType = 2
	case "TEM4":
		wordType = 3
	case "TEM8":
		wordType = 4
	case "KAOYAN":
		wordType = 5
	case "GRE":
		wordType = 6
	case "TOEFL":
		wordType = 7
	case "IELTS":
		wordType = 8
	}
	wordList, total, err = dao.GetWordPage(wordType, pageNum, pageSize)
	if err != nil {
		return nil, 0, err
	}
	return wordList, total, nil
}

/*
GetWordByWordService
@author: LJR
@Description: 输入单词进行查找业务逻辑
@param word
@return Word, total, err
*/
func GetWordByWordService(word string) (Word []*model.Word, total int, err error) {
	Word, err = dao.GetWordByWord(word)
	if err != nil {
		return nil, 0, err
	}
	total = len(Word)
	return Word, total, nil
}

func GetWordListService(wordType string) (wordIDlist []string, err error) {
	num, _ := strconv.Atoi(wordType)
	wordInt := int8(num)
	wordList, err := dao.GetWordByWordType(wordInt)
	wordIDlist = make([]string, 0, len(wordList))
	if err != nil {
		return nil, err
	}
	for _, word := range wordList {
		word := strconv.FormatUint(word.WordId, 10)
		wordIDlist = append(wordIDlist, word)
	}
	return wordIDlist, nil
}
