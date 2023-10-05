package internal

func CheckDuplicateZero(dataN []int32) []int32 {
	total := len(dataN)
	numZeros := 0
	for i := 0; i < len(dataN); i++ {
		if dataN[i] == 0 {
			numZeros++
		}
	}
	for i := total - 1; i >= 0; i-- {
		if dataN[i] == 0 {
			if i+numZeros < total {
				dataN[i+numZeros] = 0
			}
			numZeros--
		}
		if i+numZeros < total {
			dataN[i+numZeros] = dataN[i]
		}
	}

	return dataN
}
