package main

import (
	"fmt"
	"reflect"
)

type Itemset struct {
	data []string
}

func (d *Itemset) getItemset() []string {
	return d.data
}
func appendItemset(d1 Itemset, d2 Itemset) Itemset {
	data := append(d1.data, d2.data...)
	return Itemset{data}
}

type ItemsetList struct {
	itemset_list []Itemset
}

func appendItemsetList(itemlist1 ItemsetList, itemlist2 ItemsetList) ItemsetList {
	data := append(itemlist1.itemset_list, itemlist2.itemset_list...)
	return ItemsetList{data}
}

func (il *ItemsetList) getHeadItemset(k int) (ItemsetList, ItemsetList) {
	result1 := []Itemset{}
	result2 := []Itemset{}
	for _, itemset := range il.itemset_list {
		if k < len(itemset.data) {
			result1 = append(result1, Itemset{itemset.data[:k]})
			result2 = append(result2, Itemset{itemset.data[k:]})
		} else {
			fmt.Println("k:", k, "is over itemset size:", len(itemset.data))
			return *il, ItemsetList{result2}
		}
	}
	return ItemsetList{result1}, ItemsetList{result2}
}

func (itemset_list *ItemsetList) generateCandidates() ItemsetList {
	k := len(itemset_list.itemset_list[0].data) + 1
	if k == 2 {
		return itemset_list.generateCandidate1()
	} else if k > 2 {
		return itemset_list.generateCandidates2(k)
	} else {
		fmt.Println("generate candidates fail")
		return ItemsetList{}
	}
}

func (itemset_list *ItemsetList) generateCandidates2(k int) ItemsetList {
	head, tail := itemset_list.getHeadItemset(k - 2)
	result := []Itemset{}
	for i, head_itemset := range head.itemset_list {
		for j := i + 1; j < len(head.itemset_list); j++ {
			if reflect.DeepEqual(head.itemset_list[j], head_itemset) {
				merge_itemset := appendItemset(appendItemset(head_itemset, tail.itemset_list[i]), tail.itemset_list[j])
				result = append(result, merge_itemset)
			}
		}
	}
	return ItemsetList{result}
}

func (itemset_list *ItemsetList) generateCandidate1() ItemsetList {
	candidate := []Itemset{}

	for i, itemset := range itemset_list.itemset_list {
		for j := i + 1; j < len(itemset_list.itemset_list); j++ {
			candidate = append(candidate, appendItemset(itemset, itemset_list.itemset_list[j]))
		}
	}

	return ItemsetList{candidate}
}

type Transaction struct {
	id   int
	data []string
}

func (t *Transaction) set(id int, data []string) {
	t.id = id
	t.data = data
}

func (t *Transaction) isInItemset(target_itemset Itemset) bool {
	for _, target_item := range target_itemset.getItemset() {
		if t.isInItem(target_item) {
			continue
		} else {
			return false
		}
	}
	return true
}
func (t *Transaction) isInItem(target_item string) bool {
	for _, transaction_item := range t.data {
		if transaction_item == target_item {
			return true
		}
	}
	return false
}

type TransactionData struct {
	transaction_data []Transaction
}

func (td *TransactionData) set(transaction_data []Transaction) {
	td.transaction_data = transaction_data
}
func (td *TransactionData) append(transaction Transaction) {
	td.transaction_data = append(td.transaction_data, transaction)
}

func (td *TransactionData) countItemset(itemset Itemset) int {
	count := 0
	for _, transaction := range td.transaction_data {
		if transaction.isInItemset(itemset) {
			count++
		}
	}
	return count
}

type PathTable struct {
	itemset_list []Itemset
	count_list   []int
}

func (pt *PathTable) addItemset(itemset Itemset, count int) {
	pt.itemset_list = append(pt.itemset_list, itemset)
	pt.count_list = append(pt.count_list, count)
}

func (pt *PathTable) getFrequentItemsetList(min_frequency int) ItemsetList {
	itemset_list := []Itemset{}
	for i, count := range pt.count_list {
		if count >= min_frequency {
			itemset_list = append(itemset_list, pt.itemset_list[i])
		}
	}
	return ItemsetList{itemset_list}
}

func (td *TransactionData) path(itemset_list ItemsetList) PathTable {
	var count int
	var path_table PathTable
	for _, itemset := range itemset_list.itemset_list {
		count = td.countItemset(itemset)
		path_table.addItemset(itemset, count)
	}

	fmt.Println(path_table)

	return path_table
}

func (td *TransactionData) getMinFrequency(min_support float64) int {
	return int(float64(len(td.transaction_data)) * min_support)
}

func (td *TransactionData) pickupFrequencyItemset(min_support float64) ItemsetList {
	frequecy_itemset_list := ItemsetList{}

	min_frequency := td.getMinFrequency(min_support)
	fmt.Println(min_support, min_frequency)

	// path == 1
	candidates := initItemsetList()
	path_table := td.path(candidates)
	frequent_itemset := path_table.getFrequentItemsetList(min_frequency)
	frequecy_itemset_list = appendItemsetList(frequecy_itemset_list, frequent_itemset)

	// path >= 2
	i := 2
	for len(frequent_itemset.itemset_list) != 0 {
		candidates = frequent_itemset.generateCandidates()
		fmt.Println("path", i, "candidates", candidates)
		path_table = td.path(candidates)
		frequent_itemset = path_table.getFrequentItemsetList(min_frequency)
		fmt.Println("path", i, "frequent_itemset", frequent_itemset)

		frequecy_itemset_list = appendItemsetList(frequecy_itemset_list, frequent_itemset)
		i++
	}

	return frequecy_itemset_list
}

func sampleTrasactionData() TransactionData {
	var td TransactionData
	var t Transaction
	t.set(10, []string{"A", "B", "D"})
	td.append(t)
	t.set(20, []string{"A", "B", "C"})
	td.append(t)
	t.set(30, []string{"A", "B", "C", "E"})
	td.append(t)
	t.set(40, []string{"B", "E"})
	td.append(t)
	return td
}

func initItemsetList() ItemsetList {
	itemset_list := []Itemset{}
	itemlist := []string{"A", "B", "C", "D", "E"}
	for _, item := range itemlist {
		itemset_list = append(itemset_list, Itemset{[]string{item}})
	}
	return ItemsetList{itemset_list}
}

func main() {
	min_support := 0.5

	trasaction_data := sampleTrasactionData()
	fmt.Println(trasaction_data)

	fmt.Println(trasaction_data.pickupFrequencyItemset(min_support))
}
