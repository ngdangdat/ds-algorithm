package main

import (
	"testing"
)

func TestEdgeCase(t *testing.T) {
	selfPointingNode := &ListNode{Val: 1}
	selfPointingNode.Next = selfPointingNode
	tests := []struct {
		name     string
		list     *ListNode
		expected bool
	}{
		{
			name:     "nil list",
			list:     nil,
			expected: false,
		},
		{
			name:     "single node list",
			list:     &ListNode{Val: 1},
			expected: false,
		},
		{
			name:     "cycle self pointing node",
			list:     selfPointingNode,
			expected: true,
		},
	}
	for _, tCase := range tests {
		t.Run(tCase.name, func(t *testing.T) {
			got := hasCycle(tCase.list)
			if got != tCase.expected {
				t.Errorf("failed, expected: %v, got: %v", tCase.expected, got)
			}
		})
	}
}
