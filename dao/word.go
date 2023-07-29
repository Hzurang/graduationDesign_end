package dao

import (
	"errors"
	"ginStudy/global"
	"ginStudy/model"
	"ginStudy/utils"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func CreateWord(word *model.Word) (err error) {
	err = global.Db.Create(word).Error
	if err != nil {
		zap.L().Error("CreateWord", zap.Error(err))
		return utils.GetStrAndError("单词", 10010)
	}
	return
}

func GetWordByWordId(wordId uint64) (word *model.Word, err error) {
	word = &model.Word{WordId: wordId}
	if err := global.Db.Where("word_id = ?", word.WordId).First(word).Error; err != nil {
		zap.L().Error("GetWordByWordId", zap.Error(err))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.GetError(40009)
		}
		return nil, utils.GetStrAndError("单词", 10009)
	}
	return word, nil
}

func GetWordByWordAndWordType(wordEng string, wordType int8) (word *model.Word, err error) {
	word = &model.Word{Word: wordEng}
	if err := global.Db.Where("word = ? AND word_type = ?", wordEng, wordType).First(word).Error; err == nil {
		return word, utils.GetError(40015)
	}
	return nil, nil
}

func GetWordByWordType(wordType int8) (word []*model.Word, err error) {
	word = make([]*model.Word, 0, 50)
	err = global.Db.Where("word_type = ?", wordType).Find(&word).Error
	if err != nil {
		return nil, utils.GetError(40015)
	}
	return word, nil
}

func UpdateWordType(word *model.Word) (err error) {
	res := global.Db.Model(word).Where("word_type = ?", word.WordType).Updates(map[string]interface{}{
		"phonetic_trans_eng": word.PhoneticTransEng,
		"phonetic_trans_ame": word.PhoneticTransAme,
	})
	if res.Error != nil {
		zap.L().Error("UpdateWordType", zap.Error(err))
		return utils.GetError(40016)
	}
	return nil
}

func UpdateWordByWordId(word *model.Word) (err error) {
	res := global.Db.Model(word).Where("word_id = ?", word.WordId).Updates(map[string]interface{}{
		"word_id":            word.WordId,
		"word":               word.Word,
		"phonetic_trans_eng": word.PhoneticTransEng,
		"phonetic_trans_ame": word.PhoneticTransAme,
		"word_meaning":       word.WordMeaning,
		"mnemonic_aid":       word.MnemonicAid,
		"chi_etymology":      word.ChiEtymology,
		"sentence_eng_1":     word.SentenceEng1,
		"sentence_chi_1":     word.SentenceChi1,
		"sentence_eng_2":     word.SentenceEng2,
		"sentence_chi_2":     word.SentenceChi2,
		"sentence_eng_3":     word.SentenceEng3,
		"sentence_chi_3":     word.SentenceChi3,
		"word_type":          word.WordType,
	})
	if res.Error != nil {
		zap.L().Error("UpdateWordByWordId", zap.Error(err))
		return utils.GetError(40016)
	}
	return nil
}

func DeleteWordByWordId(wordId uint64) (err error) {
	word := &model.Word{WordId: wordId}
	global.Db.Model(word).Where("word_id = ?", wordId).Update("delete_isok", 1)
	num := global.Db.Where("word_id = ?", wordId).Delete(word).RowsAffected
	if num == 0 {
		zap.L().Error("DeleteWordByWordId", zap.Error(err))
		return utils.GetError(40017)
	}
	return nil
}

func GetWordPage(wordType int8, pageNum int, pageSize int) (wordList []model.Word, total int64, err error) {
	wordList = make([]model.Word, 0, 50)
	offset := (pageNum - 1) * pageSize
	if wordType != 0 {
		err = global.Db.Select("created_at, word_id, word, word_meaning, word_type").Limit(pageSize).Offset(offset).Where("word_type = ?", wordType).Order("created_at desc").Find(&wordList).Error
		global.Db.Model(&wordList).Where("word_type = ?", wordType).Count(&total)
		if err != nil {
			return nil, 0, utils.GetError(40021)
		}
		return wordList, total, nil
	}
	err = global.Db.Select("created_at, word_id, word, word_meaning, word_type").Limit(pageSize).Offset(offset).Order("created_at desc").Find(&wordList).Error
	global.Db.Model(&wordList).Count(&total)
	if err != nil {
		return nil, 0, utils.GetError(40021)
	}
	return wordList, total, nil
}

func GetWordByWord(word string) (Word []*model.Word, err error) {
	Word = make([]*model.Word, 0, 5)
	if err := global.Db.Where("word = ?", word).Find(&Word).Error; err != nil {
		zap.L().Error("GetWordByWord", zap.Error(err))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.GetError(40009)
		}
		return nil, utils.GetStrAndError("单词", 10009)
	}
	return Word, nil
}
