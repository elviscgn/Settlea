package data

import (
	"fmt"
	"math/rand"
)

type Port struct {
	Text  string  `json:"text"`
	Coord [2]int  `json:"coord"`
	Size  float64 `json:"size,omitempty"`
}

type PortData struct {
	ExchangeRate Port `json:"exchangeRate"`
	PortType     Port `json:"portType"`
}

func ShufflePorts(portMappings map[string]PortData) map[string]PortData {
	keys := make([]string, 0, len(portMappings))
	for key := range portMappings {
		keys = append(keys, key)
	}

	rand.Shuffle(len(keys), func(i, j int) { keys[i], keys[j] = keys[j], keys[i] })

	shuffledPorts := make(map[string]PortData)
	for i, key := range keys {
		shuffledPorts[fmt.Sprintf("port%d", i+1)] = portMappings[key]
	}
	return shuffledPorts
}

func GeneratePorts(numPorts int) map[string]PortData {
	portMappings := map[string]PortData{
		"port1": {
			ExchangeRate: Port{
				Text:  "3:1",
				Coord: [2]int{770, 490},
			},
			PortType: Port{
				Text:  "?",
				Coord: [2]int{775, 465},
			},
		},
		"port2": {
			ExchangeRate: Port{
				Text:  "2:1",
				Coord: [2]int{933, 215},
			},
			PortType: Port{
				Text:  "wood",
				Coord: [2]int{945, 205},
				Size:  0.039,
			},
		},
		"port3": {
			ExchangeRate: Port{
				Text:  "3:1",
				Coord: [2]int{1255, 210},
			},
			PortType: Port{
				Text:  "brick",
				Coord: [2]int{1267, 200},
				Size:  0.35,
			},
		},
		"port4": {
			ExchangeRate: Port{
				Text:  "3:1",
				Coord: [2]int{1505, 355},
			},
			PortType: Port{
				Text:  "sheep",
				Coord: [2]int{1515, 345},
				Size:  0.4,
			},
		},
		"port5": {
			ExchangeRate: Port{
				Text:  "2:1",
				Coord: [2]int{1675, 649},
			},
			PortType: Port{
				Text:  "ore",
				Coord: [2]int{1688, 635},
				Size:  0.04,
			},
		},
		"port6": {
			ExchangeRate: Port{
				Text:  "2:1",
				Coord: [2]int{1492, 935},
			},
			PortType: Port{
				Text:  "wheat",
				Coord: [2]int{1505, 925},
				Size:  0.045,
			},
		},
		"port7": {
			ExchangeRate: Port{
				Text:  "3:1",
				Coord: [2]int{1250, 1060},
			},
			PortType: Port{
				Text:  "?",
				Coord: [2]int{1255, 1033},
			},
		},
		"port8": {
			ExchangeRate: Port{
				Text:  "3:1",
				Coord: [2]int{947, 1055},
			},
			PortType: Port{
				Text:  "?",
				Coord: [2]int{950, 1027},
			},
		},
		"port9": {
			ExchangeRate: Port{
				Text:  "3:1",
				Coord: [2]int{761, 788},
			},
			PortType: Port{
				Text:  "?",
				Coord: [2]int{765, 760},
			},
		},
	}

	return ShufflePorts(portMappings)
}
