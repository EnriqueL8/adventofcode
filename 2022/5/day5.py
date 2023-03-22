from array import *
import re
import math


def prettyPrint(stacks):
    print('\n'.join(['\t'.join([str(cell) for cell in row]) for row in stacks]))

def parseInput(lines):
    isStack = True
    stackLines = []
    operations = []
    for line in lines:
        if line == '\n':
            isStack = False
        if isStack:
            stackLines.append(line)
        else:
            operations.append(line)

    sizeOfStack = int(stackLines[-1].strip()[-1])
    stacks = []
    for i in range(0, sizeOfStack):
        stacks.insert(i, [])

    size = len(stackLines[-1])

    for line in stackLines[:-1]:
        for i, char in enumerate(line):
            if char.isalpha():
                pos = math.ceil(i/size * sizeOfStack)-1
                stacks[pos].insert(0, char)
    return stacks, operations


def rearrangeStacks(lines, reverse=True):
    stacks, operations = parseInput(lines)
    for operation in operations:
        numbers = [int(s) for s in operation.split() if s.isdigit()]
        if len(numbers) != 3:
            continue
        amount = numbers[0]
        fromStack = numbers[1] - 1
        toStack = numbers[2] - 1
        if reverse:
            for ch in reversed(stacks[fromStack][-amount:]):
                stacks[toStack].append(ch)
        else:
            for ch in stacks[fromStack][-amount:]:
                stacks[toStack].append(ch)

        del stacks[fromStack][-amount:]


    return stacks

def findTopItems(stacks):
    items = ""
    for stack in stacks:
        items += stack[-1]

    return items



if __name__ == '__main__':
    with open('input.txt', 'r') as file:
        lines = file.readlines()
    stacks = rearrangeStacks(lines)
    print(findTopItems(stacks))
    stacks = rearrangeStacks(lines, False)
    print(findTopItems(stacks))
