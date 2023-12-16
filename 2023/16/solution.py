import itertools
import queue


with open("./16/input") as f:
    lines = f.read().strip().split("\n")

# grid represented in the complex plane.
# +-----------> real
# |
# |
# |
# |
# |
#   imaginary part


grid = {i + 1j * j: x for j, l in enumerate(lines) for i, x in enumerate(l)}


def compute_energized(entry, dir):
    beams = queue.Queue()
    beams.put((entry, dir))
    beams_seen = set()

    while not beams.empty():
        pos, dir = beams.get()
        if (pos, dir) in beams_seen:
            continue
        if pos not in grid:
            continue
        beams_seen.add((pos, dir))
        new_directions = []
        match grid[pos]:
            case "|":
                if dir.imag == 0:
                    new_directions = [1j, -1j]
            case "-":
                if dir.real == 0:
                    new_directions = [1 + 0j, -1 + 0j]
            case "/":
                # dir * 1j -> -90deg rotation. conjugate as our origin is inverted
                new_directions = [(dir * 1j).conjugate()]
            case "\\":
                # dir * -1j -> 90deg rotation. conjugate as our origin is inverted
                new_directions = [(dir * -1j).conjugate()]

        if len(new_directions) == 0:
            new_directions.append(dir)
        for d in new_directions:
            beams.put((pos + d, d))
    energized = {pos for pos, _ in beams_seen}
    return len(energized)


print(compute_energized(0 + 0j, 1))

possible_entries = []
rows, cols = len(lines), len(lines[0])

for c in range(cols):
    possible_entries.append((c + 0j, 1j))
    possible_entries.append((c + (rows - 1) * 1j, -1j))
for r in range(rows):
    possible_entries.append((0 + r * 1j, 1))
    possible_entries.append(((cols - 1) + r * 1j, -1))

print(max(compute_energized(*x) for x in possible_entries))