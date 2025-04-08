package utils

import (
    "golang.org/x/text/cases"
    "golang.org/x/text/language"
)

func Capitalize(message string) string {
	return cases.Title(language.English, cases.Compact).String(message)
}