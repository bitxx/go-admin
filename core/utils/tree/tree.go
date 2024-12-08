package tree

type nodeInfo[T any] struct {
	item     *T
	parentId int64
	id       int64
}

// GenTree 只通过parentId来构建树
func GenTree[T any](menuList *[]T, getId func(T) int64, getParentId func(T) int64, setChildren func(*T, []*T)) []*T {
	// 缓存节点信息

	nodeInfos := make([]nodeInfo[T], len(*menuList))
	menuMap := make(map[int64][]*T)

	for index := range *menuList {
		item := &(*menuList)[index]
		id := getId(*item)
		parentId := getParentId(*item)

		nodeInfos[index] = nodeInfo[T]{item: item, parentId: parentId, id: id}
		menuMap[parentId] = append(menuMap[parentId], item)
	}

	// 确定所有根节点
	allIds := make(map[int64]struct{})
	for _, info := range nodeInfos {
		allIds[info.id] = struct{}{}
	}

	rootIds := []int64{}
	for parentId, _ := range menuMap {
		if parentId <= 0 || !existsInMap(allIds, parentId) {
			rootIds = append(rootIds, parentId)
		}
	}

	// 构建树
	rootItems := make([]*T, 0, len(rootIds))
	stack := make([]*T, 0, len(*menuList))

	for _, rootId := range rootIds {
		for _, item := range menuMap[rootId] {
			rootItems = append(rootItems, item)
			stack = append(stack, item)
		}
	}

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if children, exists := menuMap[getId(*current)]; exists {
			setChildren(current, children)
			stack = append(stack, children...)
		}
	}

	if rootItems == nil {
		rootItems = []*T{}
	}
	return rootItems
}

func existsInMap(m map[int64]struct{}, key int64) bool {
	_, exists := m[key]
	return exists
}

// GenTree 通过parentIds构建，做了存储和性能优化，但代码有问题，暂不使用
/*func GenTree[T any](
	menuList *[]T,
	getId func(T) int64,
	getParentIds func(T) string,
	setChildren func(*T, []*T),
) []*T {
	menuMap := make(map[int64][]*T)
	allIds := make(map[int64]*T)               // 用于快速查找节点
	parsedParentIds := make(map[int64][]int64) // 缓存解析后的 ParentIds
	var rootItems []*T

	// 构建 menuMap 和 allIds
	for index := range *menuList {
		item := &(*menuList)[index]
		allIds[getId(*item)] = item

		// 使用 ParentIds 构建关系
		parentIdsStr := getParentIds(*item)
		itemId := getId(*item)
		parentIds, exists := parsedParentIds[itemId]
		if !exists {
			parentIds = parseParentIds(parentIdsStr)
			parsedParentIds[itemId] = parentIds
		}
		if len(parentIds) > 0 {
			for _, parentId := range parentIds {
				menuMap[parentId] = append(menuMap[parentId], item)
			}
		}
	}

	// 找出根节点
	for _, item := range *menuList {
		parentIdsStr := getParentIds(item)
		parentIds := parseParentIds(parentIdsStr)
		if isRootNode(parentIds) {
			rootItems = append(rootItems, &item)
		}
	}

	// 构建树
	stack := rootItems

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if children, exists := menuMap[getId(*current)]; exists {
			setChildren(current, children)
			stack = append(stack, children...)
		}
	}

	if len(rootItems) == 0 {
		return []*T{}
	}
	return rootItems
}

func parseParentIds(parentIdsStr string) []int64 {
	if parentIdsStr == "" {
		return nil
	}
	parts := strings.Split(strings.TrimSpace(parentIdsStr), ",")
	result := make([]int64, 0, len(parts))
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		if id, err := strconv.ParseInt(part, 10, 64); err == nil {
			result = append(result, id)
		}
	}
	return result
}

func isRootNode(parentIds []int64) bool {
	return len(parentIds) == 0 || parentIds[len(parentIds)-1] == 0
}
*/
