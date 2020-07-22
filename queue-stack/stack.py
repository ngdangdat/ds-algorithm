class Stack(object):
    class Node(object):
        def __init__(self, val):
            self.val = val
            self.next_node = None

    def __init__(self):
        self.head = None

    def add(self, item):
        node = Stack.Node(val=item)
        node.next_node = self.head
        self.head = node

    def remove(self):
        if not self.is_empty():
            self.head = self.head.next_node

    def peek(self):
        if not self.is_empty():
            return self.head.val

    def is_empty(self):
        return self.head is None
