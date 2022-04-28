package chord

import (
	"fmt"
	"reflect"
	"testing"
)

var node1 = ChordNord{1, nil}
var node2 = ChordNord{8, nil}
var node3 = ChordNord{16, nil}
var hash_list = []int{8, 14, 21, 32, 38, 42, 48, 53, 66}

func TestChord(t *testing.T) {
	chord_network := ChordNetwork{}

	t.Run("constructorChordNode", func(t *testing.T) {
		construct_node := ConstructorChordNord(node1.end)
		if !reflect.DeepEqual(construct_node, node1) {
			t.Error("constructChordNord Error")
		}
	})
	t.Run("addNode", func(t *testing.T) {
		construct_node1 := ConstructorChordNord(1)
		construct_node2 := ConstructorChordNord(8)
		chord_network = chord_network.AddNode(construct_node2)
		t.Run("length 1", func(t *testing.T) {
			if !reflect.DeepEqual(chord_network.node_list, []ChordNord{construct_node1, construct_node2}) {
				t.Error(chord_network)
			}
		})

		construct_node3 := ConstructorChordNord(16)
		chord_network = chord_network.AddNode(construct_node3)
		t.Run("length > 2", func(t *testing.T) {
			if !reflect.DeepEqual(chord_network.node_list,
				[]ChordNord{construct_node1, construct_node2, construct_node3}) {
				t.Error(chord_network)
			}
		})
	})
	t.Run("consisitency test", func(t *testing.T) {
		if !chord_network.checkConsistency() {
			t.Error(chord_network, chord_network.checkConsistency())
		}
	})
	t.Run("construct from list", func(t *testing.T) {
		int_list := []int{8, 16}
		cn_from_list := ChordNetworkFromList(int_list)
		if !reflect.DeepEqual(chord_network, cn_from_list) {
			t.Error(cn_from_list)
		}
		int_list = []int{16, 8}
		cn_from_list = ChordNetworkFromList(int_list)
		if !reflect.DeepEqual(chord_network, cn_from_list) {
			t.Error(cn_from_list)
		}
	})

	chord_network = ChordNetworkFromList(hash_list)
	fmt.Println(chord_network)

	t.Run("Search Node", func(t *testing.T) {
		index := 14
		if chord_network.SearchNodeIndex(index) != 14 {
			t.Error(chord_network.SearchNodeIndex(index))
		}
	})

}
