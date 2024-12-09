package aoc

import "aoc2024/util"

func Nine(lines []string) uint {
	arr := util.ParseSingleDigitArray(lines[0])
	zeroes := make([]int, 0)
	files := make([]int, 0)
	for i, n := range arr {
		if i%2 == 0 {
			files = append(files, n)
		} else {
			zeroes = append(zeroes, n)
		}
	}
	lastPos := len(files) - 1
	zeroesPos := 0
	fileId := 0
	memoryPos := 0

	var sum uint
	for files[fileId] > 0 {
		fileLength := files[fileId]
		sum += checksum(memoryPos, memoryPos+fileLength, fileId)
		files[fileId] = 0
		memoryPos += fileLength
		fileId += 1
		zeroLength := zeroes[zeroesPos]
		if zeroLength == 0 {
			zeroesPos += 1
			continue
		}
		lastFileLength := files[lastPos]
		for zeroes[zeroesPos] > 0 && lastFileLength > 0 {
			if lastFileLength > zeroLength {
				sum += checksum(memoryPos, memoryPos+zeroLength, lastPos)
				memoryPos += zeroLength
				zeroes[zeroesPos] = 0
				files[lastPos] -= zeroLength
				zeroesPos += 1
				break
			} else {
				sum += checksum(memoryPos, memoryPos+lastFileLength, lastPos)
				memoryPos += lastFileLength
				zeroes[zeroesPos] -= lastFileLength
				zeroLength -= lastFileLength
				files[lastPos] = 0
				lastPos -= 1
				lastFileLength = files[lastPos]
				if zeroes[zeroesPos] == 0 {
					zeroesPos += 1
					break
				}
			}

		}
	}
	return sum
}

func checksum(start int, end int, id int) uint {
	var checksum uint
	for i := start; i < end; i++ {
		checksum += uint(id * i)
	}
	return checksum
}
