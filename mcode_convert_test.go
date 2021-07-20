/*
	Copyright (c) 2019-2021 Ivan B. Rybko
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
	"reflect"
)

func TestGetBytes(t *testing.T) {
	//
	t.Run("Convert int8 value to byte slice", func(t *testing.T) {
		//
		var num int8 = 1
		got  := func() byte {
			return GetBytes(num)[0]
		}()
		expected := func() byte {
			//
			step := make([]byte,1)
			step[0] = byte(num)
			return step[0]
		}()
		//
		if got != expected {
			t.Errorf("got %q, expected %q", got, expected)
		}
	})
	//
	t.Run("Convert float32 value to byte slice", func(t *testing.T) {
		var num float32 = 1.1
		got  := func() byte {
			return GetBytes(num)[0]
		}()
		expected := func() byte {
			//
			step := make([]byte,1)
			step[0] = byte(num)
			return step[0]
		}()
		//
		if got != expected {
			t.Errorf("got %q, expected %q", got, expected)
		}

	})
	//
	t.Run("Convert float64 value to byte slice", func(t *testing.T) {
		var num float64 = 1.1
		got  := func() byte {
			return GetBytes(num)[0]
		}()
		expected := func() byte {
			//
			step := make([]byte,1)
			step[0] = byte(num)
			return step[0]
		}()
		//
		if got != expected {
			t.Errorf("got %q, expected %q", got, expected)
		}

	})
	//
	t.Run("Convert string value to byte slice", func(t *testing.T) {
		var num string = "1.1"
		got  	 := GetBytes(num)
		expected := []byte(num)
		//
		if !reflect.DeepEqual(got,expected) {
			t.Errorf("got %q, expected %q", got, expected)
		}
	})

}
//
func TestToInt64(t *testing.T) {
	//
	var strnum string = "1"
	var expected int64 = 1
	got  := ToInt64(strnum)

	if got != expected {
		t.Errorf("got %q, expected %q", got, expected)
	}
}
//
func TestToFloat64(t *testing.T) {
	//
	var strnum string = "1.1"
	var expected float64 = 1.1
	got  := ToFloat64(strnum)

	if got != expected {
		t.Errorf("got %q, expected %q", got, expected)
	}
}
