import string

alfa = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

def priorityOf(item):
    return alfa.index(item) + 1

def findIntersectionOf(rucksacks):
    result = set(rucksacks[0])
    for r in rucksacks[1:]:
        result.intersection_update(r)
    print(result)
    return list(result)[0]

def calculatePriorities(input):
    sumOfPriorities = 0
    for line in iter(input.splitlines()):
        compartmentOne, compartmentTwo = line[:len(line)//2], line[len(line)//2:]
        commonItem = findIntersectionOf([compartmentOne, compartmentTwo])
        sumOfPriorities += priorityOf(commonItem)
    return sumOfPriorities

def calculateGroupPriorities(input):
    sumOfPriorities = 0
    groupIndex = 0
    groupBags = []
    for line in iter(input.splitlines()):
        groupBags.append(line)
        if groupIndex < 2:
            groupIndex += 1
            continue
        commonItem = findIntersectionOf(groupBags)
        sumOfPriorities += priorityOf(commonItem)
        groupIndex = 0
        groupBags = []
    return sumOfPriorities

if __name__ == '__main__':
    with open('input.txt', 'r') as file:
        input = file.read()
    print(calculatePriorities(input))
    print(calculateGroupPriorities(input))
