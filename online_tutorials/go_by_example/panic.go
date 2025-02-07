package main

import "os"

func main() {
	panic("Prismatic core FAILING!")

	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}