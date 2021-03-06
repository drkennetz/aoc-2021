package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// ToInt will case a given arg into an int type.
// Supported types are:
//    - string
func ToInt(arg interface{}) int {
	var val int
	switch arg.(type) {
	case string:
		var err error
		val, err = strconv.Atoi(arg.(string))
		if err != nil {
			panic("error converting string to int " + err.Error())
		}
	case uint8:
		val = int(arg.(uint8))
	case uint:
		val = int(arg.(uint))
	default:
		panic(fmt.Sprintf("unhandled type for int casting %T", arg))
	}
	return val
}

// Max finds the max between 2 numbers
func Max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

// Abs finds the absolute value between 2 numbers
func Abs(a, b int) int {
	if a - b < 0 {
		return (a-b)*-1
	} else {
		return a - b
	}
}
// ToString will case a given arg into an int type.
// Supported types are:
//    - int
//    - byte
//    - rune
func ToString(arg interface{}) string {
	var str string
	switch arg.(type) {
	case int:
		str = strconv.Itoa(arg.(int))
	case byte:
		b := arg.(byte)
		str = string(rune(b))
	case rune:
		str = string(arg.(rune))
	default:
		panic(fmt.Sprintf("unhandled type for string casting %T", arg))
	}
	return str
}

const (
	ASCIICodeCapA   = int('A') // 65
	ASCIICodeCapZ   = int('Z') // 65
	ASCIICodeLowerA = int('a') // 97
	ASCIICodeLowerZ = int('z') // 97
)

// ToASCIICode returns the ascii code of a given input
func ToASCIICode(arg interface{}) int {
	var asciiVal int
	switch arg.(type) {
	case string:
		str := arg.(string)
		if len(str) != 1 {
			panic("can only convert ascii Code for string of length 1")
		}
		asciiVal = int(str[0])
	case byte:
		asciiVal = int(arg.(byte))
	case rune:
		asciiVal = int(arg.(rune))
	}

	return asciiVal
}

// PopString pops the list char from a []string
func PopString(s *[]string) string {
	i := len(*s)
	pop := (*s)[i-1]
	*s=(*s)[:i-1]
	return pop
}

// ASCIIIntToChar returns a one character string of the given int
func ASCIIIntToChar(code int) string {
	return string(rune(code))
}

// FileReader
func FileReader(s string) []string {
	parsed := make([]string, 0)

	file, err := os.Open(s)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		parsed = append(parsed, line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
	return parsed
}