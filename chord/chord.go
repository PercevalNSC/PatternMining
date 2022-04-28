package chord

type ChordNord struct {
	end   int
	cache []int
}

func ConstructorChordNord(end int) ChordNord {
	cn := ChordNord{}
	cn.end = end

	return cn
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

func ChordNetworkFromList(nord_num_list []int) ChordNetwork {
	var end int
	var new_nord ChordNord
	var chord_network ChordNetwork
	for _, nord_num := range nord_num_list {
		end = nord_num
		new_nord = ConstructorChordNord(end)
		chord_network = chord_network.AddNode(new_nord)
	}

	return chord_network
}

func (cn *ChordNetwork) SearchNodeIndex(target_index int) int {
	for _, node := range cn.node_list {
		if target_index <= node.end {
			return node.end
		}
	}
	return 1
}
