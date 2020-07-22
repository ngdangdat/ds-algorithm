class Queue(object):
    class Node(object):
        def __init__(self, val):
            self.val = val
            self.next_node = None

    def __init__(self):
        self.first = None
        self.last = None

    def add(self, item):
        node = Queue.Node(val=item)
        if self.first is None:
            self.first = node
            self.last = node
        else:
            self.last.next_node = node
            self.last = node

    def remove(self):
        if self.is_empty():
            return
        if self.first.next_node is None:
            self.last = None
        self.first = self.first.next_node

    def peek(self):
        if self.first is None:
            return None
        val = self.first.val

        return val

    def is_empty(self):
        return self.first is None
