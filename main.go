package main

import (
	"fmt"

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

func main() {
	min_support := 0.5

	trasaction_data := sampleTrasactionData()
	fmt.Println(trasaction_data)

	fmt.Println(trasaction_data.PickupFrequencyItemset(min_support))
}
