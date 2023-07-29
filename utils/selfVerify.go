package utils

import (
	"fmt"
	"ginStudy/model"
	"github.com/go-playground/validator/v10"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
)

/*
SignUpParamStructLevelValidation
@author: LJR
@Description: 自定义 ParamMobilePasswordSignUp 结构体校验函数
@param sl
*/
func SignUpParamStructLevelValidation(sl validator.StructLevel) {
	su := sl.Current().Interface().(model.ParamMobilePasswordSignUp)
	if su.Password != su.RePassword {
		// 输出错误提示信息，最后一个参数就是传递的param
		sl.ReportError(su.RePassword, "re_password", "RePassword", "eqfield", "password")
	}
}

/*
VerifyMobileParam
@author: LJR
@Description: 校验手机号码
@param fl
@return bool
*/
func VerifyMobileParam(fl validator.FieldLevel) bool {
	ok, _ := regexp.MatchString(`^(13|14|15|17|18|19)[0-9]{9}$`, fl.Field().String())
	return ok
}

/*
VerifyIDCardParam
@author: LJR
@Description: 校验身份证号码
@param fl
@return bool
*/
func VerifyIDCardParam(fl validator.FieldLevel) bool {
	id := fl.Field().String()

	var a1Map = map[int]int{
		0:  1,
		1:  0,
		2:  10,
		3:  9,
		4:  8,
		5:  7,
		6:  6,
		7:  5,
		8:  4,
		9:  3,
		10: 2,
	}

	var idStr = strings.ToUpper(string(id))
	var reg, err = regexp.Compile(`^[0-9]{17}[0-9X]$`)
	if err != nil {
		return false
	}
	if !reg.Match([]byte(idStr)) {
		return false
	}
	var sum int
	var signChar = ""
	for index, c := range idStr {
		var i = 18 - index
		if i != 1 {
			if v, err := strconv.Atoi(string(c)); err == nil {
				var weight = int(math.Pow(2, float64(i-1))) % 11
				sum += v * weight
			} else {
				return false
			}
		} else {
			signChar = string(c)
		}
	}
	var a1 = a1Map[sum%11]
	var a1Str = fmt.Sprintf("%d", a1)
	if a1 == 10 {
		a1Str = "X"
	}
	return a1Str == signChar
}

/*
VerifyDate
@author: LJR
@Description: 校验时间
@param fl
@return bool
*/
func VerifyDate(fl validator.FieldLevel) bool {
	date, err := time.Parse("2006-01-02", fl.Field().String())
	if err != nil {
		return false
	}
	if date.Before(time.Now()) {
		return false
	}
	return true
}
