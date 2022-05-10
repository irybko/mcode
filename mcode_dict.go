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

	"sort"
	"strings"
	"regexp"

	"reflect"
	"time"
	"sync"
	"database/sql"

	"net"
	"net/url"
	"net/http"

	"io/ioutil"
	"encoding/json"
)
//
var     rmut	sync.Mutex
var     rwmut 	sync.RWMutex
//
type Dict map[string]interface{} // relevant JSON object
//
func NewDict() Dict {
	return make(Dict)
}
//
// Get Dict`s instance keys
//
func (inst Dict) GetKeys() (int, []string) {
	//
    	var keys = []string{}
    	//
    	for k, _ := range inst {
		keys = append(keys, k)
    	}
    	//
	sort.Strings(keys)
	log.Println("inst keys: ", keys)
	//
    	return len(keys), keys
}
//
// Get Dict`s instance`s value by key 
//
func (inst Dict) GetValue(key string) interface{} {
	value, ok := inst[key]
	if !ok {
		errmsg := fmt.Sprintf("Error was occured: %s is incorrect key",key)
		log.Println(errmsg)
	}
	return value
}
//
// Check Dict`s instance keys
//
func (inst Dict) CheckKeys(hdrs []string) bool {
	var flag bool
	step := []string{}
	//
	_, keys := inst.GetKeys()
	//
	//
	if len(hdrs) == 0 {
		for _, k := range keys {
			step = append(step, k)
		}
		flag = true
	} else {
		if reflect.DeepEqual(hdrs, keys) {
			flag = true
		} else {
			flag = false
		} 
	}
	return flag
}
//
// Sort Dict`s instance by key
//
func (inst Dict) SortDict() {
	//
	sorted := make(Dict)
	//
	mlen, keys := inst.GetKeys()
	//
	index := 0
	for index < mlen {
		//
		k := keys[index]
		sorted[k] = inst[k]
		index++
	}
	inst = sorted
}
//
// Convert Dict`s instance to list 
//
func (inst Dict) ToList() [][]interface{} {
	//
	mlen, keys  := inst.GetKeys()
	index := 0
	log.Println("max: ", mlen)
	list  := make([][]interface{}, mlen)
	//
	for index < mlen {
		//
		k := keys[index]
		v := inst.GetValue(k)
		p := []interface{}{ k, v }
		list[index] = p
		index++
	}
	log.Println("current list content: ",list)
	return list
}
//
// Convert Dict`s instance to strings formatted as key=value array
//
func (inst Dict) ToStrings() []string {
	//
	var res = []string{}
	//
	inst.SortDict()
	list := inst.ToList()
	//
	idx := 0
	for idx < len(list) {
		//
		elem := list[idx]
		k := elem[0]
		v := elem[1]
		step := fmt.Sprintf("%s=%s", k, v)
		res = append(res, step)
		idx++
	}
	//
	return res
}
//
// 
//
func (inst Dict) GetArg(key string) string {
	//
	return inst.GetValue(key).(string)
}
//
// ParseURL to map
//
func (inst Dict) ParseURL(urlpath string) {
	//
    	step, err := url.Parse(urlpath) 	// Parse the URL and ensure there are no errors.
    	if err != nil {
        	panic(err)
    	}
	// 
    	host, port,_    := net.SplitHostPort(step.Host)
	inst["scheme"] 	= step.Scheme
	inst["host"] 	= host
	inst["port"] 	= port
    	inst["path"] 	= step.Path
	queryargsmap,_  := url.ParseQuery(step.RawQuery)
	if len(queryargsmap) != 0 {
		//
	    	for key, values := range queryargsmap {
        		inst[key] = strings.Join(values,",")
		}
	}
}
//
// Run file server by Dict`s instance which has keys which are file system`s paths and relevant values which are port numbers
//
func (inst Dict) RunFileServer() {
	//
	mlen, keys := inst.GetKeys()
	idx  := 0
	var wg sync.WaitGroup
	//
	wg.Add(mlen)
	//
	for idx < len(keys) {
		//
		path := keys[idx]
		port := inst.GetValue(path).(string)
		// check path and port by regular expreessions
		check, err := regexp.Match(`^(/[a-z][A-Z]*/)*$`, []byte(path))
		if err != nil {
			log.Fatal("Path is not correct: ", err)
		}
		if check {
			//
			check, err = regexp.Match(`^:[a-z]*$`, []byte(port)) 
			if err != nil {
				log.Fatal("Path is not correct: ", err)
			}
			//
			if check {
				go func() {
					//
					wg.Done()
					//
					err := http.ListenAndServe(port, http.FileServer(http.Dir(path)))
					if err != nil {
						log.Fatal(err)
					}
				}()
			}
		}
		idx++
	}
	wg.Wait()
}
//
// Convert any struct instance to Dict`s instance 
//
func (inst Dict) GetStructFields(pntr interface{}) {
	//
    	elem := reflect.ValueOf(pntr)//.Elem()
	ElemType := elem.Type()
	//
	for index:=0; index< elem.NumField(); index++ {
		field := elem.Field(index)
		key   := ElemType.Field(index).Name
		value := field.Interface()
		inst[key] = value 
	}
}
//
// Copy content of map[string]string to Dict`s instance
//
func (inst Dict) CopyContent(dct map[string]string) {
	//
	idx := 0
	mlen, keys := inst.GetKeys()
	//
	for idx < mlen {
		k1 := keys[idx]
		for k2, val := range dct {
			//
			if k1 == k2 {
				inst[k1] = val
			}
			inst[k2] = val
		}
		idx++
	}
}
//
// Substitutes templates with values dictionaries  
//
func (inst Dict) DictSubst(tmpl string) string {
	//
	for k, v := range inst {
		//
		if strings.Index(tmpl, k) != -1 {
			//
			var step string = "" 
			switch tvalue := v.(type) {
			case int: 
				step = fmt.Sprintf("%d", tvalue)
			case string: 
				step = tvalue
			} 
			tmpl = strings.ReplaceAll(tmpl, k, step)
		}
	}
	return tmpl
}
//
// Read request`s JSON formatted message string into Dict`s instance
//
func (inst Dict) UnmarshalRequest(request *http.Request) {
    	//
    	if request.Method == http.MethodPost {
		MemoryToRead, err := ioutil.ReadAll(request.Body)
	        if err != nil {
			log.Fatal("Error was occured during an attempt to read request Body: ", err)
	        }
        	//
	        err = json.Unmarshal(MemoryToRead, &inst)
        	if err != nil {
			log.Fatal("Error was occured during an attempt to parse json object: ", err)
	        }
    	}
}
//
// Convert http request`s form to Dict`s inst using available field names
//
func (inst Dict) HtmlFormToDict(request *http.Request, fnames []string) {
	//
	var store = map[string]string{}
	count 	:= 0
	//
    	if request.Method == http.MethodPost {
        	for count < len(fnames) {
            		key := fnames[count]
            		store[key] = request.FormValue(key)
            		count++
        	}
   	}
	//
	inst.CopyContent(store)
}
//
// Send request to DBMS by dbengine, dbpath, sqlreq  
//
func (inst Dict) DBRequest(dbengine string, dbpath string, sqlreq string) {
	//
	var sqlcode string = "" 
	if inst == nil {
		sqlcode = sqlreq
	} else {
		for k, _ := range inst { 
			k = fmt.Sprintf(":%s",k) 
		}
		sqlcode = inst.DictSubst(sqlreq)
	}
	//
 	var err		error
	var resultset 	=[]string{}
	var mu 		sync.Mutex
	var result 	sql.Result
	//
	dbpntr  	:= new(sql.DB)
	rows    	:= new(sql.Rows)
	stmt		:= new(sql.Stmt)
	trnsct  	:= new(sql.Tx)
	//
	log.Println("\t\tpath: ",dbpath)
	log.Println("\t\tsqlcode: ",sqlcode)
	//
	mu.Lock()
	defer mu.Unlock()
	//
	dbpntr, err = sql.Open(dbengine, dbpath)
	if err != nil {
       		log.Fatal(err)
    	}
	defer dbpntr.Close()
	//
	if dbpntr != nil {
		//
		if strings.Split(sqlcode," ")[0] == "select" {
			//
			func() {
				//
				rows, err = dbpntr.Query(sqlcode)
				if err != nil {
	        			log.Fatal(err)

				}
				defer rows.Close()
				//
				for {
					switch rows.Next() {
					case true:
						var name string
						if err = rows.Scan(&name); err != nil {
							log.Fatal(err)
						}
						resultset = append(resultset, name)
					case false:	
						// Check for errors from iterating over rows.
						iterr := rows.Err()
						if iterr != nil {
							log.Println(err)
						}
					}
				}
				//
			}() 
			//
			time.Sleep(5*time.Millisecond)
		} else {
			func() {
				trnsct, err = dbpntr.Begin()
				if err != nil {
			       		log.Fatal(err)
	    			}
				//
			   	stmt, err   = trnsct.Prepare(sqlcode)
				if err!= nil{
        				log.Fatal(err)
			    	}
				result, err = stmt.Exec(inst)
				if err != nil{
        				log.Fatal(err)
			    	}
				trnsct.Commit()
				//
				_, err      = result.RowsAffected()
				if err != nil{
        				log.Fatal(err)
			    	}
			}()
			time.Sleep(5*time.Millisecond)
		} 
	}
	inst["resultset"] = resultset 
}

