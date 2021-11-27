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

def addSong(day, currentTime, usedSongs, song, songTime):
	day.append(song)
	currentTime += songTime
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
	addSong(day, currentTime, [], randomSong, songTime[randomnum])
	currentTime += kidsTime[randomnum]
	currentArtistList.append(kidsArtistList[randomnum])

	# Second kids song
	newNumber = random.randint(0,len(kidsList)-1)
	while newNumber == randomnum:
		newNumber = random.randint(0,len(kidsList)-1)
	randomSong = kidsList[newNumber]
	addSong(day, currentTime, [], randomSong, songTime[randomnum])
	currentTime += kidsTime[newNumber]
	currentArtistList.append(kidsArtistList[newNumber])

	# Random numbers for rest of the songs
	randomNumbers = random.sample(range(0, len(songList)-1), 5)
	for x in range(0, len(randomNumbers)-1):
		randomnum = randomNumbers[x]
		while randomnum in usedList:
			randomnum = random.randint(0,len(songList)-1)
		randomNumbers[x] = randomnum

	# Add rest of songs
	i = 0
	while currentTime < minTime:
		randomnum = randomNumbers[i]
		currentTime += songTime[randomnum]
		addSong(day, currentTime, usedSongs, songList[randomnum], songTime[randomnum])
		currentArtistList.append(artistList[randomnum])
		usedList.append(randomnum)
		i = i+1

	print(name + ": ")
	for i in range(0, len(day)-1):
		print(day[i] + " - " + currentArtistList[i])
	print("\n")


generateDay("Monday")
generateDay("Tuesday")
generateDay("Wednsday")
generateDay("Thursday")
generateDay("Friday")
generateDay("Saturday")
generateDay("Sunday")