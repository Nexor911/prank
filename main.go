package main

import (
	    "fmt"
		"math/rand"
		"os/exec"
		"runtime"
		"time"
)

func Openbrowser(url string) {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "rundll132"
		args = []string{"url.dll,FileProtocolHandler", url}
	case "darwin":
		cmd = "open"
		args = []string{url}
	default:
		cmd = "xdg-open"
		args = []string{url}
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
		Openbrowser("https://media4.giphy.com/media/v1.Y2lkPTc5MGI3NjExOXM0dzlldXE1MHczeXg2dXU5M3dlbW4ycGtuM2VyOWZ4bTF5YmJ3NiZlcD12MV9pbnRlcm5hbF9naWZfYnlfaWQmY3Q9Zw/mupHrOkUJOAEidyAp8/giphy.gif")
	}
}
