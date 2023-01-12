package datastructure

type TreeConfig struct {
	ID       string
	PID      string
	Children string
}

func (c *TreeConfig) GenerateMapTree(nodes []map[string]interface{}) []map[string]interface{} {
	idMap := make(map[interface{}]interface{})
	jsonTree := make([]map[string]interface{}, 0)

	var id, pid, children string
	if c.ID != "" {
		id = c.ID
	} else {
		id = "id"
	}

	if c.PID != "" {
		pid = c.PID
	} else {
		pid = "pid"
	}

	if c.Children != "" {
		children = c.Children
	} else {
		children = "children"
	}

	for _, node := range nodes {
		for k, v := range node {
			if k == id {
				value, ok := v.(interface{})
				if ok {
					idMap[value] = node
				}
			}
		}
	}

	for _, node := range nodes {
		pid, _ := node[pid].(interface{})
		parent, ok := idMap[pid].(map[string]interface{})
		if ok {
			p, _ := parent[children].([]map[string]interface{})
			parent[children] = append(p, node)
		} else {
			jsonTree = append(jsonTree, node)
		}
	}

	return jsonTree
}
