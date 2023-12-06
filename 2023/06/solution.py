import functools
import itertools
import math
import re
from collections import defaultdict

with open("06/input") as f:
# with open("06/example_input") as f:
    lines = f.read().strip().split("\n")
    ns = [map(int, re.findall(r"\d+", l)) for l in lines]


td = list(zip(ns[0], ns[1]))

res1 = []
for time, record in td:
    ways = 0
    for speed in range(1, time):
        dist = speed * (time - speed)
        if dist > record:
            ways += 1
    if ways > 0:
        res1.append(ways)

print(functools.reduce(lambda x, y: x * y, res1))

# part 2

race2 = [map(int, re.findall(r"\d+", l.replace(" ", ""))) for l in lines]
td2 = list(zip(race2[0], race2[1]))
for time, record in td2:
    ways = 0
    for speed in range(1, time):
        dist = speed * (time - speed)
        if dist > record:
            ways += 1
    if ways > 0:
        res1.append(ways)

print(ways)