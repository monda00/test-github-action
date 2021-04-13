package main

import (
	"os"

	aw "github.com/deanishe/awgo"
)

var wf *aw.Workflow

func init() {
	wf = aw.New()
}

func run() {
	wf.NewItem("hoge hoge").Arg("hoge arg").Valid(true)
	wf.NewItem("new item").Arg("new item arg").Valid(true)

	args := os.Args
	if len(args) > 1 {
		wf.NewItem(args[1]).Arg(args[1] + " arg").Valid(true)
	}
	wf.SendFeedback()
}

func main() {
	wf.Run(run)
}
