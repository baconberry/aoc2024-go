package aoc

import "aoc2024/util"

type FileFragment struct {
	fileId       int
	memoryPos    int
	memoryLength int
}

func Nine(lines []string) int {
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

	var sum int
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
func NineSecond(lines []string) int {
	arr := util.ParseSingleDigitArray(lines[0])
	memory := make([]FileFragment, 0)
	memoryOffset := 0
	for i, n := range arr {
		if i%2 == 0 {
			memory = append(memory, FileFragment{i / 2, memoryOffset, n})
		} else {
			memory = append(memory, FileFragment{-1, memoryOffset, n})
		}
		memoryOffset += n
	}
	var sum int
	var zeroPos int
	newMemory := make([]FileFragment, 0)
	for zeroPos < len(memory) {
		fragment := memory[zeroPos]
		if fragment.fileId >= 0 {
			newMemory = append(newMemory, fragment)
			zeroPos += 1
			continue
		}
		zeroLength := fragment.memoryLength
		lastFragmentPos := queryLastItem(memory[zeroPos+1:], zeroLength, zeroPos+1)
		if lastFragmentPos < 0 {
			zeroPos += 1
			if fragment.memoryLength > 0 {
				newMemory = append(newMemory, fragment)
			}
			continue
		}
		lastFragment := memory[lastFragmentPos]
		memory[lastFragmentPos] = FileFragment{-1, lastFragment.memoryPos, lastFragment.memoryLength}
		lastFragment.memoryPos = fragment.memoryPos
		newMemory = append(newMemory, lastFragment)
		memory[zeroPos].memoryLength -= lastFragment.memoryLength
		memory[zeroPos].memoryPos += lastFragment.memoryLength
	}
	for _, fragment := range newMemory {
		if fragment.fileId >= 0 {
			sum += checksum(fragment.memoryPos, fragment.memoryLength+fragment.memoryPos, fragment.fileId)
		}
	}
	return sum
}
func queryLastItem(fragments []FileFragment, length int, offset int) int {
	if len(fragments) == 0 {
		return -1
	}
	lastPos := len(fragments) - 1
	if fragments[lastPos].memoryLength <= length && fragments[lastPos].fileId >= 0 {
		return lastPos + offset
	}
	return queryLastItem(fragments[:lastPos], length, offset)
}

func checksum(start int, end int, id int) int {
	var checksum int
	for i := start; i < end; i++ {
		checksum += int(id * i)
	}
	return checksum
}
