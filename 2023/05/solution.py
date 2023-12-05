import math
import re
from collections import defaultdict

# with open("05/input") as f:
with open("05/example_input") as f:
    parts = f.read().strip().split("\n\n")

seeds = list(map(int, re.findall(r"\d+", parts[0])))
seedstuple = list(zip(seeds[::2], seeds[1::2]))
seeds2range = [range(start, ln) for start, ln in seedstuple]

for paragraph in parts[1:]:
    lines = paragraph.split("\n")
    # currMapStr = lines[0].split(" ")[0]
    maps: list[(int, int, int)] = []
    for l in lines[1:]:
        m = list(map(int, re.findall(r"\d+", l)))
        dst, src, ln = m[0], m[1], m[2]
        maps.append((src, dst, ln))

    newSeeds = []
    for x in seeds:
        mapped = False
        for src, dst, ln in maps:
            if x >= src and x < src + ln:
                newSeeds.append(dst + (x - src))
                mapped = True
                break

        if not mapped:
            newSeeds.append(x)
    # newSeeds = list(map(lambda x: currMap.get(x, x), seeds))
    seeds = newSeeds


res1, res2 = 0, 0
res1 = min(seeds)

print(f"res1: {res1}")
print(f"res2: {res2}")
