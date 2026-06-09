package tempconv

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5) + FreezingF
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - FreezingF) * 5 / 9)
}

func KToC(k Kelvin) Celsius {
	return Celsius(k) + AbsoluteZeroC
}

func CToK(c Celsius) Kelvin {
	return Kelvin(c - AbsoluteZeroC)
}

func KToF(k Kelvin) Fahrenheit {
	return CToF(KToC(k))
}

func FToK(f Fahrenheit) Kelvin {
	return CToK(FToC(f))
}
