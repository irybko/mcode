package main

import (
    "fmt"
    "time"
    "github.com/irybko/mcode/mcode"
)
//
func test_listsubst(loopsz int32) string {
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
    kv := make(mcode.Dict)
    kv["id"] = 0
    kv["currentdate"]  = time.Now().String()
    kv["actualdate"]   = time.Now().String()
    kv["productname"]  = "internet"
    kv["providername"] = "interzet.ru"
    kv["unitname"]     = "month"
    kv["costofunit"]   = "10.43$"
    kv["quantity"]     = 1
    kv["position"]     = "10.43$"


    kvarr := make([]mcode.Dict,0, loopsz)

    for count < loopsz {
        kvarr = append(kvarr, kv)
        count++
    }
    return mcode.Listsubst(content,tmpl,"content", kvarr)
}
//
func main() {
     fmt.Println(test_listsubst(10))
}



