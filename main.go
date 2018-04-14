/*
title           :menu.go
author          :Jordan Robison
version         :0.1
usage           :Golang menu.go

					Comments
This script will pull the winning lotto numbers from the history and compare numbers it generated.
I hope to find a combination of winning numbers.
By cross examing old numbers with new numbers I am hoping to slim down the odds.
Think Card counting.

*/

package main

import "fmt"

func main() {
	fmt.Println(" Hello ....Are we feeling risky today? ")
	/* At this point the script does an update to comon-numbers.txt also last winning number.*/

	/* Websites with datato convert to csv
	   curl -o https://data.ny.gov/api/views/d6yy-54nr/rows.csv?accessType=DOWNLOAD >> winningnumbers.txt
	   curl -o https://www.wilottery.com/lottogames/powerballhistoryOD.aspx
	*/

	fmt.Println("Starting Calulaution.. .. .. ..")
	/* Searches for best combinations against comon-numberes.txt , winning.txt vs combolist.txt*/

	fmt.Println(" Winner Winner Chicken Dinner !!!")
	/* Provides top 20 likely matches from combolist.txt*/

	fmt.Println(" Here are your results... try not to loose them")
}
