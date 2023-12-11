from functools import reduce
from itertools import combinations
import re

# with open("11/example_input") as f:
with open("11/input") as f:
    lines = f.read().strip().split("\n")

galaxies = {(i, j) for i, l in enumerate(lines) for j, x in enumerate(l) if x == "#"}

gal_rows = set(x for y, x in galaxies)
gal_cols = set(y for y, x in galaxies)
empty_rows = set(range(len(lines))) - gal_rows
empty_cols = set(range(len(lines[0]))) - gal_cols


pairs = list(combinations(galaxies, 2))


def manh_dist(y1, x1, y2, x2, expansion=2):
    d = abs(x2 - x1) + abs(y2 - y1)
    d += (expansion - 1) * len(empty_rows & set(range(min(x1, x2), max(x1, x2))))
    d += (expansion - 1) * len(empty_cols & set(range(min(y1, y2), max(y1, y2))))
    return d


d1 = [manh_dist(*p1, *p2) for p1, p2 in pairs]

print(sum(d1))

d2 = [manh_dist(*p1, *p2, expansion=int(1e6)) for p1, p2 in pairs]
print(sum(d2))
