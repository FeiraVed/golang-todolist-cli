package helper

import "sort"

func SortTodolistKey(data map[string]string) []string {
	keys := make([]string, 0, len(data))
	for k := range data {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	return keys
}
