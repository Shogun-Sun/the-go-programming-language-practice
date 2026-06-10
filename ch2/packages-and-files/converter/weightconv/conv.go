package weightconv

func PtoKg(p Pound) Kilogram {
	return Kilogram(p * KgPerP)
}

func KgToP(kg Kilogram) Pound {
	return Pound(kg / KgPerP)
}
