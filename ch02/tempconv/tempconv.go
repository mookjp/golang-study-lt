package tempconv

import "fmt"

// Celsius は摂氏をあらわします
type Celsius float64

// Fahrenheit は華氏をあらわします
type Fahrenheit float64

// KelvinScale は絶対温度をあらわします
type KelvinScale float64

// 温度計算に必要なconstantsです
const (
	DiffOfKAndC   float64 = 273.15
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
func (k KelvinScale) String() string {
	return fmt.Sprintf("%gK", k)
}
