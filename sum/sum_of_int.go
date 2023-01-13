package sum

var ByteToIntMap = map[byte]int64{
	'0': 0,
	'1': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
}

func SumIntegers(s string) int64 {
	sum := int64(0)
	for i := 0; i < len(s); i++ {
		if v, ok := ByteToIntMap[s[i]]; ok {
			sum += v
		}
	}

	return sum
}
