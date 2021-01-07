package algorithm

import "fmt"

//输入：
// equations = [["a","b"],["b","c"]], values = [2.0,3.0],
// queries = [["a","c"],["b","a"],["a","e"],["a","a"],["x","x"]]
//输出：[6.00000,0.50000,-1.00000,1.00000,-1.00000]
//解释：
//条件：a / b = 2.0, b / c = 3.0
//问题：a / c = ?, b / a = ?, a / e = ?, a / a = ?, x / x = ?
//结果：[6.0, 0.5, -1.0, 1.0, -1.0 ]
// a =2.0b b=3.0c  => a=6.0c
// a,d a,e => e,d
func calcEquation1(equations [][]string, values []float64, queries [][]string) []float64 {
	// 给方程组中的每个变量编号
	id := map[string]int{}
	for _, eq := range equations {
		a, b := eq[0], eq[1]
		if _, has := id[a]; !has {
			id[a] = len(id)
		}
		if _, has := id[b]; !has {
			id[b] = len(id)
		}
	}

	// 建图
	type edge struct {
		to     int
		weight float64
	}
	graph := make([][]edge, len(id))
	for i, eq := range equations {
		v, w := id[eq[0]], id[eq[1]]
		graph[v] = append(graph[v], edge{w, values[i]})
		graph[w] = append(graph[w], edge{v, 1 / values[i]})
	}

	bfs := func(start, end int) float64 {
		ratios := make([]float64, len(graph))
		ratios[start] = 1
		queue := []int{start}
		for len(queue) > 0 {
			v := queue[0]
			queue = queue[1:]
			if v == end {
				return ratios[v]
			}
			for _, e := range graph[v] {
				if w := e.to; ratios[w] == 0 {
					ratios[w] = ratios[v] * e.weight
					queue = append(queue, w)
				}
			}
		}
		return -1
	}

	ans := make([]float64, len(queries))
	for i, q := range queries {
		start, hasS := id[q[0]]
		end, hasE := id[q[1]]
		if !hasS || !hasE {
			ans[i] = -1
		} else {
			ans[i] = bfs(start, end)
		}
	}
	return ans
}

var allEquation = map[string]map[string]float64{}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	allEquation = map[string]map[string]float64{}
	l := len(equations)
	for i := 0; i < l; i++ {
		key := equations[i][0]
		if _, ok := allEquation[key]; !ok {
			allEquation[key] = map[string]float64{}
		}
		allEquation[key][equations[i][1]] = values[i]
	}

	fmt.Println(allEquation)

	res := make([]float64, len(queries))
	for i := range queries {
		key := queries[i][0]
		target := queries[i][1]
		res[i] = -1.0
		if _, ok := allEquation[key]; ok {
			if key == target {
				res[i] = 1.0
			} else {
				res[i] = backCalcEquation(allEquation[key], target, 1.0)
			}
		}
		if res[i] == -1.0 {
			if _, ok := allEquation[target]; res[i] == -1.0 && ok {
				res[i] = 1 / backCalcEquation(allEquation[target], key, 1.0)
			}
		}
		if res[i] == -1.0 {
			for _, item := range allEquation {
				if _, ok := item[key]; ok {
					if _, ok := item[target]; ok {
						res[i] = item[target] / item[key]
					}
				}
			}
		}
	}
	return res
}

func backCalcEquation(m map[string]float64, key string, res float64) float64 {
	for k := range m {
		if k == key {
			return res * m[k]
		}
		res1 := -1.0
		if _, ok := allEquation[k]; ok {
			res1 = backCalcEquation(allEquation[k], key, res*m[k])
			if res1 != -1.0 {
				return res1
			}
		}
	}
	return -1.0
}
