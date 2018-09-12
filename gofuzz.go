package bug

import "strconv"

func FaintHeart(i int) {
	if i > 1000000 {
		str := strconv.Itoa(i)
		panic("wohooo:" + str)
	}
}
