# 234. Palindrome Linked List

**Difficulty**: Medium

## Problem Description
Given the head of a singly linked list, return true if it is a palindrome or false otherwise.

## Approach
Use slow and fast pointers to cut it into halves. Reverse the second half and compare with the first half.

## Key Points
- Find middle using slow/fast pointers
- Reverse the second half of the list
- Compare first half with reversed second half
- Optional: restore the original list structure

## Time Complexity
O(n) where n is the number of nodes

## Space Complexity
O(1)