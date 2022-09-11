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
	"os"
	"log"
	"strings" 
	"testing"
)
/*
	Prepare data for tests
*/
var str1 string = "Lorem ipsum dolor sit amet"
var str2 string = "Lorem ipsum dolor sit amet, consectetur adipiscing elit"
//
var bstr1 []byte = []byte(str1)
var bstr2 []byte = []byte(str2)
//
var sstr1 = func() []string {	return strings.Split(str1," ") 	}()
var sstr2 = func() []string { 	return strings.Split(str2," ") 	}()
//
func TestCmpPair(t *testing.T) {
	//
	t.Run("Compare two same strings", func(t *testing.T) {
		got  := CmpPair(str1,str1)
		expected := true
		//
		if got != expected {
			t.Errorf("got %t, expected %t", got, expected)
		}
	})
	//
	t.Run("Compare two different strings", func(t *testing.T) {
		//
		got  := CmpPair(str1,str2)
		expected := false
		//
		if got != expected {
			t.Errorf("got %t, expected %t", got, expected)
		}
	})
	//
	t.Run("Compare two same byte slices", func(t *testing.T) {
		//
		got := CmpPair(bstr1,bstr1)
		expected := true
		//
		if got != expected {
			t.Errorf("got %t, expected %t", got, expected)
		}
	})
	//
	t.Run("Compare two different byte slices", func(t *testing.T) {
		//
		got := CmpPair(bstr1,bstr2)
		expected := false
		//
		if got != expected {
			t.Errorf("got %t, expected %t", got, expected)
		}
	})
	//
	t.Run("Compare two same string slices", func(t *testing.T) {
		//
		got := CmpPair(sstr1,sstr1)
		expected := true
		//
		if got != expected {
			t.Errorf("got %t, expected %t", got, expected)
		}
	})
	//
	t.Run("Compare two different string slices", func(t *testing.T) {
		//
		got := CmpPair(sstr1,sstr2)
		expected := false
		//
		if got != expected {
			t.Errorf("got %t, expected %t", got, expected)
		}
	})
}
//
func ExampleCmpPair() {
	//
	var usage string = `
		=========================================
		Available datatypes:
		=========================================
			1. string; 
			2. bytes; 
			3. strings;
			4. exit.
		=========================================
		Type your number as command line argument
	`
	//
	for {
		choice := ReadStdinInt(usage)
		//
		switch choice {
		case 1:	
			log.Println("Compare two same strings: ",str1,str1,CmpPair(str1,str1))
			log.Println("Compare two different strings: ",str1,str2,CmpPair(str1,str2))

		case 2:
			log.Println("Compare two same byte slices: ",bstr1,bstr1,CmpPair(bstr1,bstr1))
			log.Println("Compare two different byte slices: ",bstr1,bstr2,CmpPair(bstr1,bstr2))

		case 3:
			log.Println("Compare two same string slices: ",sstr1,sstr2,CmpPair(sstr1,sstr1))
			log.Println("Compare two different string slices: ",sstr1,sstr2,CmpPair(sstr1,sstr2))

		case 4: os.Exit(0)	
		}
	}
}


