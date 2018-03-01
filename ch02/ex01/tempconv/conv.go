package tempconv

// TODO: test

// CToF は摂氏の温度を華氏に変換します。
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// FToC は華氏の温度を摂氏にに変換します。
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func CToK(c Celsius) KelvinScale {
	return KelvinScale(-c / AbsoluteZeroC)
}

func FToK(f Fahrenheit) KelvinScale {
	c := FToC(f)
	return KelvinScale(CToK(c))
}

func KToC(k KelvinScale) Celsius {
	return Celsius(float64(-k) * float64(AbsoluteZeroC))
}
