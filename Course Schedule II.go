//210. Course Schedule II
//https://leetcode.com/problems/course-schedule-ii/
func findOrder(numCourses int, prerequisites [][]int) []int {
	graph, incomings := buildMap(numCourses, prerequisites)

	var starts []int
	// loop over our incomings list to check which nodes don't have previous requeriments and add them to a new list
	for i, incoming := range incomings {
		if incoming == 0 {
			starts = append(starts, i)
		}
	}

	var res []int
	var visitNum int
	// this loop will iterate over our possible starting nodes
	for _, start := range starts {
		// create a slice as a queue with our starts
		queue := []int{start}
		// add our starts to the result
		res = append(res, start)
		//add 1 to the number of nodes visited for each node in our starts
		visitNum += 1
		//now we iterate over each node in our queue bigger than 0
		for len(queue) > 0 {
			// we slice our node out of the queue
			node := queue[0]
			queue = queue[1:]
			// we check for the node in our graph
			neighbors, ok := graph[node]
			if ok {
				// we iterate over the neighbors of the node
				for _, neighbor := range neighbors {
					// we reduce the incoming of each neighbor by 1
					incomings[neighbor] -= 1
					// once the incomings for that neighbor reach 0, we can add it to our queue as one of our starting nodes
					if incomings[neighbor] == 0 {
						queue = append(queue, neighbor)
						res = append(res, neighbor)
						visitNum += 1
					}
				}
			}
		}
	}
	// finally we check our node count to see if we did go through every course
	if visitNum == numCourses {
		return res
	}
	//else we return an empty array!
	return []int{}
}

func buildMap(numCourses int, prerequisites [][]int) (map[int][]int, []int) {
	// graph will map nodes to their neightbors
	graph := make(map[int][]int)
	// incomings will keep track of the number of required prerequisites for a node
	incomings := make([]int, numCourses)
	for _, preq := range prerequisites {
		// check if node is in graph
		if _, ok := graph[preq[1]]; !ok {
			// add node to graph
			graph[preq[1]] = []int{preq[0]}
		} else {
			// add neightbor to node in graph
			graph[preq[1]] = append(graph[preq[1]], preq[0])
		}
		// add one to the count of incomings to that node
		incomings[preq[0]] += 1
	}
	return graph, incomings
}