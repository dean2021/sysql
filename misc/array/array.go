package array

// Get 从给定的数组中获取指定索引的元素。
// 如果索引在数组的有效范围内，则返回该索引处的元素；
// 否则，返回该类型的零值。
//
// 参数：
// arr []T - 任意类型的切片。
// index int - 要获取元素的索引。
//
// 返回值：
// T - 索引处的元素或该类型的零值。
func Get[T any](arr []T, index int) T {
	// 检查索引是否在数组的有效范围内
	if index >= 0 && index < len(arr) {
		return arr[index]
	}

	// 索引不在有效范围内时，返回该类型的零值
	var null T
	return null
}
