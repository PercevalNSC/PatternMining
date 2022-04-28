package main

import (
	"fmt"

	"DB2/chord"
	tp "DB2/transactionpackage"
)

func sampleTrasactionData() tp.TransactionData {
	var td tp.TransactionData
	var t tp.Transaction
	t.Set(10, []string{"A", "B", "D"})
	td.Append(t)
	t.Set(20, []string{"A", "B", "C"})
	td.Append(t)
	t.Set(30, []string{"A", "B", "C", "E"})
	td.Append(t)
	t.Set(40, []string{"B", "E"})
	td.Append(t)
	return td
}

func pattern_mining() {
	min_support := 0.5

	trasaction_data := sampleTrasactionData()
	fmt.Println(trasaction_data)

	fmt.Println(trasaction_data.PickupFrequencyItemset(min_support))
}

func main() {
	chord_network := chord.ChordNetwork{}
	chord1 := chord.ConstructorChordNord(8)
	fmt.Println(chord1)
	chord_network = chord_network.AddNode(chord1)
	fmt.Println(chord_network)

}
