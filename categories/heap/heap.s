.section .rodata
fmt_peek:
    .string "peek: %ld\n"
fmt_extract:
    .string "extract: %ld\n"
fmt_empty:
    .string "extract: empty\n"

.section .text
.globl heap_init
.type heap_init, @function
# void heap_init(struct heap *heap)
# struct heap layout: uint64_t size; int64_t data[32]
heap_init:
    movq $0, (%rdi)
    ret

.globl heap_insert
.type heap_insert, @function
# int heap_insert(struct heap *heap, int64_t value)
# Returns 0 on success and -1 when the heap is full.
heap_insert:
    movq (%rdi), %rcx
    cmpq $32, %rcx
    jae .insert_full

    leaq 8(%rdi), %rdx
    movq %rsi, (%rdx,%rcx,8)
    incq %rcx
    movq %rcx, (%rdi)
    decq %rcx

.insert_sift_up:
    testq %rcx, %rcx
    jz .insert_done

    movq %rcx, %r8
    decq %r8
    shrq $1, %r8
    movq (%rdx,%r8,8), %r9
    cmpq %rsi, %r9
    jge .insert_done

    movq %r9, (%rdx,%rcx,8)
    movq %rsi, (%rdx,%r8,8)
    movq %r8, %rcx
    jmp .insert_sift_up

.insert_done:
    xorl %eax, %eax
    ret

.insert_full:
    movl $-1, %eax
    ret

.globl heap_peek
.type heap_peek, @function
# int heap_peek(struct heap *heap, int64_t *value)
# Returns 0 on success and -1 when the heap is empty.
heap_peek:
    cmpq $0, (%rdi)
    je .peek_empty
    movq 8(%rdi), %rax
    movq %rax, (%rsi)
    xorl %eax, %eax
    ret

.peek_empty:
    movl $-1, %eax
    ret

.globl heap_extract
.type heap_extract, @function
# int heap_extract(struct heap *heap, int64_t *value)
# Removes the maximum value, writes it to *value, and returns 0.
# Returns -1 without changing the heap when it is empty.
heap_extract:
    movq (%rdi), %rcx
    testq %rcx, %rcx
    jz .extract_empty

    movq 8(%rdi), %r8
    movq %r8, (%rsi)

    decq %rcx
    movq %rcx, (%rdi)
    jz .extract_done

    leaq 8(%rdi), %rdx
    movq (%rdx,%rcx,8), %r9
    movq %r9, (%rdx)
    xorl %r10d, %r10d

.extract_sift_down:
    leaq 1(%r10,%r10), %r11
    cmpq %rcx, %r11
    jae .extract_done

    movq %r11, %rax
    leaq 1(%r11), %r8
    cmpq %rcx, %r8
    jae .right_child_missing
    movq (%rdx,%r8,8), %rsi
    cmpq (%rdx,%r11,8), %rsi
    jle .right_child_missing
    movq %r8, %rax

.right_child_missing:
    movq (%rdx,%rax,8), %rsi
    cmpq %r9, %rsi
    jle .extract_done

    movq %rsi, (%rdx,%r10,8)
    movq %r9, (%rdx,%rax,8)
    movq %rax, %r10
    jmp .extract_sift_down

.extract_done:
    xorl %eax, %eax
    ret

.extract_empty:
    movl $-1, %eax
    ret

.globl main
.type main, @function
.extern printf
main:
    pushq %rbp
    movq %rsp, %rbp

    leaq heap(%rip), %rdi
    call heap_init

    leaq heap(%rip), %rdi
    movq $10, %rsi
    call heap_insert
    movq $4, %rsi
    call heap_insert
    movq $15, %rsi
    call heap_insert
    movq $20, %rsi
    call heap_insert
    movq $8, %rsi
    call heap_insert
    movq $12, %rsi
    call heap_insert

    leaq heap(%rip), %rdi
    leaq result(%rip), %rsi
    call heap_peek
    testl %eax, %eax
    jnz .main_extract
    leaq fmt_peek(%rip), %rdi
    movq result(%rip), %rsi
    xorl %eax, %eax
    call printf

.main_extract:
    leaq heap(%rip), %rdi
    leaq result(%rip), %rsi
    call heap_extract
    testl %eax, %eax
    jnz .main_empty
    leaq fmt_extract(%rip), %rdi
    movq result(%rip), %rsi
    xorl %eax, %eax
    call printf
    jmp .main_extract

.main_empty:
    leaq fmt_empty(%rip), %rdi
    xorl %eax, %eax
    call printf

    xorl %eax, %eax
    leave
    ret

.section .bss
.align 8
heap:
    .zero 264
result:
    .zero 8

.section .note.GNU-stack,"",@progbits
