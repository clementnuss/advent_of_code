import itertools
import math
import re
from collections import defaultdict

with open("05/input") as f:
# with open("05/example_input") as f:
    parts = f.read().strip().split("\n\n")

seeds = list(map(int, re.findall(r"\d+", parts[0])))
seedstuple = list(zip(seeds[::2], seeds[1::2]))

fns = []
finvs = []

for paragraph in parts[1:]:
    lines = paragraph.split("\n")
    maps: list[(int, int, int)] = []
    nbrs = [list(map(int, re.findall(r"\d+", l))) for l in lines[1:]]

    def f(x, nbrs=nbrs):
        for dst, src, n in nbrs:
            if src <= x and x < src + n:
                return x - src + dst
        return x

    def finv(x, nbrs=nbrs):
        for src, dst, n in nbrs:
            if src <= x and x < src + n:
                return x - src + dst
        return x

    fns.append(f)
    finvs.append(finv)


def F(x, functions):
    for f in functions:
        x = f(x)
    return x


res1 = min(map(lambda x: F(x, fns), seeds))
print(f"res1: {res1}")

for i in itertools.count():
    src = F(i, finvs[::-1])
    for start, length in seedstuple:
        if src >= start and src < start + length:
            print(f"res2: {i}")
            exit(0)
