package encrypt

import (
	"crypto/md5"
	"encoding/hex"
)

/*
GetMd5String
@author: LJR
@date: 2023-03-23 23:19:49
@Description: 生成32位md5字符串(加密)
@param value
@param md5Key
@return string
*/
func GetMd5String(value, md5Key string) string {
	h := md5.New()
	_, _ = h.Write([]byte(value + md5Key))
	return hex.EncodeToString(h.Sum(nil))
}
