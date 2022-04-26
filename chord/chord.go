package chord

type ChordNord struct {
	start int
	end   int
	cache map[int]int
}

func ConstructorChordNord(start int, end int) ChordNord {
	cn := ChordNord{}
	cn.start = start
	cn.end = end

	return cn
}

type ChordNetwork struct {
	node_list []ChordNord
}

func (chord_network *ChordNetwork) checkConsistency() bool {
	var start int
	end := chord_network.node_list[0].end
	var node ChordNord

	for i := 1; i < len(chord_network.node_list); i++ {
		node = chord_network.node_list[i]
		start = node.start
		if start != end {
			return false
		}
		end = node.end
	}
	start = chord_network.node_list[0].start
	if start == end {
		return true
	} else {
		return false
	}
}

func (cn *ChordNetwork) AddNode(add_node ChordNord) ChordNetwork {
	result := ChordNetwork{}

	if len(cn.node_list) == 0 {
		result.node_list = []ChordNord{add_node}
	} else {
		for i, chord_node := range cn.node_list {
			if chord_node.end == add_node.start {
				result = cn.insertNode(i+1, add_node)
			}
		}
	}

	special_node := ConstructorChordNord(
		result.node_list[len(result.node_list)-1].end, result.node_list[0].start)
	result.node_list = append(result.node_list, special_node)

	return result
}

func (cn *ChordNetwork) insertNode(position int, add_node ChordNord) ChordNetwork {
	result := []ChordNord{}
	result = append(result, cn.node_list[:position]...)
	result = append(result, add_node)
	result = append(result, cn.node_list[position:len(cn.node_list)-1]...)
	return ChordNetwork{result}
}

func ChordNetworkFromList(nord_num_list []int) ChordNetwork {
	var start, end int
	var new_nord ChordNord
	var chord_network ChordNetwork
	for i, nord_num := range nord_num_list {
		if i == 0 {
			start = nord_num
			continue
		}

		end = nord_num
		new_nord = ConstructorChordNord(start, end)
		chord_network = chord_network.AddNode(new_nord)

		start = nord_num
	}

	return chord_network
}
