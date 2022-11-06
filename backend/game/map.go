package game

import "math/rand"


//function for creating starting map with random placed breakable blocks.
// 0 is floor tile
// 1 is breakable block
// 2 is wall
func CreateBaseMap() []int {
	var m = make([]int, 0)
	i := 0
	for i < 121 {
		if i < 11 || i > 109 || i%11 == 0 || i == 10 || i == 21 || i == 32 || i == 43 || i == 54 || i == 65 || i == 76 || i == 87 || i == 98 || i == 109 || i == 24 || i == 26 || i == 28 || i == 30 || i == 46 || i == 48 || i == 50 || i == 52 || i == 68 || i == 70 || i == 72 || i == 74 || i == 90 || i == 92 || i == 94 || i == 96 {
			m = append(m, 2)
		} else {
			if rand.Intn(10) > 5 {
				m = append(m, 0)
			} else {
				m = append(m, 1)
			}
		}
		i++
	}
	return m
}