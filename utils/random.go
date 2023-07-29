package utils

import (
	"math/rand"
	"sync"
	"time"
)

var (
	randSeek = int64(1)
	l        sync.Mutex
)

// AlphanumericSet 字母数字集
var AlphanumericSet = []rune{
	'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
}

/*
GetRoundString
@author: LJR
@Description: 获取随机字符串（加锁版） 好处：短时间内生成的随机数不会重复  坏处：时间开销比没加锁版多些
@param size 具体需要几位
@return string 返回一个字符串
*/
func GetRoundString(size int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(getRandSeek()))
	for i := 0; i < size; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

/*
GetRoundStringNotLock
@author: LJR
@Description: 获取随机字符串（没加锁版） 好处：时间开销小  坏处：短时间内生成的随机数会重复
@param size 具体需要几位
@return string 返回一个字符串
*/
func GetRoundStringNotLock(size int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < size; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

/*
GetRoundNumber
@author: LJR
@Description: 获取随机字符串（加锁版） 好处：短时间内生成的随机数不会重复  坏处：时间开销比没加锁版多些
@param size 具体需要几位
@return string 返回一个字符串（纯数字）
*/
func GetRoundNumber(size int) string {
	str := "0123456789"
	bytes := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(getRandSeek()))
	for i := 0; i < size; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

/*
GetRoundNumberNotLock
@author: LJR
@Description: 获取随机字符串（没加锁版） 好处：时间开销小  坏处：短时间内生成的随机数会重复
@param size 具体需要几位
@return string 返回一个字符串（纯数字）
*/
func GetRoundNumberNotLock(size int) string {
	str := "0123456789"
	bytes := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < size; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func getRandSeek() int64 {
	l.Lock()
	if randSeek >= 100000000 {
		randSeek = 1
	}
	randSeek++
	l.Unlock()
	return time.Now().UnixNano() + randSeek
}

/*
RandomString
@author: LJR
@Description: 生成一段随机的字符串
@param n int 字符串长度
@return string
*/
func RandomString(n int) string {
	var letters = []byte("qwertyuioplkjhgfdsazxcvbnmQWERTYUIOPLKJHGFDSAZXCVBNM")
	result := make([]byte, n)
	rand.Seed(time.Now().Unix())
	// TODO 不断用随机字母填充字符串
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

/*
GetInvCodeByUID
@author: LJR
@date: 2023-04-06 21:49:02
@Description: 获取指定长度的邀请码
@param uid 用户ID
@param l
@return string
*/
func GetInvCodeByUID(uid uint64, l int) string {
	r := rand.New(rand.NewSource(int64(uid)))
	var code []rune
	for i := 0; i < l; i++ {
		idx := r.Intn(len(AlphanumericSet))
		code = append(code, AlphanumericSet[idx])
	}
	return string(code)
}
