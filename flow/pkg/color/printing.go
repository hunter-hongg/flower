package color 

import (
	"fmt"
	"os"
)

func Error(msg string) {
	fmt.Println(Red("Error ") + msg)
	os.Exit(1)
}

func Errore(msg error) {
	if msg == nil { return }
	fmt.Println(Red("Error ") + msg.Error())	
	os.Exit(1)
}

func Warning(msg string) {
	fmt.Println(Yellow("Warning ") + msg)
}

func Info(msg string) {
	fmt.Println(Blue("Info ") + msg)
}

func Step(msg string) {
	fmt.Println(Cyan("Step ") + msg)
}