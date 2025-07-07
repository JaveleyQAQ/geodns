package input

import (
	"bufio"
	"os"
	"strings"
)

type InputProcessor struct{}

func NewInputProcessor() *InputProcessor {
	return &InputProcessor{}
}

func (ip *InputProcessor) ReadFromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			lines = append(lines, line)
		}
	}
	return lines, scanner.Err()
}

func (ip *InputProcessor) ReadFromStdin() ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			lines = append(lines, line)
		}
	}
	return lines, scanner.Err()
}

func (ip *InputProcessor) ParseCommaSeparated(input string) []string {
	parts := strings.Split(input, ",")
	var result []string
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part != "" {
			result = append(result, part)
		}
	}
	return result
}

func (ip *InputProcessor) GetDomains(input string) ([]string, error) {
	if input == "-" {
		return ip.ReadFromStdin()
	} else if strings.Contains(input, ",") {
		return ip.ParseCommaSeparated(input), nil
	} else if _, err := os.Stat(input); err == nil {
		return ip.ReadFromFile(input)
	} else {
		return []string{input}, nil
	}
}
