package goSdk

import (
	"os"
	"text/template"
	"github.com/qianlnk/log"
)

type T struct {
	Add func(int) int
}

func (t *T) Sub(i int) (int, error) {
	//log.Println("get argument i:", i)
	return i - 1, nil
}

func Arguments() {
	ts := &T{
		Add: func(i int) int {
			return i + 1
		},
	}
	tpl := `


		// 只能使用 call 调用
		call field func Add: {{ call .ts.Add .y }}
		// 直接传入 .y 调用
		call method func Sub: {{ .ts.Sub .y }}


		{{- if and (call .ts.Add .y ) (.ts.Sub .y) 1 3 -}}
	true
{{else}}
	{{- false -}}
{{end}}

{{- $a := 1}}
{{$a = 2}}
{{- $a }}

{{1}}


	`
	t, _ := template.New("test").Parse(tpl)
	t.Execute(os.Stdout, map[string]interface{}{
		"y": 1,
		"ts": ts,
	})
}



type Info struct {
	Name string
	User User
}

type User struct {
	Name string
	Age int
	Books []string
}
func ParseFile() {


	info := Info{
		Name: "huxiaoyu",
		User: User{
			Name:"wang",
			Age:18,
			Books:[]string{"Chiese","USA","Jap"},
		},
	}

	t, err := template.ParseGlob("./*.tmpl")
	if err != nil {
		log.Error(err)
		return
	}
	err = t.ExecuteTemplate(os.Stdout,"tmp2",info)
	if err != nil {
		log.Error(err)
		return
	}

}