package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(comma("1234123456789"))
	fmt.Println(commarec("1234123456789"))
}

func comma(s string) string {
	var buf bytes.Buffer

	for i := 0; i < len(s); i++ {
		buf.WriteByte(s[i])
		if (len(s)-i-1)%3 == 0 && i < len(s)-1 {
			buf.WriteRune(',')
		}
	}

	return buf.String()
}

func commarec(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return commarec(s[:n-3]) + "," + s[n-3:]
}
