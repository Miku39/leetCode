package main

import (
	"sort"
)

// 692. Top K Frequent Words
// https://leetcode.com/problems/top-k-frequent-words/

type State struct {
	word  string
	count int
}

func topKFrequent(words []string, k int) []string {
	wordCountMap := make(map[string]int)
	for _, v := range words {
		wordCountMap[v]++
	}

	stateArray := []State{}
	for j, w := range wordCountMap {
		stateArray = append(stateArray, State{j, w})
	}

	sort.Slice(stateArray, func(i, j int) bool {
		if stateArray[i].count == stateArray[j].count {
			return stateArray[i].word < stateArray[j].word
		}

		return stateArray[i].count > stateArray[j].count
	})

	result := []string{}
	for _, x := range stateArray[:k] {
		result = append(result, x.word)
	}

	return result
}
