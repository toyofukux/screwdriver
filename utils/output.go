package utils

import "github.com/fatih/color"

var (
	warnColor  *color.Color
	errorColor *color.Color
)

func init() {
	warnColor = color.New(color.FgYellow)
	errorColor = color.New(color.FgRed)
}

// WarnOutput output message with color yellow
func WarnOutput(messages ...interface{}) {
	warnColor.Println(messages)
}

// WarnOutputf output formatted message with color yellow
func WarnOutputf(format string, messages ...interface{}) {
	warnColor.Printf(format+"\n", messages)
}

// ErrorOutput output message with color red
func ErrorOutput(messages ...interface{}) {
	errorColor.Println(messages)
}

// ErrorOutputf output formatted message with color red
func ErrorOutputf(format string, messages ...interface{}) {
	errorColor.Printf(format+"\n", messages)
}
