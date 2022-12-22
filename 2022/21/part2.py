
from collections import deque
import copy
import re

class Monkey:
    name: str
    val = None
    op: None
    depends = None
    listeners = None

    def __init__(self, name, val):
        self.name = name
        self.val = val
        self.depends = []
        self.listeners = deque()

    def notify(self):
        if self.name == "humn":
            return

        if len(self.depends) > 0:
            d1, d2 = self.depends[0], self.depends[1]
            if monkeys[d1].val == None or monkeys[d2].val == None:
                return

            v1, v2 = complex(monkeys[d1].val), complex(monkeys[d2].val)
            if self.op == "+":
                self.val = v1 + v2
            elif self.op == "*":
                self.val = v1 * v2
            elif self.op == "/":
                self.val = v1 / v2
            elif self.op == "-":
                self.val = v1 - v2

        while self.listeners:
            self.listeners.pop().notify()

monkeys = {}

# input = open("test_input", "r")
input = open("input", "r")
for l in input.readlines():
    match = re.findall(r"\w+|[+\/*-]+", l)
    if len(match) == 2:
        m = Monkey(match[0], match[1])
        if m.name == "humn":
            m.val = None
    elif len(match) == 4:
        m = Monkey(match[0], None)
        m.op = match[2]
        m.depends.append(match[1])
        m.depends.append(match[3])

    monkeys[m.name] = m

for name, m in monkeys.items():
    for d in m.depends:
        monkeys[d].listeners.append(m)

root = monkeys["root"]
root.op = "-"

humn = monkeys["humn"]
humn.val = complex(1j)

for _, m in monkeys.items():
    m.notify()

print(root.val)
print(round(-root.val.real/root.val.imag))