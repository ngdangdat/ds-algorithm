.PHONY: all run clean

all: heap

heap: heap.s
	gcc -no-pie -o $@ $<

run: heap
	./heap

clean:
	rm -f heap
