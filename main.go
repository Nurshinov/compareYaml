package main

import (
	"bytes"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	var f1, f2 map[string]interface{}
	var b bytes.Buffer
	b1, err := ioutil.ReadFile(os.Args[1])
	b2, err := ioutil.ReadFile(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	yaml.Unmarshal(b1, &f1)
	yaml.Unmarshal(b2, &f2)
	compareMaps(f1, f2, "price")
	yamlEncoder := yaml.NewEncoder(&b)
	yamlEncoder.SetIndent(2)
	yamlEncoder.Encode(&f2)
	yaml.Marshal(&f2)
	os.WriteFile("config3.yml", b.Bytes(), 0666)

}

func compareMaps(m1, m2 map[string]interface{}, fk string) {
	for k, v := range m1 {
		if keyExist(m2, k) {
			switch t := v.(type) {
			case map[string]interface{}:
				compareMaps(t, m2[k].(map[string]interface{}), fk)
			}
		} else {
			m2[k] = v
		}
	}
}

func keyExist(m1 map[string]interface{}, key string) bool {
	_, ok := m1[key]
	return ok
}
