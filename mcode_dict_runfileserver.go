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
	"log"
	"sync"
	"regexp"
	"net/http"
)
	
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

