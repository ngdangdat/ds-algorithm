# 148. Sort List

**Difficulty**: Medium

## Problem Description
Given the head of a linked list, return the list after sorting it in ascending order.

## Approach
Slow and fast pointers + merge sort

## Key Points
- Use merge sort algorithm for O(n log n) time complexity
- Find middle using slow/fast pointers to split the list
- Recursively sort both halves
- Merge the sorted halves

## Time Complexity
O(n log n) where n is the number of nodes

## Space Complexity
O(1) for iterative approach, O(log n) for recursive approach