import re

input = open("./input", "r")
bp = [ tuple(re.findall(r"\d+", l)) for l in input.readlines()]


