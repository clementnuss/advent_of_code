import re

input = open("test_input", "r")
moves_line = ""
parse_move = False

m = []

c = 0
r = 0
for l in input.readlines():
    if l == "\n":
        parse_move = True
        continue
    if parse_move:
        moves_line = l
        break

    m.append([])
    c = 0
    for dot in l.rstrip("\n"):
        m[r].append(dot)
        c += 1
    r += 1

r = 0
c = 0
INV_DIR = {0: "R", 1: "D", 2: "L", 3: "U"}
DIR = {"R": 0, "D": 1, "L": 2,  "U": 3}
dir = DIR["R"]


while m[0][c] == ' ':
    c += 1


def mv():
    global r, c, dir
    if dir == 0:   # right
        c += 1
    elif dir == 1:  # down
        r += 1
    elif dir == 2:  # left
        c -= 1
    elif dir == 3:  # up
        r -= 1


def loop_back():
    reverse()
    mv()
    while True:
        if r == 0 or c == 0 or m[r][c] == ' ':
            break
        mv()
    reverse()


def reverse():
    global dir
    dir = (dir + 2) % 4


def move(n):
    global r, c, dir
    while n != 0:
        mv()
        if c >= len(m[r]):
            loop_back()
        elif r >= len(m):
            loop_back()

        if m[r][c] == ' ':
            loop_back()
        elif m[r][c] == '.':
            n -= 1
            continue
        elif m[r][c] == '#':
            reverse()
            mv()
            reverse()
            return


def print_state():
    for row in range(len(m)):
        line = ''
        for col in range(len(m[row])):
            if r == row and c == col:
                line += INV_DIR[dir]
            else:
                line += m[row][col]

        print(line)

for mov in re.findall(r"(\d+|\w)", moves_line):
    print(mov)
    if re.match(r"\d+", mov):
        move(int(mov))
    elif re.match(r"\w", mov):
        if mov == 'R':
            dir = (dir + 1) % 4
        elif mov == 'L':
            dir = (dir - 1) % 4
    print_state()

print(f"last position: row {r} col {c} and dir {dir}")
