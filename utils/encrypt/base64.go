package encrypt

import "encoding/base64"

/*
Base64Encode
@author: LJR
@date: 2023-03-23 23:43:05
@Description: base64 编码
@param str
@return string
*/
func Base64Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

/*
Base64Decode
@author: LJR
@date: 2023-03-23 23:43:21
@Description: base64 解码
@param str
@return string
@return []byte
*/
func Base64Decode(str string) (string, []byte) {
	resBytes, _ := base64.StdEncoding.DecodeString(str)
	return string(resBytes), resBytes
}
