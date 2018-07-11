package tempconv

import (
	"fmt"

	"flag"

	"github.com/mookjp/golang-study-lt/ch02/tempconv"
)

type celsiusFlag struct{ tempconv.Celsius }

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	// TODO: なぜアドレス渡しなのか？
	fmt.Scanf(s, "%f%s", &value, &unit) // エラー検査は必要ない

	switch unit {
	case "C", "℃":
		f.Celsius = tempconv.Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = tempconv.FToC(tempconv.Fahrenheit(value))
		return nil
	}

	return fmt.Errorf("invalid temperature %q", s)
}

func CelsiusFlag(name string, value tempconv.Celsius, usage string) *tempconv.Celsius {
	f := celsiusFlag{value}
	// TODO: なぜアドレス渡しなのか？
	flag.CommandLine.Var(&f, name, usage)
	// TODO: なぜアドレス渡しなのか？
	return &f.Celsius
}
