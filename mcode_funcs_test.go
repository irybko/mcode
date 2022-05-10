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
	"os"
	"log"
	"testing"
	"reflect"
)
//
func TestArgsToMap(t *testing.T) {
	//
	var testdata = map[string]string{
		//
		"arg1": "val1",
		"arg2": "val2",
		"arg3": "val3",
	}
	//	
	go func() {
		got := NewDict()
		//
		got.ArgsToMap(true)
		got.SortDict()
		for k1, v1 := range got {
			for k2, v2 := range testdata {
				//
				if k1 != k2 && v1 != v2 {
					t.Errorf("got key: %s and value: %s, expected key: %s and value: %s", k1, v1, k2, v2)
				}
			}
		}
	}()
	for k,v := range testdata {
		//
		err := os.Setenv(k,v)
		if err != nil {
			log.Fatal(err)
		}
	}  
}
//
func TestGetKeys(t *testing.T) {
	//
	expected := []string{ "k1", "k2", "k3" } 
	//
	dct := make(Dict)
	dct["k1"] = "v1"
	dct["k2"] = "v2"
	dct["k3"] = "v3"
	//
	_, got := dct.GetKeys()
	//
	if !reflect.DeepEqual(expected, got) {
		//
		t.Errorf("got: %q, expected: %q", got, expected)
	}
}
//
func TestGetValue(t *testing.T) {
	//
	dct := NewDict()
	dct["k1"] = "v1"
	expected := "v1"
	got := dct.GetValue("k1").(string)
	if !reflect.DeepEqual(got,expected) {
		t.Errorf("got: %q, expected: %q", got, expected)
	}
}
//
func TestGetStructFields(t *testing.T) {
	//
	expected := NewDict()
	expected["Field1"] = 0
	expected["Field2"] = "John"
	expected["Field3"] = "Jamesson"
	//
	type Some struct {
		Field1 int
		Field2 string
		Field3 string
	}
	//
	rec := Some{ Field1: 0, Field2: "John", Field3: "Jamesson" }
	//
	got := NewDict()
	got.GetStructFields(rec)
	//
	if !reflect.DeepEqual(got, expected) {
		//
		t.Errorf("got: %q, expected: %q", got, expected)
	}
	//
}
//
func TestDictSubst(t *testing.T) {
	//
	expected := "insert into persons(name,lastname,age,occupation) values(John,Smith,34,mechanic);"
	//
	tmplstr := "insert into persons(name,lastname,age,occupation) values(v1,v2,v3,v4);"
	//
	kv := NewDict()
	kv["v1"] = "John"
	kv["v2"] = "Smith"
	kv["v3"] = 34
	kv["v4"] = "mechanic"
	//
	got	 := kv.DictSubst(tmplstr)
	//
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("got: %s, expected: %s", got, expected)
	}
}
//
func ExampleDictSubst() {
	//
	tmplstr := `insert into persons(name,lastname,age,occupation) values(v1,v2,v3,v4);`
	kv := NewDict()
	kv["v1"] = "John"
	kv["v2"] = "Smith"
	kv["v3"] = 34
	kv["v4"] = "mechanic"
	fmt.Printf("result of substitution %s", kv.DictSubst(tmplstr))
	// Output
	// "insert intp persons(name,lastname,age,occupation) values(John,Smith,34,mechanic)"
}
//
func TestChooseLang(t *testing.T) {
	//
	var langs = map[string]Dictionary{
		"0": { 
			Ru: "Вы в хорошем настроении?", 
			En: "Are you in good mood?", 
		},
		"1": { 
			Ru: "Как Вы себя чувствуете?", 
			En: "How are you feeling?",   
		},
	}
	//
	step 	 := ChooseLang("en",langs)
	expected := langs["0"].En
	got	 := step["0"]
	//
	if !reflect.DeepEqual(got,expected) {
		t.Errorf("got %q, expected %q", got, expected)
	}
}
//
func ExampleChooseLang() {
	//
	var words = map[string]Dictionary{
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
//
