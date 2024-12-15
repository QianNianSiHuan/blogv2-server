package unitls

func InList[T comparable](Key T, list []T) bool {
	for _, s := range list {
		if Key == s {
			return true
		}
	}
	return false
}
