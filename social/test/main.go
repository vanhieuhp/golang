package main

import (
	"fmt"
	"reflect"
	"strings"
)

type Post struct {
	Title string `json:"title" validate:"required,max=100"`
}

func main() {

	post := Post{}

	t := reflect.TypeOf(post)
	field, _ := t.FieldByName("Title")

	validateTag := field.Tag.Get("validate")

	validateRules := strings.Split(validateTag, ",")
	fmt.Println("Validate Tag:", validateTag)
	fmt.Println("Validate Rules:", validateRules)
}
