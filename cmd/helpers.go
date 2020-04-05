package cmd

import (
	"fmt"
	"os"
)

func er(msg interface{}) {
	fmt.Println("Error:", msg)
	fmt.Println("Wow we have an error")
	os.Exit(1)
}
