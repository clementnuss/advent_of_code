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

loop = set()


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
loop.add(S)

res1 = 1
while pos != S:
    pos, dir = next(pos, dir)
    loop.add(pos)
    res1 += 1
    # print(pos, dir)

print(res1 / 2)


def iterative_dfs(graph, node):
    visited = [False] * len(graph)
    queue = [node]

    while queue:
        current_node = queue.pop(0)

        if not visited[current_node]:
            visited[current_node] = True
            connected_component = [current_node]

            for neighbor in graph[current_node]:
                if not visited[neighbor]:
                    queue.append(neighbor)
                    connected_component.append(neighbor)

            yield connected_component


visited = set()

def dfs(pos):
    queue = [pos]
    connected_component = []

    while queue:
        curr = queue.pop(0)
        if curr not in visited:
            visited.add(curr)
            for neighbor in filter(
                lambda p: valid(p), [d(pos) for d in [up, down, left, right]]
            ):
                if neighbor not in visited and neighbor not in loop:
                    connected_component.append(neighbor)
                    queue.append(neighbor)

    yield connected_component


def find_connected_components():
    connected_components = []

    for i in range(len(lines)):
        for j in range(len(lines[0])):
            pos = (i, j)
            if pos not in visited and pos not in loop:
                for cc in dfs(pos):
                    connected_components.append(cc)

    return connected_components


cc = find_connected_components()
# print(cc)
print(res2)
