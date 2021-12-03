with open("input.txt") as f: lines = [x[:-1] for x in f.readlines()]

def getMostOrLeastCommon(index, discardedList, isMostCommon):
    zcount = 0
    ocount = 0
    for i, line in enumerate(lines):
        if discardedList[i]:
            continue

        if line[index] == "0":
            zcount += 1
        elif line[index] == "1":
            ocount += 1
        else:
            asset(False)

    if isMostCommon:
        return "1" if ocount >= zcount else "0"
    else:
        return "1" if ocount < zcount else "0" 

def getValue(isOx):
    isDiscarded = [False]*len(lines)
    for bitI in range(len(lines[0])):
        commonVal = getMostOrLeastCommon(bitI, isDiscarded, isOx)
        for lineI, line in enumerate(lines):
            if line[bitI] != commonVal:
                isDiscarded[lineI] = True
                if False not in isDiscarded:
                    return line
                if isDiscarded.count(False) == 1:
                    return lines[isDiscarded.index(False)]

oxRating = getValue(True)
oxRatingInt = int(oxRating, 2)
print("Oxygen rating:       "+str(oxRating)+" | "+str(oxRatingInt))

co2Rating = getValue(False)
co2RatingInt = int(co2Rating, 2)
print("CO2 rating:          "+str(co2Rating)+" | "+str(co2RatingInt))

print("Life support rating:                "+str(oxRatingInt*co2RatingInt))
