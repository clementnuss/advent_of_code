from functools import cache, reduce
from itertools import combinations
import numpy as np

m = {"#": -1, "O": 1, ".": 0}

# with open("14/example_input") as f:
with open("14/input") as f:
    p = [l.split() for l in f.read().strip().split("\n")]
    arr = np.array([[m[c] for c in l[0]] for l in p]).T

res1=0
for i in range(len(arr)):
    row :np.ndarray = arr[i]
    curr = 0
    blocks = np.where(row == -1)[0]
    blocks = np.append(blocks,[len(row)])
    while len(blocks):
        next = blocks[0]
        cnt = np.sum((row[curr:next]),0)
        if cnt > 0:
            n = len(row) - curr
            k = n - cnt
            res1 += (n*(n+1) - k*(k+1)) / 2
        curr = next + 1
        blocks = blocks[1:]


print(res1)

