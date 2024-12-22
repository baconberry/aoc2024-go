package aoc

import (
	"aoc2024/util"
)

func Twentytwo(lines []string, part int) int {
	secretNumbers := util.ParseIntGrid(lines)
	sum := 0
	monkey := MonkeyBusiness{}
	bananaMap := make(map[MonkeyPrice]MonkeySeller)
	for m, number := range secretNumbers {
		newSecret := number[0]
		monkey.ResetBuffer()
		for i := 0; i < 2000; i++ {
			newSecret = monkey.iterateSecretNumber(newSecret)
			if i > 3 {
				mp := monkey.CurrentMonkeyPrice()
				_, ok := bananaMap[mp]
				if !ok {
					bananaMap[mp] = NewMonkeySeller(len(lines))
				}
				bm := bananaMap[mp]
				if !bm.hasSold[m] {
					bm.bananas += newSecret % 10
					bm.hasSold[m] = true
					bananaMap[mp] = bm
				}
			}
		}
		sum += newSecret
	}
	if part == 2 {
		maxBananas := 0
		for _, bananas := range bananaMap {
			if maxBananas < bananas.bananas {
				//log.Println(bananas, combo)
				maxBananas = bananas.bananas
			}
		}
		return maxBananas
	}
	return sum
}

func (m *MonkeyBusiness) iterateSecretNumber(secretNumber int) int {
	initNumber := secretNumber
	secretNumber = (secretNumber * 64) ^ secretNumber
	secretNumber = secretNumber % 16777216
	secretNumber = (secretNumber / 32) ^ secretNumber
	secretNumber = secretNumber % 16777216
	secretNumber = (secretNumber * 2048) ^ secretNumber
	secretNumber = secretNumber % 16777216
	m.setNextNumber((secretNumber % 10) - (initNumber % 10))
	return secretNumber
}

type MonkeyBusiness struct {
	buffer  []int
	pointer int
}

type MonkeyPrice struct {
	c1, c2, c3, c4 int
}

type MonkeySeller struct {
	bananas int
	hasSold []bool
}

func (s *MonkeySeller) addBananas(i int) {
	s.bananas += i
}

func NewMonkeySeller(nMonkeys int) MonkeySeller {
	return MonkeySeller{0, make([]bool, nMonkeys)}
}

func (m *MonkeyBusiness) setNextNumber(n int) {
	m.pointer = (m.pointer + 1) % len(m.buffer)
	m.buffer[m.pointer] = n
}

func (m *MonkeyBusiness) CurrentMonkeyPrice() MonkeyPrice {
	mp := MonkeyPrice{}
	for i := 0; i < len(m.buffer); i++ {
		bufferPos := (m.pointer - i) % len(m.buffer)
		bufferPos = (bufferPos + 4) % len(m.buffer)
		mp.setPrice(i, m.buffer[bufferPos])
	}
	return mp
}

func (m *MonkeyBusiness) ResetBuffer() {
	m.buffer = make([]int, 4)
	m.pointer = -1
}

func (mp *MonkeyPrice) setPrice(pos, n int) {
	switch pos {
	case 0:
		mp.c4 = n
	case 1:
		mp.c3 = n
	case 2:
		mp.c2 = n
	case 3:
		mp.c1 = n
	}
}
