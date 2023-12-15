import itertools
import xxhash
import numpy as np

h = xxhash.xxh64()


def hash(arr: np.array):
    h.update(arr.copy(order="C"))
    ret = h.intdigest()
    h.reset()
    return ret


m = {"#": -1, "O": 1, ".": 0}

# with open("14/example_input") as f:
with open("14/input") as f:
    p = [l.split() for l in f.read().strip().split("\n")]
    arr = np.array([[m[c] for c in l[0]] for l in p]).T


def tilt(arr: np.array, dir):
    arr = np.rot90(arr, dir)
    newArr = np.zeros_like(arr)
    for i in range(len(arr)):
        row: np.array = arr[i]
        newRow = newArr[i]
        curr = 0
        blocks = np.where(row == -1)[0]
        newRow[blocks] = -1
        blocks = np.append(blocks, [len(row)])
        while len(blocks):
            next = blocks[0]
            cnt = np.sum((row[curr:next]), 0)
            if cnt > 0:
                newRow[curr : curr + cnt] = 1
            curr = next + 1
            blocks = blocks[1:]
    # return newArr
    return np.rot90(newArr, -1 * dir)


def cycle(arr):
    for dir in range(4):
        # print(f"direction: {dir}")
        arr = tilt(arr, dir)
        # print(arr)
    return arr


def north_weight(arr):
    w = 0
    for i in range(len(arr)):
        row: np.array = arr[i]
        n = len(row)
        for j in range(len(row)):
            if row[j] == 1:
                w += n - j
        # curr = 0
        # blocks = np.where(row == -1)[0]
        # blocks = np.append(blocks, [len(row)])
        # while len(blocks):
        #     next = blocks[0]
        #     cnt = np.sum((row[curr:next]), 0)
        #     if cnt > 0:
        #         n = len(row) - curr
        #         k = n - cnt
        #         w += (n * (n + 1) - k * (k + 1)) / 2
        #     curr = next + 1
        #     blocks = blocks[1:]
    return w


print(f"res1: {north_weight(tilt(arr,0))}")

seen = {}


def cycles(arr):
    for i in itertools.count():
        arr = cycle(arr)
        # print(arr.T)
        new_hash = hash(arr)
        if new_hash in seen:
            return arr, i - seen[new_hash]
        else:
            seen[hash(arr)] = i


arr, cycle_length = cycles(arr)
loop_start = seen[hash(arr)]
remaining = int((1e9 - (loop_start + 1)) % cycle_length)
while remaining:
    print(north_weight(arr))
    arr = cycle(arr)
    remaining -= 1

print(north_weight(arr))
