package model

import "time"

type ParamRefreshDb struct {
	UserName string `json:"user_name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ParamMobilePasswordLogin struct {
	Mobile   string `json:"mobile" binding:"required,mobile"`
	Password string `json:"password" binding:"required"`
}

type ParamMobilePasswordSignUp struct {
	Mobile         string `json:"mobile" binding:"required,mobile"`
	Code           string `json:"code" binding:"required"`
	Password       string `json:"password" binding:"required"`
	RePassword     string `json:"re_password" binding:"required,eqfield=Password"`
	InvitationCode string `json:"invitation_code"`
}

type ParamSMSCodeLogin struct {
	Mobile string `json:"mobile" binding:"required,mobile"`
	Code   string `json:"code" binding:"required"`
}

type ParamModifyPwdByOldPwd struct {
	Password    string `json:"password" binding:"required"`
	OldPassword string `json:"old_password" binding:"required"`
}

type ParamModifyPwdBySMSCode struct {
	Mobile   string `json:"mobile" binding:"required,mobile"`
	Code     string `json:"code" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ParamResetPwdBySMSCode struct {
	Mobile     string `json:"mobile" binding:"required,mobile"`
	Code       string `json:"code" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

type ParamBindEmail struct {
	Email string `json:"email" binding:"required,email"`
	Code  string `json:"code" binding:"required"`
}

type ParamInsertEssay struct {
	EssayTitle   string `json:"essay_title"`
	EssayAuthor  string `json:"essay_author"`
	EssayContent string `json:"essay_content"`
	EssayIsOk    int8   `json:"essay_isok"`
	EssayType    int8   `json:"essay_type"`
}

type ParamEssayInfo struct {
	EssayId      uint64 `json:"essay_id"`
	EssayTitle   string `json:"essay_title" binding:"required"`
	EssayAuthor  string `json:"essay_author"`
	PublishAt    string `json:"publish_at"`
	EssayContent string `json:"essay_content" binding:"required"`
	EssayIsOk    int8   `json:"essay_isok"`
	EssayType    int8   `json:"essay_type" binding:"required"`
}

type ParamUserInfo struct {
	UserId            uint64    `json:"user_id" binding:"required"`
	LeXueAppId        string    `json:"le_xue_app_id" binding:"required"`
	Gender            int8      `json:"gender" binding:"required"`
	School            string    `json:"school" binding:"required"`
	Birthday          time.Time `json:"birthday" binding:"required"`
	Area              string    `json:"area" binding:"required"`
	NickName          string    `json:"nickname" binding:"required"`
	HeadSculpture     string    `json:"head_sculpture" binding:"required"`
	Integral          uint64    `json:"integral" binding:"required"`
	WordNeedReciteNum int       `json:"word_need_recite_num"`
	LastStartTime     time.Time `json:"last_start_time"`
	EngLevel          int8      `json:"eng_level" binding:"required"`
	Role              int8      `json:"role" binding:"required"`
	InvitationCode    string    `json:"invitation_code" binding:"required"`
	Signature         string    `json:"signature" binding:"required"`
}

type ParamModifyUserInfo struct {
	LeXueAppId string    `json:"le_xue_app_id" binding:"required"`
	Gender     int8      `json:"gender" binding:"required"`
	School     string    `json:"school"`
	Birthday   time.Time `json:"birthday"`
	Area       string    `json:"area"`
	NickName   string    `json:"nickname" binding:"required"`
	Signature  string    `json:"signature"`
}

type ParamSentenceInfo struct {
	Errmsg      string `json:"errmsg"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	Note        string `json:"note"`
	Translation string `json:"translation"`
	Picture     string `json:"picture"`
	AudioPath   string `json:"tts"`
}

type ParamInsertListen struct {
	ListenTitle      string `json:"listen_title"`
	ListenSource     string `json:"listen_source"`
	ListenEditor     string `json:"listen_editor"`
	ListenContent    string `json:"listen_content"`
	ListenMediaPath  string `json:"listen_media_path"`
	ListenMp3Path    string `json:"listen_mp_3_path"`
	ListenType       int8   `json:"listen_type"`
	ListenSecondType string `json:"listen_second_type"`
}

type ParamListenInfo struct {
	ListenId         uint64 `json:"listen_id"`
	ListenTitle      string `json:"listen_title" binding:"required"`
	ListenSource     string `json:"listen_source" binding:"required"`
	ListenEditor     string `json:"listen_editor"`
	ListenContent    string `json:"listen_content" binding:"required"`
	ListenMediaPath  string `json:"listen_media_path"`
	ListenMp3Path    string `json:"listen_mp_3_path"`
	ListenType       int8   `json:"listen_type" binding:"required"`
	ListenSecondType string `json:"listen_second_type"`
}

type ParamInsertWord struct {
	Word         string `json:"word"`
	MnemonicAid  string `json:"mnemonic_aid"`
	ChiEtymology string `json:"chi_etymology"`
	WordType     int8   `json:"word_type"`
}

type ParamWordInfo struct {
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
}

type ParamUserDate struct {
	Date             time.Time `json:"date" binding:"required"`
	WordLearnNumber  int       `json:"word_learn_number" binding:"required"`
	WordReviewNumber int       `json:"word_review_number" binding:"required"`
	Remark           string    `json:"remark"`
}

type ParamMine struct {
	NickName string `json:"nickname"`
	Integral string `json:"integral"`
	Days     string `json:"days"`
	Words    string `json:"words"`
}

type ParamDate struct {
	Year             int    `json:"year"`
	Month            int    `json:"month"`
	Date             int    `json:"date"`
	WordLearnNumber  int    `json:"word_learn_number"`
	WordReviewNumber int    `json:"word_review_number"`
	Remark           string `json:"remark"`
}

type ParamUserEnglevel struct {
	WordNeedReciteNum int  `json:"word_need_recite_num"`
	EngLevel          int8 `json:"eng_level"`
}
