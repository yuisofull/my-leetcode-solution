// https://leetcode.com/problems/string-compression

import "strconv"

func compress(chars []byte) int {
	tmp := chars
	chars = make([]byte, 0)
	var sum int
	chars = append(chars, tmp[0])
	sum = 1
	for i := 1; i < len(tmp); i++ {
		if tmp[i] != tmp[i-1] {
			if sum != 1 {
				chars = append(chars, []byte(strconv.Itoa(sum))...)
				sum = 1
			}
			chars = append(chars, tmp[i])
		} else {
			sum++
		}
	}
	if sum != 1 {
		chars = append(chars, []byte(strconv.Itoa(sum))...)
	}
	return len(chars)
}