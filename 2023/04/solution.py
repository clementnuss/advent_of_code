import math
import re
from collections import defaultdict

with open("04/input") as f:
# with open("04/example_input") as f:
    lines = f.read().strip().split("\n")

res1, res2 = 0, 0
copies = [1] * len(lines)

for i, l in enumerate(lines):
    spl = l.split(":")[1].split("| ")
    count = 0
    winning = {match.group(0) for match in re.finditer(r"\d+", spl[0])}
    # print(winning, type(winning))
    for match in re.finditer(r"\d+", spl[1]):
        if match.group(0) in winning:
            count += 1
    if count > 0:
        res1 += 1 << count - 1
        for j in range(i+1,i+count+1):
            copies[j] += copies[i]



print(f"res1: {res1}")
print(f"res2: {sum(copies)}")
