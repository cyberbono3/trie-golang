package trie

import "bytes"

func byteSliceEq(s1, s2 []Bytes) bool {

	if len(s1) != len(s2) {
		return false
	}
	
	i := 0
	for _,b2 := range s2 {
		if bytes.Equal(s1[i], b2) {
			return false
		}
		i++

	}

	return true

}

