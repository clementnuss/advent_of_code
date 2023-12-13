from functools import cache, reduce
from itertools import combinations
import re

# with open("13/example_input") as f:
with open("13/input") as f:
    patterns = [l.split() for l in f.read().strip().split("\n\n")]


def check(p, l):
    return p[: l + 1] == p[l + 1 :][::-1]


def detect_pattern(p, vertical=False):
    ret = 0
    if len(p) % 2 != 0:
        ret += detect_pattern(p[:-1], vertical)  # ignore last line
        right_hack = p + p[:1]  # make last line equal to first
        ret += detect_pattern(right_hack, vertical)
        return ret

    i = 0
    while i < len(p) - 1:
        if p[i] == p[i + 1]:
            if check(p, i):
                return (i+1) * (100 if not vertical else 1)
        i += 1

    return 0

res1 = 0
for p in patterns:
    res1 += detect_pattern(p)
    p_t = ["".join(x) for x in list(zip(*p))]
    res1 += detect_pattern(p_t, vertical=True)

print(res1)
