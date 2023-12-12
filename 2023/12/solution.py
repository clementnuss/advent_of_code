from functools import cache, reduce
from itertools import combinations
import re

# with inspiration from https://github.com/fuglede/adventofcode/blob/master/2023/day12/solutions.py

# with open("12/example_input") as f:
with open("12/input") as f:
    lines = [l.split() for l in f.read().strip().split("\n")]

rec_grp = [(record, tuple(map(int, groups.split(",")))) for record, groups in lines]
print(rec_grp)


@cache
def dp(rec, groups, curr_group=0):
    if len(rec) == 0:
        if len(groups) == 0 and curr_group == 0:
            return 1
        else:
            return 0

    ret = 0
    match rec[0]:
        case "?":
            ret += dp("#" + rec[1:], groups, curr_group)
            ret += dp("." + rec[1:], groups, curr_group)
        case "#":
            ret += dp(rec[1:], groups, curr_group + 1)
        case ".":
            if curr_group > 0:
                if len(groups) and curr_group == groups[0]:
                    ret += dp(rec[1:], groups[1:])
            else:
                ret += dp(rec[1:], groups)
    return ret


res1 = 0
res2 = 0
for record, group in rec_grp:
    res1 += dp(record + ".", group)
    res2 += dp("?".join([record]*5) + ".", 5 * group)
print(res1)
print(res2)
