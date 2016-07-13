package str

func convert(s string, numRows int) string {
	length := len(s)
	if length <= numRows {
		return s
	}
	if numRows == 1 {
		return s
	}
	z := (numRows) / 2
	offset := numRows + 1
	offsetNeighbor := numRows/2 + 1
	tmp := make([]byte, 0, length)
	j := 0
	for cursor := 0; cursor < numRows; cursor++ {
		i := cursor
		var currerntOffset int
		if i != z && numRows != 2 {
			currerntOffset = offset
		} else {
			currerntOffset = offsetNeighbor
		}
		for i < length {
			// tmp[j] = s[i]
			tmp = append(tmp, s[i])
			j++
			i = i + currerntOffset
		}
	}
	return string(tmp)
}
