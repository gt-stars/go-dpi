package classifiers

import (
	"github.com/google/gopacket/layers"
	"github.com/mushorg/go-dpi"
)

// RPCClassifier struct
type RPCClassifier struct{}

// HeuristicClassify for RPCClassifier
func (classifier RPCClassifier) HeuristicClassify(flow *godpi.Flow) bool {
	if len(flow.Packets) == 0 {
		return false
	}
	for _, packet := range flow.Packets {
		if layer := (*packet).Layer(layers.LayerTypeTCP); layer != nil {
			srcPort := layer.(*layers.TCP).SrcPort
			dstPort := layer.(*layers.TCP).DstPort
			if srcPort != 135 && dstPort != 135 {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

// GetProtocol returns the corresponding protocol
func (classifier RPCClassifier) GetProtocol() godpi.Protocol {
	return godpi.Rpc
}
