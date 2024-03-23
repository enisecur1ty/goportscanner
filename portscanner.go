package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	banner := `

	▓█████  ███▄    █  ██▓  ██████ ▓█████  ▄████▄   █    ██  ██▀███   ██▓▄▄▄█████▓▓██   ██▓
	▓█   ▀  ██ ▀█   █ ▓██▒▒██    ▒ ▓█   ▀ ▒██▀ ▀█   ██  ▓██▒▓██ ▒ ██▒▓██▒▓  ██▒ ▓▒ ▒██  ██▒
	▒███   ▓██  ▀█ ██▒▒██▒░ ▓██▄   ▒███   ▒▓█    ▄ ▓██  ▒██░▓██ ░▄█ ▒▒██▒▒ ▓██░ ▒░  ▒██ ██░
	▒▓█  ▄ ▓██▒  ▐▌██▒░██░  ▒   ██▒▒▓█  ▄ ▒▓▓▄ ▄██▒▓▓█  ░██░▒██▀▀█▄  ░██░░ ▓██▓ ░   ░ ▐██▓░
	░▒████▒▒██░   ▓██░░██░▒██████▒▒░▒████▒▒ ▓███▀ ░▒▒█████▓ ░██▓ ▒██▒░██░  ▒██▒ ░   ░ ██▒▓░
	░░ ▒░ ░░ ▒░   ▒ ▒ ░▓  ▒ ▒▓▒ ▒ ░░░ ▒░ ░░ ░▒ ▒  ░░▒▓▒ ▒ ▒ ░ ▒▓ ░▒▓░░▓    ▒ ░░      ██▒▒▒ 
	 ░ ░  ░░ ░░   ░ ▒░ ▒ ░░ ░▒  ░ ░ ░ ░  ░  ░  ▒   ░░▒░ ░ ░   ░▒ ░ ▒░ ▒ ░    ░     ▓██ ░▒░ 
	   ░      ░   ░ ░  ▒ ░░  ░  ░     ░   ░         ░░░ ░ ░   ░░   ░  ▒ ░  ░       ▒ ▒ ░░  
	   ░  ░         ░  ░        ░     ░  ░░ ░         ░        ░      ░            ░ ░     
										  ░                                        ░ ░     
	
`

	fmt.Println(banner)

	fmt.Print("Hedef adresi (örn: scanme.nmap.org): ")
	targetAddr, _ := reader.ReadString('\n')
	targetAddr = strings.TrimSpace(targetAddr)

	fmt.Print("Başlangıç portu: ")
	startPortStr, _ := reader.ReadString('\n')
	startPortStr = strings.TrimSpace(startPortStr)
	startPort, _ := strconv.Atoi(startPortStr)

	fmt.Print("Bitiş portu: ")
	endPortStr, _ := reader.ReadString('\n')
	endPortStr = strings.TrimSpace(endPortStr)
	endPort, _ := strconv.Atoi(endPortStr)

	ports := 0
	openPorts := []int{}

	for i := startPort; i <= endPort; i++ {
		addr := fmt.Sprintf("%s:%d", targetAddr, i)
		conn, err := net.DialTimeout("tcp", addr, 2*time.Second)
		if err == nil {
			fmt.Println(i, " Portu açık!!")
			conn.Close()
			ports++
			openPorts = append(openPorts, i)
		}
	}

	fmt.Println("=========================================")
	fmt.Println("=========================================")
	fmt.Println("Tarama tamamlandı.")
	fmt.Println(ports, " Port açık!")
	fmt.Println("Açık olan portlar:", openPorts)
	fmt.Println("=========================================")
	fmt.Println("=========================================")
}
