package domain

import (
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v2"
)

type Item struct {
	ID          int    `yaml:"id"`
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	RawHeight   int    `yaml:"height"`
	RawWidth    int    `yaml:"width"`
	RawShape    string `yaml:"rawshape"`
	Shape       Shape
	Value       int `yaml:"value"`
}

type Items struct {
	Items []*Item `yaml:"items"`
}

func (i *Items) parseYAML(filename string) error {
	var items Items

	// create reader for file
	fileHandle, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}

	// read the file
	fileContent, err := io.ReadAll(fileHandle)
	if err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	// unmarshal the file content into the item struct
	err = yaml.Unmarshal(fileContent, &items)
	if err != nil {
		return fmt.Errorf("error unmarshalling file: %w", err)
	}

	for _, item := range items.Items {
		// parse the raw shape into a matrix
		parsedShape := &Shape{}
		parsedShape.parseShape(item.RawShape, item.RawWidth, item.RawHeight)
		item.Shape = *parsedShape
		if err != nil {
			return fmt.Errorf("error parsing shape: %w", err)
		}
	}

	return nil
}
