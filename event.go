package main

import (
	"bytes"
	"fmt"
	"os"
	"text/template"
)

func addUtf8Bom(filename string) {
	newString := "\ufeff"
	if _, err := os.Stat(filename); err == nil {
		fileOld, err := os.ReadFile(filename)
		if err != nil {
			panic(err)
		}
		newString += string(fileOld)
	}
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	_, err = file.WriteString(newString)
	if err != nil {
		panic(err)
	}
}

type Event struct {
	Namespace        string
	EventId          int
	EventType        string
	Picture          string
	Hidden           bool
	FireOnlyOnce     bool
	GotoLocation     string
	LeftPortrait     string
	RightPortrait    string
	WeightMultiplier string
	Trigger          string
	Immediate        string
	Option           string
	Option2          string
	Option3          string
	Option4          string
	Option5          string
	After            string
}

func (e Event) String() string {
	tmpl, err := template.New("event").Parse("{{.Namespace}}.{{.EventId}} = {\n" +
		"\ttype = {{.EventType}}\n" +
		"\ttitle = {{.Namespace}}_{{.EventId}}_title\n" +
		"\tdesc = {{.Namespace}}_{{.EventId}}_desc\n" +
		"\tpicture = {{.Picture}}\n" +
		"{{if .Hidden}}\thidden = yes\n{{end}}" +
		"{{if .FireOnlyOnce}}\tfire_only_once = yes\n{{end}}" +
		"{{if .GotoLocation}}\tgoto_location = {{.GotoLocation}}\n{{end}}" +
		"{{if .LeftPortrait}}\tleft_portrait = {{.LeftPortrait}}\n{{end}}" +
		"{{if .RightPortrait}}\tright_portrait = {{.RightPortrait}}\n{{end}}\n" +
		"{{if .WeightMultiplier}}\t{{.WeightMultiplier}}\n\n{{end}}" +
		"{{if .Trigger}}\t{{.Trigger}}\n\n{{end}}" +
		"{{if .Immediate}}\t{{.Immediate}}\n\n{{end}}" +
		"{{if .Option}}\t{{.Option}}\n{{end}}" +
		"{{if .Option2}}\n\t{{.Option2}}\n{{end}}" +
		"{{if .Option3}}\n\t{{.Option3}}\n{{end}}" +
		"{{if .Option4}}\n\t{{.Option4}}\n{{end}}" +
		"{{if .Option5}}\n\t{{.Option5}}\n{{end}}" +
		"{{if .After}}\n\t{{.After}}\n{{end}}" +
		"}")
	if err != nil {
		panic(err)
	}
	var tpl bytes.Buffer
	if err := tmpl.Execute(&tpl, e); err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	return tpl.String()
}

type EventFile struct {
	File      string
	Namespace string
}

func (e EventFile) makeFile(events []Event) {
	e.File = "namespace = " + e.Namespace + "\n\n"
	for i, event := range events {
		event.EventId = i + 1
		e.File += event.String() + "\n\n"
	}
	e.write()
}

func (e EventFile) write() {
	filePath := e.Namespace + ".txt"
	fileOutput, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer fileOutput.Close()
	_, err = fileOutput.WriteString(e.File)
	if err != nil {
		panic(err)
	}
	addUtf8Bom(filePath)
}
