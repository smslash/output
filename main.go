package main

import (
	"os"

	"git/ssengerb/ascii-art-output/logic"
)

func main() {
	length := len(os.Args)

	if length == 2 {
		if !logic.Characters(os.Args[1]) {
			logic.Error(2)
			return
		}
		logic.Default(os.Args[1])
	} else if length == 4 {
		if os.Args[1][:9] != "--output=" || os.Args[1][len(os.Args[1])-4:] != ".txt" {
			logic.Error(1)
			return
		}

		if !logic.Characters(os.Args[3]) {
			logic.Error(2)
			return
		}

		if os.Args[3] != "standard" && os.Args[3] != "shadow" && os.Args[3] != "thinkertoy" {
			logic.Error(3)
			return
		}

		logic.Check(os.Args)
	} else {
		logic.Error(1)
		return
	}
}
