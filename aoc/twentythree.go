package aoc

import (
	"aoc2024/util"
)

type ComputerNode struct {
	id          string
	connections map[string]bool
}

func Twentythree(lines []string) int {
	grid := util.ParseStringGrid(lines, "-")
	network := make(map[string]ComputerNode)
	notCountedMap := make(map[string]bool)
	for _, row := range grid {
		device1 := row[0]
		device2 := row[1]
		connectDevice(&network, device1, device2)
		connectDevice(&network, device2, device1)
		notCountedMap[device1] = false
		notCountedMap[device2] = false
	}
	sum := 0
	allDevices := make([]string, len(notCountedMap))
	i := 0
	for device, _ := range notCountedMap {
		allDevices[i] = device
		i += 1
	}
	for i := 0; i < len(allDevices); i++ {
		for j := i + 1; j < len(allDevices); j++ {
			for k := j + 1; k < len(allDevices); k++ {
				di := allDevices[i]
				dj := allDevices[j]
				dk := allDevices[k]
				if anyStartsWith('t', di, dj, dk) {
					connected := isConnected(&network, di, dj)
					connected = connected && isConnected(&network, dj, dk)
					connected = connected && isConnected(&network, di, dk)
					if connected {
						sum += 1
					}
				}
			}
		}
	}

	return sum
}

func anyStartsWith(toFind rune, ds ...string) bool {
	for _, d := range ds {
		if rune(d[0]) == toFind {
			return true
		}
	}
	return false
}

func isConnected(network *map[string]ComputerNode, a string, b string) bool {
	return (*network)[a].connections[b]
}

func connectDevice(n *map[string]ComputerNode, d1 string, d2 string) {
	_, ok := (*n)[d1]
	if !ok {
		(*n)[d1] = ComputerNode{d1, make(map[string]bool)}
	}
	(*n)[d1].connections[d2] = true
}
