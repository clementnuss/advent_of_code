from functools import reduce
import re

with open("09/input") as f:
# with open("09/example_input") as f:
    lines = f.read().strip().split("\n")

res1 = 0
res2 = 0
for l in lines:
    orig_seq = list(map(int, l.split(" ")))
    seq = []
    seq.append(orig_seq)
    while not reduce(lambda x, y: x == 0 and y == 0, seq[-1]):
        prev_seq = seq[-1]
        new_seq = []
        for i in range(1, len(seq[-1])):
            new_seq.append(prev_seq[i] - prev_seq[i - 1])
        seq.append(new_seq)
    tmp = 0
    for s in seq[::-1]:
        res1 += s[-1]
        tmp = s[0] - tmp
    res2 += tmp

print(res1)
print(res2)
