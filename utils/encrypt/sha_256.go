package encrypt

import (
	"crypto/sha256"
	"encoding/hex"
)

/*
GetSHA256HashCode
@author: LJR
@date: 2023-03-23 21:43:51
@Description: SHA256加密算法
@param stringMsg 传入需要加密的字符串
@return string 返回加密后的密文
*/
func GetSHA256HashCode(stringMsg string) string {
	message := []byte(stringMsg)
	// 创建一个基于SHA256算法的hash.Hash接口的对象
	hash := sha256.New()
	// 输入数据
	hash.Write(message)
	// 计算哈希值
	bytes := hash.Sum(nil)
	// 将字符串编码为16进制格式,返回字符串
	hashCode := hex.EncodeToString(bytes)
	// 返回哈希值
	return hashCode
}
