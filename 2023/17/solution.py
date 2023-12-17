import functools
import sys

sys.setrecursionlimit(10000000)


with open("./17/example_input") as f:
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
INVALID = 10000000


@functools.cache
def dp(z: complex, dir: complex, rem):
    if z not in grid:
        return INVALID

    if z == (COLS - 1) + (ROWS - 1) * 1j:
        return grid[z]

    a = []
    if rem > 0:
        a.append(grid[z] + dp(z + dir, dir, rem - 1))

    newdir = (
        dir * 1j
    ).conjugate()  # *  1j == 90˚ counter-clockwise. conjugate as Im origin upside down
    a.append(grid[z] + dp(z + newdir, newdir, 3))
    newdir = (
        dir * -1j
    ).conjugate()  # * -1j == 90˚ clockwise. conjugate as Im origin upside down
    a.append(grid[z] + dp(z + newdir, newdir, 3))

    return min(a)


right = dp(0 + 0j, 1 + 0j, 3)
print(right)
