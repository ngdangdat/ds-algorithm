package main

import "testing"

func TestReversedLinkedList(t *testing.T) {
	input := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: nil}}
	expected := &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: nil}}
	output := reverseList(input)
	equal := expected.IsEqual(output)
	if !equal {
		t.Errorf("Not equal")
	}
}

func TestReverseEmptyList(t *testing.T) {
	input := (*ListNode)(nil)
	expected := (*ListNode)(nil)
	output := reverseList(input)
	if output != expected {
		t.Errorf("Expected nil, got %v", output)
	}
}

func TestReverseSingleNode(t *testing.T) {
	input := &ListNode{Val: 42, Next: nil}
	expected := &ListNode{Val: 42, Next: nil}
	output := reverseList(input)
	equal := expected.IsEqual(output)
	if !equal {
		t.Errorf("Single node reversal failed")
	}
}

func TestReverseLongerList(t *testing.T) {
	input := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}}
	expected := &ListNode{Val: 5, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: nil}}}}}
	output := reverseList(input)
	equal := expected.IsEqual(output)
	if !equal {
		t.Errorf("Longer list reversal failed")
	}
}

func TestReverseTwoNodeDifferent(t *testing.T) {
	input := &ListNode{Val: 5, Next: &ListNode{Val: 3, Next: nil}}
	expected := &ListNode{Val: 3, Next: &ListNode{Val: 5, Next: nil}}
	output := reverseList(input)
	equal := expected.IsEqual(output)
	if !equal {
		t.Errorf("Two node reversal failed")
	}
}

func TestReverseListRecursion(t *testing.T) {
	input := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: nil}}
	expected := &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: nil}}
	output := reverseListRecursion(input)
	equal := expected.IsEqual(output)
	if !equal {
		t.Errorf("Recursive reversal failed")
	}
}

func TestReverseEmptyListRecursion(t *testing.T) {
	input := (*ListNode)(nil)
	expected := (*ListNode)(nil)
	output := reverseListRecursion(input)
	if output != expected {
		t.Errorf("Expected nil, got %v", output)
	}
}

func TestReverseSingleNodeRecursion(t *testing.T) {
	input := &ListNode{Val: 42, Next: nil}
	expected := &ListNode{Val: 42, Next: nil}
	output := reverseListRecursion(input)
	equal := expected.IsEqual(output)
	if !equal {
		t.Errorf("Single node recursive reversal failed")
	}
}

func TestReverseLongerListRecursion(t *testing.T) {
	input := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}}
	expected := &ListNode{Val: 5, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: nil}}}}}
	output := reverseListRecursion(input)
	equal := expected.IsEqual(output)
	if !equal {
		t.Errorf("Longer list recursive reversal failed")
	}
}
