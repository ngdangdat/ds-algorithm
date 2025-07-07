N = 3
mat = [[0 for _ in range(N)] for _ in range(N)]
mat[N - 1][N - 1] = 1

for i in range(N - 1, -1, -1):
    for j in range(N - 1, -1, -1):
        if i < N - 1:
            mat[i][j] += mat[i + 1][j]
        if j < N - 1:
            mat[i][j] += mat[i][j + 1]

print("result:", mat[0][0])
