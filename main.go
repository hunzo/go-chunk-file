package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	bigfile := "./file5M"

	file, err := os.Open(bigfile)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	fileInfo, _ := file.Stat()
	fmt.Printf("%v", fileInfo.Size())

	var fileSize int64 = fileInfo.Size()

	const fileChunk = 1 * (1 << 20)

	fmt.Printf("%v\n", fileChunk)
	fmt.Printf("%v\n", fileSize)

	totalPartsNum := uint64(math.Ceil(float64(fileSize) / float64(fileChunk)))
	fmt.Printf("Splitting to %d pieces.\n", totalPartsNum)

	start := 0

	for i := uint64(0); i < totalPartsNum; i++ {

		partSize := int(math.Min(fileChunk, float64(fileSize-int64(i*fileChunk))))
		partBuffer := make([]byte, partSize)

		file.Read(partBuffer)

		fmt.Printf("bytes %d-%d/%d\n", start, (start + partSize - 1), fileSize)
		fmt.Printf("%v\n", partSize)
		fmt.Printf("%s\n", partBuffer)

		start = start + partSize

		// write to disk
		fileName := "somebigfile_" + strconv.FormatUint(i, 10)
		_, err := os.Create(fileName)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// write/save buffer to disk
		ioutil.WriteFile(fileName, partBuffer, os.ModeAppend)

		fmt.Println("Split to : ", fileName)
	}
}
