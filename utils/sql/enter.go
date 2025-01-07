package sql

import "fmt"

func ConvertSliceSql(list []uint) (s string) {
	s += "("
	for k, v := range list {
		if k == len(list)-1 {
			s += fmt.Sprintf("%d", v)
			break
		}
		s += fmt.Sprintf("%d,", v)
	}
	s += ")"
	return
}
func ConvertSliceOrderSql(list []uint) (s string) {
	for k, v := range list {
		if k == len(list)-1 {
			s += fmt.Sprintf("id = %d desc ", v)
			break
		}
		s += fmt.Sprintf("id = %d desc ,", v)
	}
	return
}
