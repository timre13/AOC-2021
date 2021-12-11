def loadInput():
    with open("input.txt") as f:
        inp = f.read()
    fishes = [int(x) for x in inp.strip().split(",")]
    return fishes

def printFishes(fishes):
    print("[{}]: ".format(len(fishes))+",".join([str(x) for x in fishes]))

def simulateDay(fishes):
    length = len(fishes)
    for i in range(length):
        if fishes[i] == 0:
            fishes[i] = 6
            fishes.append(8)
        else:
            fishes[i] -= 1

def doTest():
    fishes = [3, 4, 3, 1, 2]
    states = [
        [2,3,2,0,1],
        [1,2,1,6,0,8],
        [0,1,0,5,6,7,8],
        [6,0,6,4,5,6,7,8,8],
        [5,6,5,3,4,5,6,7,7,8],
        [4,5,4,2,3,4,5,6,6,7],
        [3,4,3,1,2,3,4,5,5,6],
        [2,3,2,0,1,2,3,4,4,5],
        [1,2,1,6,0,1,2,3,3,4,8],
        [0,1,0,5,6,0,1,2,2,3,7,8],
        [6,0,6,4,5,6,0,1,1,2,6,7,8,8,8],
        [5,6,5,3,4,5,6,0,0,1,5,6,7,7,7,8,8],
        [4,5,4,2,3,4,5,6,6,0,4,5,6,6,6,7,7,8,8],
        [3,4,3,1,2,3,4,5,5,6,3,4,5,5,5,6,6,7,7,8],
        [2,3,2,0,1,2,3,4,4,5,2,3,4,4,4,5,5,6,6,7],
        [1,2,1,6,0,1,2,3,3,4,1,2,3,3,3,4,4,5,5,6,8],
        [0,1,0,5,6,0,1,2,2,3,0,1,2,2,2,3,3,4,4,5,7,8],
        [6,0,6,4,5,6,0,1,1,2,6,0,1,1,1,2,2,3,3,4,6,7,8,8,8,8],
    ]

    for i in range(18):
        simulateDay(fishes)
        if fishes == states[i]:
            print("Day {} passed: {}".format(i+1, fishes))
        else:
            print("Test failed: day {}, expected: {}, got: {}".format(i+1, states[i], fishes))

#doTest()

fishes = loadInput()
print("There are {} fishes".format(len(fishes)))
for i in range(80):
    simulateDay(fishes)
    print("After {} days there are {} fishes".format(i+1, len(fishes)))