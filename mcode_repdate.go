/*
	Copyright (c) 2020 Ivan B. Rybko
	================================

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

	"time"
	"fmt"
)
//
type ReportDate struct {
	ReportYear	string
	ReportMonth	string
	ReportDay	string
	ReportTime	string
}
//
func NewReportDate() ReportDate {
	curtime := time.Now() 
	return ReportDate{ 
		ReportYear: 	fmt.Sprintf("%d",	curtime.Year()), 
		ReportMonth: 	fmt.Sprintf("%02d",	curtime.Month()), 
		ReportDay: 	fmt.Sprintf("%02d",	curtime.Day()), 
		ReportTime: 	fmt.Sprintf("%02d:%02d:%02d", curtime.Hour(), curtime.Minute(), curtime.Second()), 
	}
}
//
func (inst ReportDate) PrintReportDate() {
	fmt.Printf("%s-%s-%s-%s",inst.ReportYear, inst.ReportMonth, inst.ReportDay, inst.ReportTime)
}
//
func (inst ReportDate) GetReportDate() string { 
	return fmt.Sprintf("%s-%s-%s-%s", inst.ReportYear, inst.ReportMonth, inst.ReportDay, inst.ReportTime)
}
//
func (inst ReportDate) GetReportYear() string {
	return fmt.Sprintf("%s", inst.ReportYear)
}
//
func (inst ReportDate) GetReportMonth() string {
	return fmt.Sprintf("%s", inst.ReportMonth)
}
//
func (inst ReportDate) GetReportDay() string {
	return fmt.Sprintf("%s", inst.ReportDay)
}
//
func (inst ReportDate) GetReportTime() string {
	return inst.ReportTime
}
//
func (inst ReportDate) GetReportDatetime() string {
	//
	return fmt.Sprintf("%s_%s",inst.GetReportDate(), inst.GetReportTime())
}
//
func (inst ReportDate) PrintErr(msg string) {
    errmsg  := "\x1b[6;31;40m \t %s : %s \x1b[0m" // red string
    repdate := inst.GetReportDate()
    //
    fmt.Println(fmt.Sprintf(errmsg,repdate,msg))
}
//
func (inst ReportDate) GetFieldMap() (res map[string]string) {
	//
	dct := NewDict()
	dct.GetStructFields(inst)
	//
	mlen, keys := dct.GetKeys()
	idx  := 0
	//
	for idx < mlen {
		//
		k := keys[idx]
		v := dct.GetValue(k).(string)
		//
		res[k] = v
	}
	return 
}
