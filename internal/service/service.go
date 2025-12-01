package service

import (
	"fmt"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func isMorse(text string) bool {

	runes := []rune(text)

	switch runes[0] {
	case '-', '.':
		return true
	}

	return false
}

func ConverText(text string) (string, error) {

	text = strings.TrimSpace(text)

	if text == "" {
		return "", fmt.Errorf("empty value")
	}

	if isMorse(text) {
		return morse.ToText(text), nil
	}

	return morse.ToMorse(text), nil

}
