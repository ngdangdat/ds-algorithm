#include "iostream"

using namespace std;

struct ListNode {
  int val;
  ListNode *next;
  ListNode(): val(0), next(nullptr) {}
  ListNode(int x): val(x), next(nullptr) {}
  ListNode(int x, ListNode *next): val(x), next(next) {}
};


class Solution {
  public:
    ListNode* reverseLList(ListNode* head) {
      ListNode* prev = nullptr;
      ListNode* current = head;
      ListNode* next = nullptr;

      while (current != nullptr) {
        next = current->next;
        current->next = prev;
        prev = current;
        current = next;
      }

      return prev;
    }
};

void iterateLinkedList(ListNode* head) {
  ListNode* itr_head = head;
  while (itr_head != nullptr) {
    cout << itr_head->val << " ";
    itr_head = itr_head->next;
  }
  cout << endl;
}

int main() {
  // 1, 2, 3, 4, 5
  int inputs[] = {1, 2, 3, 4, 5};
  ListNode* head = new ListNode(inputs[0]);
  ListNode* cur_head = head;
  for (int i = 1; i < (sizeof(inputs) / sizeof(int)); i++) {
    ListNode* new_head = new ListNode(inputs[i]);
    cur_head->next = new_head;
    cur_head = new_head;
  }

  cout << "Standard iteration." << endl;
  iterateLinkedList(head);

  cout << "Reversed iteration." << endl;
  Solution* sol = new Solution();
  ListNode* new_head = sol->reverseLList(head);
  iterateLinkedList(new_head);
  return 0;
}

