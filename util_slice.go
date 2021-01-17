package gnet

import (
	"fmt"
	"strconv"
	"strings"
)

//将有序的数字切片转换为数字范围表示的字符串数组
func NumToRange(numList []int) []string {
	if len(numList) == 0 {
		return []string{}
	}
	var retList []string
	var start = numList[0]
	var end int
	var listEnd = numList[len(numList)-1]
	for i := 0; i < len(numList)-1; i++ {
		if numList[i+1] > numList[i]+1 {
			end = numList[i]
			if start == end {
				retList = append(retList, strconv.Itoa(start))
			} else {
				retList = append(retList, strconv.Itoa(start)+"-"+strconv.Itoa(end))
			}
			start = numList[i+1]
		}
	}
	if start == listEnd {
		retList = append(retList, strconv.Itoa(start))
	} else {
		retList = append(retList, strconv.Itoa(start)+"-"+strconv.Itoa(listEnd))
	}
	return retList
}


func NumToString(numList []int) string{
	return strings.Replace(strings.Trim(fmt.Sprint(numList), "[]"), " ", ",", -1)
}