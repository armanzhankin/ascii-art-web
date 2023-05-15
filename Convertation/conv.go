package Convertation

import (
	"os"
	"strings"
)

func Valid(s string) bool {
	runes := []rune(s)
	for _, ch := range runes {
		if !(ch <= 126) {
			return false
		}
	}
	return true
}

func Convert(text []string, typeOfFormat string) (string, error) {
	res := ""
	font, err := os.ReadFile(string(typeOfFormat + ".txt"))
	if err != nil {
		return "", err
	}
	cont := strings.ReplaceAll(string(font), "\r", "")
	mass := strings.Split(string(cont), "\n\n")

	for _, t := range text {
		if t == "" {
			res += "\n"
			continue
		} else {
			for col := 0; col < 8; col++ {
				for _, v := range t {
					res += strings.Split(mass[v-32], "\n")[col]
				}
				res += "\n"
			}
		}
	}
	return res, nil
}

func SplitS(s string) []string {
	var result []string
	s = strings.ReplaceAll(s, "\\n", "\n")
	s = strings.ReplaceAll(s, string(rune(13)), "")
	result = strings.Split(s, "\n")
	for i := 0; i < len(result); i++ {
		if len(s)+1 == len(result) && result[i] == "" {
			temp := result[i+1:]
			result = result[:i]
			result = append(result, temp...)
		}
	}
	return result
}
