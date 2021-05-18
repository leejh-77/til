package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	deleteReadme()

	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	buffer := bytes.Buffer{}
	buffer.WriteString("# til\nToday I learned\n")

	for _, f := range files {
		if f.IsDir() && f.Name()[0] != '.' {
			generate(f.Name(), &buffer)
		}
	}

	err = ioutil.WriteFile("./README.md", buffer.Bytes(), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func deleteReadme() {
	if f, _ := os.Stat("./README.md"); f != nil {
		err := os.Remove("./README.md")
		if err != nil {
			log.Fatal(err)
		}
	}
}

func generate(path string, buffer *bytes.Buffer) {
	buffer.WriteString("### " + path + "\n")

	fs, err := ioutil.ReadDir("./" + path)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range fs {
		fn := f.Name()
		log.Println(fn)
		n := fn[:len(fn) - 3] // slice '.md' extension
		buffer.WriteString("* [" + n + "](" + path + "/" + strings.ReplaceAll(fn, " ", "$20") + ")\n")
	}
}


