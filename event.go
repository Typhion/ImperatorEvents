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

type Option struct {
	Name    string
	Effects string
}

func (o Option) String() string {
	tmpl, err := template.New("option").Parse("option = {\n" +
		"\t\tname = {{.Name}}\n" +
		"{{.Effects}}" +
		"\t}\n\n")
	if err != nil {
		panic(err)
	}
	var tpl bytes.Buffer
	if err := tmpl.Execute(&tpl, o); err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	return tpl.String()
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
	Options          []Option
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
		"{{range .Options}}{{if .}}\t{{.}}{{end}}{{end}}" +
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

func (e Event) LocString() string {
	tmpl, err := template.New("loc").Parse(" {{.Namespace}}_{{.EventId}}_title:0 \"\"\n" +
		" {{.Namespace}}_{{.EventId}}_desc:0 \"\"\n" +
		"{{range .Options}}{{if .Name}} {{.Name}}{{end}}:0 \"\"\n{{end}}")
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
	LocFile   string
	Namespace string
}

func (e EventFile) makeFile(events []Event) {
	e.File = "namespace = " + e.Namespace + "\n\n"
	for i, event := range events {
		events[i].Namespace = e.Namespace
		events[i].EventId = i + 1
		for j := range event.Options {
			event.Options[j].Name = fmt.Sprintf("%s_%d.%s", e.Namespace, event.EventId, string(rune(j+97)))
		}
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

func (e EventFile) makeLocFile(events []Event) {
	e.LocFile = "l_english:\n\n# Generated File\n\n"
	for _, event := range events {
		e.LocFile += event.LocString() + "\n"
	}
	e.writeLoc()
}

func (e EventFile) writeLoc() {
	filePath := e.Namespace + "_l_english.yml"
	fileOutput, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer fileOutput.Close()
	_, err = fileOutput.WriteString(e.LocFile)
	if err != nil {
		panic(err)
	}
	addUtf8Bom(filePath)
}
