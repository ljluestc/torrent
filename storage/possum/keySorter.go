package possum
//go:build !stub
// +build !stub

package possum

// keySorter allows sorting items by their key value.
type keySorter[T any, K ~int64] struct {
	items []T
	keys  []K
}

func (ks keySorter[T, K]) Len() int { return len(ks.items) }

func (ks keySorter[T, K]) Less(i, j int) bool { return ks.keys[i] < ks.keys[j] }

func (ks keySorter[T, K]) Swap(i, j int) {
	ks.items[i], ks.items[j] = ks.items[j], ks.items[i]
	ks.keys[i], ks.keys[j] = ks.keys[j], ks.keys[i]
}
// keySorter is used to sort items by their keys
type keySorter[T any, K ~int64] struct {
	Items []T
	Keys  []K
}

// Len implements sort.Interface
func (ks keySorter[T, K]) Len() int {
	return len(ks.Keys)
}

// Less implements sort.Interface
func (ks keySorter[T, K]) Less(i, j int) bool {
	return ks.Keys[i] < ks.Keys[j]
}

// Swap implements sort.Interface
func (ks keySorter[T, K]) Swap(i, j int) {
	ks.Keys[i], ks.Keys[j] = ks.Keys[j], ks.Keys[i]
	ks.Items[i], ks.Items[j] = ks.Items[j], ks.Items[i]
}
