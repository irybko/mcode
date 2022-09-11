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
	"time"
	"sync"
	"database/sql"
	"strings"
)
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
