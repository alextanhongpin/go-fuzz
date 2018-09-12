// +build gofuzz

package bug

func ParseStrings(input []byte) (result []string) {
	for len(input) > 0 {
		str := input[0]
		input = input[1:]
		result = append(result, string(input[:str]))
		input = input[str:]
	}
	return
}

func Fuzz(input []byte) int {
	ParseStrings(input)
	return 1
}
