package main

import (
	"bytes"
	"image/png"
	"io/ioutil"
	"os"
	"strings"

	"github.com/toteki/wiz"
	"golang.org/x/image/webp"
)

func main() {
	wiz.MkDir("./output")
	/////////////////////////////////
	files, err := ioutil.ReadDir("./input")
	if err != nil {
		wiz.Red(err.Error())
		return
	}
	s := []string{}
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".webp") {
			s = append(s, file.Name())
		}
	}
	wiz.Green(s)
	/////////////////////////////////
	i := 0
	for i < len(s) {
		if wiz.SilentPrompt(s[i]) == "" {
			nemu := strings.TrimSuffix(s[i], ".webp")
			convert(nemu)
			i++
		} else {
			return
		}
	}
}

func convert(f string) {
	b, err := wiz.ReadFile("./input/" + f + ".webp")
	if err != nil {
		wiz.Red(err)
		return
	}
	i, err := webp.Decode(bytes.NewReader(b))
	if err != nil {
		wiz.Red(err)
		return
	}
	////////////
	_, chk := wiz.ReadFile("./output/" + f + ".png")
	if chk == nil {
		wiz.Red("Already exists")
		return
	}
	fi, err := os.Create("./output/" + f + ".png")
	if err != nil {
		wiz.Red(err)
		return
	}
	defer fi.Close()
	err = png.Encode(fi, i)
	if err != nil {
		wiz.Red(err)
		return
	}
	err = wiz.DeleteFile("./input/" + f + ".webp")
	if err != nil {
		wiz.Purple(err)
		return
	}
}
