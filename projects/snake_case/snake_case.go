package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main()  {
	fmt.Println(ToSnakeCase("     Arief    Samuel"))
}

type SnakeCaseState string

const (
    Start   SnakeCaseState = "Start"
    Lower SnakeCaseState = "Lower"
    Upper SnakeCaseState = "Upper"
    NewWord SnakeCaseState = "NewWord"
)

func ToSnakeCase(s string) string{
	if s == "" {
		return s;
	}
	var sb strings.Builder
	state := Start

	for i, c := range s {
		if c == ' ' {
			if state != Start {
				state = NewWord
			}
		} else if unicode.IsUpper(c) {
			switch state {
			case Upper:
				hasNext := (i + 1 < len(s))
				if i > 0 && hasNext {
					nextChar := s[i+1]
					if unicode.IsUpper(rune(nextChar)) && nextChar != '_' {
						sb.WriteRune('_')
					}
				}
			case Lower:
				sb.WriteRune('_')
			case NewWord:
				sb.WriteRune('_')
			}
			str := unicode.ToLower(c)
			sb.WriteRune(str)
			state = Upper 
		} else if c == '_' {
			sb.WriteRune('_')
			state = Start
		} else {
			if state == NewWord {
				sb.WriteRune('_')
			}
			sb.WriteRune(c)
			state = Lower
		}
	}
	return sb.String()
}