package util

import (
	"encoding/json"
	"io/fs"
	"os"
	"strings"

	"github.com/mlange-42/beecs/comp"
)

// PatchesFromFile reads patch configurations from a JSON file.
func PatchesFromFile(path string) ([]comp.PatchConfig, error) {
	file, err := os.Open(path)
	if err != nil {
		return []comp.PatchConfig{}, err
	}
	defer file.Close()

	var patches []comp.PatchConfig

	decoder := json.NewDecoder(file)
	decoder.DisallowUnknownFields()
	if err = decoder.Decode(&patches); err != nil {
		return []comp.PatchConfig{}, err
	}

	return patches, nil
}

// FloatArrayFromFile reads a slice of floats from a space-separated file.
func FloatArrayFromFile(f fs.FS, path string) ([]float64, error) {
	content, err := fs.ReadFile(f, path)
	if err != nil {
		return nil, err
	}
	strCont := string(content)
	strCont = strings.ReplaceAll(strCont, " \r\n", ",")
	strCont = strings.ReplaceAll(strCont, " \n", ",")
	strCont = strings.ReplaceAll(strCont, "\r\n", ",")
	strCont = strings.ReplaceAll(strCont, "\n", ",")
	strCont = strings.ReplaceAll(strCont, " ", ",")
	strCont = strings.TrimSuffix(strCont, ",")
	strCont = "[" + strCont + "]"

	result := []float64{}
	if err = json.Unmarshal([]byte(strCont), &result); err != nil {
		return nil, err
	}
	return result, nil
}
