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
	"os"
	"fmt"
	"log"
	"sync"
	"regexp"
	"net/http"
)
//
type FileCache struct {
	lock	 	sync.Mutex
	filename 	string	
	header		http.Header
	content 	[]byte
}
//
type DirCache struct {
	dirpath 	string
	portnum		int
	cachefiles	[]FileCache	
}
//
func (dcpntr *DirCache) SetFields(dirpath string, portnum string) {
	//
	dcpntr.dirpath 	  = dirpath
	dcpntr.portnum 	  = portnum	
	//
	fnames := GetFileNames(dcpntr.dirpath)
	maxlen := len(fnames)
	//
	index  := 0
	for index < maxlen {
		//
		curpath := fnames[index]
		//
		dcpntr.cachefiles[index].lock.Lock()
		dcpntr.cachefiles[index].filename = curpath
		dcpntr.cachefiles[index].header   = {}
		dcpntr.cachefiles[index].Set("Content-Length", fmt.Sprintf("%d",GetFSize(curpath)))
		dcpntr.cachefiles[index].Set("Content-Type", GetContentType(curpath))
		dcpntr.cachefiles[index].content = ReadBytes(curpath)
		//
		dcpntr.cachefiles[index].lock.UnLock()	
		//
		index++
	} 
}
//
func (dcpntr *DirCache) RunFileServer(netaddr string) {
	//
	check1, err1 := regexp.Match(`^:[a-z]*$`, []byte(dcpntr.portnum)) 
	if err1 != nil {
		log.Fatal("Port is not correct: ", err1)
	}
	//
	mlen 	:= len(dcpntr.cachefiles)
	index   := 0
	var wg sync.WaitGroup
	wg.Add(mlen)
	//
	for index < mlen {
		//
		path 	:= fmt.Sprintf("%s/%s", dcpntr.dirpath,dcpntr.cachefiles[index].filename) 
 		check2, err2 := regexp.Match(`^(/[a-z][A-Z]*/)*$`, []byte(path))
		if err2 != nil {
			log.Fatal("Path is not correct: ", err2)
		}
		addr := fmt.Sprintf("%s:%d/%s", netaddr, dcpntr.portnum, path)
		go func() {
			defer wg.Done()
			//
			fs := http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
				//
				Corsheaders("GET", resp)
				resp.WriteHeader(dcpntr.cachefiles[index].header)
				resp.Write(dcpntr.cachefiles[index].content) 
			})
			err2 := http.ListenAndServe(addr,fs)
			CheckError(err2)
		}()
		//
		index++
	}
	wg.Wait()
}
//
func GetFSize(path string) int64 {
	//
	fi, err := os.Lstat(path)
	if err != nil {
		log.Fatal(err)
	}
	return fi.Size() 
}

//
func ReadBytes(path string) []byte {
	//
	fd, err1 := os.Open(path)
	CheckError(err1)
	fc, err2 := ioutil.ReadAll(fd)
	CheckError(err2)
	return fc
}
//
func GetFileNames(dirpath string) []string {
	//
	var result = []string{}
	files, err := ioutil.ReadDir(dirpath)
	checkError(err)
	for i:=0; i < len(files); i++ {
		path := files[i].Name()
		result = append(result, path)
	}
	return result
}

