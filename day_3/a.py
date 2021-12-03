with open("input.txt") as f:
    lines = f.readlines()

gamma = ""
for i in range(len(lines[0])-1):
    zcount = 0
    ocount = 0
    for line in lines:
        if line[i] == "0":
            zcount += 1
        elif line[i] == "1":
            ocount += 1

    if zcount > ocount:
        gamma += "0"
    elif ocount > zcount:
        gamma += "1"
    else:
        assert(False)

gammaDec = int(gamma, 2)
epsilonDec = int(gamma.replace("0", "-").replace("1", "0").replace("-", "1"), 2)
print(gammaDec*epsilonDec)
