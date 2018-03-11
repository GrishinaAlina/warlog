package logparser

import (
	"regexp"
	"github.com/go-yaml/yaml"
	"os"
	"fmt"
	"bytes"
)

//структура конфига объекта
type ObjectConfig struct {
	ObjectAttribute        *regexp.Regexp
	NewlineObjectBeginning bool
	Objects                map[string]map[string]string
}

//буффер для сохранения части последнего объекта
var buffer = make([]byte, 0)

//функция для парсинга конфига объекта
func parseObjectConfig(pathToCfg string) (*ObjectConfig, error) {
	cfgFile, err := os.Open(pathToCfg)
	if err != nil {
		return nil, err
	}
	decoder := yaml.NewDecoder(cfgFile)

	type parseObjectConfig struct {
		ObjectAttribute        Regexp                       `yaml:"object_predicate"`
		NewlineObjectBeginning Bool                         `yaml:"newline_object_beginning"`
		Objects                map[string]map[string]string `yaml:"objects"`
	}

	var parseResult parseObjectConfig
	err = decoder.Decode(&parseResult)
	if err != nil {
		return nil, err
	}

	return &ObjectConfig{parseResult.ObjectAttribute.Value,
		parseResult.NewlineObjectBeginning.Value,
		parseResult.Objects}, nil
}

//функция для отсечения части последнего объекта и сохранения ее буффер
func saveLastObjectPartInBuffer(splitBytes *[][]byte, objectConfig *ObjectConfig) {
	i := 0
	matchAllIndex := make([]int, 0)
	for i = len(*splitBytes) - 1; i >= 0; i-- {
		matchAllIndex = objectConfig.ObjectAttribute.FindIndex((*splitBytes)[i])
		if len(matchAllIndex) != 0 {
			break
		}
	}

	if objectConfig.NewlineObjectBeginning {
		buffer = bytes.Join((*splitBytes)[i:len(*splitBytes)], []byte(""))
		*splitBytes = (*splitBytes)[0:i]
		return
	}

	buffer = append(buffer, (*splitBytes)[i][matchAllIndex[1]:len((*splitBytes)[i])]...)
	buffer = append(buffer, bytes.Join((*splitBytes)[i+1:len(*splitBytes)], nil)...)

	(*splitBytes)[i] = (*splitBytes)[i][0:matchAllIndex[0]]
	*splitBytes = (*splitBytes)[0:i+1]
	//fmt.Println("buffer = " + string(buffer))
}

//функция для получения итогового регулярного выражения для каждого объекта из конфига
func getTotalRegularExpression(objects *map[string]map[string]string) (*map[string]*regexp.Regexp, error) {
	var result = make(map[string]*regexp.Regexp)
	var err error
	for object := range *objects {
		var regStr string
		for field := range (*objects)[object] {
			regStr += fmt.Sprintf("(%s)", (*objects)[object][field])
		}
		result[object], err = regexp.Compile(regStr)
		if err != nil {
			return nil, err
		}
	}
	return &result, nil
}

//функция для разделения исходного набора байт на объекты
func getLogObjects(splitBytes *[][]byte, objectConfig *ObjectConfig) (*[][]byte) {
	matchAllIndex := make([]int, 0)
	result := make([][]byte, 0)
	lastI := -1
	lastMatchAllIndex := make([]int, 0)

	if len(*splitBytes) == 1 {
		result = append(result, (*splitBytes)[0])
		return &result
	}

	if objectConfig.NewlineObjectBeginning {
		for i := 1; i < len(*splitBytes); i++ {
			matchAllIndex = objectConfig.ObjectAttribute.FindIndex((*splitBytes)[i])
			if len(matchAllIndex) == 0 {
				continue
			}
			result = append(result, bytes.Join((*splitBytes)[lastI+1:i+1], nil))
			lastI = i
		}
		return &result
	}

	for i := 1; i < len(*splitBytes); i++ {
		if i == len(*splitBytes) - 1 {
			result = append(result,
				append((*splitBytes)[i-1][lastMatchAllIndex[1]:len((*splitBytes)[i-1])], (*splitBytes)[i]...))
			break
		}
		matchAllIndex = objectConfig.ObjectAttribute.FindIndex((*splitBytes)[i])
		if len(matchAllIndex) == 0 {
			continue
		}
		if i == 1 {
			result = append(result, append(bytes.Join((*splitBytes)[lastI+1:i], nil),
				(*splitBytes)[i][0:matchAllIndex[0]]...))
			lastI = i
			lastMatchAllIndex = matchAllIndex
			continue
		}
		result = append(result, append((*splitBytes)[i-1][lastMatchAllIndex[1]:len((*splitBytes)[i-1])],
			append(bytes.Join((*splitBytes)[lastI+1:i], nil),
			(*splitBytes)[i][0:matchAllIndex[0]]...)...))

		lastI = i
		lastMatchAllIndex = matchAllIndex
	}

	return &result
}

func ParseLogObjects(logObjects []byte, pathToCfg string) (*map[string][]byte, error) {
	objectConfig, err := parseObjectConfig(pathToCfg)
	if err != nil {
		return nil, err
	}

	if len(buffer) != 0 {
		logObjects = append(logObjects, buffer...)
		buffer = make([]byte, 0)
	}

	newlineSplitBytes := bytes.Split(logObjects, []byte("\n"))
	saveLastObjectPartInBuffer(&newlineSplitBytes, objectConfig)
	for i:= range newlineSplitBytes {
		fmt.Println(string(newlineSplitBytes[i]))
	}

	var totalReg *map[string]*regexp.Regexp
	totalReg, err = getTotalRegularExpression(&objectConfig.Objects)
	fmt.Println(totalReg)

	splitLogObjects := getLogObjects(&newlineSplitBytes, objectConfig)
	for i := range *splitLogObjects {
		fmt.Println(string((*splitLogObjects)[i]))
	}

	return nil, nil
}