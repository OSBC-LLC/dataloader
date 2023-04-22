package string

import (
	"path/filepath"

	"github.com/OSBC-LLC/dataloader/pkg/number"
)

// GetRandomFirstname: Get a random male, female, or either first name.
func GetRandomMaleFirstName() (string, error) {
	rootPath, err := getRootPath()
	if err != nil {
		return "", err
	}
	fileName := filepath.Join(rootPath, "data", "firstNames", "male.txt")
	data, err := readFileLines(fileName)
	if err != nil {
		return "", err
	}
	return data[number.GetRandomInt(0, len(data)-1)], nil
}

func GetRandomFemaleFirstName() (string, error) {
	rootPath, err := getRootPath()
	if err != nil {
		return "", err
	}
	fileName := filepath.Join(rootPath, "data", "firstNames", "female.txt")
	data, err := readFileLines(fileName)
	if err != nil {
		return "", err
	}
	return data[number.GetRandomInt(0, len(data)-1)], nil
}

func GetRandomMaleOrFemaleFirstName() (string, error) {
	rootPath, err := getRootPath()
	if err != nil {
		return "", err
	}
	fileName := filepath.Join(rootPath, "data", "firstNames", "male.txt")
	allData, err := readFileLines(fileName)
	if err != nil {
		return "", err
	}
	fileName = filepath.Join(rootPath, "data", "firstNames", "female.txt")
	femaleData, err := readFileLines(fileName)
	if err != nil {
		return "", err
	}
	allData = append(allData, femaleData...)
	return allData[number.GetRandomInt(0, len(allData)-1)], nil
}

func GetRandomLastName() (string, error) {
	rootPath, err := getRootPath()
	if err != nil {
		return "", err
	}
	fileName := filepath.Join(rootPath, "data", "lastNames", "lastNames.txt")
	data, err := readFileLines(fileName)
	if err != nil {
		return "", err
	}
	return data[number.GetRandomInt(0, len(data)-1)], nil
}
