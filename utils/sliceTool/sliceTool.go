package sliceTool

import "cmp"

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

// 切片去重升级版 泛型参数 利用map的key不能重复的特性+append函数  一次for循环搞定
func Unique[T cmp.Ordered](ss []T) []T {
	size := len(ss)
	if size == 0 {
		return []T{}
	}
	newSlices := make([]T, 0) //这里新建一个切片,大于为0, 因为我们不知道有几个非重复数据,后面都使用append来动态增加并扩容
	m1 := make(map[T]byte)
	for _, v := range ss {
		if _, ok := m1[v]; !ok { //如果数据不在map中,放入
			m1[v] = 1                        // 保存到map中,用于下次判断
			newSlices = append(newSlices, v) // 将数据放入新的切片中
		}
	}
	return newSlices
}
