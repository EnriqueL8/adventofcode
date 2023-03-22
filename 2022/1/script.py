def keepTrackOfCalories(currentCalories, newCalorie, numberToTrack):
    currentCalories.append(newCalorie)
    currentCalories.sort()
    return currentCalories[-numberToTrack:]

def findBiggestCalories(elves, numberOfCalories):
    biggestCalories = [0] * numberOfCalories
    for elf in elves:
        calories = elf.rstrip().split('\n')
        calories = [int(c) for c in calories]
        totalCalories = sum(calories)
        biggestCalories = keepTrackOfCalories(biggestCalories, totalCalories, numberOfCalories)

    return biggestCalories

f = open("input.txt", "r")
elves = f.read().split('\n\n')


print(findBiggestCalories(elves, 1))
print(sum(findBiggestCalories(elves, 3)))
