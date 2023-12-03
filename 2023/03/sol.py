import math
import re
from collections import defaultdict

with open("input") as f:
    lines = f.read().strip().split("\n")


def findAdjacentSymbol(i, j) -> bool:
    for m in range(i - 1, i + 2):
        for n in range(j - 1, j + 2):
            if m >= 0 and m < len(lines) and n >= 0 and n < len(lines[0]):
                c = lines[m][n]
                if not (c >= "0" and c <= "9" or c == "."):
                    return True
    return False


res1 = 0

for i in range(len(lines)):
    currentValue = 0
    parsing = False
    adjSymb = False
    for j in range(len(lines[i])):
        c = lines[i][j]
        if c >= "0" and c <= "9":
            if parsing:
                currentValue *= 10
                currentValue += int(c)
                if not adjSymb:
                    adjSymb = findAdjacentSymbol(i, j)
            else:
                adjSymb = False
                currentValue = int(c)
                parsing = True
                if not adjSymb:
                    adjSymb = findAdjacentSymbol(i, j)
        else:
            if parsing:
                parsing = False
                if adjSymb:
                    res1 += currentValue
    if parsing and adjSymb:
        res1 += currentValue

print(f"part 1: {res1}")