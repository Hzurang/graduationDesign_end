package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"ginStudy/utils"
	"github.com/jordan-wright/email"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"net/smtp"
	"net/url"
	"strings"
	"time"
)

type APISMSResultBody struct {
	Status    string `json:"status"`
	Reason    string `json:"reason"`
	RequestID string `json:"request_id"`
}

const (
	ALIYUN_THIRD_API_SMSCODE_URL = "https://dfsmsv2.market.alicloudapi.com/data/send_sms_v2"
)

/*
SMSCode
@author: LJR
@Description: 短信发送的流程
@param mobile
@return string
@return error
*/
func SMSCode(mobile string) (string, error) {
	code := utils.GetRoundNumber(6)
	v := url.Values{}
	v.Set("content", fmt.Sprintf("code:%s,expire_at:2", code))
	v.Set("phone_number", mobile)
	v.Set("template_id", "TPL_0001")
	body := strings.NewReader(v.Encode())
	client := &http.Client{}
	req, err := http.NewRequest("POST", ALIYUN_THIRD_API_SMSCODE_URL, body)
	req.Header.Add("Authorization", "APPCODE"+" "+Config.APISMSCodeConfig.APPCode)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	resp, err := client.Do(req)
	if err != nil {
		zap.L().Error(utils.GetCodeMsg(10013))
		return "", errors.New(utils.GetCodeMsg(30006))
	}
	defer resp.Body.Close()
	resCode := resp.StatusCode
	switch resCode {
	case 200:
		data, _ := ioutil.ReadAll(resp.Body)
		jsonStr := string(data)
		bytes := []byte(jsonStr)
		testT := new(APISMSResultBody)
		_ = json.Unmarshal(bytes, testT)
		if testT.Status == "OK" {
			return code, nil
		}
	case 403:
		zap.L().Error("触发限发机制 / 套餐余额用完")
		return "", errors.New(utils.GetCodeMsg(30006))
	}
	return "", errors.New(utils.GetCodeMsg(30006))
}

/*
EmailCode
@author: LJR
@Description: 邮箱发送的流程
@param em
@return string
@return error
*/
func EmailCode(em string) (string, error) {
	code := utils.GetRoundNumber(6)
	mailUserName := Config.QQEmailConfig.UserName
	mailSecret := Config.QQEmailConfig.Secret
	e := email.NewEmail()
	e.From = fmt.Sprintf("乐学英语APP <%s>", mailUserName)
	e.To = []string{em}
	e.Subject = "乐学英语APP邮箱验证信息"
	t := time.Now().Format("2006-01-02 15:04:05")
	//设置文件发送的内容
	content := fmt.Sprintf(`
	<div>
		<div>
			尊敬的%s，您好！
		</div>
		<div style="padding: 8px 40px 8px 50px;">
			<p>您于 %s 提交的邮箱验证，本次验证码为<u><strong>%s</strong></u>，为了保证账号安全，验证码有效期为2分钟。请确认为本人操作，切勿向他人泄露，感谢您的理解与使用。</p>
		</div>
		<div>
			<p>此邮箱为系统邮箱，请勿回复。</p>
		</div>
	</div>
	`, em, t, code)
	e.HTML = []byte(content)
	//设置服务器相关的配置
	err := e.Send("smtp.qq.com:25", smtp.PlainAuth("", mailUserName, mailSecret, "smtp.qq.com"))
	if err != nil {
		return "", errors.New(utils.GetCodeMsg(30003))
	}
	return code, nil
}
