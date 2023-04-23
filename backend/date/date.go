package date

import (
	"fmt"
	"strings"
	"strconv"
)

const SUCCESS int = 0
const NOT_INTEGER_ERR int = 1
const STRING_PARSING_ERR int = 2
const ZERO_OR_LOWER_INPUT_ERR int = 3
const INVALID_DATE_ERR int = 4

var ErrorMessageList = []string {
	"",
	"Not an integer",
	"String parsing error",
	"<= 0 input", 
	"Invalid date",
}

var DayList = []string {"Minggu", "Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu"}
var DaysInMonthsList = []int {0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}


type Date struct {
	input string;
	d int;
	m int;
	y int;
	valid bool;
	dayResult string;
	errorCode int;
	errorMessage string;
}

func StringToDate( s string) (bool, int, int, int, int) {
	arr := strings.Split(s, "/")
	intArr := []int{}

	if len(arr) == 3 {
		for _, n := range arr{
			val, err := strconv.Atoi(n)

			if err != nil {
				return false, -999, -999, -999, NOT_INTEGER_ERR
			}
			intArr = append(intArr, val)
		}

		v, e := IsValidDate(intArr[0], intArr[1], intArr[2])
		if (v){
			return true, intArr[0], intArr[1], intArr[2], SUCCESS
		} else {
			return false, -999,-999,-999, e
		}
	} else {
		return false, -999, -999, -999, STRING_PARSING_ERR
	}

}

func IsValidDate( d int, m int, y int) (bool, int) {
	if d < 1 || m < 1 || y < 1 {
		return false, ZERO_OR_LOWER_INPUT_ERR
	} else if d > 31 || m > 12{
		return false, INVALID_DATE_ERR
	} else {
		if (m == 2){
			if IsLeapYear(y){
				if (d > 29){
					return false, INVALID_DATE_ERR
				}
			} else {
				if ( d> DaysInMonthsList[2]){
					return false, INVALID_DATE_ERR
				}
			}
			return true, SUCCESS
		} else {
			if (d > DaysInMonthsList[m]) {
				return false, INVALID_DATE_ERR
			}
			return true, SUCCESS
		}
		
	}
}

func IsLeapYear(y int) bool {
	if y % 4 != 0 {
		return false
	} else {
		if y % 400 == 0 {
			return true
		} else if y % 100 == 0{
			return false
		} else {
			return true
		}
	}
}

func (date *Date) GetDayFromDate(s string) {
	// Reset
	date.ResetDate()

	// Validate Input
	date.input = s
	date.valid, date.d, date.m, date.y, date.errorCode = StringToDate(s)
	date.errorMessage = ErrorMessageList[date.errorCode]

	// Progress further if Input is valid
	if (date.valid){
		// Get amount of days so far from year
		var days int = 365 * (date.y-1)
		days += AmountOfLeapYearSoFar(date.y)


		// Get amount of days so far from month
		for i := 1; i < date.m; i++{
			days += DaysInMonthsList[i]
		}

		if (IsLeapYear(date.y) && date.m > 2) { days++ }


		// Get amount of days so far from date
		days += date.d

		// Get the specific day
		// Some references said adjustment is needed because 1 January 0001 is Saturday
		// ADJUSTMENT := 5
		var res int = (days ) % 7 
		date.dayResult = DayList[res]

	}
}

func (date *Date) DisplayDate() {
	if (date.valid){
		fmt.Println("========")
		fmt.Println("Input: ", date.input)
		fmt.Println("Date: ", date.d)
		fmt.Println("Month: ", date.m)
		fmt.Println("Year: ", date.y)
		fmt.Println("Valid Status: ", date.valid)
		fmt.Println("Result: ", date.dayResult)
		fmt.Println("Error Code: ", date.errorCode)
		fmt.Println("Error Message: ", date.errorMessage)
		fmt.Println("========")
	} else {
		fmt.Println("========")
		fmt.Println("Input: ", date.input)
		fmt.Println("Valid Status: ", date.valid)
		fmt.Println("Error Code: ", date.errorCode)
		fmt.Println("Error Message: ", date.errorMessage)
		fmt.Println("========")
		
	}
}

func (date *Date) ResetDate() {
	date.input = ""
	date.d = 0;
	date.m = 0;
	date.y = 0;
	date.dayResult = "";
	date.valid = true;
	date.errorCode = 0;
	date.errorMessage = "";
}

func AmountOfLeapYearSoFar(y int) int {
	count := 0
	for i := 1; i < y; i++ {
		if (IsLeapYear(i)){
			count++
		}
	}		
	return count
}

func (date *Date) GetResult() string {
	return date.dayResult
}