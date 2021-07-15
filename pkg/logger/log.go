package logger

import (
	"fmt"

	"github.com/fatih/color"
)

func Info(input interface{}) {
	d := color.New(color.FgBlue)
	_, err := d.Printf("# ℹ %v\n", input)
	if err != nil {
		fmt.Println(err)
	}
}

func Good(input interface{}) {
	d := color.New(color.FgGreen)
	_, err := d.Printf("# ✓ %v\n", input)
	if err != nil {
		fmt.Println(err)
	}
}

func Bad(input interface{}) {
	d := color.New(color.FgRed)
	_, err := d.Printf("# ❌ %v\n", input)
	if err != nil {
		fmt.Println(err)
	}
}

func Warn(input interface{}) {
	d := color.New(color.FgYellow)
	_, err := d.Printf("# ⚠ %v\n", input)
	if err != nil {
		fmt.Println(err)
	}
}

func Raw(input interface{}) {
	fmt.Println(input)
}
