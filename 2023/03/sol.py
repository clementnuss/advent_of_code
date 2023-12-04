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

def findAdjacentStar(i, j) -> (bool,tuple()):
    for m in range(i - 1, i + 2):
        for n in range(j - 1, j + 2):
            if m >= 0 and m < len(lines) and n >= 0 and n < len(lines[0]):
                c = lines[m][n]
                if c == "*":
                    return True,(m,n)
    return False,(-1,-1)

res1 = 0
res2 = 0

partsMap = {}

for i in range(len(lines)):
    currentValue = 0
    parsing = False
    adjSymb = False
    adjStar = False
    adjStarCoord = (0,0)
    for j in range(len(lines[i])):
        c = lines[i][j]
        if c >= "0" and c <= "9":
            if parsing:
                currentValue *= 10
                currentValue += int(c)
                if not adjSymb:
                    adjSymb = findAdjacentSymbol(i, j)
                if not adjStar:
                    adjStar,adjStarCoord = findAdjacentStar(i, j)
            else:
                adjStar,adjStarCoord = findAdjacentStar(i, j)
                adjSymb = findAdjacentSymbol(i, j)
                currentValue = int(c)
                parsing = True
        else:
            if parsing:
                parsing = False
                if adjSymb:
                    res1 += currentValue
                if adjStar:
                    if adjStarCoord in partsMap:
                        res2 += currentValue * partsMap[adjStarCoord]
                    else:
                        partsMap[adjStarCoord] = currentValue

    if parsing and adjSymb:
        res1 += currentValue
        if adjStar:
            if adjStarCoord in partsMap:
                res2 += currentValue * partsMap[adjStarCoord]
            else:
                partsMap[adjStarCoord] = currentValue

print(f"part 1: {res1}")
print(f"part 2: {res2}")