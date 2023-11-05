package main

import (
	"Compression/fileEntropy"
)

func main() {
	// testing file entropy calculation
	file5 := fileEntropy.FileInformation{FilePath: "C:\\repos\\Compression\\src\\text_files\\text5.txt"}
	file5.GetFileInformation()
	file5.ShowFileInformation()
}
