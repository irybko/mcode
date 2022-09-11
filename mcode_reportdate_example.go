/*
	Copyright (c) 2021 Ivan Rybko
	=============================

	This program is a part of golang middleware functions library package mcode.

	This program is free software: you can redistribute it and/or modify 
	it under the terms of the GNU General Public License as published by 
	the Free Software Foundation, either version 3 of the License, or 
	(at your option) any later version. 

	This program is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of 
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. 
	See the GNU General Public License for more details.

	You should have received a copy of the GNU General Public License 
	along with this program. If not, see <https://www.gnu.org/licenses/>.

*/
package mcode 

import (
	"fmt"
	"os"
)

const usage = `Usage: %s getdate | getmonth | getyear | getday | gettime | printerr`

func ExampleReportDate() {
	if len(os.Args) < 2 {
		//
		fmt.Println(fmt.Sprintf(usage, os.Args[0])) 
	}
	if len(os.Args) == 2 {
		//
		repdate := NewReportDate()
		switch os.Args[1] {
		case "getdate":
			fmt.Println(repdate.GetReportDate())
		case "getmonth":
			fmt.Println(repdate.GetReportMonth())
		case "getyear":
			fmt.Println(repdate.GetReportYear())
		case "getday":
			fmt.Println(repdate.GetReportDay())
		case "gettime":
			fmt.Println(repdate.GetReportTime())
		case "printerr":
			repdate.PrintErr("Attention.During the code execution an error was occured")
		}
	}
}
