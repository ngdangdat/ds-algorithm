#include <cassert>
#include <unordered_map>
#include <vector>

using namespace std;

class Node {
public:
    int val;
    vector<Node*> neighbors;

    Node() : val(0), neighbors() {}
    Node(int _val) : val(_val), neighbors() {}
    Node(int _val, vector<Node*> _neighbors)
        : val(_val), neighbors(_neighbors) {}
};

class Solution {
public:
    Node* _cloneGraph(Node* node, unordered_map<Node*, Node*>* old_new_map) {
      if (node == nullptr) {
        return node;
      }
      auto it = old_new_map->find(node);
      if (it != old_new_map->cend()) {
        return it->second;
      }
      Node* n = new Node(node->val);
      old_new_map->insert({ node, n });
      for (Node* _neighbor : node->neighbors) {
        Node* child = _cloneGraph(_neighbor, old_new_map);
        n->neighbors.push_back(child);
      }
      return n;
    }

    Node* cloneGraph(Node* node) {
      unordered_map<Node*, Node*> old_new_map = {};
      return _cloneGraph(node, &old_new_map);
    }
};

int main() {
    Solution solution;

    // Empty graph.
    assert(solution.cloneGraph(nullptr) == nullptr);

    // Single node with no neighbors.
    Node single(1);
    Node* single_clone = solution.cloneGraph(&single);
    assert(single_clone != nullptr);
    assert(single_clone != &single);
    assert(single_clone->val == 1);
    assert(single_clone->neighbors.empty());

    // Four-node cycle: [[2,4],[1,3],[2,4],[1,3]].
    Node node1(1);
    Node node2(2);
    Node node3(3);
    Node node4(4);
    node1.neighbors = {&node2, &node4};
    node2.neighbors = {&node1, &node3};
    node3.neighbors = {&node2, &node4};
    node4.neighbors = {&node1, &node3};

    Node* cycle_clone = solution.cloneGraph(&node1);
    assert(cycle_clone != nullptr);
    assert(cycle_clone != &node1);
    assert(cycle_clone->val == 1);
    assert(cycle_clone->neighbors.size() == 2);

    return 0;
}
