"""
1 2 4
9 1 2
3 1 5
"""

mat = []
mat.append([1, 2, 4])
mat.append([9, 8, 2])
mat.append([3, 1, 5])

dp = [[0 for _ in range(3)] for _ in range(3)]

dp[0][0] = mat[0][0]
for i in range(0, 3):
    for j in range(0, 3):
        if j > 0 and dp[i][j - 1] + mat[i][j] > dp[i][j]:
            dp[i][j] = dp[i][j - 1] + mat[i][j]
        if i > 0 and dp[i - 1][j] + mat[i][j] > dp[i][j]:
            dp[i][j] = dp[i - 1][j] + mat[i][j]

print(dp[2][2])
