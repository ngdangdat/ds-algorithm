# Linked List Problems

This directory contains solutions to various linked list algorithm problems.

## Common Helpers

The `000-common-helpers/` folder contains reusable utilities for linked list operations:

### ListNode Structure
```go
type ListNode struct {
    Val  int
    Next *ListNode
}
```

### Available Methods
- `IsEqual(head *ListNode) bool` - Compare two linked lists for equality
- `Print()` - Print linked list values to console
- `String() string` - String representation of the linked list
- `ListNodeFromIntList(l []int) *ListNode` - Create linked list from integer slice

### Usage
Copy the helper functions from `000-common-helpers/node.go` into your solution files, or reference them for testing and debugging your linked list implementations.

## Problem Solutions
Each numbered folder contains a complete solution with:
- `main.go` - The main solution implementation
- `node.go` - Local copy of helper functions
- `main_test.go` - Test cases
- `go.mod` - Go module file