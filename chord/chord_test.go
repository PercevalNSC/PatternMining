package chord

import (
	"reflect"
	"testing"
)

func TestChord(t *testing.T) {
	chord_network := ChordNetwork{}

	t.Run("constructorChordNode", func(t *testing.T) {
		nord1 := ConstructorChordNord(1, 8)
		if !reflect.DeepEqual(nord1, ChordNord{1, 8, nil}) {
			t.Error("constructChordNord Error")
		}
	})
	t.Run("addNode", func(t *testing.T) {
		nord1 := ConstructorChordNord(1, 8)
		nord2 := ConstructorChordNord(8, 1)
		chord_network = chord_network.AddNode(nord1)
		t.Run("length 1", func(t *testing.T) {
			if !reflect.DeepEqual(chord_network.node_list, []ChordNord{nord1, nord2}) {
				t.Error(chord_network)
			}
		})

		nord2 = ConstructorChordNord(16, 1)
		nord3 := ConstructorChordNord(8, 16)
		chord_network = chord_network.AddNode(nord3)
		t.Run("length > 2", func(t *testing.T) {
			if !reflect.DeepEqual(chord_network.node_list,
				[]ChordNord{nord1, nord3, nord2}) {
				t.Error(chord_network)
			}
		})
	})

}