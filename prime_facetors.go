package prime

const testVersion = 2

func Factors(n int64) (factors []int64) {
	factors = make([]int64, 0)
	for n%2 == 0 {
		n /= 2
		factors = append(factors, 2)
	}
	for i := int64(3); i <= n; i = i + 2 {
		for n%i == 0 {
			n /= i
			factors = append(factors, i)
		}
	}
	return
}
