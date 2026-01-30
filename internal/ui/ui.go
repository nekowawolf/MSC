package ui

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

func ShowBanner() {
	fmt.Print("\033[H\033[2J")

	cyan := color.New(color.FgCyan).SprintFunc()
	white := color.New(color.FgWhite).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()

	banner := `
 __  ___   _____   ______
/  |/  /  / ___/  / ____/   MultiSendChain v1.0
/ /|_/ /   \__ \  / /        Fast • Reliable • Efficient
/ /  / /   ___/ / / /___     Developed by: nekowawolf
/_/  /_/   /____/  \____/
`
	fmt.Println(cyan(banner))
	fmt.Println(white("------------------------------------------------------"))
	fmt.Printf("> [MSC] Status   : %s\n", green("READY"))
	fmt.Printf("> [MSC] Network  : %s\n", white("Multi-Chain Enabled"))
	fmt.Printf("> [MSC] Security : %s\n", green("Private Key Loaded"))
	fmt.Println(white("------------------------------------------------------"))
	fmt.Println()
    time.Sleep(500 * time.Millisecond)
}