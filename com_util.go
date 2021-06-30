/* Copyright 2021 CoderLee Inc. All Rights Reserved. */
/* com_util.go - 公共方法类 */
/* modification history
-----------------------
2021/6/30, by coderlee, 创建
*/
/*
DESCRIPTION
实现了常用的公共方法
*/

package BaiDu_Ai

import "sort"

/*
In - 判断元素是否在数组中
PARAMS:
	- target: 要匹配的元素
	- keys: 被检索的数组
RETURNS:
	- true: 在数组中
	- false: 不在数组中
*/
func In(target string, keys []string) bool {
	sort.Strings(keys)
	index := sort.SearchStrings(keys, target)
	if index < len(keys) && keys[index] == target {
		return true
	}
	return false
}

/*
GetKeys - 获取map的key集合
PARAMS:
	- m: 被处理的map
RETURNS:
	- keys: map的key集合
*/
func GetKeys(m map[string]string) []string {
	keys := make([]string, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
