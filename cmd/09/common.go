package common

import (
	"fmt"
	"os"
	"strconv"
)

/*
file, free space, alternating , length in blocks
each file has uniq ID based on order before rearrangement, starting with ID 0

move file blocks at a time from end of the disk to the leftmost free space

checksum - sum += (block position * fileID), skip for empty
*/

type Block struct {
	data int
}

type Disk struct {
	data []Block
}

func (d *Disk) Print() {
	for _, b := range d.data {
		if b.data != -1 {
			fmt.Printf("%d", b.data)
		} else {
			fmt.Print(".")
		}
	}
	fmt.Println()
}

func (d *Disk) swap(a, b int) {
	tmp := Block{data: d.data[a].data}
	d.data[a] = d.data[b]
	d.data[b] = tmp
}

func (d *Disk) getLastFileID() int {
	lastFile := 0
	for _, block := range d.data {
		if block.data > lastFile {
			lastFile = block.data
		}
	}
	return lastFile
}

func (d *Disk) getFilePosition(id int) int {
	for i, block := range d.data {
		if block.data == id {
			return i
		}
	}
	return -1
}

func (d *Disk) getFileSize(id int) int {
	sum := 0
	for _, block := range d.data {
		if block.data == id {
			sum++
		}
	}
	return sum
}

func (d *Disk) getFirstFreeSpace(wantedSize int) int {
	size := 0
	for i, block := range d.data {
		if block.data == -1 {
			size++
			if size >= wantedSize {
				return i + 1 - size
			}
		} else {
			size = 0
		}
	}
	return -1
}

func (d *Disk) getFirstFreePosition() int {
	return d.getFilePosition(-1)
}

func (d *Disk) getLastDatablockPosition() int {
	for i := len(d.data) - 1; i >= 0; i-- {
		if d.data[i].data != -1 {
			return i
		}
	}
	return -1
}

func (d *Disk) Compress() {
	for {
		firstFree := d.getFirstFreePosition()
		lastData := d.getLastDatablockPosition()
		if firstFree > lastData {
			break
		}
		d.swap(firstFree, lastData)
	}
}

func (d *Disk) Defrag() {
	biggestFileID := d.getLastFileID()
	for currentFile := biggestFileID; currentFile >= 0; currentFile-- {
		fileSize := d.getFileSize(currentFile)
		freeSpacePos := d.getFirstFreeSpace(fileSize)
		filePos := d.getFilePosition(currentFile)
		if freeSpacePos < 0 || freeSpacePos > filePos {
			continue
		}
		// file will fit, swap them
		for i := range fileSize {
			d.swap(filePos+i, freeSpacePos+i)
		}
	}
}

func (d *Disk) GetChecksum() int {
	checksum := 0
	for i, block := range d.data {
		if block.data != -1 {
			checksum += block.data * i
		}
	}
	return checksum
}

func ParseInput(filename string) (Disk, error) {
	compressed, err := getCompressed(filename)
	if err != nil {
		return Disk{}, fmt.Errorf("failed to get compressed: %s\n", err)
	}
	// decompress
	diskData := make([]Block, 0)
	id := 0
	file := true
	for _, r := range compressed {
		size, _ := strconv.Atoi(string(r))
		if file {
			diskData = append(diskData, generate(id, size)...)
			id++
		} else {
			//empty space
			diskData = append(diskData, generate(-1, size)...)
		}
		file = !file
	}

	return Disk{data: diskData}, nil
}

func generate(id, size int) []Block {
	blocks := make([]Block, 0)
	for i := 0; i < size; i++ {
		blocks = append(blocks, Block{data: id})
	}
	return blocks
}

func getCompressed(filename string) ([]rune, error) {
	f, err := os.ReadFile(filename)
	if err != nil {
		return []rune{}, fmt.Errorf("failed to load input: %s\n", err)
	}
	runes := make([]rune, 0)
	for _, r := range string(f) {
		runes = append(runes, r)
	}

	return runes, nil
}
