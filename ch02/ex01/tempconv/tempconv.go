package tempconv

import "fmt"

// Celsius は摂氏をあらわします
type Celsius float64

// Fahrenheit は華氏をあらわします
type Fahrenheit float64

// 温度計算に必要なconstantsです
const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g℃", c)
}
func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}
