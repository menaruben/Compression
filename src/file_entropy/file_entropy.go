package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

type FileInformation struct {
	FilePath               string
	CharProbabilities      [256]float64
	CharOccurences         [256]int
	CharInformationContent [256]float64
	AmountChars            int
	Entropy                float64
	FileContent            string
}

func (f *FileInformation) readFile() {
	content, err := os.ReadFile(f.FilePath)
	if err != nil {
		panic(err)
	}

	f.FileContent = string(content)
	f.AmountChars = len(f.FileContent)
}

func (f *FileInformation) getCharProbabilities() {
	// charProbabilities := [256]float64{}
	// charOccurences := [256]int{}

	var charAscii int = 0
	for _, char := range f.FileContent {
		charAscii = int(char)
		f.CharOccurences[charAscii]++
	}

	for i, charOccurence := range f.CharOccurences {
		f.CharProbabilities[i] = float64(charOccurence) / float64(f.AmountChars)
	}
}

func (f FileInformation) showProbabilitesNotEqualZero() {
	fmt.Println("Probabilities: " + f.FilePath)
	for i, probability := range f.CharProbabilities {
		if probability != 0 {
			fmt.Printf("%q: %f\n", i, probability)
		}
	}
	fmt.Println()
}

func (f *FileInformation) getCharInformationContent() {
	for i, probability := range f.CharProbabilities {
		if probability > 0.0 {
			f.CharInformationContent[i] = math.Log2(1 / probability)
		} else {
			f.CharInformationContent[i] = 0.0
		}
	}
}

func (f FileInformation) showCharInformationContentNotEqualZero() {
	fmt.Println("Information content: " + f.FilePath)
	for i, informationContent := range f.CharInformationContent {
		if informationContent != 0 {
			fmt.Printf("%q: %f\n", i, informationContent)
		}
	}
	fmt.Println()
}

func (f *FileInformation) getFileEntropy() {
	for i, probability := range f.CharProbabilities {
		if probability != 0 {
			f.Entropy += probability * f.CharInformationContent[i]
		}
	}
}

func (f FileInformation) showFileInformation() {
	fmt.Println("File: " + f.FilePath + ": ")
	fmt.Println("Amount of chars: ", f.AmountChars)
	entropyString := strconv.FormatFloat(f.Entropy, 'f', 6, 64)
	fmt.Println("File Entropy" + entropyString)

	for i := 0; i < 256; i++ {
		if f.CharProbabilities[i] != 0 {
			// uses q instead of c because there may be non-printable characters like \n etc. which
			// we would not see in the console if printed out. with q we can see them :)
			fmt.Printf("%q: P = %f, I =  %f\n", i, f.CharProbabilities[i], f.CharInformationContent[i])
		}
	}
	fmt.Println("Entropy: ", f.Entropy)
}

func (f *FileInformation) getFileInformation() {
	f.readFile()
	f.getCharProbabilities()
	f.getCharInformationContent()
	f.getFileEntropy()
}

func main() {
	// file1 := FileInformation{FilePath: "C:\\repos\\Compression\\src\\text_files\\text1.txt"}
	// file1.getFileInformation()
	// file1.showFileInformation()

	// file2 := FileInformation{FilePath: "C:\\repos\\Compression\\src\\text_files\\text2.txt"}
	// file2.getFileInformation()
	// file2.showFileInformation()

	// file3 := FileInformation{FilePath: "C:\\repos\\Compression\\src\\text_files\\text3.txt"}
	// file3.getFileInformation()
	// file3.showFileInformation()

	// file4 := FileInformation{FilePath: "C:\\repos\\Compression\\src\\text_files\\text4.txt"}
	// file4.getFileInformation()
	// file4.showFileInformation()

	file5 := FileInformation{FilePath: "C:\\repos\\Compression\\src\\text_files\\text5.txt"}
	file5.getFileInformation()
	file5.showFileInformation()
}
