package main

import (
	"flag"
	"fmt"
	"os"
)

var posts = []string{"Using Docker and Docker Compose for Local Development and Small Deployments", "Trying Clean Architecture on Golang", "Understanding Depenedency Management and Package Management in Golang", "Dancing with Go's Mutexes", "Selenium: easy as pie", "Go Best Practises - Testing"}

var options struct {
	postCount int
	version   bool
}

func main() {
	var Usage = func() {
		fmt.Fprintf("os.Stderr, Usage: %s[post_count]\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.IntVar(&options.postCount, "count", 1, "Enter the number of post you would love to see")
	flag.BoolVar(&options.version, "version", false, "View the version of this application")
	flag.Parse()

	var intPtr *int
	intPtr = &options.postCount

	if intPtr != nil {
		if *intPtr > len(posts) {
			fmt.Printf("%s\n", "Number of post specified is greater than what we currently have...")
			os.Exit(0)
		}
		buildPost := getPostByCount(posts, *intPtr)
		for post := range buildPost {
			fmt.Printf("%s\n", buildPost[post])
		}
		os.Exit(0)
	}

	if options.version {
		fmt.Printf("%s", "goCmd installed version 0.0.1\n")
		os.Exit(0)
	}

	if len(flag.Args()) == 0 {
		Usage()
		os.Exit(1)
	}

}

func getPostByCount() {

}
