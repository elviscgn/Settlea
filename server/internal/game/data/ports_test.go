package data

import (
	"testing"
)

func TestShufflePorts(t *testing.T) {
	ports := GeneratePorts(9)

	originalPorts := make(map[string]PortData)
	for key, port := range ports {
		originalPorts[key] = port
	}

	shuffledPorts := ShufflePorts(ports)

	for key, original := range originalPorts {
		if original.ExchangeRate == shuffledPorts[key].ExchangeRate {
			t.Errorf("Expected exchange rate to be randomized for %s, got %v", key, shuffledPorts[key].ExchangeRate)
		}
		if original.PortType == shuffledPorts[key].PortType {
			t.Errorf("Expected port type to be randomized for %s, got %v", key, shuffledPorts[key].PortType)
		}
	}
}
