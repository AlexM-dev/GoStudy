package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	f1, err := os.OpenFile("./in.txt", os.O_RDONLY, 0777)
	if err != nil {
		log.Fatal(err)
	}
	defer f1.Close()

	fileReader := bufio.NewReader(f1)

	strings := make([]string, 0)

	for {
		line, _, err := fileReader.ReadLine()
		if err != nil {
			break
		}
		if regexp.MustCompile(`(\d+\+\d+)`).MatchString(string(line)) {
			sub := regexp.MustCompile(`(\d+)`).FindAllStringSubmatch(string(line), 2)
			a, _ := strconv.Atoi(sub[0][1])
			b, _ := strconv.Atoi(sub[1][1])
			str := fmt.Sprintf("%d+%d=%d\n", a, b, a+b)
			strings = append(strings, str)
		}
		if regexp.MustCompile(`(\d+\-\d+)`).MatchString(string(line)) {
			sub := regexp.MustCompile(`(\d+)`).FindAllStringSubmatch(string(line), 2)
			a, _ := strconv.Atoi(sub[0][1])
			b, _ := strconv.Atoi(sub[1][1])
			str := fmt.Sprintf("%d+%d=%d\n", a, b, a-b)
			strings = append(strings, str)
		}
	}

	_ = os.Remove("./out.txt")
	f2, err := os.OpenFile("./out.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
	if err != nil {
		log.Fatal(err)
	}
	defer f2.Close()

	fileWriter := bufio.NewWriter(f2)

	for _, i := range strings {
		fileWriter.Write([]byte(i))
	}
	fileWriter.Flush()
}
