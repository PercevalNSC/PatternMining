package chord

import (
	"sort"
)

type ChordNord struct {
	end   int
	next  int
	cache []int
}

func ConstructorChordNord(end int) ChordNord {
	cn := ChordNord{}
	cn.end = end

	return cn
}

func (node *ChordNord) isIndexInNode(index int) int {
	if node.end < index || node.end == 1 {
		return node.next
	} else {
		return node.end
	}
}

type ChordNetwork struct {
	node_list []ChordNord
}

func (chord_network *ChordNetwork) checkConsistency() bool {
	return true
}

func (cn *ChordNetwork) AddNode(add_node ChordNord) ChordNetwork {
	result := cn

	if len(cn.node_list) == 0 {
		result.node_list = []ChordNord{ConstructorChordNord(1), add_node}
	} else {
		for i, chord_node := range cn.node_list {
			if chord_node.end < add_node.end {
				if i == len(cn.node_list)-1 {
					result.node_list = append(result.node_list, add_node)
				} else if add_node.end < result.node_list[i+1].end {
					*result = cn.insertNode(i+1, add_node)
				}
			} else {
				break
			}
		}
	}

	result.linkNode()

	return *result
}

func (cn *ChordNetwork) PushNode(add_node ChordNord) ChordNetwork {
	result := cn

	if len(cn.node_list) == 0 {
		result.node_list = []ChordNord{ConstructorChordNord(1), add_node}
	} else {
		result.node_list = append(result.node_list, add_node)
	}

	return *result
}

func (cn *ChordNetwork) insertNode(position int, add_node ChordNord) ChordNetwork {
	result := []ChordNord{}
	result = append(result, cn.node_list[:position]...)
	//fmt.Println(result)
	result = append(result, add_node)
	result = append(result, cn.node_list[position:len(cn.node_list)]...)
	return ChordNetwork{result}
}

func ChordNetworkFromList(node_num_list []int) ChordNetwork {
	var end int
	var new_nord ChordNord
	var chord_network ChordNetwork

	sort.Slice(node_num_list, func(i, j int) bool { return node_num_list[i] < node_num_list[j] })

	for _, nord_num := range node_num_list {
		end = nord_num
		new_nord = ConstructorChordNord(end)
		chord_network = chord_network.PushNode(new_nord)
	}

	chord_network.linkNode()

	return chord_network
}

func (cn *ChordNetwork) linkNode() {
	linked_cn := cn
	for i := range linked_cn.node_list {
		if i == len(cn.node_list)-1 {
			linked_cn.node_list[i].next = linked_cn.node_list[0].end
			break
		}
		linked_cn.node_list[i].next = linked_cn.node_list[i+1].end
	}
}

func (cn *ChordNetwork) SearchNodeIndex(target_index int) int {
	var old_index, new_index int
	old_index = cn.node_list[0].isIndexInNode(target_index)

	for i, node := range cn.node_list {
		if i == 0 {
			continue
		}

		new_index = node.isIndexInNode(target_index)
		if old_index == new_index {
			return old_index
		} else {
			old_index = new_index
		}
	}
	return cn.node_list[0].end
}
