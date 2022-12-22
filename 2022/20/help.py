import sys

class Node:
    def __init__(self, value):
        self.value = value
        self.original = None
        self.prev = None
        self.next = None
    def move(self, n):
        self.prev.next = self.next
        self.next.prev = self.prev
        for i in range(abs(self.value)%(n-1)):
            if(self.value>0):
                self.prev = self.next
                self.next = self.next.next
            else:
                self.next = self.prev
                self.prev = self.prev.prev
        self.prev.next = self
        self.next.prev = self

input = open("test_input","r")
l = [int(k) for k in input.readlines()]
root = Node(l[0])
current = root
for i in l[1:]:
    past = current
    current  = Node(i)
    current.prev = past
    past.next = current
    past.original = current
current.next = root
root.prev = current

current = root
while current:
    current.move(len(l))
    current = current.original
current = root

c =0
k = []
while c<2:
    if current.value == 0:
        c =c+1
    if(c == 1):
        k.append(current.value)
    current = current.next

a,b,c =  k[1000%len(k)], k[2000%len(k)], k[3000%len(k)]
print(a+b+c)