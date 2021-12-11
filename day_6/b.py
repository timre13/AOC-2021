import matplotlib.pyplot as plt

def loadInput():
    with open("input.txt") as f:
        inp = f.read()
    days = [0]*9
    for fish in inp.strip().split(","):
        days[int(fish)] += 1
    return days

def simulateDay(days):
    first = days.pop(0)
    days[6] += first
    days.append(first)

days = loadInput()
print("There are {} fishes".format(sum(days)))
counts = [sum(days)]
for i in range(256):
    simulateDay(days)
    print("After {} days there are {} fishes".format(i+1, sum(days)))
    counts.append(sum(days))

plt.title("Advent of Code 2021 :: Day 6")
plt.ylabel("Number of fishes")
plt.xlabel("Number of days elapsed")
plt.plot(counts)
plt.show()