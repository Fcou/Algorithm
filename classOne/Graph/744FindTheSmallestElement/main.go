/*
给定一个只包含小写字母的有序数组 letters 和一个目标字母 target
寻找有序数组里面比目标字母大的最小字母
*/

package main

import (
	"fmt"
	"sort"
)

func findTheSmallest(letters []byte, target byte) (sma byte) {
	zcha := 26
	fcha := 0
	for _, l := range letters {
		if int(l-target) > 0 && int(l-target) < zcha {
			zcha = int(l - target)
			sma = l
		}
	}
	if zcha == 26 {
		for _, l := range letters {
			lt := int(l) - int(target)
			if lt < 0 && lt < fcha {
				fcha = int(l - target)
				sma = l
			}
		}
	}

	return
}

//运用了sort包里的函数
func nextGreatestLetter(letters []byte, target byte) byte {
	n := len(letters)
	i := sort.Search(n, func(i int) bool {
		return target < letters[i]
	})
	return letters[i%n]
}

func main() {
	letters := []byte{'a', 'b', 'e', 'f', 'h'}
	var target byte = 'h'
	sma := findTheSmallest(letters, target)
	fmt.Printf("Find the smallest is:%c", sma)
}
