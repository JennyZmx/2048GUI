package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

//Record up to 6 best records; ByScore is actually []Record
type ScoreRecord struct {
	Records ByScore
	Num     int
	Date1   string
	Date2   string
	Date3   string
	Date4   string
	Date5   string
	Date6   string
	Score1  int
	Score2  int
	Score3  int
	Score4  int
	Score5  int
	Score6  int
}

//This field is implemented for sorting function(sort.Sort)
type ByScore []Record

//One record corresponds to a single record
type Record struct {
	Date  string
	Score int
}

//Read record from local file
//File is arranged as the following or not exit
//First line: "Please keep this file, otherwise all score records will be removed! DO NOT change the text!"
//Following lines: Date(yyyy-mm-dd) score. In the sequence of decreasing score
func (s *ScoreRecord) ReadRecordFile() {
	if _, err := os.Stat("scoreboard.txt"); os.IsNotExist(err) { //The record file does not exist
		s.FillEmpty()
		return
	} else {
		file, err1 := os.Open("scoreboard.txt")
		if err1 != nil { //file is not valid
			s.FillEmpty()
			return
		}
		//read the file line by line as a slice of string (one line a string)
		fileLines := make([]string, 0)
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			fileLines = append(fileLines, scanner.Text())
		}
		if scanner.Err() != nil || len(fileLines) < 2 { //can only handle with at lease one record
			s.FillEmpty()
			return
		}
		file.Close()
		//Subroutine function to read the slice of string
		s.Num = len(fileLines) - 1
		s.ReadRecord(fileLines)
	}
}

//Fill s.Records with empty records
func (s *ScoreRecord) FillEmpty() {
	for i := len(s.Records); i < 6; i++ {
		temp := Record {
			Date:  "",
			Score: 0,
		}
		s.Records = append(s.Records, temp)
	}
}

//Read file lines. If this function is called, that means there is at least one record
func (s *ScoreRecord) ReadRecord(fileLines []string) {
	for i := 1; i < len(fileLines); i++ {
		line := strings.Split(fileLines[i], " ")
		a, err2 := strconv.Atoi(line[1]) //convert the score to int
		if err2 != nil {                 //file is not valid
			fmt.Println("Error: File disrupted! Please delete!")
			return
		}
		temp := Record{
			Date:  line[0],
			Score: a,
		}
		s.Records = append(s.Records, temp) //append to main ScoreRecord struct
	}
	//Make it to exactly 6 records, for the convenience of assigning individual fields
	//as we cannot use for loop in that function
	s.FillEmpty()
	//assign scores to individual fields (for display purpose)
	s.UpdateRecords()
}

//update the individual fields for each record
func (s *ScoreRecord) UpdateRecords() {
	sort.Sort(s.Records)
	s.Date1 = s.Records[0].Date
	s.Date2 = s.Records[1].Date
	s.Date3 = s.Records[2].Date
	s.Date4 = s.Records[3].Date
	s.Date5 = s.Records[4].Date
	s.Date6 = s.Records[5].Date
	s.Score1 = s.Records[0].Score
	s.Score2 = s.Records[1].Score
	s.Score3 = s.Records[2].Score
	s.Score4 = s.Records[3].Score
	s.Score5 = s.Records[4].Score
	s.Score6 = s.Records[5].Score
}

//Write record to the original file. Only called when about to quit the game
func (s *ScoreRecord) WriteRecord() {
	//Sort it before write to the file
	sort.Sort(s.Records)
	//overwrite the file by creating one
	outFile, err := os.Create("scoreboard.txt")
	defer outFile.Close()
	if err != nil {
		fmt.Println("ERROR: cannot generate score record file!")
		return
	}
	//Print the file
	fmt.Fprint(outFile,
		"Please keep this file, otherwise all score records will be removed! DO NOT change the text!")
	for i := 0; i < s.Num; i++ {
		fmt.Fprintln(outFile)
		fmt.Fprint(outFile, s.Records[i].Date, " ", s.Records[i].Score)
	}
}

//The following 3 functions are required for sorting records by score field using sort.Sort()
//Return the length of slice
func (r ByScore) Len() int {
	return len(r)
}

//Swap swaps the elements with indexes i and j.
func (r ByScore) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

//Less reports whether the element[i] should sort before the element[j]
//This is actually "more"--higher score for our purpose
func (r ByScore) Less(i, j int) bool {
	return r[i].Score > r[j].Score
}

//Display the score board in window
func (s *ScoreRecord) Render() string {
	return `
<div class = "scoreboardheader">SCORE BOARD</div>
<table>
  <tr class = "headline">
    <td class = "rank">RANK</td>
    <td class = "date">DATE</td>
    <td class = "score">SCORE</td>
  </tr>
  <tr class = "oddline">
    <td class = "rank">1</td>
    <td class = "date">{{if .Date1}}{{ .Date1}}{{end}}</td>
    <td class = "score">{{if .Score1}}{{ .Score1}}{{end}}</td>
  </tr>
  <tr class = "evenline">
    <td class = "rank">2</td>
    <td class = "date">{{if .Date2}}{{ .Date2}}{{end}}</td>
    <td class = "score">{{if .Score2}}{{ .Score2}}{{end}}</td>
  </tr>
  <tr class = "oddline">
    <td class = "rank">3</td>
    <td class = "date">{{if .Date3}}{{ .Date3}}{{end}}</td>
    <td class = "score">{{if .Score3}}{{ .Score3}}{{end}}</td>
  </tr>
  <tr class = "evenline">
    <td class = "rank">4</td>
    <td class = "date">{{if .Date4}}{{ .Date4}}{{end}}</td>
    <td class = "score">{{if .Score4}}{{ .Score4}}{{end}}</td>
  </tr>
  <tr class = "oddline">
    <td class = "rank">5</td>
    <td class = "date">{{if .Date5}}{{ .Date5}}{{end}}</td>
    <td class = "score">{{if .Score5}}{{ .Score5}}{{end}}</td>
  </tr>
  <tr class = "evenline">
    <td class = "rank">6</td>
    <td class = "date">{{if .Date6}}{{ .Date6}}{{end}}</td>
    <td class = "score">{{if .Score6}}{{ .Score6}}{{end}}</td>
  </tr>
</table>
<button class = "MainButton closescoreboard" onclick = "CloseWin">Close</button>
 `
}

//Generate date string to be used in record
func GenerateDate() string {
	Y, M, D := time.Now().Date()
	StringY := strconv.Itoa(Y)
	StringD := strconv.Itoa(D)
	StringM := strconv.Itoa(int(M))
	return (StringY + "-" + StringM + "-" + StringD)
}

//Add a record to score record srtuct
func (s *ScoreRecord) AddRecord(current int) {
	sort.Sort(s.Records)
	if s.Records[5].Score < current {
		r := Record{
			Date:  GenerateDate(),
			Score: current,
		}
		s.Records = append(s.Records[:5], r)
		if s.Num < 6 {
			s.Num += 1
		}
	}
}
