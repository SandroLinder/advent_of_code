package main

import (
	"bytes"
	"html/template"
	"os"
	"strconv"
	"strings"
)

func showErr(err error) {
	if err != nil {
		println("An error happened: ")
		println(err.Error())
	}
}

func createInputFiles() {
	showErr(os.RemoveAll("resources"))
	showErr(os.Mkdir("resources", 0755))

	filesToCreate := [4]string{"test_ex_01.txt", "test_ex_02.txt", "input_1.txt", "input_2.txt"}

	for i := 1; i <= 24; i++ {
		nameBuilder := strings.Builder{}
		nameBuilder.WriteString("resources/dec_")
		if i < 10 {
			nameBuilder.WriteString("0")
		}
		nameBuilder.WriteString(strconv.Itoa(i))

		showErr(os.Mkdir(nameBuilder.String(), 0755))

		for _, file := range filesToCreate {
			showErr(os.WriteFile(nameBuilder.String()+"/"+file, []byte{}, 0755))
		}
	}
}

func createGoFiles() {
	base := ReadFileIntoString("base.tmpl")

	temp := template.New("base")

	baseTemplate, err := temp.Parse(base)

	if err != nil {
		println(err.Error())
	}

	for i := 1; i <= 24; i++ {
		nameBuilder := strings.Builder{}
		dayBuilder := strings.Builder{}
		nameBuilder.WriteString("dec_")

		if i < 10 {
			dayBuilder.WriteString("0")
		}
		dayBuilder.WriteString(strconv.Itoa(i))

		nameBuilder.WriteString(dayBuilder.String())
		nameBuilder.WriteString(".go")

		var content bytes.Buffer

		showErr(baseTemplate.Execute(&content, struct {
			Day string
		}{dayBuilder.String()}))

		showErr(os.WriteFile(nameBuilder.String(), content.Bytes(), 0755))

	}
}
