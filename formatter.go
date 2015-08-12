package main

import (
	"fmt"
	"strings"
)

const (
	successColor = "\x1b[32m"
	failColor    = "\x1b[31m"
	varColor     = "\x1b[33m"
	reset        = "\x1b[m"
)

// Formatter formats the messages, and optionally disabling colors. See
// FormatMessage for the structure of messages.
type Formatter struct {
	options *GlobalOptions
}

// Info uses no color.
func (f *Formatter) Info(messages ...string) string {
	return FormatMessage("", f.options.ShowColors, messages...)
}

// Success uses successColor (green) as color.
func (f *Formatter) Success(messages ...string) string {
	return FormatMessage(successColor, f.options.ShowColors, messages...)
}

// Fail uses failColor (red) as color.
func (f *Formatter) Fail(messages ...string) string {
	return FormatMessage(failColor, f.options.ShowColors, messages...)
}

// FormatMessage handles one or two messages. If more messages are used, those
// are ignore. If no messages are used, than it will return an empty string.
// 1 message : --> message[0]
// 2 messages: --> message[0]: message[1]
// color will be applied to the first message, varColor will be used for the
// second message. If useColors is false, than color will be ignored.
func FormatMessage(color string, useColors bool, messages ...string) string {
	segments := []string{}

	l := len(messages)

	if l > 0 {
		segments = append(segments, "-->")
	}

	if l >= 1 {
		if useColors {
			segments = append(segments, fmt.Sprintf(" %s%s%s", color, messages[0], reset))
		} else {
			segments = append(segments, fmt.Sprintf(" %s", messages[0]))
		}
	}

	if l >= 2 {
		if useColors {
			segments = append(segments, fmt.Sprintf(": %s%s%s", varColor, messages[1], reset))
		} else {
			segments = append(segments, fmt.Sprintf(": %s", messages[1]))
		}
	}

	if l > 2 {
		for _, m := range messages[2:] {
			segments = append(segments, fmt.Sprintf(" %s", m))
		}
	}

	return strings.Join(segments, "")
}
