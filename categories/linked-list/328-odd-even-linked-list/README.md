# 328. Odd Even Linked List

**Difficulty**: Medium

## Problem Description
Given the head of a singly linked list, group all the nodes with odd indices together followed by the nodes with even indices, and return the reordered list.

## Approach
Use odd and even pointers, and a node to save the head of the even list.

## Key Points
- First node is odd, second node is even, and so on
- Maintain the relative order within odd and even groups
- Connect odd tail to even head at the end

## Time Complexity
O(n) where n is the number of nodes

## Space Complexity
O(1)