package comutil

import "sort"

/*
	In
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
	GetKeys
*/
func GetKeys(m map[string]string) []string {
	keys := make([]string, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
