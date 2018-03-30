package main

import "regexp"

const text = "my email is chore.jujuju@gmail.com"

func main() {
	regexp.MustCompile("chore.jujuju@gmail.com")
}
