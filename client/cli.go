package main

import (
	"banco_de_espana/client/usecases"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	createClientCommand := flag.NewFlagSet("create", flag.ExitOnError)
	showClientCommand := flag.NewFlagSet("show", flag.ExitOnError)
	listClientCommand := flag.NewFlagSet("list", flag.ExitOnError)

	clientIDPtr := showClientCommand.Int64("id", 0, "Client id to fetch. (Required)")

	newClientNamePtr := createClientCommand.String("name", "", "New client name (Required)")
	newClientEmailPtr := createClientCommand.String("email", "", "New client email address. (Required)")
	newClientPhonePtr := createClientCommand.String("phone", "", "New client phone number")

	flag.Usage = cliUsage

	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "client":

		if len(os.Args) == 2 {
			entitieUsage()
			os.Exit(1)
		}

		switch os.Args[2] {
		case "create":
			createClientCommand.Parse(os.Args[3:])
		case "show":
			showClientCommand.Parse(os.Args[3:])
		case "list":
			listClientCommand.Parse(os.Args[3:])
			fmt.Println(os.Args[1:])
		}
	default:
		flag.Usage()
		os.Exit(1)
	}

	switch {
	case createClientCommand.Parsed():
		if err := usecases.CreateClient(*newClientNamePtr, *newClientEmailPtr, *newClientPhonePtr); err != nil {
			fmt.Printf("Can't create a client with params name %v, email: %v, phone %v. Got an error: %v\n", *newClientNamePtr, *newClientEmailPtr, *newClientPhonePtr, err)
			createClientCommand.PrintDefaults()
			os.Exit(1)
		}
	case showClientCommand.Parsed():
		if *clientIDPtr == 0 {
			println("Please provide a client ID.\n")
			showClientCommand.PrintDefaults()
			os.Exit(1)
		}

		fmt.Printf("showClientCommand.Parsed for id %v\n", *clientIDPtr)
	case listClientCommand.Parsed():
		listClientCommand.PrintDefaults()
		fmt.Println("listClientCommand.Parsed")
	}
}

func output() io.Writer {
	return flag.CommandLine.Output()
}

func entitieUsage() {
	avaliableCommands := []string{"create", "show", "list"}
	fmt.Fprintf(output(), "usage: %s %s <command> <options>\n\n", os.Args[0], os.Args[1])
	fmt.Fprintf(output(), "Allowed commands: %v\n", avaliableCommands)
}

func cliUsage() {
	fmt.Fprintln(output(), "~~~Banco de Espana~~~")
	fmt.Fprintf(output(), "Usage of %s: <entitie> <command> <args>\n", os.Args[0])
	fmt.Fprintln(output(), "Allowed entities: client|account|transfer")
}
