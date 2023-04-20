package os

import (
	"log"
	"os"
)

func CreateFile() {
	filePath, err := os.Create("./test/creating.txt")
	if err != nil {
		log.Println(err)
	}

	defer filePath.Close()
}

func OpenFile() {
	filePath, err := os.Open("./test/creating.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer filePath.Close()
}

func ReadFileContents() {
	bytes, err := os.ReadFile("./test/creating.txt")
	if err != nil {
		log.Fatal(err)
	}
	fileText := string(bytes[:])
	log.Println(fileText)
}

func ReadFileChunkWise() {
	chunkSize := 2
	b := make([]byte, chunkSize)
	file, err := os.Open("./test/creating.txt")
	if err != nil {
		log.Fatal(err)
	}
	for {
		bytesRead, _ := file.Read(b)
		if bytesRead == 0 {
			break
		}

		log.Println(bytesRead, string(b))
	}
	file.Close()
}

func WriteFile() {
	content := "Hello, world!"
	err := os.WriteFile("./test/creating.txt", []byte(content), 0666)
	if err != nil {
		log.Fatal(err)
	}
}

func AppendToAFile() {
	content := "\nAdding a new line"
	file, err := os.OpenFile("./test/creating.txt", os.O_APPEND|os.O_WRONLY, 0644)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	file.Write([]byte(content))
}

func RemoveFile() {
	err := os.Remove("./test/remove.txt");
	if err != nil{
		log.Fatal(err);
	}
}