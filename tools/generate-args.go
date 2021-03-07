package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var err error
	defer func(err *error) {
		if *err != nil {
			log.Println((*err).Error())
			os.Exit(1)
		}
	}(&err)

	var buf []byte
	if buf, err = ioutil.ReadFile("args.txt"); err != nil {
		return
	}

	out := &bytes.Buffer{}
	out.WriteString("//go:generate go run tools/generate-args.go\n")
	out.WriteString("package cardanocli\n")

	lines := strings.Split(string(buf), "\n")
	for i, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		splits := strings.Split(line, " ")
		if len(splits) != 2 {
			err = fmt.Errorf("invalid line: %d=%s", i+1, line)
			return
		}
		var arg string
		var num int
		arg = strings.TrimSpace(splits[0])
		if num, err = strconv.Atoi(strings.TrimSpace(splits[1])); err != nil {
			return
		}

		out.WriteString(fmt.Sprintf("\nfunc (c *Cmd) %s(%s) *Cmd {\n", ArgToFuncName(arg), NumToArgsIn(num)))
		out.WriteString(fmt.Sprintf("    return c.Append(%s%s)\n", strconv.Quote(arg), NumToArgsOut(num)))
		out.WriteString("}\n")

	}

	err = ioutil.WriteFile("args.go", out.Bytes(), 0640)
}

func ArgToFuncName(arg string) string {
	name := ""
	if strings.HasPrefix(arg, "--") {
		name = "Opt"
		arg = strings.TrimPrefix(arg, "--")
	}
	splits := strings.Split(arg, "-")
	for _, split := range splits {
		split = strings.TrimSpace(split)
		if split == "" {
			continue
		}
		name = name + strings.Title(split)
	}
	return name
}

func NumToArgsIn(num int) string {
	if num == 0 {
		return ""
	}
	var ins []string
	for i := 0; i < num; i++ {
		ins = append(ins, fmt.Sprintf("arg%d string", i+1))
	}
	return strings.Join(ins, ", ")
}

func NumToArgsOut(num int) string {
	if num == 0 {
		return ""
	}
	var ins []string
	for i := 0; i < num; i++ {
		ins = append(ins, fmt.Sprintf("arg%d", i+1))
	}
	return ", " + strings.Join(ins, ", ")
}
