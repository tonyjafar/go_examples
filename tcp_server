package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	lsn, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}
	defer lsn.Close()

	for {
		conn, err := lsn.Accept()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(conn, "Connected to Server\nWelcome to DBOnMem\nType quitNow to exit\n")
		go handel(conn)
	}
}

func handel(conn net.Conn) {
	input := make(map[string]string)
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		text := strings.ToUpper(scanner.Text())
		if strings.Contains(text, "QUITNOW") {
			conn.Close()
		}
		s := strings.Split(text, " ")
		switch t := s[0]; t {
		case "GET":
			if input[s[1]] == "" {
				fmt.Fprintf(conn, "No value for %s\n", s[1])
			}
			fmt.Fprintf(conn, "%s\n", input[s[1]])
		case "SET":
			if len(s) != 3 {
				fmt.Fprintf(conn, "usage: SET KEY VALUE\n")
				continue
			}
			input[s[1]] = s[2]
		case "DEL":
			if len(s) != 2 {
				fmt.Fprintf(conn, "usage: DEL KEY\n")
				continue
			}
			delete(input, s[1])
		case "VIEW":
			fmt.Fprintf(conn, "Key     Value\n")
			for k, v := range input {
				fmt.Fprintf(conn, "%s     %s\n", k, v)
			}
		case "DELALL":
			input = make(map[string]string)
		default:
			fmt.Fprintf(conn, "USE SET/GET/DEL/DELLALL/VIEW\n")
		}

	}
}
