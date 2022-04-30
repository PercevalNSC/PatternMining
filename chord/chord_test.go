package chord

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

var node1 = ChordNord{1, 0, nil}
var node2 = ChordNord{8, 0, nil}
var node_1_2_push = ChordNetwork{[]ChordNord{node1, node2}}
var node_1_2_link = ChordNetwork{[]ChordNord{{1, 8, nil}, {8, 1, nil}}}
var node3 = ChordNord{3, 0, nil}
var node_1_2_3 = ChordNetwork{[]ChordNord{{1, 3, nil}, {3, 8, nil}, {8, 1, nil}}}
var list1 = []int{3, 8}
var list2 = []int{8, 3}
var hash_list = []int{8, 14, 21, 32, 38, 42, 48, 53, 66}
var index_pairs = []map[string]int{
	{"index": 1, "node": 8},
	{"index": 14, "node": 14},
	{"index": 43, "node": 48},
	{"index": 77, "node": 1},
}

func TestChord(t *testing.T) {
	chord_network := ChordNetwork{}

	t.Run("constructorChordNode", func(t *testing.T) {
		construct_node := ConstructorChordNord(node1.end)
		if !reflect.DeepEqual(construct_node, node1) {
			t.Error("constructChordNord Error", construct_node, node1)
		}
	})
	t.Run("pushNode", func(t *testing.T) {
		chord_network = chord_network.PushNode(node2)
		t.Run("length 1", func(t *testing.T) {
			if !reflect.DeepEqual(chord_network, node_1_2_push) {
				t.Error(chord_network)
			}
		})
	})
	t.Run("linkNode", func(t *testing.T) {
		chord_network.linkNode()
		if !reflect.DeepEqual(chord_network, node_1_2_link) {
			t.Error("chord_network", chord_network, "test_cn", node_1_2_link)
		}
	})
	t.Run("addNode", func(t *testing.T) {
		chord_network = chord_network.AddNode(node3)
		if !reflect.DeepEqual(chord_network, node_1_2_3) {
			t.Error(chord_network, ":", node_1_2_3)
		}
	})
	t.Run("consisitency test", func(t *testing.T) {
		if !chord_network.checkConsistency() {
			t.Error(chord_network, chord_network.checkConsistency())
		}
	})
	t.Run("construct from list", func(t *testing.T) {
		t.Run("ordered list", func(t *testing.T) {
			cn_from_list := ChordNetworkFromList(list1)
			if !reflect.DeepEqual(chord_network, cn_from_list) {
				t.Error(chord_network, ":", cn_from_list)
			}
		})
		t.Run("unorderd_list", func(t *testing.T) {
			cn_from_list := ChordNetworkFromList(list2)
			if !reflect.DeepEqual(chord_network, cn_from_list) {
				t.Error(chord_network, ":", cn_from_list)
			}
		})

	})

	chord_network = ChordNetworkFromList(hash_list)
	fmt.Println(chord_network)

	t.Run("Search Node", func(t *testing.T) {
		for i, index_pair := range index_pairs {
			t.Run("pair"+strconv.Itoa(i), func(t *testing.T) {
				searched_index := chord_network.SearchNodeIndex(index_pair["index"])
				if searched_index != index_pair["node"] {
					t.Error(searched_index, index_pair["node"])
				}
			})
		}
	})

}
