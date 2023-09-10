package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"

	"github.com/VictorMilhomem/golox/cmd/vm"
	"github.com/chzyer/readline"
)

func runFile(fpath string, vm *vm.VM) {
	bytes, err := ioutil.ReadFile(path.Base(fpath))
	source := string(bytes)
	if err != nil {
		log.Fatalf("could not open file %s", fpath)
		return
	}
	vm.Interpret(source)
}

func emitLine(line string, builder strings.Builder) strings.Builder {
	builder.WriteString(line + "\n")
	return builder
}

func repl(vm *vm.VM) {
	rl, err := readline.New("glox > ")
	if err != nil {
		log.Panic("failed to readline")
	}
	defer rl.Close()

	for {
		source := strings.Builder{}
		line, err := rl.Readline()
		if err != nil {
			log.Println("failed to readline")
			continue
		}
		source = emitLine(line, source)
		vm.Interpret(source.String())
	}
}

func main() {
	args := os.Args[1:]
	vm := vm.NewVm()
	vm.InitVM()
	if len(args) > 1 {
		fmt.Println("Usage: glox [script]")
		os.Exit(64)
	} else if len(args) == 1 {
		runFile(args[0], vm)
	} else {
		repl(vm)
	}
	vm.FreeVM()
}
