
from collections import deque


input = open("./input", "r")

droplet = []

for l in input.readlines():
    i, j, k = l.strip().split(",")
    droplet.append(tuple([int(i), int(j), int(k)]))

# print(f"droplet: {droplet}")

# inspired by https://github.com/matheusstutzel/adventOfCode/blob/main/2022/18/p1.py


def sides(x, y, z):
    yield (x+1, y, z)
    yield (x-1, y, z)
    yield (x, y+1, z)
    yield (x, y-1, z)
    yield (x, y, z+1)
    yield (x, y, z-1)


total = 0

max_droplet = tuple([0, 0, 0])


def max_z(x, y, z):
    global max_droplet
    if z > max_droplet[2]:
        for s in sides(x, y, z):
            if s not in droplet:
                max_droplet = s
                break


for p in droplet:
    max_z(*p)

max = max(max_droplet)

def is_inside(x,y,z):
    for v in x,y,z:
        if v < -2 or v > max + 2:
            return False
    return True

seen = {max_droplet}
droplet_envelope = [max_droplet]
q = deque([max_droplet])
while len(q) > 0:
    currentDroplet = q.pop()
    for s in sides(*currentDroplet):
        if (s not in seen and s not in droplet and is_inside(*s)):
            droplet_envelope.append(s)
            seen.add(s)
            q.append(s)

for p in droplet:
    for s in sides(*p):
        if s in droplet_envelope:
            total += 1

print("total: ", total)
