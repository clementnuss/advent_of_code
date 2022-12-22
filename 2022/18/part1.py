
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

for p in droplet:
    for s in sides(*p):
        if s not in droplet:
            total += 1


print("total: ", total)
