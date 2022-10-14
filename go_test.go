package main
 
import (
	"fmt"
    "sort"
    "strings"
	"unicode"
	"strconv"
)

const sIn string = "aCaaCbabbbCbccZZЙЙ"
 
func main() {
	if IsLetter(sIn) {
		fmt.Println(SortString(sIn))
		fmt.Println(Second(SortString(sIn)))
	} else {
		fmt.Println("ТОЛЬКО БУКВЫ!")
	}
	
}

func SortString(oStr string) string {
	s := strings.Split(oStr, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func Second(sSort string) string {
	var res string
	var tempS string
	oDict := make(map[string] int)
	var dTemp int

	for pos, char := range sSort {
		var sChar string = string(char)
		if pos == 0 {
			oDict[sChar] = pos +1
			dTemp += 1
			tempS = sChar
		} else {
			if tempS != sChar {
				dTemp = 1
				oDict[sChar] = dTemp
				tempS = sChar
			} else {
				dTemp += 1
				oDict[sChar] = dTemp
			}
		}
	}
	fmt.Println(oDict)
	for k, d := range oDict {
		res += k + strconv.Itoa(d)
		
	}
	return res

}

func IsLetter(s string) bool {
    for _, r := range s {
        if !unicode.IsLetter(r) {
            return false
        }
    }
    return true
}