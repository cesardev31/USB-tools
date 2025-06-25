package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Prompt(msg string) string {
	fmt.Printf("%s: ", msg)
	return ReadLine()
}

func ReadLine() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func ExpandHome(path string) string {
	if strings.HasPrefix(path, "~") {
		home, err := os.UserHomeDir()
		if err != nil {
			return path
		}
		return strings.Replace(path, "~", home, 1)
	}
	return path
}

func ParseChoice(input string, max int) int {
	var idx int
	_, err := fmt.Sscanf(input, "%d", &idx)
	if err != nil || idx < 1 || idx > max {
		return -1
	}
	return idx - 1
}
