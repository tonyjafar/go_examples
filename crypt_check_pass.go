package main

import (
	"flag"
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	pass := flag.String("password", "", "Enter Password to encrypt - used with --create")
	passP := flag.String("plain_pass", "", "Enter Plain Password to compare - used with --compare")
	passE := flag.String("hashed_pass", "", "Enter hashed Password to compare (use '' around it) - used with --compare")
	cost := flag.Int("cost", 10, "cost for encryption - used with --create")
	create := flag.Bool("create", false, "Create new hashed password")
	compare := flag.Bool("compare", false, "compare plain and hashed passwords")
	flag.Parse()
	if (*create && *compare) || (!*create && !*compare) {
		flag.PrintDefaults()
		fmt.Println("You should just one using from --create with  (--password) or --compare with (--hashed_pass and --plain_pass)")
		os.Exit(2)
	}

	if *create {
		createPass(*pass, *cost)
	}

	if *compare {
		compPass(*passP, *passE)
	}

}

func createPass(p string, c int) {

	if p == "" {
		flag.PrintDefaults()
		os.Exit(2)
	}

	encPass, err := bcrypt.GenerateFromPassword([]byte(p), c)
	if err != nil {
		fmt.Printf("Error creyting the password Error: %s", err)
		os.Exit(2)
	}
	fmt.Println(string(encPass))
}

func compPass(p1, p2 string) {
	if p1 == "" || p2 == "" {
		flag.PrintDefaults()
		os.Exit(2)
	}
	test := bcrypt.CompareHashAndPassword([]byte(p2), []byte(p1))
	if test == nil {
		fmt.Println("Matched")
	} else {
		fmt.Printf("%s", test)
	}
}
