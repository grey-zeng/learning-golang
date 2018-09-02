package main

import (
	"fmt"
	"sort"
)

func main() {

	var prereqs = map[string][]string{
		"algorithms": {"data structres"},
		"calculus":   {"linear algebra"},
		"compilers": {
			"data structres",
			"formal languages",
			"computer organization",
		},
		"data structres":        {"discrete math"},
		"databases":             {"data structres"},
		"discrete math":         {"intro to programming"},
		"formal languages":      {"discrete math"},
		"networks":              {"operating systems"},
		"operating systems":     {"data structres", "computer organization"},
		"programming languages": {"data structres", "computer organization"},
	}

	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	// 可以做类型推导来简化
	var visitAll func(items []string)
	visitAll = func(iterms []string) {
		for _, item := range iterms {
			if !seen[item] {
				seen[item] = true
				// 先访问
				visitAll(m[item])
				// 没问题再进行塞入
				order = append(order, item)
			}
		}
	}
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	// 只是为了能统一输出
	sort.Strings(keys)
	visitAll(keys)
	return order
}
