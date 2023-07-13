package helper

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func InputName(input *os.File) string {
	if input == nil {
		input = os.Stdin
	}

	reader := bufio.NewReader(input)
	text, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	text = strings.Replace(text, "\n", "", -1)

	return text
}

func InputId(input *os.File) int {
	if input == nil {
		input = os.Stdin
	}

	reader := bufio.NewReader(input)
	text, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	text = strings.Replace(text, "\n", "", -1)

	id, _ := strconv.Atoi(text)

	return id
}

func InputMain(text string, input *os.File) string {
	if input == nil {
		input = os.Stdin
	}

	fmt.Print(text)
	data := ""
	_, err := fmt.Fscanln(input, &data)
	if err != nil {
		panic(err)
	}

	return data

}
