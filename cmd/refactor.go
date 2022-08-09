package main

//TODO: refactor this function

import (
	"log"
	"os"
	"strings"
)

const directoryStructure = "client\\repository_management"

var clientFile string

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
		idx2 := strings.Index(s2, "RepositoryAPIRequest")
		if idx2 == -1 {
			log.Printf("Could not find model name")
		}
		modelName := s2[0:idx2]

		modelName = handleSpecials(modelName)

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

		newName := "Create" + modelName + "Repository"
		log.Printf("Old name: %s", oldName)
		log.Printf("New name: %s", newName)

		nS := strings.ReplaceAll(s, oldName, newName)

		clientFile = strings.ReplaceAll(clientFile, oldName, newName)

		log.Printf("New string: %s", nS)

		sO := []byte(nS) // convert string to byte slice
		os.WriteFile(directoryStructure+"\\tmp\\"+fileName, sO, 0644)

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

		modelName = handleSpecials(modelName)

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

		clientFile = strings.ReplaceAll(clientFile, oldName, newName)

		log.Printf("New string: %s", nS)

		sO := []byte(nS) // convert string to byte slice
		os.WriteFile(directoryStructure+"\\tmp\\"+fileName, sO, 0644)

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

		modelName = handleSpecials(modelName)

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

		clientFile = strings.ReplaceAll(clientFile, oldName, newName)

		log.Printf("New string: %s", nS)
		sO := []byte(nS) // convert string to byte slice
		os.WriteFile(directoryStructure+"\\tmp\\"+fileName, sO, 0644)

		// repeat replace for params file
		fileName = "get_" + sType[1] + "_parameters.go"
		dataParams, err := os.ReadFile(directoryStructure + "\\" + fileName)
		if err != nil {
			log.Fatal(err)
		}

		s = string(dataParams)
		nS = strings.ReplaceAll(s, oldName, newName)

		clientFile = strings.ReplaceAll(clientFile, oldName, newName)

		log.Printf("New string: %s", nS)
		sO = []byte(nS) // convert string to byte slice
		os.WriteFile(directoryStructure+"\\tmp\\"+fileName, sO, 0644)
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
		idx2 := strings.Index(s2, "RepositoryAPIRequest")
		if idx2 == -1 {
			log.Printf("Could not find model name")
		}
		modelName := s2[0:idx2]

		modelName = handleSpecials(modelName)

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

		newName := "Update" + modelName + "Repository"
		log.Printf("Old name: %s", oldName)
		log.Printf("New name: %s", newName)

		nS := strings.ReplaceAll(s, oldName, newName)

		clientFile = strings.ReplaceAll(clientFile, oldName, newName)

		log.Printf("New string: %s", nS)

		sO := []byte(nS) // convert string to byte slice
		os.WriteFile(directoryStructure+"\\tmp\\"+fileName, sO, 0644)

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

		modelName = handleSpecials(modelName)

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

		clientFile = strings.ReplaceAll(clientFile, oldName, newName)

		log.Printf("New string: %s", nS)

		sO := []byte(nS) // convert string to byte slice
		os.WriteFile(directoryStructure+"\\tmp\\"+fileName, sO, 0644)

	}
}

func handleSpecials(modelName string) string {
	var outType string
	if strings.Contains(modelName, "Hosted") {
		outType = "Hosted"
	} else if strings.Contains(modelName, "Proxy") {
		outType = "Proxy"
	} else if strings.Contains(modelName, "Group") {
		outType = "Group"
	} else {
	}

	// go = golang = Golang
	// gitlfs = Gitlfs = GitLfs
	// rubygems = Rubygems = RubyGems
	if strings.Contains(modelName, "go") || strings.Contains(modelName, "Go") || strings.Contains(modelName, "golang") {
		return "Golang" + outType
	}
	if strings.Contains(modelName, "gitlfs") || strings.Contains(modelName, "Gitlfs") {
		return "GitLfs" + outType
	}
	if strings.Contains(modelName, "rubygems") || strings.Contains(modelName, "Rubygems") {
		return "RubyGems" + outType
	}
	return modelName
}

func main() {
	files, err := os.ReadDir(directoryStructure)
	if err != nil {
		log.Fatal(err)
	}

	clientData, err := os.ReadFile(directoryStructure + "\\" + "repository_management_client.go")
	if err != nil {
		log.Fatal(err)
	}
	clientFile = string(clientData)

	for _, file := range files {

		if !file.IsDir() {
			sType := strings.Split(file.Name(), "_")
			switch sType[0] {
			case "create":
				if sType[2] == "responses.go" {
					reworkCreateRepoResp(file.Name(), sType)
				} else {
					reworkCreateRepoParams(file.Name(), sType)
				}
			case "get":
				if sType[2] == "responses.go" {
					reworkGetRepoUpdate(file.Name(), sType)
				}
			case "update":
				if sType[2] == "responses.go" {
					reworkUpdateRepoResp(file.Name(), sType)
				} else {
					reworkUpdateRepoParams(file.Name(), sType)
				}
			}

		}
	}

	sO := []byte(clientFile) // convert string to byte slice
	err = os.WriteFile(directoryStructure+"\\tmp\\"+"repository_management_client.go", sO, 0644)
	if err != nil {
		log.Fatal(err)
	}

}
