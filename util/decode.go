package util

import (
	"encoding/json"
	"os"

	"github.com/mlange-42/beecs/comp"
)

func PatchesFromFile(path string) ([]comp.PatchConfig, error) {
	file, err := os.Open(path)
	if err != nil {
		return []comp.PatchConfig{}, err
	}

	var patches []comp.PatchConfig

	decoder := json.NewDecoder(file)
	decoder.DisallowUnknownFields()
	if err = decoder.Decode(&patches); err != nil {
		return []comp.PatchConfig{}, err
	}

	return patches, nil
}
