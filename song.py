import csv
import random

minTime = 1100
maxTime = 1300

firstSong = "At Christmas"
firstTime = 102
firstArtist = "Intro"

songList = []
songTime = []
artistList = []
usedList = []

kidsList = []
kidsTime = []
kidsArtistList = []

weekData = []

with open('./songs.csv', 'r') as file:
	reader = csv.reader(file)
	for row in reader:
		songList.append(row[0])
		songTime.append(int(row[1]))
		artistList.append(row[2])

with open('./kids.csv', 'r') as file:
	reader = csv.reader(file)
	for row in reader:
		kidsList.append(row[0])
		kidsTime.append(int(row[1]))
		kidsArtistList.append(row[2])

def addSong(day, usedSongs, song):
	day.append(song)
	usedSongs.append(song)

def generateDay(name):
	day = []
	currentTime = 0
	usedSongs = []
	currentArtistList = []

	# First song
	currentTime += firstTime
	day.append(firstSong)
	currentArtistList.append(firstArtist)

	# First kids song
	randomnum = random.randint(0,len(kidsList)-1)
	randomSong = kidsList[randomnum]
	addSong(day, [], randomSong)
	currentTime += kidsTime[randomnum]
	currentArtistList.append(kidsArtistList[randomnum])

	# Second kids song
	newNumber = random.randint(0,len(kidsList)-1)
	while newNumber == randomnum:
		newNumber = random.randint(0,len(kidsList)-1)
	randomSong = kidsList[newNumber]
	addSong(day, [], randomSong)
	currentTime += kidsTime[newNumber]
	currentArtistList.append(kidsArtistList[newNumber])

	# Random numbers for rest of the songs
	randomNumbers = random.sample(range(0, len(songList)-1), 8)
	for x in range(0, len(randomNumbers)-1):
		randomnum = randomNumbers[x]
		while randomnum in usedList or randomnum in randomNumbers:
			randomnum = random.randint(0,len(songList)-1)
		randomNumbers[x] = randomnum

	# Add rest of songs
	i = 0
	while currentTime < minTime:
		randomnum = randomNumbers[i]
		addSong(day, usedSongs, songList[randomnum])
		currentArtistList.append(artistList[randomnum])
		usedList.append(randomnum)
		i = i+1
		currentTime += songTime[randomnum]

	songplusartist = []

	print(name + ": ")
	for i in range(0, len(day)):
		songplusartist.append(day[i] + " - <i>" + currentArtistList[i] + "<i>")
		print(day[i] + " - " + currentArtistList[i])
	if currentTime%60 < 10:
		print("Total runtime: " + str(currentTime/60) + ":0" + str(currentTime%60))
	else:
		print("Total runtime: " + str(currentTime/60) + ":" + str(currentTime%60))
	print("\n")

	weekData.append(songplusartist)

generateDay("Monday")
generateDay("Tuesday")
generateDay("Wednsday")
generateDay("Thursday")
generateDay("Friday")
generateDay("Saturday")
generateDay("Sunday")

dayList = ["Monday","Tuesday","Wednsday","Thursday","Friday","Saturday","Sunday"]
htmloutput = "<figure><table><tbody><tr>"
for x in range(0, 7):
	htmloutput += "<td style='background-color: #1d1e1e;color:#fff; width: 100px;'><b>"
	htmloutput += dayList[x]
	htmloutput += "</b></td>"
htmloutput += "</tr>"

maxSongs = 0
for x in range(0, len(weekData)):
	if maxSongs < len(weekData[x]):
		maxSongs = len(weekData[x])

for x in range(0, maxSongs):
	htmloutput += "<tr>"
	for y in range(0, 7):
		htmloutput += "<td style = 'width: 100px;'>"
		try:
			htmloutput += weekData[y][x]
		except IndexError:
			htmloutput += ""
		htmloutput += "</td>"
	htmloutput += "</tr>"

htmloutput += "</tbody></table></figure>\n"

html_file = open("table.html", "w")
html_file.write(htmloutput)
html_file.close()