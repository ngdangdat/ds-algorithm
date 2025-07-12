# 725. Split Linked List in Parts

**Difficulty**: Medium

## Problem Description
Given the head of a singly linked list and an integer k, split the linked list into k consecutive linked list parts.

## Approach
Find length n, and use n//k and n%k to determine the size for each part.

## Key Points
- Each part should be as equal as possible
- Earlier parts should be larger than later parts if division is not exact
- If there are fewer nodes than k, some parts will be empty (null)

## Time Complexity
O(n) where n is the number of nodes

## Space Complexity
O(1) excluding the output array