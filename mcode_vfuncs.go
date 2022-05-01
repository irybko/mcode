package mcode

import (
    "text/template"
    "bytes"
)
/*
    VIEW INTERFACE FUNCTIONS
*/
var (
        jsontmpl        map[string]map[string]string
        prepstmts       map[string]map[string]string
)
//
type View interface {
        Subst(tmpltext string, kv Dict) string
        Listsubst(maintmpl string, embedtmpl string, key string, lstkv []Dict) string
}
//
func Subst(tmpltext string, kv Dict) string {
        t, _ := template.New("").Parse(tmpltext)
        buf := new(bytes.Buffer)
        t.Execute(buf, &kv)
        return buf.String()
}
//
func Listsubst(maintmpl string, embedtmpl string, key string, lstkv []Dict) string {
        var count int = 0
        arr := make([]string, 0, len(lstkv))

        for count < len(lstkv) {
                arr = append(arr, Subst(embedtmpl, lstkv[count]))
                count++
        }
        step := make(Dict)
        step[key] = arr
        return Subst(maintmpl, step)
}
