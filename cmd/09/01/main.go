package main

import (
	"fmt"
	"os"

	common "halamix2.pl/advent_of_code_24/cmd/09"
)

func main() {
	disk, err := common.ParseInput("cmd/09/input.txt")
	if err != nil {
		fmt.Printf("failed to load data:%s\n", err)
		os.Exit(1)
	}

	disk.Compress()

	checksum := disk.GetChecksum()

	fmt.Printf("Checksum: %d\n", checksum)
}
