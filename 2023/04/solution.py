import math
import re
from collections import defaultdict

with open("04/input") as f:
# with open("04/example_input") as f:
    lines = f.read().strip().split("\n")

res1, res2 = 0, 0

for l in lines:
    spl = l.split(":")[1].split("| ")
    power = -1
    winning = {match.group(0) for match in re.finditer(r"\d+", spl[0])}
    # print(winning, type(winning))
    for match in re.finditer(r"\d+", spl[1]):
        print(match.group(0))
        if match.group(0) in winning:
            power += 1
    if power >= 0:
        res1 += 1 << power

print(f"res1: {res1}")
