package main

import (
    "fmt"
    "math"
    "strings"
    "sort"
)

func findCommonPrefix(str1 string, str2 string) string {
	result := ""
	n := int(math.Min(float64(len(str1)), float64(len(str2))))
	for i := 0; i < n; i++ {
		if str1[i] != str2[i] {
			break
		}
		result += string(str1[i])
	} 
	return result
}

func commonPrefixPreprocess(arr []string) map[string][]string {
	dict := make(map[string][]string)
	var prefixes []string
	var i int
	prefix := arr[0]
	for i = 1; i < len(arr); i++ {
		prefix = findCommonPrefix(prefix, arr[i])
		prefixes = append(prefixes,prefix)
		prefixes = append(prefixes,arr[i])
	}
	for _, p := range prefixes {
		if len(p) > 0 {
			for i = 0; i < len(arr); i++ {
				if strings.HasPrefix(arr[i],p) {
					dict[p] = append(dict[p], arr[i])
				}
			}
		}
	}
	return dict
}

func findWithPrefix(dict map[string][]string, prefix string) []string {
	result, _ := dict[prefix]
	return result
}

func main() {
	stringSet := []string{"dog", "deer", "deal", "ea"}
	sort.Strings(stringSet)
	dict := commonPrefixPreprocess(stringSet)
	r := findWithPrefix(dict, "d")
	fmt.Println(dict)
	fmt.Println(r)
	r = findWithPrefix(dict, "de")
	fmt.Println(r)
}