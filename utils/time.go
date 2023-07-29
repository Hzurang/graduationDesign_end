package utils

import "time"

/*
DateTimeToTimestamp
@author: LJR
@date: 2023-03-23 23:02:23
@Description: 日期时间字符串转时间戳（秒）
@param datetime
@return int64
*/
func DateTimeToTimestamp(datetime string) int64 {
	local, _ := time.LoadLocation("Local") //获取时区
	tmp, _ := time.ParseInLocation("2006-01-02 15:04:05", datetime, local)
	return tmp.Unix() //转化为时间戳 类型是int64

}

/*
DateToTimestamp
@author: LJR
@date: 2023-03-23 23:17:21
@Description: 纯日期字符串转时间戳（秒）
@param datetime
@return int64
*/
func DateToTimestamp(datetime string) int64 {
	local, _ := time.LoadLocation("Local") //获取时区
	tmp, _ := time.ParseInLocation("2006-01-02", datetime, local)
	return tmp.Unix() //转化为时间戳 类型是int64
}

/*
TimestampToDate
@author: LJR
@date: 2023-03-23 23:17:49
@Description: 时间戳(秒)转时间字符串
@param timestamp
@return string
*/
func TimestampToDate(timestamp int64) string {
	return time.Unix(timestamp, 0).Format("2006-01-02 15:04:05")
}

/*
FormatSecond
@author: LJR
@date: 2023-03-23 23:18:05
@Description: 秒转换为时分秒
@param seconds
@return day
@return hour
@return minute
@return second
*/
func FormatSecond(seconds int64) (day, hour, minute, second int64) {
	day = seconds / (24 * 3600)
	hour = (seconds - day*3600*24) / 3600
	minute = (seconds - day*24*3600 - hour*3600) / 60
	second = seconds - day*24*3600 - hour*3600 - minute*60
	return day, hour, minute, second
}
