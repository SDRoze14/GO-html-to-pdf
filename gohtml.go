package main

import (
	"net/http"
	"text/template"
)

type Section1 struct {
	Head  string
	Head2 string
}

type Section2 struct {
	check_box1 string
	check_box2 string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	s1 := Section1{
		Head:  "แบบสรุปแผนงาน/โครงการที่สอดคล้องกับยุทธศาสตร์ชาติ และการปฏิรูป",
		Head2: "การเพิ่มประสิทธิภาพการขับเคลื่อนนโยบายด้านการเปลี่ยนแปลงสถาพภูมิอากาศของประเทศ",
	}
	t, _ := template.ParseFiles("test.html")
	t.Execute(w, s1)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8000", nil)
}
