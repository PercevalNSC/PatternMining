package transactionpackage

import (
	"reflect"
	"testing"
)

func TestItemset(t *testing.T) {
	itemset1 := Itemset{[]string{"A", "B", "C"}}
	itemset2 := Itemset{[]string{"B", "C", "E"}}
	sum_itemset := appendItemset(itemset1, itemset2)
	example_itemset := Itemset{[]string{"A", "B", "C", "B", "C", "E"}}
	if !reflect.DeepEqual(example_itemset, sum_itemset) {
		t.Error("Itemset append error")
	}
}

func TestItemsetList(t *testing.T) {
	itemset1 := Itemset{[]string{"A", "B", "C"}}
	itemset2 := Itemset{[]string{"A", "B", "E"}}
	item_list := ItemsetList{[]Itemset{itemset1, itemset2}}
	item_list1 := ItemsetList{[]Itemset{itemset1}}
	item_list2 := ItemsetList{[]Itemset{itemset2}}
	if !reflect.DeepEqual(item_list, appendItemsetList(item_list1, item_list2)) {
		t.Error("Itemset append error")
	}

	head, tail := item_list1.getHeadItemset(1)
	if head.itemset_list[0].data[0] != "A" {
		t.Error(head)
	}
	if tail.itemset_list[0].data[0] != "B" && tail.itemset_list[0].data[1] != "C" {
		t.Error(tail)
	}
	head, tail = item_list1.getHeadItemset(4)
	if !reflect.DeepEqual(head, item_list1) {
		t.Error(head)
	}
	if !reflect.DeepEqual(tail, ItemsetList{[]Itemset{}}) {
		t.Error(tail)
	}

	if !reflect.DeepEqual(item_list.generateCandidates().itemset_list[0].data, []string{"A", "B", "C", "E"}) {
		t.Error(item_list.generateCandidates())
	}
}
