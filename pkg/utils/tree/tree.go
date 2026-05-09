package tree

func GenTree[T any](menuList *[]T, getId func(T) int64, getParentId func(T) int64, setChildren func(*T, []*T)) []*T {
	// 创建一个 map 来存储每个 ParentId 对应的子菜单
	menuMap := make(map[int64][]*T)
	allIds := make(map[int64]struct{})
	for index := range *menuList {
		item := &(*menuList)[index]
		menuMap[getParentId(*item)] = append(menuMap[getParentId(*item)], item)
		allIds[getId(*item)] = struct{}{}
	}

	// 汇总所有根节点
	var rootIds []int64
	for id := range menuMap {
		// 如果 parentId <= 0 并且存在子节点，则标记为根节点
		if id <= 0 && len(menuMap[id]) > 0 {
			rootIds = append(rootIds, id)
			continue
		}

		// 如果 parentId 在节点列表中不存在，并且存在子节点，则标记为根节点
		if _, exists := allIds[id]; !exists && len(menuMap[id]) > 0 {
			rootIds = append(rootIds, id)
		}
	}

	// 汇总根节点的数据
	var rootItems []*T
	var stack []*T

	for _, rootId := range rootIds {
		for _, item := range menuMap[rootId] {
			stack = append(stack, item)
			rootItems = append(rootItems, item)
		}
	}

	// 使用栈模拟递归来构建树形结构
	for len(stack) > 0 {
		// 弹出栈顶元素
		currentItem := stack[len(stack)-1]
		stack = stack[:len(stack)-1] // 移除栈顶元素

		// 获取当前节点的子节点
		if children, exists := menuMap[getId(*currentItem)]; exists {
			// 给当前菜单项设置子菜单
			setChildren(currentItem, children)

			// 将子节点加入栈中进行处理
			stack = append(stack, children...)
		}
	}
	return rootItems
}
