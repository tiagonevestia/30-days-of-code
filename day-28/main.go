package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	NTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	N := int32(NTemp)

	matched := regexp.MustCompile(`^[\w.+\-]+@gmail\.com$`)
	names := []string{}

	for NItr := 0; NItr < int(N); NItr++ {
		firstNameEmailID := strings.Split(readLine(reader), " ")

		firstName := firstNameEmailID[0]
		emailID := firstNameEmailID[1]

		if matched.MatchString(emailID) {
			names = append(names, firstName)
		}
	}

	sort.Strings(names)

	for _, v := range names {
		fmt.Println(v)
	}
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
