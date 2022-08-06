package algorithm

import "sync"

// QuickSort 简单版本 对重复数据无优化
func QuickSort(arr []int, start, end int) {
	// 判断是否需要排序，小于1则不需要排序
	if end-start < 1 {
		return
	}

	// 开始从左向右找，找到大于目标元素则开始从右向左找
	j := end
	for i := start + 1; i < j; i++ {
		// 从左向右，找大于目标值
		if arr[i] > arr[start] {
			for ; j > i; j-- {
				// 从右向左，找小于目标值
				if arr[j] < arr[start] {
					arr[i], arr[j] = arr[j], arr[i]
					break
				}
			}
		}
	}
	// 判断是否大于
	// 当前 arr[j] 可能会出现两种情况，arr[j] > 目标值。这种情况不能直接进行替换
	// 所以我们使用if判断一下
	if arr[j] > arr[start] {
		j--
	}
	// 交换
	arr[j], arr[start] = arr[start], arr[j]
	var wg sync.WaitGroup
	// 多线程优化
	wg.Add(1)
	go func() {
		QuickSort(arr, start, j-1)
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		QuickSort(arr, j+1, end)
		wg.Done()
	}()
	wg.Wait()
}

// QuickSort1 优化数据 ，增加内存使用量
func QuickSort1(arr []int, start, end int) {
	// 判断是否需要排序，小于1则不需要排序
	if end-start < 1 {
		return
	}

	// 当前数组用于记录与目标值相同的元素下标
	var recordSameValue []int

	// 开始从左向右找，找到大于目标元素则开始从右向左找
	j := end
	for i := start + 1; i < j; i++ {
		// 从左向右，找大于目标值
		if arr[i] > arr[start] {
			for ; j > i; j-- {
				// 从右向左，找小于目标值
				if arr[j] <= arr[start] {
					arr[i], arr[j] = arr[j], arr[i]
					break
				}
			}
		}
		// 判断与目标值是否相等
		if arr[i] == arr[start] {
			recordSameValue = append(recordSameValue, i)
		}
	}
	// 判断是否大于
	// 当前 arr[j] 可能会出现两种情况，arr[j] > 目标值。这种情况不能直接进行替换
	// 所以我们使用if判断一下
	if arr[j] > arr[start] {
		j--
	}
	// 交换
	arr[j], arr[start] = arr[start], arr[j]

	// 处理相同数据
	for i, x := 0, j-len(recordSameValue); x < j; x++ {
		// 如果第一个数重复，我们对其进行替换
		if arr[start] == arr[j] {
			// 找到与目标数不相同的值
			if arr[x] != arr[j] {
				// 替换
				arr[x], arr[start] = arr[start], arr[x]
			}
			continue
		}
		// 由于我们的相同元素下标是有序的，并且小于 j
		// 所以，当recordSame[i] > j- len(recordSame)时
		// 剩下的元素已经全部在正确的位置，我们不需要继续判断
		if recordSameValue[i] > j-len(recordSameValue) {
			break
		}
		// 找到不同的元素，越过已经存在于正确位置的相同元素
		if arr[x] != arr[j] {
			arr[x], arr[recordSameValue[i]] = arr[recordSameValue[i]], arr[x]
			i++
		}
	}

	var wg sync.WaitGroup
	// 多线程优化
	wg.Add(1)
	go func() {
		QuickSort1(arr, start, j-len(recordSameValue)-1)
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		QuickSort1(arr, j+1, end)
		wg.Done()
	}()
	wg.Wait()
}
