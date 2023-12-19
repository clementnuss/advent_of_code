# with open("19/example_input") as file:
with open("19/input") as file:
    rules, parts = file.read().strip().split("\n\n")

elems = []
for p in parts.split("\n"):
    elem = {}
    for category in p[1:-1].split(","):
        elem[category[:1]] = int(category[2:])
    elems.append(elem)


accepted = []
lambda x: x + 1
def act(elem, actions):
    for action in actions[:-1]:
        e = action[:1]
        threshold = int(action[2 : action.find(":")])
        new_action = action[action.find(":") + 1 :]
        cmp = action[1:2]
        if cmp == ">":
            if elem[e] > threshold:
                return r[new_action](elem)
        if cmp == "<":
            if elem[e] < threshold:
                return r[new_action](elem)
    r[actions[-1]](elem)

r = {}
for rule in rules.split("\n"):
    name = rule[: rule.find("{")]
    actions = rule[rule.find("{") + 1 : -1].split(",")
    r[name] = lambda elem, actions=actions: act(elem, actions)
r["A"] = lambda elem: accepted.append(elem)
r["R"] = lambda _: None


for elem in elems:
    r["in"](elem)

print(accepted)
res = 0
for elem in accepted:
    for cat in ["x", "m", "a", "s"]:
        res += int(elem[cat])

print(res)
