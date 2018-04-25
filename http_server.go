package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	t := time.Now()
	// Mon Jan 2 15:04:05 -0700 MST 2006
	logFileName := "access-" + t.Format("2006-01-02T150405") + ".log"
	f, err := os.OpenFile(logFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	lsn, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}
	defer lsn.Close()
	for {
		conn, err := lsn.Accept()
		if err != nil {
			log.Println(err.Error())
			continue
		}
		go handel(conn)
	}

}

func handel(conn net.Conn) {
	defer conn.Close()
	request(conn)
}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		log.Println(ln)
		if i == 0 {
			handelReq(conn, ln)
		}
		if ln == "" {
			break
		}
		i++
	}
}

func handelReq(conn net.Conn, ln string) {
	m := strings.Fields(ln)[0]
	u := strings.Fields(ln)[1]

	if m == "GET" && u == "/hello" {
		procGo(conn)
	}

	if m == "GET" {
		proc(conn, u)
	}

	if m == "POST" && u == "/apply" {
		applyProcess(conn)
	}
}

func proc(conn net.Conn, u string) {
	fmt.Fprintf(conn, "you request %s", u)
}

func procGo(conn net.Conn) {
	body := `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title>Hello World!</title>
	</head>
	<body>
	<h1>Hello World Try GO!</h1>
	<a href="/apply">apply</a><br>
	<form method="POST" action="/apply">
	<input type="submit" value="apply">
	</form>
	</body>
	</html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)

}

func applyProcess(conn net.Conn) {

	body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body>
	<strong>APPLY PROCESS</strong><br>
	<a href="/hello">hello</a><br>
	</body></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
