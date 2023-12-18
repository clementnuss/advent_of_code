from collections import defaultdict
import heapq


# with open("./17/example_input") as f:
with open("./17/input") as f:
    lines = f.read().strip().split("\n")

# grid represented in the complex plane.
# +-----------> real
# |
# |
# |
# |
# |
#   imaginary part


grid = {i + 1j * j: int(x) for j, l in enumerate(lines) for i, x in enumerate(l)}
ROWS, COLS = len(lines), len(lines[0])
TARGET = (COLS - 1) + (ROWS - 1) * 1j
INVALID = 10000000


def heuristic(z1: complex):
    z = z1 - TARGET
    return (z.real**2 + z.imag**2) ** 0.5


def a_star(p2=False):
    heap = [(heuristic(0 + 0j), 0, -1, 0 + 0j, 1 + 0j, 3 if not p2 else 10)]
    best = defaultdict(lambda: 1_000_000_000_000)  # z, dir, remaining

    step = 0
    while heap:
        h, heat_loss, _, z, dir, rem = heapq.heappop(heap)
        if z not in grid:
            continue

        if z == TARGET:
            if not p2 or rem < 7:
                print(h, step)
                return heat_loss

        newdirs = []  # (new dir, remaining moves)
        if rem > 0:
            newdirs.append((dir, rem))
        if not p2 or rem < 7:
            newdirs.append(
                (dir * 1j.conjugate(), 3 if not p2 else 10)
            )  # *  1j == 90˚ counter-clockwise. conjugate as Im origin upside down
            newdirs.append(
                (dir * -1j.conjugate(), 3 if not p2 else 10)
            )  # * -1j == 90˚ clockwise. conjugate as Im origin upside down
        for d, rem in newdirs:
            z2 = z + d
            rem -= 1
            if z2 not in grid:
                continue
            dist2 = heat_loss + grid[z2]
            if dist2 >= best[z2, d, rem]:
                continue

            best[z2, d, rem] = dist2
            heapq.heappush(heap, (dist2 + 2 * heuristic(z2), dist2, step, z2, d, rem))
            step += 1


print(a_star())
print(a_star(p2=True))
