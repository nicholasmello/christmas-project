package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"time"
	"strconv"
	"io"
	"log"
	"os"
	"runtime"
)

func readFile(file_name string) ([]string, []int, []string) {
	file, err := os.Open(file_name)

	if err != nil {
		log.Fatal(err)
	}

	reader := csv.NewReader(file)

	var songList []string
	var artistList []string
	var songTime []int

	for {
		record, err := reader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		songTimeStr := record[1]
		songTimeInt, err := strconv.Atoi(songTimeStr)

		if err != nil {
			log.Fatal(err)
		}

		songList = append(songList, record[0])
		songTime = append(songTime, songTimeInt)
		artistList = append(artistList, record[2])
	}
	return songList, songTime, artistList
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func addSongs(day []string, dayArtists []string, usedSongs []string, list []string, artistList []string, listTime []int, numRandom int, dayTime int) ([]string, []string, int, []string) {
	for i := 0; i < numRandom; i++ {
		for {
			rand.Seed(time.Now().UnixNano())
			randomnum := rand.Intn(len(list))
			if !contains(usedSongs, list[randomnum]) {
				day = append(day, list[randomnum])
				dayArtists = append(dayArtists, artistList[randomnum])
				usedSongs = append(usedSongs, list[randomnum])
				dayTime += listTime[randomnum]
				break
			}
		}
	}

	return day, dayArtists, dayTime, usedSongs
}

func fillDay(day []string, dayArtists []string, usedSongs []string, list []string, artistList []string, listTime []int, numRandom int, dayTime int) ([]string, []string, int, []string) {
	for {
		day, dayArtists, dayTime, usedSongs = addSongs(day, dayArtists, usedSongs, list, artistList, listTime, 1, dayTime)
		if dayTime > 1100 {
			break
		}
	}
	return day, dayArtists, dayTime, usedSongs
}

func pprintDay(dayName string, songs []string, artists []string, totalTime int) {
	fmt.Print("---")
	fmt.Print(dayName)
	fmt.Println("---")
	for i := 0; i < len(songs); i++ {
		fmt.Print(songs[i])
		fmt.Print(" by ")
		fmt.Println(artists[i])
	}
	fmt.Print(totalTime/60)
	fmt.Print(":")
	if totalTime % 60 < 10 {
		fmt.Print("0")
	}
	fmt.Println(totalTime%60)
	fmt.Println("")
}

func swap_days(day []string, dayArtists []string, first int, second int) ([]string, []string) {
	var firstDay = day[first]
	var firstArtist = dayArtists[first]
	day[first] = day[second]
	dayArtists[first] = dayArtists[second]
	day[second] = firstDay
	dayArtists[second] = firstArtist
	return day, dayArtists
}

func main() {
	var usedSongs []string
	var kidsUsedSongs []string

	var monday []string
	var mondayArtists []string
	var tuesday []string
	var tuesdayArtists []string
	var wednsday []string
	var wednsdayArtists []string
	var thursday []string
	var thursdayArtists []string
	var friday []string
	var fridayArtists []string
	var saturday []string
	var saturdayArtists []string
	var sunday []string
	var sundayArtists []string

	mondayTime := 0;
	tuesdayTime := 0;
	wednsdayTime := 0;
	thursdayTime := 0;
	fridayTime := 0;
	saturdayTime := 0;
	sundayTime := 0;

	songList, songTime, artistList := readFile("songs.csv");
	kidsList, kidsTime, kidsArtistList := readFile("kids.csv");

	// Adds the first song to every day
	monday = append(monday, songList[0])
	mondayArtists = append(mondayArtists, artistList[0])
	mondayTime += songTime[0]
	tuesday = append(tuesday, songList[0])
	tuesdayArtists = append(tuesdayArtists, artistList[0])
	tuesdayTime += songTime[0]
	wednsday = append(wednsday, songList[0])
	wednsdayArtists = append(wednsdayArtists, artistList[0])
	wednsdayTime += songTime[0]
	thursday = append(thursday, songList[0])
	thursdayArtists = append(thursdayArtists, artistList[0])
	thursdayTime += songTime[0]
	friday = append(friday, songList[0])
	fridayArtists = append(fridayArtists, artistList[0])
	fridayTime += songTime[0]
	saturday = append(saturday, songList[0])
	saturdayArtists = append(saturdayArtists, artistList[0])
	saturdayTime += songTime[0]
	sunday = append(sunday, songList[0])
	sundayArtists = append(sundayArtists, artistList[0])
	sundayTime += songTime[0]
	usedSongs = append(usedSongs, songList[0])

	// Adding first kids songs to each day
	monday, mondayArtists, mondayTime, kidsUsedSongs = addSongs(monday, mondayArtists, kidsUsedSongs, kidsList, kidsArtistList, kidsTime, 2, mondayTime)
	tuesday, tuesdayArtists, tuesdayTime, kidsUsedSongs = addSongs(tuesday, tuesdayArtists, kidsUsedSongs, kidsList, kidsArtistList, kidsTime, 2, tuesdayTime)
	wednsday, wednsdayArtists, wednsdayTime, kidsUsedSongs = addSongs(wednsday, wednsdayArtists, kidsUsedSongs, kidsList, kidsArtistList, kidsTime, 2, wednsdayTime)
	thursday, thursdayArtists, thursdayTime, kidsUsedSongs = addSongs(thursday, thursdayArtists, kidsUsedSongs, kidsList, kidsArtistList, kidsTime, 2, thursdayTime)
	friday, fridayArtists, fridayTime, kidsUsedSongs = addSongs(friday, fridayArtists, kidsUsedSongs, kidsList, kidsArtistList, kidsTime, 2, fridayTime)
	saturday, saturdayArtists, saturdayTime, kidsUsedSongs = addSongs(saturday, saturdayArtists, kidsUsedSongs, kidsList, kidsArtistList, kidsTime, 2, saturdayTime)
	sunday, sundayArtists, sundayTime, kidsUsedSongs = addSongs(sunday, sundayArtists, kidsUsedSongs, kidsList, kidsArtistList, kidsTime, 2, sundayTime)

	// Adding to fill the rest of the days
	monday, mondayArtists, mondayTime, usedSongs = fillDay(monday, mondayArtists, usedSongs, songList, artistList, songTime, 5, mondayTime)
	tuesday, tuesdayArtists, tuesdayTime, usedSongs = fillDay(tuesday, tuesdayArtists, usedSongs, songList, artistList, songTime, 5, tuesdayTime)
	wednsday, wednsdayArtists, wednsdayTime, usedSongs = fillDay(wednsday, wednsdayArtists, usedSongs, songList, artistList, songTime, 5, wednsdayTime)
	thursday, thursdayArtists, thursdayTime, usedSongs = fillDay(thursday, thursdayArtists, usedSongs, songList, artistList, songTime, 5, thursdayTime)
	friday, fridayArtists, fridayTime, usedSongs = fillDay(friday, fridayArtists, usedSongs, songList, artistList, songTime, 5, fridayTime)
	saturday, saturdayArtists, saturdayTime, usedSongs = fillDay(saturday, saturdayArtists, usedSongs, songList, artistList, songTime, 5, saturdayTime)
	sunday, sundayArtists, sundayTime, usedSongs = fillDay(sunday, sundayArtists, usedSongs, songList, artistList, songTime, 5, sundayTime)

	// Making kids song not back to back
	monday, mondayArtists = swap_days(monday, mondayArtists, 2, 3)
	tuesday, tuesdayArtists = swap_days(tuesday, tuesdayArtists, 2, 3)
	wednsday, wednsdayArtists = swap_days(wednsday, wednsdayArtists, 2, 3)
	thursday, thursdayArtists = swap_days(thursday, thursdayArtists, 2, 3)
	friday, fridayArtists = swap_days(friday, fridayArtists, 2, 3)
	saturday, saturdayArtists = swap_days(saturday, saturdayArtists, 2, 3)
	sunday, sundayArtists = swap_days(sunday, sundayArtists, 2, 3)

	fmt.Println("\n")

	pprintDay("Monday", monday, mondayArtists, mondayTime)
	pprintDay("Tuesday", tuesday, tuesdayArtists, tuesdayTime)
	pprintDay("Wednsday", wednsday, wednsdayArtists, wednsdayTime)
	pprintDay("Thursday", thursday, thursdayArtists, thursdayTime)
	pprintDay("Friday", friday, fridayArtists, fridayTime)
	pprintDay("Saturday", saturday, saturdayArtists, saturdayTime)
	pprintDay("Sunday", sunday, sundayArtists, sundayTime)

	weekData := [][]string{monday,tuesday,wednsday,thursday,friday,saturday,sunday}
	weekDataArtists := [][]string{mondayArtists,tuesdayArtists,wednsdayArtists,thursdayArtists,fridayArtists,saturdayArtists,sundayArtists}
	dayList := []string{"Monday","Tuesday","Wednsday","Thursday","Friday","Saturday","Sunday"}
	htmloutput := "<figure><table><tbody><tr>"
	for i := 0; i < 7; i++ {
		htmloutput += "<td style='background-color: #1d1e1e;color:#fff; width: 100px;'><b>"
		htmloutput += dayList[i]
		htmloutput += "</b></td>"
	}
	htmloutput += "<tr>"
	
	maxSongs := 0
	for i := 0; i < len(weekData); i++ {
		if maxSongs < len(weekData[i]) {
			maxSongs = len(weekData[i])
		}
	}

	for i := 0; i < maxSongs; i++ {
		htmloutput += "<tr>"
		for j := 0; j < 7; j++ {
			htmloutput += "<td style = 'width: 100px;'>"
			if (len(weekData[j]) > i) {
				htmloutput += weekData[j][i]
				htmloutput += " - <i>"
				htmloutput += weekDataArtists[j][i]
				htmloutput += "<i>"
			}
			htmloutput += "</td>"
		}
		htmloutput += "</tr>"
	}
	htmloutput += "</tbody></table></figure>\n"

	file, err := os.Create("table.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := file.WriteString(htmloutput)
	if err != nil {
		fmt.Println(err)
		file.Close()
		return
	}
	fmt.Println(l, "bytes written successfully")
	err = file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	if runtime.GOOS == "windows" {
		fmt.Scanln()
	}
}