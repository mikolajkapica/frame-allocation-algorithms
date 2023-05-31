package utils

func SumOfTrue(a []bool) int {
	sum := 0
	for i := 0; i < len(a); i++ {
		if a[i] {
			sum++
		}
	}
	return sum
}
