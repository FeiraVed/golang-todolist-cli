package core

import "fmt"

func Banner() {
	banner := `%s
██╗  ██╗███████╗██╗     ██╗      █████╗
██║  ██║██╔════╝██║     ██║     ██╔══██╗
███████║█████╗  ██║     ██║     ██║  ██║%s
██╔══██║██╔══╝  ██║     ██║     ██║  ██║
██║  ██║███████╗███████╗███████╗╚█████╔╝
╚═╝  ╚═╝╚══════╝╚══════╝╚══════╝ ╚════╝%s
╔═══════════════════════════════════════╗
║  %s❒ %sGolang Todolist CLI%s                ║
╚═══════════════════════════════════════╝%s

	`
	fmt.Printf(banner, Red, White, Yellow, Green, Cyan, Yellow, White)
}
