package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)


func main() {
	var index int
	newLines := map[int]string{}
	for {
		file1, err := ioutil.ReadFile(os.Args[1])
		file2, err := ioutil.ReadFile(os.Args[2])
		if err != nil {
			log.Fatalln(err)
		}
		lines1 := strings.Split(string(file1), "\n")
		lines2 := strings.Split(string(file2), "\n")
		index = changeSecondFile(lines1,lines2,index)
		if len(lines1) == len(lines2) {
			break
		}
		newLines[index - 1] = lines1[index - 1]
	}

	for k,v := range newLines {

	}


}

func changeSecondFile(lines1,lines2 []string, lineIndex int) int {
		for i := lineIndex; i < len(lines2); i++ {
		if lines2[i] != lines1[i] {
			lines2[i] = "empty line\n" + lines2[i]
			lineIndex = i + 1
			break
		}
	}
	output := strings.Join(lines2, "\n")
	err := ioutil.WriteFile(os.Args[2], []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}
	return lineIndex
}
