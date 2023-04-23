package string

import (
	"path/filepath"

	"github.com/OSBC-LLC/dataloader/pkg/number"
)

func buildFileName(path []string) (string, error) {
	rootPath, err := getRootPath()
	if err != nil {
		return "", err
	}
	return filepath.Join(rootPath, path[0], path[1], path[2]), nil
}

// GetRandomFirstname: Get a random male, female, or either first name.
func GetRandomMaleFirstName() (string, error) {
	fileName, err := buildFileName([]string{"data", "firstNames", "male.txt"})
	if err != nil {
		return "", err
	}
	data, err := readFileLines(fileName)
	if err != nil {
		return "", err
	}
	return data[number.GetRandomInt(0, len(data)-1)], nil
}

func GetRandomFemaleFirstName() (string, error) {
	fileName, err := buildFileName([]string{"data", "firstNames", "female.txt"})
	if err != nil {
		return "", err
	}
	data, err := readFileLines(fileName)
	if err != nil {
		return "", err
	}
	return data[number.GetRandomInt(0, len(data)-1)], nil
}

func GetRandomMaleOrFemaleFirstName() (string, error) {
	fileName, err := buildFileName([]string{"data", "firstNames", "male.txt"})
	if err != nil {
		return "", err
	}
	allData, err := readFileLines(fileName)
	if err != nil {
		return "", err
	}
	fileName, err = buildFileName([]string{"data", "firstNames", "female.txt"})
	if err != nil {
		return "", err
	}
	femaleData, err := readFileLines(fileName)
	if err != nil {
		return "", err
	}
	allData = append(allData, femaleData...)
	return allData[number.GetRandomInt(0, len(allData)-1)], nil
}

func GetRandomLastName() (string, error) {
	fileName, err := buildFileName([]string{"data", "lastNames", "lastNames.txt"})
	if err != nil {
		return "", err
	}
	data, err := readFileLines(fileName)
	if err != nil {
		return "", err
	}
	return data[number.GetRandomInt(0, len(data)-1)], nil
}

func GetRandomCountry() (string, error) {
	fileName, err := buildFileName([]string{"data", "countries", "countries.txt"})
	if err != nil {
		return "", err
	}
	data, err := readFileLines(fileName)
	if err != nil {
		return "", err
	}
	return data[number.GetRandomInt(0, len(data)-1)], nil
}

func GetRandomCountryCode(l int) (string, error) {
	var fileName string
	var err error
	switch l {
	case 2:
		fileName, err = buildFileName([]string{"data", "countries", "alpha2codes.txt"})
		if err != nil {
			return "", err
		}
	default:
		fileName, err = buildFileName([]string{"data", "countries", "alpha3codes.txt"})
		if err != nil {
			return "", err
		}
	}
	data, err := readFileLines(fileName)
	if err != nil {
		return "", err
	}
	return data[number.GetRandomInt(0, len(data)-1)], nil
}
