class Node(object):
    def __init__(self, key, val):
        self.key = key
        self.val = val
        self.prev_node = None
        self.next_node = None


class LRUCache(object):
    def __init__(self, capacity):
        self.capacity = capacity
        self.head = None
        self.tail = None
        self.size = 0
        self.node_map = dict()

    def use_node(self, node):
        if node is self.head:
            # do nothing
            return

        if node.prev_node is not None:
            node.prev_node.next_node = node.next_node
        if node.next_node is not None:
            node.next_node.prev_node = node.prev_node

        if node is self.tail:
            node.prev_node.next_node = None
            self.tail = node.prev_node

        # move node to head
        self.head.prev_node = node
        node.next_node = self.head
        node.prev_node = None
        self.head = node

    def get(self, key):
        node = self.node_map.get(key)
        if node:
            self.use_node(node)
            return node.val
        return -1

    def put(self, key, val):
        node = None
        if key in self.node_map:
            node = self.node_map[key]
            node.val = val
            self.node_map[key] = node
        else:
            node = Node(key, val)
            self.node_map[key] = node

            if self.size < self.capacity:
                if self.size == 0:
                    self.head = node
                    self.tail = node

                self.size += 1
            elif self.size == self.capacity:
                tail_key = self.tail.key

                if self.size == 1:
                    self.head = node
                    self.tail = node
                else:
                    self.tail = self.tail.prev_node
                    self.tail.next_node = None

                del self.node_map[tail_key]
        self.use_node(node)


if __name__ == "__main__":
    """
    ["LRUCache","put","put","get","put","get","put","get","get","get"]
    [[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]
    """
    CAPACITY = 2
    lru = LRUCache(CAPACITY)
    lru.put(1, 1)
    lru.put(2, 2)
    print(lru.get(1))
    lru.put(3, 3)
    print(lru.get(2))
    lru.put(4, 4)
    print(lru.get(1))
    print(lru.get(3))
    print(lru.get(4))
