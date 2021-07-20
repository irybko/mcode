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
	//
	"fmt"
	"sync"
	"log"
	"bytes"
	"reflect"
	"net/http"
	"encoding/json"
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
	expected := NewDict()
	expected.CopyContent(testdata)
	//
	expected.SortDict()
	expectedkeys := expected.GetKeys()
	expectedlist := expected.ToStrings()
	//
	//
	checkproc  := func(tt *testing.T) {
		//
		got  := ArgsToMap()
		//
		gotdct := NewDict()
		gotdct.CopyContent(got)
		//
		gotdct.SortDict()
		//
		gotkeys := gotdct.GetKeys()
		//
		idx := 0
		for idx < len(gotkeys) {
			//
			k1 := expectedkeys[idx]
			v1 := expected.GetValue(k1)
			//
			k2 := gotkeys[idx]
			v2 := gotdct.GetValue(k2)
			//
			fmtstr := fmt.Sprintf("got key: %q and value: %q, expected key: %q and value: %q", k1, v1, k2, v2)
			if k1 != k2 && v1 != v2 {
				tt.Errorf(fmtstr )
			}
			tt.Log(fmtstr)
			//
			idx++
		}
	}
	//
	t.Run("subtest1", func(t1 *testing.T) {
		//
		os.Args = expectedlist
		//
		log.Println("cmdline args: ", os.Args)
		//
		checkproc(t1)
	})
	//
	t.Run("subtest2", func(t2 *testing.T) {
		//
		idx := 0
		//
		for idx < len(expectedkeys) {
			//
			k := expectedkeys[idx]
			v := expected[k].(string)
			os.Setenv(k,v)
			idx++
		}
		//
		checkproc(t2)
	})
}
//
func TestGetKeys(t *testing.T) {
	//
	expected := []string{ "k1", "k2", "k3" } 
	//
	mss := make(map[string]string)
	mss["k1"] = "v1"
	mss["k2"] = "v2"
	mss["k3"] = "v3"
	//
	dct := make(Dict)
	dct.CopyContent(mss)
	//
	got := GetBytes(dct.GetKeys())
	//
	if !bytes.Equal(expected, got) {
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
	if got != expected {
		t.Errorf("got: %q, expected: %q", got, expected)
	}
}
//
func TestGetStructFields(t *testing.T) {
	//
	expected := NewDict()
	expected["fld1"] = 0
	expected["fld2"] = "John"
	expected["fld3"] = "Jamesson"
	//
	type Some struct {
		fld1 int
		fld2 string
		fld3 string
	}
	//
	rec := Some{ fld1: 0, fld2: "John", fld3: "Jamesson" }
	//
	got := NewDict()
	got.GetStructFields(rec)
	//
	b1 := GetBytes(expected.EncodeJson())
	b2 := GetBytes(got.EncodeJson())
	//
	if !bytes.Equal(b1,b2) {
		t.Errorf("got: %q, expected: %q", b1, b2)
	}
	//
}
//
func TestDictSubst(t *testing.T) {
	//
	tmplstr := `insert into persons(name,lastname,age,occupation) values(v1,v2,v3,v4);`
	kv := NewDict()
	kv["v1"] = "John"
	kv["v2"] = "Smith"
	kv["v3"] = 34
	kv["v4"] = "mechanic"
	expected := "insert intp persons(name,lastname,age,occupation) values(John,Smith,34,mechanic)"
	got	 := DictSubst(tmplstr, kv)
	b1 	 := GetBytes(expected)
	b2       := Getbytes(got)
	if !bytes.Equal(b1,b2) {
		t.Errorf("got: %q, expected: %q", got, expected)
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
	fmt.Println("result of substitution %s", DictSubst(tmplstr, kv))
	// Output
	// "insert intp persons(name,lastname,age,occupation) values(John,Smith,34,mechanic)"
}
//
func TestUnmarshalRequest(t *testing.T) {
	//
	val1 := RandomString(3)
	val2 := RandomString(3)
	val3 := RandomString(3)
	//
	expected := make(map[string]string{
		"k1": val1,
		"k2": val2,
		"k3": val3,
	})
	//
	got     := make(map[string]string)
        kv	:= NewDict()
	//
	formpath := "/testunmarshalrequest"
	//
	go func() {
		http.HanleFunc(formpath, func(req *http.Request, resp http.ResponseWriter) {
			//
			kv.UnmarshalRequest(req)
		})
		http.ListenAndServe(":9000",nil)
	}()
	//
	formdata, err := json.Marshal(expected)
	if err != nil {
		//
		log.Fatal(err)
	}
	//

	_, err = http.Post(formpath,"application/json", formdata)
	if err != nil {
		log.Fatal("During http.PostForm execution an error was occured: \n",err)
	}
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("got: %q, expected: %q", got, expected)
	}
}
//
func ExampleUnmarshalRequest() {
	//
	const usage = `
		readrequest port=number 
		===================================================
		port 	- number of the port listen to

		Below you can see examples of usage curl to test \
		correct work and relevant console output:
		===================================================
		
		curl --data \
			'{"fname":"james",\
			"lname":"williams"}'\
			-H "Content-Type: application/json"\
			-X POST http://localhost:port/readrequest

		curl --data \
			'{"fname":"john",\
			"lname":"johnson"}'\
			-H "Content-Type: application/json"\
			-X POST http://localhost:11000/readrequest
		===================================================
	`
	//
	requests := make(chan mcode.Dict)
	defer close(requests)
	//
	var grp sync.WaitGroup
	//
	fmt.Println(usage)
	portnum := fmt.Sprintf(":%s", ReadStdinString("port number"))
	//
	grp.Add(2)
	go func() {
		defer grp.Done()
		//
		http.HandleFunc("/readrequest", func(resp http.ResponseWriter, req *http.Request) {
		//
			data := mcode.UnmarshalRequest(req)
			//
			requests <- data
			resp.Header().Set("Content-Type", "text/plain; charset=utf-8")
			resp.Write([]byte("Your request is accepted\n"))
		})
		http.ListenAndServe(portnum, nil)
	}()
	//
	go func() {
		defer grp.Done()
		//
		fmt.Println("Request's message keys and values") 
		for req := range requests {
			//
			for key, value := range req {
				fmt.Println(key,"\t\t",value.(string))
			}
		}
	}()
	grp.Wait()
	// Output:
	//		Request's message keys and values
	//		=================================
	//		fname 		 john
	//		lname 		 johnson
	//		fname 		 james
	//		lname 		 williams
}

