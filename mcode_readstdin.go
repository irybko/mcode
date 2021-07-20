/*
	Copyright (c) 2020-2021 Ivan B. Rybko
	=====================================

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
	"log"
	"os"
	"bufio"
)
//
//
//
func ReadStdinString(question string) string {
	//
	fmt.Print(fmt.Sprintf("\t\t%s:\t", question))
	in := bufio.NewReader(os.Stdin)
	line, err := in.ReadString('\n') 
	//	
	if err != nil {
		log.Println(err)
	}
	line = string(line)
	return line
	
}
//
//
//
func ReadStdinInt(question string) int {
	fmt.Print(fmt.Sprintf("\t\t%s:\t", question))
	var v int = 0
	fmt.Scanf("%d\n", &v)
	return v
}
//
//
//
func ReadStdinInt64(question string) int64 {
	fmt.Print(fmt.Sprintf("\t\t%s:\t", question))
	var v int64 = 0
	fmt.Scanf("%d\n", &v)
	return v
}
//
//
//
func ReadStdinFloat64(question string) float64 {
	fmt.Print(fmt.Sprintf("\t\t%s:\t", question))
	var v float64 = 0
	fmt.Scanf("%f\n", &v)
	return v
}

