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
	"testing"
)
//
func TestReadStdinString(t *testing.T) {
	//
	expected := "I am fine, thank you!"
	got    := ReadStdinString("How do you do?")
	//
	if expected != got {
		t.Errorf("got: %q, expected: %q", got, expected)
	}
}
//
func TestReadStdinInt(t *testing.T) {
	//
	expected := 40
	got	 := ReadStdinInt("How age are you?")
	//
	if expected != got {
		t.Errorf("got: %q, expected: %q", got, expected)
	}
}
//
func TestReadinInt64(t *testing.T) {
	//
	expected := 40
	got	 := ReadStdinInt64("How age are you?")
	//
	if expected != got {
		t.Errorf("got: %q, expected: %q", got, expected)
	}

}
//
func TestReadinFloat64(t *testing.T) {
	//
	expected := 120.5
	got	 := ReadStdinFloat64("How weight are you?")
	//
	if expected != got {
		t.Errorf("got: %q, expected: %q", got, expected)
	}

}

