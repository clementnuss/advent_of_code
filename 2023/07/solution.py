from functools import cmp_to_key
import itertools
import math
import re
from collections import OrderedDict, defaultdict

hands = []

with open("07/input") as f:
# with open("07/example_input") as f:
    lines = f.read().strip().split("\n")
    for l in lines:
        spl = l.split(" ")
        hand = spl[0]
        bid = int(spl[1])
        hands.append((hand, bid))

high_cards = OrderedDict()
one_pair = OrderedDict()
two_pair = OrderedDict()
three_of_a_kind = OrderedDict()
full_house = OrderedDict()
four_of_a_kind = OrderedDict()
five_of_a_kind = OrderedDict()

for hand, bid in hands:
    d = {}
    for l in hand:
        if l in d:
            d[l].append(l)
        else:
            d[l] = [l]

    if len(d) == 5:
        high_cards[hand] = bid
    elif len(d) == 4:
        one_pair[hand] = bid
    elif len(d) == 3:
        parsed = False
        for card, lst in d.items():
            if len(lst) == 3:
                three_of_a_kind[hand] = bid
                parsed = True
        if not parsed:
            two_pair[hand] = bid
    elif len(d) == 2:
        parsed = False
        for card, lst in d.items():
            if len(lst) == 4:
                four_of_a_kind[hand] = bid
                parsed = True
        if not parsed:
            full_house[hand] = bid
    elif len(d) == 1:
        five_of_a_kind[hand] = bid


cards = ["A", "K", "Q", "J", "T", "9", "8", "7", "6", "5", "4", "3", "2"]
card_values = {}
for i, c in enumerate(cards):
    card_values[c] = len(cards) - i


def sort_hands(item1, item2):
    for i in range(len(item1[0])):
        c1 = item1[0][i]
        c2 = item2[0][i]
        if c1 == c2:
            continue
        else:
            return card_values[c2] - card_values[c1]

    print("error", item1, item2)


res = []
for d in [
    five_of_a_kind,
    four_of_a_kind,
    full_house,
    three_of_a_kind,
    two_pair,
    one_pair,
    high_cards,
]:
    for _, bid in sorted(d.items(), key=cmp_to_key(sort_hands)):
        res.append(bid)

print(sum(map(lambda item: item[1] * (len(res) - item[0]), enumerate(res))))

