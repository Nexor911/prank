package main

import (
        "fmt"
	"math/rand"
	"os/exec"
	"runtime"
	"time"
)

func Openvideo(path string) {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "rundll32"
		args = []string{"url.dll,FileProtocolHandler", path}
	case "darwin":
		cmd = "open"
		args = []string{path}
	default:
		cmd = "xdg-open"
		args = []string{path}
	}

	exec.Command(cmd, args...).Start()
}

func main() {
	rand.Seed(time.Now().UnixNano())
	secret := rand.Intn(5) + 1

	var guess int
	fmt.Print("Угадай число от 1 до 5: ")
	fmt.Scanln(&guess)

	if guess == secret {
		fmt.Println("Ты угадал!")
	} else {
		fmt.Println("Не угадал!")
		Openvideo("screamer.mp4")
	}
}
