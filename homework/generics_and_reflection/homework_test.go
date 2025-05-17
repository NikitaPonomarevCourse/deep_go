package main

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type Person struct {
	Name     string  `properties:"name"`
	Address  string  `properties:"address,omitempty"`
	Position *string `properties:"position,omitempty"`
	Age      int     `properties:"age"`
	Married  bool    `properties:"married"`
}

func Serialize[T any](person T) string {
	answer := ""
	t := reflect.TypeOf(person)
	v := reflect.ValueOf(person)
	if t.Kind() != reflect.Struct {
		return ""
	}
	for i := 0; i < t.NumField(); i++ {
		data := strings.Split(t.Field(i).Tag.Get("properties"), ",")
		if len(data) > 1 {
			if data[1] == "omitempty" {
				if v.Field(i).IsZero() {
					continue
				}
			}
		}
		if v.Field(i).Kind() == reflect.Ptr && !v.Field(i).IsNil() {
			answer += fmt.Sprint(data[0], "=", v.Field(i).Elem().Interface())
		} else if v.Field(i).Kind() != reflect.Ptr {
			answer += fmt.Sprint(data[0], "=", v.Field(i).Interface())
		} else {
			continue
		}
		if i != t.NumField()-1 {
			answer += "\n"
		}
	}
	return answer
}

func TestSerialization(t *testing.T) {
	tests := map[string]struct {
		person Person
		result string
	}{
		"test case with empty fields": {
			result: "name=\nage=0\nmarried=false",
		},
		"test case with fields": {
			person: Person{
				Name:    "John Doe",
				Age:     30,
				Married: true,
			},
			result: "name=John Doe\nage=30\nmarried=true",
		},
		"test case with omitempty field": {
			person: Person{
				Name:    "John Doe",
				Age:     30,
				Married: true,
				Address: "Paris",
			},
			result: "name=John Doe\naddress=Paris\nage=30\nmarried=true",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := Serialize(test.person)
			assert.Equal(t, test.result, result)
		})
	}
}
