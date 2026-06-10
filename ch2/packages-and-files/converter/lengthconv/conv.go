package lengthconv

func MToFt(m Meter) Foot {
	return Foot(m / MetersPerFoot)
}

func FtToM(f Foot) Meter {
	return Meter(f * MetersPerFoot)
}
