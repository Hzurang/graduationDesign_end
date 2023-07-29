package convert

import (
	"bytes"
	"github.com/axgle/mahonia"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
)

/*
GbkToUtf8
@author: LJR
@date: 2023-03-23 22:40:28
@Description: 将 Gbk 转成 Utf8
@param GBKStr
@return UTF8Str
@return err
*/
func GbkToUtf8(GBKStr []byte) (UTF8Str []byte, err error) {
	r := transform.NewReader(bytes.NewReader(GBKStr), simplifiedchinese.GBK.NewDecoder())
	UTF8Str, err = io.ReadAll(r)
	if err != nil {
		return
	}
	return
}

/*
Utf8ToGbk
@author: LJR
@date: 2023-03-23 22:41:03
@Description: 将 Utf8 转成 Gbk
@param UTF8Str
@return GBKStr
@return err
*/
func Utf8ToGbk(UTF8Str []byte) (GBKStr []byte, err error) {
	r := transform.NewReader(bytes.NewReader(UTF8Str), simplifiedchinese.GBK.NewEncoder())
	GBKStr, err = io.ReadAll(r)
	if err != nil {
		return
	}
	return
}

func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}
