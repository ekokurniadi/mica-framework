package formatter

import "mica-framework/entity"

type ExampleFormatter struct {
	ID          int    `json:"id"`
	ExampleName string `json:"example_name"`
}

func FormatExample(example entity.Example) ExampleFormatter {
	exampleFormatter := ExampleFormatter{}
	exampleFormatter.ID = example.ID
	exampleFormatter.ExampleName = example.ExampleName
	return exampleFormatter
}
func FormatExamples(examples []entity.Example) []ExampleFormatter {
	exampleFormatters := []ExampleFormatter{}
	for _, ex := range examples {
		exampleFormatter := FormatExample(ex)
		exampleFormatters = append(exampleFormatters, exampleFormatter)
	}
	return exampleFormatters
}
