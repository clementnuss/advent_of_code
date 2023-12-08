from functools import cmp_to_key
import itertools
import math
import re
from collections import OrderedDict, defaultdict

with open("08/input") as f:
# with open("08/example_input") as f:
    lines = f.read().strip().split("\n")

instr = lines[0]
nodes = {}

for l in lines[2:]:
    node, left, right = re.findall(r"[A-Z]{3}", l)
    nodes[node] = (left,right)

node = "AAA"
i = 0
while node != "ZZZ":
    match instr[i % len(instr)]:
        case "L":
           node = nodes[node][0]
        case "R":
           node = nodes[node][1]
    i+=1

print(i)

