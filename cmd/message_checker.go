package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const (
	minTitleLength = 5
)

type PackageChecker struct {
	filePath            string
	bodyPatternIncluded bool
}

func (checker *PackageChecker) eval() bool {
	if !checker.evalEmail() {
		return false
	}

	file, err := os.Open(checker.filePath)
	if err != nil {
		println(err)
		return false
	}
	defer file.Close()

	if !checker.evalFile(file) {
		return false
	}

	if *bodyPattern != "" && !checker.bodyPatternIncluded {
		fmt.Printf("Text `%s` not included in body.\n", *bodyPattern)
		return false
	}

	return true
}

func (checker *PackageChecker) evalFile(file *os.File) bool {
	scanner := bufio.NewScanner(file)
	line := 0
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())

		if text != "" && text[0] == '#' {
			continue
		}

		if line == 0 && !checker.evalTitle(text) { // title
			return false
		}

		if line == 2 && text == "" {
			return false
		}

		if strings.Contains(text, *bodyPattern) {
			checker.bodyPatternIncluded = true
		}

		line++
	}

	return true
}

func (checker *PackageChecker) evalTitle(title string) bool {
	if title == "" {
		fmt.Printf("Title is empty\n")
		return false
	}

	if len(title) > *titleLength {
		fmt.Printf("Title is longer than %s\n", *titleLength)
		return false
	} else if len(title) < minTitleLength {
		fmt.Printf("Title is too short\n")
		return false
	}

	firstChar := string([]byte{title[0]})
	if *titleCapitalized && firstChar != strings.ToUpper(firstChar) {
		fmt.Printf("Title is not capitalized: %s\n", title)
		return false
	}

	return true
}

func (checker *PackageChecker) evalEmail() bool {
	if *emailDomain == "" {
		return true
	}

	cmd := exec.Command("git", "config", "user.email")
	data, err := cmd.Output()
	if err != nil {
		fmt.Printf("Error getting user email: %s\n", err)
		return false
	}

	email := strings.TrimSpace(string(data))
	if strings.Contains(email, *emailDomain) {
		return true
	}

	fmt.Printf("Email `%s` doesn't contain `%s`\n", email, *emailDomain)
	return false
}
