# 133. Clone Graph

## Manual test cases

The `main` function includes assertions for:

1. Empty graph: pass `nullptr` and expect `nullptr`.
2. Single node with no neighbors: `[[]]`.
3. Four-node cycle: `[[2,4],[1,3],[2,4],[1,3]]`, including a distinct cloned root.

## Run

```bash
g++ -std=c++17 -Wall -Wextra -pedantic main.cpp -o clone-graph && ./clone-graph
```
