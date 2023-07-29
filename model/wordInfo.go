package model

import "time"

type WordTranslationInfo struct {
	Word              string              `json:"word"`
	PhoneticTrans1    string              `json:"phonetic_trans_1"`
	WordMp3Path1      string              `json:"word_mp_3_path_1"`
	PhoneticTrans2    string              `json:"phonetic_trans_2"`
	WordMp3Path2      string              `json:"word_mp_3_path_2"`
	WordMeaning       string              `json:"word_meaning"`
	WordForm          string              `json:"word_form"`
	NetworkDefinition []NetworkDefinition `json:"network_definition"`
	WordPhrase        []WordPhrase        `json:"word_phrase"`
	WordNearSynonym   []WordNearSynonym   `json:"word_near_synonym"`
	WordSentence      []WordSentence      `json:"word_sentence"`
}

type NetworkDefinition struct {
	Meaning  string `json:"meaning"`
	Sentence string `json:"sentence"`
}

type WordPhrase struct {
	Phrase  string `json:"phrase"`
	Meaning string `json:"meaning"`
}

type WordNearSynonym struct {
	Meaning string `json:"meaning"`
	English string `json:"english"`
}

type WordSentence struct {
	Sentence string `json:"sentence"`
	Mp3Path  string `json:"mp3Path"`
	Meaning  string `json:"meaning"`
}

type WordDetail struct {
	WordId           uint64 `json:"word_id"`
	Word             string `json:"word"`
	PhoneticTransEng string `json:"phonetic_trans_eng"`
	PhoneticTransAme string `json:"phonetic_trans_ame"`
	WordMeaning      string `json:"word_meaning"`
	MnemonicAid      string `json:"mnemonic_aid"`
	ChiEtymology     string `json:"chi_etymology"`
	SentenceEng1     string `json:"sentence_eng_1"`
	SentenceChi1     string `json:"sentence_chi_1"`
	SentenceEng2     string `json:"sentence_eng_2"`
	SentenceChi2     string `json:"sentence_chi_2"`
	SentenceEng3     string `json:"sentence_eng_3"`
	SentenceChi3     string `json:"sentence_chi_3"`
	WordType         int8   `json:"word_type"`
	WordMp3Path1     string `json:"word_mp_3_path_1"`
	WordMp3Path2     string `json:"word_mp_3_path_2"`
	SentenceMp3Path1 string `json:"sentence_mp_3_path_1"`
	SentenceMp3Path2 string `json:"sentence_mp_3_path_2"`
	SentenceMp3Path3 string `json:"sentence_mp_3_path_3"`
}

type WordList struct {
	WordId      uint64    `json:"word_id"`
	Word        string    `json:"word"`
	WordMeaning string    `json:"word_meaning"`
	WordType    int8      `json:"word_type"`
	CreatedAt   time.Time `json:"created_at"`
}
