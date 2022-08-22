package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/fatih/color"
)

var (
	chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	ascii_art = `
 _______  .__  __                        
 \      \ |__|/  |________  ____
 /   |   \|  \   __\_  __ \/  _ \
/    |    \  ||  |  |  | \(  <_> )
\____|__  /__||__|  |__|   \____/
        \/
  _________      .__
 /   _____/ ____ |__|_____   ___________
 \_____  \ /    \|  \____ \_/ __ \_  __ \
 /        \   |  \  |  |_> >  ___/|  | \/
/_______  /___|  /__|   __/ \___  >__|
        \/     \/   |__|        \/
`

	red    = color.New(color.FgRed)
	green  = color.New(color.FgGreen)
	cyan   = color.New(color.FgCyan).Add(color.Bold)
	yellow = color.New(color.FgYellow)
)

func randomNumber() int {
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(len(chars)-0) - 0
}

func NitroGen() string {
	code := []rune("https://discord.gift/****************")
	for i := 0; i < 16; i++ {
		code[i+21] = chars[randomNumber()]
	}

	return string(code)
}

func CheckNitro(code string) bool {
	r, err := http.Get("https://discord.com/api/v6/entitlements/gift-codes/" + code + "?with_application=false&with_subscription_plan=true")
	if err != nil {
		panic(err)
	}

	if r.StatusCode == 200 {
		return true
	} else {
		return false
	}
}

func main() {
	color.Red(ascii_art)

	pwd, _ := os.Getwd()

	fmt.Printf("Generating nitro codes... \n(I will write them to %snitro.txt)\n", pwd+string(os.PathSeparator))

	for {
		code := NitroGen()
		valid := CheckNitro(code)

		if valid {
			red.Print("* ")
			cyan.Print(code)
			yellow.Print(" -> ")
			green.Println("[VALID]")

			if _, err := os.Stat("nitro.txt"); err == nil {
				f, err := os.OpenFile("nitro.txt", os.O_APPEND|os.O_WRONLY, 0644)
				defer f.Close()

				if err != nil {
					red.Add(color.Bold).Println("No se pudo guardar el codigo", code, "en nitro.txt")
					panic(err)
				}

				f.Write([]byte(code + "\n"))
			} else {
				f, err := os.Create("nitro.txt")
				defer f.Close()

				if err != nil {
					red.Add(color.Bold).Println("No se pudo guardar el codigo", code, "en nitro.txt")
					panic(err)
				}

				f.Write([]byte(code + "\n"))
			}
		} else {
			green.Print("* ")
			cyan.Print(code)
			yellow.Print(" -> ")
			red.Println("[INVALID]")
		}

	}
}
