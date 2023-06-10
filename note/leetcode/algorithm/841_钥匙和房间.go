package algorithm

func canVisitAllRooms(rooms [][]int) bool {
	n := len(rooms)
	num := 0
	vis := make([]bool, n)
	queue := []int{0}
	vis[0] = true
	for i := 0; i < len(queue); i++ { // 切片在追加的过程中len(queue)是动态变化的
		x := queue[i]
		num++
		for _, it := range rooms[x] {
			if !vis[it] {
				vis[it] = true
				queue = append(queue, it)
			}
		}
	}
	return num == n
}

var (
	canVisitAllRoomsNum int
	existArray          []bool
)

func canVisitAllRooms1(rooms [][]int) bool {
	canVisitAllRoomsNum = 0
	existArray = make([]bool, len(rooms))
	dfs(rooms, 0)
	return canVisitAllRoomsNum == len(rooms)
}

func dfs(rooms [][]int, room int) {
	existArray[room] = true
	canVisitAllRoomsNum++
	for _, item := range rooms[room] {
		if !existArray[item] {
			dfs(rooms, item)
		}
	}
}
