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
func TestChooseLang(t *testing.T) {
	//
	var langs = map[string]Dictionary{
		"0": { Ru: "Вы в хорошем настроении?", En: "Are you in good mood?" },
		"1": { Ru: "Как Вы себя чувствуете?", En: "How are you feeling?"   },
	}
	//
	step 	 := ChooseLang("en",langs)
	//
	extected := langs["0"].En
	got	 := step["0"].En
	//
	if got != extected {
		t.Errorf("got %q, expected %q", got, expected)
	}
}
//
func ExampleChooseLang() {
	//
	var words = map[string]mcode.Dictionary{
		"d1": { Ru: "Немецкая овчарка",		En: "German shepherd" },
		"d2": { Ru: "Шотландская овчарка",      En: "Collie" },
		"d3": { Ru: "Cреднеазиатская овчарка",  En: "Central asian shepherd dog" },
	}
	//
	step := ChooseLang("en", words)
	for _, v := range step {
		fmt.Println(v)
		}
	}
	// Output: 
	// "German shepherd"
	// "Collie"
	// "Central asian shepherd dog"

}


