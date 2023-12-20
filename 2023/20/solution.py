with open("20/example_input") as file:
# with open("20/input") as file:
    lines = file.read().strip().split("\n")

registers = {}

for l in lines:
    name, *_ = l.split()
    match name[0]:
        case "%": # flip flop
            registers[name[1:]] = lambda x: not x
        case "&": # flip flop
            registers[name[1:]] = lambda x: x and True
        case _:
            registers[name] = None


print(registers)