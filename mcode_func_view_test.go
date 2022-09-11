/*
	Copyright (c) 2019-2022 Ivan B. Rybko
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
	"time"
)

//
//
//
func ExampleListsubst(loopsz int32) string {
    var count int32 = 0

    tmpl :=  `{
        "id":            {{ .id           }},
        "currentdate":   {{ .currentdate  }},
        "actualdate":    {{ .actualdate   }},
        "productname":   {{ .productname  }},
        "providername":  {{ .providername }},
        "unitname":      {{ .unitname     }},
        "costofunit":    {{ .costofunit   }},
        "quantity":      {{ .quantity     }},
        "position":      {{ .position }}
        }`
    content :=  `{
        "appto": null,
        "elem": "table",
        "props": {
            "id": "expenses",
            "className": "table",
            "name": "expensestable" 
        },
        "content": {{ .content }}
    }`
    kv := NewDict()
    kv["id"] = 0
    kv["currentdate"]  = time.Now().String()
    kv["actualdate"]   = time.Now().String()
    kv["productname"]  = "internet"
    kv["providername"] = "interzet.ru"
    kv["unitname"]     = "month"
    kv["costofunit"]   = "10.43$"
    kv["quantity"]     = 1
    kv["position"]     = "10.43$"


    kvarr := make([]Dict,0, loopsz)

    for count < loopsz {
        kvarr = append(kvarr, kv)
        count++
    }
    return Listsubst(content,tmpl,"content", kvarr)
}


