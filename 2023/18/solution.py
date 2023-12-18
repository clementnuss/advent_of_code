from collections import defaultdict
import re


# with open("./18/example_input") as f:
with open("./18/input") as f:
    lines = f.read().strip().split("\n")

# grid represented in the complex plane.
# +-----------> real
# |
# |
# |
# |
#   imaginary part

dir_to_z = {"L": -1 + 0j, "R": 1 + 0j, "U": 0 - 1j, "D": 0 + 1j}
dir2_to_z = {"2": -1 + 0j, "0": 1 + 0j, "3": 0 - 1j, "1": 0 + 1j}
vertices1 = []
vertices2 = []
z1 = 0 + 0j
z2 = 0 + 0j
for instr in lines:
    dir, n, color = re.findall(r"(\w)\s(\d+)\s\(#(\w{6})\)", instr)[0]
    n = int(n)

    z1 += (dir_to_z[dir] * n)
    vertices1.append(z1)
    z2 += dir2_to_z[color[-1:]] * int(color[:5], base=16)
    vertices2.append(z2)

def compute_area(loop):
    n = len(loop)
    interior = 0
    boundary = 0
    for i in range(n):
        # https://en.wikipedia.org/wiki/Shoelace_formula
        z1 = loop[i]
        z2 = loop[(i + 1) % n]
        boundary += abs(z2-z1)
        det = (z1.imag + z2.imag) * (z1.real - z2.real)
        interior += 0.5 * det


    # Pick's theorem: https://en.wikipedia.org/wiki/Pick%27s_theorem
    # area = interior + (boudary / 2) - 1
    return interior + boundary/2 + 1

print(compute_area(vertices1))
print(compute_area(vertices2))




