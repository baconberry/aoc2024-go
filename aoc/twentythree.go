package aoc

import (
	"aoc2024/util"
	"log"
	"sort"
	"strings"
)

type DeviceSet map[string]bool

type ComputerNode struct {
	id          string
	connections DeviceSet
}

func Twentythree(lines []string, part int) int {
	grid := util.ParseStringGrid(lines, "-")
	network := make(map[string]ComputerNode)
	notCountedMap := make(DeviceSet)
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
	if part == 1 {
		return setsOfThreeDevices(allDevices, network)
	}

	if part == 2 {
		p := make(DeviceSet, len(allDevices))
		for _, device := range allDevices {
			p[device] = true
		}
		maxCompleteSubgraph := bronKerbosch(&network, make(DeviceSet), p, make(DeviceSet))
		maxArray := make([]string, 0)
		for s, _ := range maxCompleteSubgraph {
			maxArray = append(maxArray, s)
		}
		sort.Strings(maxArray)
		log.Println(strings.Join(maxArray, ","))
		return len(maxCompleteSubgraph)
	}

	return sum
}
func bronKerbosch(network *map[string]ComputerNode, r, p, x DeviceSet) DeviceSet {
	if len(p) == 0 && len(x) == 0 {
		return r
	}
	maximal := make(DeviceSet)
	for device, _ := range p {
		np := intersect(p, (*network)[device].connections)
		nx := intersect(x, (*network)[device].connections)
		nr := union(r, device)
		result := bronKerbosch(network, nr, np, nx)
		if len(result) > len(maximal) {
			maximal = result
		}
		delete(*network, device)
		x[device] = true
	}
	return maximal
}

func union(r DeviceSet, device string) DeviceSet {
	dst := cloneSet(r)
	dst[device] = true
	return dst
}

func cloneSet(src DeviceSet) DeviceSet {
	dst := make(DeviceSet, len(src))
	for s, b := range src {
		dst[s] = b
	}
	return dst
}

func intersect(a DeviceSet, b DeviceSet) DeviceSet {
	intersection := make(DeviceSet)
	for d, _ := range a {
		if b[d] {
			intersection[d] = true
		}
	}
	return intersection
}

func setsOfThreeDevices(allDevices []string, network map[string]ComputerNode) int {
	sum := 0
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
		(*n)[d1] = ComputerNode{d1, make(DeviceSet)}
	}
	(*n)[d1].connections[d2] = true
}
