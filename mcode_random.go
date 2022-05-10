/*
	Copyright (c) 2020-2022 Ivan B. Rybko
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
	"time"
	"strconv"
	"strings"
	rand  "math/rand"
	crand "crypto/rand"
)
//
var IntCharMap = map[int]string{
	 0: "a", 1: "b", 2: "c", 3: "d", 4: "e", 5: "f", 6: "g", 7: "h", 8: "i", 9: "j",10: "k",11: "l",12: "m",13: "n",14: "o",15: "p",16: "q",17: "r",
	18: "s",19: "t",20: "u",21: "v",22: "w",23: "x",24: "y",25: "z",26: "0",27: "1",28: "2",29: "3",30: "4",31: "5",32: "6",33: "7",34: "8",35: "9",
}
//
var CharIntMap = map[string]int {
	"a": 0,	"b": 1,	"c": 2,	"d": 3,	"e": 4,	"f": 5,	"g": 6,	"h": 7,	"i": 8,	"j": 9,	"k": 10,"l": 11,"m": 12,"n": 13,"o": 14,"p": 15,"q": 16,"r": 17,
	"s": 18,"t": 19,"u": 20,"v": 21,"w": 22,"x": 23,"y": 24,"z": 25,"0": 26,"1": 27,"2": 28,"3": 29,"4": 30,"5": 31,"6": 32,"7": 33,"8": 34,"9": 35,
}
//
func RandomValue(r int) int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return r1.Intn(r)
}
//
func RandomNumber(sz1 int, sz2 int) string {
	var result string = ""
	if sz1 == 0 && sz2 != 0 {
		result = fmt.Sprintf("0.%d", RandomValue(sz2))
	}
	if sz1 != 0 && sz2 == 0 {
		result = fmt.Sprintf("%d", RandomValue(sz1))
	}
	if sz1 != 0 && sz2 != 0 {
		result = fmt.Sprintf("%d.%d", RandomValue(sz1), RandomValue(sz2))
	}
	return result
}
//
func RandomString(sz int) string {
	var result string = ""
	var count  int    = 0
	for count < sz {
		index  := RandomValue(sz)
		result += IntCharMap[index]
		count++
	}
	return result
}
// GenerateID creates a prefixed random identifier.
func GenerateID(prefix string, length int) string {
	
	idSource := (func() string {
		//
		var res string = ""
		for _, ch := range IntCharMap {
			res += ch
		}
		return res
	}())  
	id := make([]byte, length)
	// Fill our array with random numbers
	crand.Read(id)
	//
	for i, b := range id {
		id[i] = idSource[b%byte(len(idSource))]
	}
	//
	return fmt.Sprintf("%s_%s", prefix, string(id))
}
//
func RandomPath() string {
	//
	repdata  := NewReportDate()
	randpath := repdata.GetReportTime()
	return strings.Replace(randpath, ":","_",-1)
}
//
func RandomMember(lst []string) string {
	//
	sz := len(lst)
	index, err := strconv.Atoi(RandomNumber(0,sz))
	if err != nil {
		log.Fatal(err)
	}
	return lst[index]
}
/*
	RandomFullName function generates fake fullname of some person to test some application code.
	This function uses subset of most popular surnames(last names), male and female names at the USA to work.
*/
func RandomFullName() (string,string,string) {
	//
	lastnames   := []string{ "Smith","Johnson","Williams","Jones","Brown","Davis","Miller"  }
	middlenames := []string{ "James","Mary","John","Patricia","Robert","Linda","Michael","Barbara","David","Elizabeth","Richard","Jennifer","Charles","Maria","Joseph","Susan","Thomas","Lisa","Daniel","Betty" }
	//
        return RandomMember(middlenames),RandomMember(middlenames),RandomMember(lastnames) 
}
