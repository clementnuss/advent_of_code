from functools import reduce
import re

with open("10/input") as f:
    # with open("10/example_input") as f:
    lines = f.read().strip().split("\n")

res1 = 0
res2 = 0


def up(pos):
    i, j = pos
    return i - 1, j


def down(pos):
    i, j = pos
    return i + 1, j


def right(pos):
    i, j = pos
    return i, j + 1


def left(pos):
    i, j = pos
    return i, j - 1


S = (-1, -1)

for i in range(len(lines)):
    if S != (-1, -1):
        break
    for j in range(len(lines[0])):
        if lines[i][j] == "S":
            S = (i, j)
            break


def next(pos, dir):
    x, y = pos
    match lines[x][y]:
        case "|":
            if dir == up or dir == down:
                return dir(pos), dir
        case "-":
            if dir == left or dir == right:
                return dir(pos), dir
        case "L":
            if dir == down:
                return right(pos), right
            elif dir == left:
                return up(pos), up
        case "J":
            if dir == right:
                return up(pos), up
            elif dir == down:
                return left(pos), left
        case "7":
            if dir == right:
                return down(pos), down
            elif dir == up:
                return left(pos), left
        case "F":
            if dir == up:
                return right(pos), right
            elif dir == left:
                return down(pos), down

    return Exception("invalid move")


pos = S
dir = right

loop = list()


def valid(pos):
    i, j = pos
    return i >= 0 and i < len(lines) and j >= 0 and j < len(lines[0])


for d in [left, right, up, down]:
    i, j = d(S)
    if not valid(d(S)):
        continue
    for ds in [left, right, up, down]:
        ret = next((i, j), ds)
        if type(ret) != Exception and ret[0] == S:
            dir = d
            print((i, j), ds)

pos = dir(S)
loop.append(S)

while pos != S:
    pos, dir = next(pos, dir)
    loop.append(pos)

print(len(loop) / 2)

for i in range(len(loop)):
    # https://en.wikipedia.org/wiki/Shoelace_formula
    yi1, xi1 = loop[i]
    yi1 = len(lines) - yi1
    yi2, xi2 = loop[(i + 1) % len(loop)]
    yi2 = len(lines) - yi2
    det = (yi1 + yi2) * (xi1 - xi2)
    res2 += det

# Pick's theorem: https://en.wikipedia.org/wiki/Pick%27s_theorem
# area = interior + boundary/2 - 1
area = res2 / 2
interior = area - (len(loop) / 2) + 1

print(interior)
