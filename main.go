package main

import (
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	var woDeleted []string
	file1 := readFile(os.Args[1])
	file2 := readFile(os.Args[2])
	file3 := readFile(os.Args[3])
	lines1 := strings.Split(string(file1), "\n")
	lines2 := strings.Split(string(file2), "\n")
	lines3 := strings.Split(string(file3), "\n")
	removeComments(lines1, lines2)
	m1, m2 := getChangedLines(lines1, lines2)
	deletedKeys := sortMaps(m2, m1)
	for index, v := range lines3 {
		if _, ok := deletedKeys[getPathForLine(lines3, uint(index))]; !ok {
			woDeleted = append(woDeleted, v)
		}
	}
	os.WriteFile(os.Args[3], []byte(strings.Join(woDeleted, "\n")), 0644)
	file3 = readFile(os.Args[3])
	lines3 = strings.Split(string(file3), "\n")
	newKeys := sortMaps(m1, m2)
	for k, v := range newKeys {
		keys := strings.Split(k, ".")
		if strings.Count(k, ".") < 2 {
			for i, line := range lines1 {
				if strings.Contains(getPathForLine(lines1, uint(i)), k) {
					lines3[len(lines3)-1] += "\n" + line
				}
			}
		}
		keys = keys[:len(keys)-1]
		keyPath := strings.Join(keys, ".")
		for k3, v3 := range lines3 {
			path := getPathForLine(lines3, uint(k3))
			if path == keyPath {
				for _, val := range v {
					lines3[k3] = v3 + "\n" + val
				}
			}
		}
		//log.Printf("Ключ %s был добавлен", k)
	}
	os.WriteFile(os.Args[3], []byte(strings.Join(lines3, "\n")), 0644)
}

func sortMaps(m1, m2 map[string][]string) map[string][]string {
	resultMap := map[string][]string{}
	for k, v := range m1 {
		if _, ok := m2[k]; !ok {
			resultMap[k] = append(resultMap[k], v...)
		}
	}
	return resultMap
}

func getChangedLines(lines1, lines2 []string) (map[string][]string, map[string][]string) {
	m1 := map[string][]string{}
	m2 := map[string][]string{}
	for k, v := range lines1 {
		path := getPathForLine(lines1, uint(k))
		m1[path] = append(m1[path], v)
	}
	for k, v := range lines2 {
		path := getPathForLine(lines2, uint(k))
		m2[path] = append(m2[path], v)
	}
	return m1, m2
}

func getPathForLine(lines []string, index uint) string {
	var pathArr []string
	var path string
	curIndent := getIndent(lines[index])
	for i := index; i > 0; i-- {
		if getIndent(lines[i]) == 0 {
			break
		} else if getIndent(lines[i-1]) < getIndent(lines[i]) && getIndent(lines[i]) <= curIndent {
			pathArr = append(pathArr, getKey(lines[i-1]))
			curIndent = getIndent(lines[i-1])
		}
	}
	for i := len(pathArr) - 1; i >= 0; i-- {
		path = path + "." + pathArr[i]
	}
	return path + "." + getKey(lines[index])
}

func getIndent(line string) int {
	var indent int
	for _, b := range []byte(line) {
		if b == 32 {
			indent++
		} else {
			break
		}
	}
	return indent
}

func getKey(line string) string {
	return strings.TrimSpace(strings.Split(line, ":")[0])
}

func removeComments(arrs ...[]string) {
	re := regexp.MustCompile(`\#.*$`)
	for _, arr := range arrs {
		for k, v := range arr {
			arr[k] = re.ReplaceAllString(v, "")
		}
	}
}

func readFile(filename string) []byte {
	f1, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return f1
}
