import functools
import itertools
import math
import re
from collections import defaultdict

with open("06/input") as f:
# with open("06/example_input") as f:
    lines = f.read().strip().split("\n")
    ns = [map(int, re.findall(r"\d+", l)) for l in lines]


rd = list(zip(ns[0], ns[1]))

res1 = 1
for duration, record in rd:
    record += 1
    h_min = math.ceil((float(duration) - math.sqrt(duration**2 - 4 * record)) / 2.0)
    h_max = (
        math.floor((float(duration) + math.sqrt(duration**2 - 4 * record)) / 2.0) + 1
    )
    res1 *= h_max - h_min

print(res1)

# part 2

race2 = [map(int, re.findall(r"\d+", l.replace(" ", ""))) for l in lines]
dr = list(zip(race2[0], race2[1]))

for duration, record in dr:
    record += 1
    h_min = math.ceil((float(duration) - math.sqrt(duration**2 - 4 * record)) / 2.0)
    h_max = (
        math.floor((float(duration) + math.sqrt(duration**2 - 4 * record)) / 2.0) + 1
    )
    res2 = h_max - h_min
    print(res2)

