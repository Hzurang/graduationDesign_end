package utils

import "strings"

/*
RemoveTopStruct
@author: LJR
@Description: 去除结构体名称前缀 如 SignUpParam.Email 中的 SignUpParam
@param fields
@return map[string]string
*/
func RemoveTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}
