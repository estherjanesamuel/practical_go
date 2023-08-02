package lengthoflastword

import (
	"fmt"
	"strings"
)


func lengthOfLastWord(s string) (int, string){
	s = strings.Trim(s, " ")
	str := strings.Split(s, " ")
	fmt.Println(str[len(str)-1])
	return len(str[len(str)-1]), str[len(str)-1]
}