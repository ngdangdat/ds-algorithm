# Heap (x86-64 GAS)

This directory contains a max-heap implemented in x86-64 GNU assembler using AT&T syntax.

## Build and run

```sh
make
./heap
```

The demo inserts six values, prints the maximum, then extracts values until the heap is empty.

## Heap API

The heap stores up to 32 signed 64-bit integers in an array-backed binary heap. Its structure is:

- offset `0`: `uint64_t size`
- offset `8`: `int64_t data[32]`

The exported routines use the System V AMD64 calling convention:

- `heap_init(struct heap *heap)` initializes an empty heap.
- `heap_insert(struct heap *heap, int64_t value)` returns `0` on success or `-1` when full.
- `heap_peek(struct heap *heap, int64_t *value)` returns `0` on success or `-1` when empty.
- `heap_extract(struct heap *heap, int64_t *value)` removes the maximum and returns `0`, or returns `-1` when empty.

Insertion and extraction take `O(log n)` time; peeking takes `O(1)` time. The heap uses `O(n)` storage.
