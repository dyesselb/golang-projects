package main

import (
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	if len(os.Args) == 3 {
		file, err := os.ReadFile(os.Args[1])
		if err != nil {
			log.Fatalln("ERROR, wrong input file")
		}
		file1 := Split(string(file))
		text := CheckModOfStrings(file1)
		text = FixRuneA(text)
		text = FixMArks(text, "'")
		text = FixMArks(text, "\"")
		text = FixMArks(text, "'")
		text = FixMArks(text, "\"")
		text = DeleteSpaces(text)
		if strings.HasSuffix(os.Args[2], ".txt") {
			final, err := os.Create(os.Args[2])
			if err != nil {
				log.Fatalln("ERROR, wrong name")
			}
			for i, name := range text {
				final.WriteString(name)
				if i != len(text)-1 && string(name[len(name)-1]) != "\n" {
					final.WriteString(" ")
				}

			}
		} else {
			log.Fatalln("ERROR, wrong format")
		}
	} else {
		log.Fatalln("ERROR, wrong number of arguments")
	}
}

func Split(s string) []string {
	var str string
	var result []string
	head := false
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			str += string(s[i])
			head = true
		} else if s[i] == ')' {
			str += string(s[i])
			head = false
		} else if IsPunct(rune(s[i])) && head == false {
			if str != "" {
				result = append(result, str)
			}
			result = append(result, string(s[i]))
			str = ""
		} else if s[i] != ' ' && s[i] != '\n' {
			str += string(s[i])
		} else {
			if head == false {
				if s[i] == '\n' {
					result = append(result, str)
					result = append(result, string(s[i]))
					str = ""
				} else if len(str) > 0 {
					result = append(result, str)
					str = ""
				}
			} else {
				str += string(s[i])
			}
		}
	}
	if len(str) > 0 {
		result = append(result, str)
	}
	return result
}

func CheckModOfStrings(s []string) []string {
	var res []string
	for _, name := range s {
		if len(name) > 5 {
			if name == "(cap)" || (name[:6] == "(cap, " && ValidName(name)) {
				StringChange(name, res, strings.Title)
			} else if name == "(low)" || (name[:6] == "(low, " && ValidName(name)) {
				StringChange(name, res, strings.ToLower)
			} else if name[:5] == "(up, " && ValidName(name) {
				StringChange(name, res, strings.ToUpper)
			} else if name == "(bin)" {
				IntChange(res, 2)
			} else if name == "(hex)" {
				IntChange(res, 16)
			} else {
				res = append(res, name)
			}
		} else if len(name) > 3 {
			if name == "(up)" {
				for i := len(res) - 1; i >= 0; i-- {
					if WrongStrings(res[i]) && res[i] != "\n" {
						res[i] = strings.ToUpper(res[i])
						break
					}
				}
			} else {
				res = append(res, name)
			}
		} else {
			res = append(res, name)
		}
	}
	return res
}

func FixRuneA(s []string) []string {
	var rev []string
	for i := 0; i < len(s); i++ {
		if i != len(s)-1 && s[i] == "a" {
			word := s[i+1]
			if IsVowel(rune(word[0])) {
				s[i] = "an"
				rev = append(rev, s[i])
			} else {
				rev = append(rev, s[i])
			}
		} else if i != len(s)-1 && s[i] == "A" {
			word := s[i+1]
			if IsVowel(rune(word[0])) {
				s[i] = "An"
				rev = append(rev, s[i])
			} else {
				rev = append(rev, s[i])
			}
		} else {
			rev = append(rev, s[i])
		}
	}
	return rev
}

func StringChange(Mod string, str []string, f func(s string) string) {
	var NumberFromMod []rune
	for _, word := range Mod {
		if word >= 48 && word <= 57 || word == '-' {
			NumberFromMod = append(NumberFromMod, word)
		}
	}
	NumberOfModifications, _ := strconv.Atoi(string(NumberFromMod))
	if NumberOfModifications == 0 {
		NumberOfModifications = 1
	}
	for i := len(str) - 1; i >= 0; i-- {
		if NumberOfModifications > 0 && WrongStrings(str[i]) {
			str[i] = f(str[i])
			NumberOfModifications = NumberOfModifications - 1
		}
	}
}

func IntChange(str []string, base int) {
	for i := len(str) - 1; i >= 0; i-- {
		if WrongStrings(str[i]) {
			num, err := strconv.ParseInt(str[i], base, 64)
			if err != nil {
				log.Fatalln("ERROR, wrong values")
			}
			str[i] = strconv.Itoa(int(num))
			break
		}
	}
}

func ValidName(name string) bool {
	if len(name) < 2 {
		return false
	}
	return name[len(name)-1] == ')' && (name[len(name)-2] >= '0' && name[len(name)-2] <= '9')
}

func IsVowel(word rune) bool {
	word = unicode.ToLower(word)
	a := "aeiouyh"
	for _, name := range a {
		if word == name {
			return true
		}
	}
	return false
}

func WrongStrings(str string) bool {
	WrongStrings := []string{".", ",", "!", "?", ":", ";", "...", "?!", "'", "\"", "\n", ""}
	for _, word := range WrongStrings {
		if word == str {
			return false
		}
	}
	return true
}

func FixMArks(str []string, mark string) []string {
	var s []string
	var rev []string
	Apostrophe := false

	for _, name := range str {
		if name != "" {
			s = append(s, name)
		}
	}
	for i := 0; i < len(s); i++ {
		if len(s) > 1 {
			if (i == 0 && s[i] == mark) || (s[i] == mark && Apostrophe == false) {
				if i < len(s)-1 {
					if WrongStrings(s[i+1]) {
						s[i+1] = s[i] + s[i+1]
						Apostrophe = true
					} else {
						rev = append(rev, s[i])
					}
				} else {
					if WrongStrings(rev[len(rev)-1]) {
						rev[len(rev)-1] += s[i]
					} else {
						rev = append(rev, s[i])
					}
				}
			} else if s[i] == mark && Apostrophe == true {
				if len(rev) != 0 && WrongStrings(rev[len(rev)-1]) {
					if WrongStrings(rev[len(rev)-1]) {
						rev[len(rev)-1] += s[i]
						Apostrophe = false
					} else {
						continue
					}
				} else {
					rev = append(rev, s[i])
				}
			} else {
				rev = append(rev, s[i])
			}
		} else {
			rev = append(rev, s[i])
		}
	}
	return rev
}

func DeleteSpaces(s []string) []string {
	var str []string
	for i, name := range s {
		if name[0] == '(' {
			word := ""
			for i := 0; i < len(name); i++ {
				if name[i] != ' ' {
					word += string(name[i])
				} else {
					if len(word) > 0 {
						str = append(str, word)
						word = ""
					} else {
						continue
					}
				}
			}
			if len(word) > 0 {
				str = append(str, word)
			}
		} else if name == "\n" && s[i-1] == "\n" {
			continue
		} else {
			str = append(str, name)
		}
	}
	var res []string
	for i := 0; i < len(str); i++ {
		if i != 0 && str[i] == "." || str[i] == "," || str[i] == ":" || str[i] == ";" || str[i] == "?" || str[i] == "!" || str[i] == "'" || str[i] == "\"" {
			res[len(res)-1] += str[i]
		} else if str[i] == "(" {
			str[i+1] = str[i] + str[i+1]
		} else if str[i] == ")" {
			res[len(res)-1] += str[i]
		} else if i == 0 && str[i] == "." || str[i] == "," || str[i] == ":" || str[i] == ";" || str[i] == "?" || str[i] == "!" {
			continue
		} else {
			res = append(res, str[i])
		}
	}
	return res
}

func IsPunct(word rune) bool {
	a := ".,:;!?\""
	for _, name := range a {
		if word == name {
			return true
		}
	}
	return false
}
