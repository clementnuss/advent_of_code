from functools import cmp_to_key
import itertools
import math
import re
from collections import OrderedDict, defaultdict

hands = []

cards = ["A", "K", "Q", "J", "T", "9", "8", "7", "6", "5", "4", "3", "2","J2"]
card_values = {}
for i, c in enumerate(cards):
    card_values[c] = len(cards) - i


with open("07/input") as f:
    # with open("07/example_input") as f:
    lines = f.read().strip().split("\n")
    for l in lines:
        spl = l.split(" ")
        hand = spl[0]
        bid = int(spl[1])
        hands.append((hand, bid))

typed_hands = []

def caracterize_hand(hand, bid, p2):
    d = {}
    for l in hand:
        if l in d:
            d[l].append(l)
        else:
            d[l] = [l]

    if p2:
        jokers = len(d.pop("J", []))
        d = sorted(list(map(lambda x: len(x), d.values())))
        if jokers == 5:
            d = [5]
        elif jokers > 0:
            d[-1] += jokers
    else:
        d = sorted(list(map(lambda x: len(x), d.values())))

    hand_type = 0
    match d:
        case *_, 5:
            hand_type = 10
        case *_, 4:
            hand_type = 9
        case *_, 2, 3:  # full house
            hand_type = 8
        case *_, 3:  # three of a kind
            hand_type = 7
        case *_, 2, 2:
            hand_type = 6
        case *_, 2:
            hand_type = 5
        case *_, 1:
            hand_type = 4

    typed_hands.append((hand_type, *map(card_values.get, hand), bid))

for hand, bid in hands:
    caracterize_hand(hand, bid, False)
sorted_hands = sorted(typed_hands)
print(sum(map(lambda item: item[1][-1] * (item[0] + 1), enumerate(sorted_hands))))

typed_hands.clear()
card_values["J"] = 1
for hand, bid in hands:
    caracterize_hand(hand, bid, True)
sorted_hands = sorted(typed_hands)
print(sum(map(lambda item: item[1][-1] * (item[0] + 1), enumerate(sorted_hands))))
