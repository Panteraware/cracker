package util

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var Printer = message.NewPrinter(language.English)
