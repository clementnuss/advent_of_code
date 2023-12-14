import numpy as np

# with open("13/example_input") as f:
with open("13/input") as f:
    patterns = [l.split() for l in f.read().strip().split("\n\n")]


def detect_pattern(p, vertical=False,dist=0):
    i = 1
    while i < len(p):
        comp_size = min(i, len(p) - i)
        left = p[i - comp_size : i]
        right = p[i : i + comp_size][::-1]

        if np.sum(left ^ right) == dist:
            return (i) * (100 if not vertical else 1)
        i += 1

    return 0


for dist in [0, 1]:
    res = 0
    for p in patterns:
        a = np.array([[c == "#" for c in l] for l in p])
        res += detect_pattern(a,dist=dist)
        res += detect_pattern(a.T, vertical=True, dist=dist)

    print(res)
