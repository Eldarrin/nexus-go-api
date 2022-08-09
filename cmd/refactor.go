package main

import (
	"log"
	"os"
	"strings"
)

const directoryStructure = "client\\repository_management"

func reworkCreateRepoParams(fileName string, sType []string) {

	// Read file to byte slice
	data, err := os.ReadFile(directoryStructure + "\\" + fileName)
	if err != nil {
		log.Fatal(err)
	}

	s := string(data)

	// find model name
	idx := strings.Index(s, "Body *models.")
	if idx > 0 {
		s2 := s[idx+len("Body *models."):]
		idx2 := strings.Index(s2, "APIRequest")
		if idx2 == -1 {
			log.Printf("Could not find model name")
		}
		modelName := s2[0:idx2]

		log.Printf("Model name: %s", modelName)

		oldName := "CreateRepository"
		switch len(sType[1]) {
		case 10:
		case 11:
			oldName = "CreateRepository" + sType[1][len(sType[1])-1:len(sType[1])]
		case 12:
			oldName = "CreateRepository" + sType[1][len(sType[1])-2:len(sType[1])]
		case 13:
			oldName = "CreateRepository" + sType[1][len(sType[1])-3:len(sType[1])]
		}

		newName := "Create" + modelName
		log.Printf("Old name: %s", oldName)
		log.Printf("New name: %s", newName)

		nS := strings.ReplaceAll(s, oldName, newName)

		log.Printf("New string: %s", nS)

	}
}

func reworkCreateRepoResp(fileName string, sType []string) {

	// Read file to byte slice
	data, err := os.ReadFile(directoryStructure + "\\" + fileName)
	if err != nil {
		log.Fatal(err)
	}

	s := string(data)

	idx := strings.Index(s, "POST /v1/repositories/")
	if idx > 0 {
		s2 := s[idx+len("POST /v1/repositories/"):]
		idx2 := strings.Index(s2, "]")
		if idx2 == -1 {
			log.Printf("Could not find model name")
		}
		modelName := s2[0:idx2]

		modelName = strings.Title(modelName)

		modelName = strings.ReplaceAll(modelName, "/", "")

		log.Printf("Model name: %s", modelName)
		//sModel := strings.Split(modelName, "/")

		oldName := "CreateRepository"
		switch len(sType[1]) {
		case 10:
		case 11:
			oldName = "CreateRepository" + sType[1][len(sType[1])-1:len(sType[1])]
		case 12:
			oldName = "CreateRepository" + sType[1][len(sType[1])-2:len(sType[1])]
		case 13:
			oldName = "CreateRepository" + sType[1][len(sType[1])-3:len(sType[1])]
		}

		newName := "Create" + modelName + "Repository"
		log.Printf("Old name: %s", oldName)
		log.Printf("New name: %s", newName)

		nS := strings.ReplaceAll(s, oldName, newName)

		log.Printf("New string: %s", nS)

	}
}

func reworkGetRepoUpdate(fileName string, sType []string) {
	// Read file to byte slice
	data, err := os.ReadFile(directoryStructure + "\\" + fileName)
	if err != nil {
		log.Fatal(err)
	}

	s := string(data)

	idx := strings.Index(s, "GET /v1/repositories/")
	if idx > 0 {
		s2 := s[idx+len("GET /v1/repositories/"):]
		idx2 := strings.Index(s2, "{")
		if idx2 == -1 {
			log.Printf("Could not find model name")
		}
		modelName := s2[0:idx2]

		modelName = strings.Title(modelName)

		modelName = strings.ReplaceAll(modelName, "/", "")

		log.Printf("Model name: %s", modelName)
		//sModel := strings.Split(modelName, "/")

		oldName := "GetRepository"
		switch len(sType[1]) {
		case 10:
		case 11:
			oldName = "GetRepository" + sType[1][len(sType[1])-1:len(sType[1])]
		case 12:
			oldName = "GetRepository" + sType[1][len(sType[1])-2:len(sType[1])]
		case 13:
			oldName = "GetRepository" + sType[1][len(sType[1])-3:len(sType[1])]
		}

		newName := "Get" + modelName + "Repository"
		log.Printf("Old name: %s", oldName)
		log.Printf("New name: %s", newName)

		nS := strings.ReplaceAll(s, oldName, newName)

		log.Printf("New string: %s", nS)

		// repeat replace for params file

	}
}

func reworkUpdateRepoParams(fileName string, sType []string) {

	// Read file to byte slice
	data, err := os.ReadFile(directoryStructure + "\\" + fileName)
	if err != nil {
		log.Fatal(err)
	}

	s := string(data)

	// find model name
	idx := strings.Index(s, "Body *models.")
	if idx > 0 {
		s2 := s[idx+len("Body *models."):]
		idx2 := strings.Index(s2, "APIRequest")
		if idx2 == -1 {
			log.Printf("Could not find model name")
		}
		modelName := s2[0:idx2]

		log.Printf("Model name: %s", modelName)

		oldName := "UpdateRepository"
		switch len(sType[1]) {
		case 10:
		case 11:
			oldName = "UpdateRepository" + sType[1][len(sType[1])-1:len(sType[1])]
		case 12:
			oldName = "UpdateRepository" + sType[1][len(sType[1])-2:len(sType[1])]
		case 13:
			oldName = "UpdateRepository" + sType[1][len(sType[1])-3:len(sType[1])]
		}

		newName := "Update" + modelName
		log.Printf("Old name: %s", oldName)
		log.Printf("New name: %s", newName)

		nS := strings.ReplaceAll(s, oldName, newName)

		log.Printf("New string: %s", nS)

	}
}

func reworkUpdateRepoResp(fileName string, sType []string) {

	// Read file to byte slice
	data, err := os.ReadFile(directoryStructure + "\\" + fileName)
	if err != nil {
		log.Fatal(err)
	}

	s := string(data)

	idx := strings.Index(s, "PUT /v1/repositories/")
	if idx > 0 {
		s2 := s[idx+len("PUT /v1/repositories/"):]
		idx2 := strings.Index(s2, "{")
		if idx2 == -1 {
			log.Printf("Could not find model name")
		}
		modelName := s2[0:idx2]

		modelName = strings.Title(modelName)

		modelName = strings.ReplaceAll(modelName, "/", "")

		log.Printf("Model name: %s", modelName)
		//sModel := strings.Split(modelName, "/")

		oldName := "UpdateRepository"
		switch len(sType[1]) {
		case 10:
		case 11:
			oldName = "UpdateRepository" + sType[1][len(sType[1])-1:len(sType[1])]
		case 12:
			oldName = "UpdateRepository" + sType[1][len(sType[1])-2:len(sType[1])]
		case 13:
			oldName = "UpdateRepository" + sType[1][len(sType[1])-3:len(sType[1])]
		}

		newName := "Update" + modelName + "Repository"
		log.Printf("Old name: %s", oldName)
		log.Printf("New name: %s", newName)

		nS := strings.ReplaceAll(s, oldName, newName)

		log.Printf("New string: %s", nS)

	}
}

func main() {
	files, err := os.ReadDir(directoryStructure)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {

		if !file.IsDir() {
			sType := strings.Split(file.Name(), "_")
			switch sType[0] {
			case "create":
				if sType[2] == "responses.go" {
					//reworkCreateRepoResp(file.Name(), sType)
				} else {
					//reworkCreateRepoParams(file.Name(), sType)
				}
			case "get":
				if sType[2] == "responses.go" {
					//reworkGetRepoUpdate(file.Name(), sType)
				} else {
					//reworkCreateRepoParams(file.Name(), sType)
				}
			case "update":
				if sType[2] == "responses.go" {
					//reworkUpdateRepoResp(file.Name(), sType)
				} else {
					reworkUpdateRepoParams(file.Name(), sType)
				}
			}

		}
	}

}
