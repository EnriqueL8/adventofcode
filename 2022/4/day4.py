
def getRangeForPair(pair):
    ranges = pair.split("-")
    return int(ranges[0]), int(ranges[1])



def calculateOverlappingPairs(input, completeOverlap=False):
    total = 0
    for line in iter(input.splitlines()):
        pairs = line.split(",")
        pairOneLow, pairOneHigh = getRangeForPair(pairs[0])
        pairTwoLow, pairTwoHigh = getRangeForPair(pairs[1])

        if pairOneLow >= pairTwoLow and  pairOneHigh <= pairTwoHigh:
            total += 1
        elif pairTwoLow >= pairOneLow and pairTwoHigh <= pairOneHigh:
            total += 1

        if(not completeOverlap):
            if pairOneLow >= pairTwoLow and pairOneLow <= pairTwoHigh:
                total += 1
            elif pairOneHigh >= pairTwoLow and pairOneHigh <= pairTwoHigh:
                total += 1
    return total

if __name__ == '__main__':
    with open('input.txt', 'r') as file:
        input = file.read()
    print(calculateOverlappingPairs(input, True))
    print(calculateOverlappingPairs(input))
