memo = dict()


def fib_naive(n):
    if n == 1 or n == 2:
        return 1
    else:
        return fib_naive(n - 1) + fib_naive(n - 2)


def fib_2(n):
    if n in memo.keys():
        return memo[n]
    if n == 1 or n == 2:
        result = 1
    else:
        result = fib_2(n - 1) + fib_2(n - 2)
    memo[n] = result
    return result


def fib_bottom(n):
    if n == 1 or n == 2:
        return 1
    memo = [1, 1]
    for i in range(2, n):
        next = memo[i - 1] + memo[i - 2]
        memo.append(next)
    return memo[n - 1]


if __name__ == '__main__':
    # print(fib_naive(50))
    print(fib_2(50))
    print(fib_bottom(50))
