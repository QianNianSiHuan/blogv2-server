package sliceTool

func DeleteSliceElems[T comparable](s []T, elem ...T) []T {
	var result []T
	elemMap := make(map[T]struct{}, len(elem))
	for _, e := range elem {
		elemMap[e] = struct{}{}
	}

	for _, v := range s {
		if _, found := elemMap[v]; !found {
			result = append(result, v)
		}
	}
	return result
}
