from functools import cmp_to_key, reduce
import re
from collections import OrderedDict, defaultdict

with open("08/input") as f:
    # with open("08/example_input") as f:
    lines = f.read().strip().split("\n")

instr = lines[0]
nodes = {}
nodes2 = []

for l in lines[2:]:
    node, left, right = re.findall(r"[A-Z\d]{3}", l)
    nodes[node] = (left, right)
    if node[-1] == "A":
        nodes2.append(node)

i = 0
node = "AAA"
while node != "ZZZ":
    node = nodes[node][0 if instr[i % len(instr)] == "L" else 1]
    i += 1
print(i)

cycles = []
for node in nodes2:
    i = 0
    while node[-1] != "Z":
        node = nodes[node][0 if instr[i % len(instr)] == "L" else 1]
        i += 1
    cycles.append(i)

def gcd(a,b):
    while a > 0 and b > 0:
        return gcd(b, a % b)
    return max(a,b)
def lcm(a,b):
    return a * int(b / gcd(a,b))

print(reduce(lambda a,b: lcm(a,b), cycles))