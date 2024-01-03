package utils

import "strings"

func GetAncestors(path string) []string{
	 segments := strings.Split(path, "/")
	 segments = segments[0 : len(segments)-2]
	 med := ""
	 result := make([]string, 10, 10)
	 for _, value := range segments {
		 med = med + "/" + value
		 result = append(result, med)
	 }
	 return result
}
