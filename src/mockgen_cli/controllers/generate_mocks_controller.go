package controllers

import "github.com/FournyP/deepsearch-mockgen-cli/src/mockgen/use_cases"

type GenerateMocksController struct {
	generateMocks *use_cases.GenerateMocks
}

func NewGenerateMocksController(generateMocks *use_cases.GenerateMocks) *GenerateMocksController {
	return &GenerateMocksController{generateMocks: generateMocks}
}

func (c *GenerateMocksController) Execute() error {
	return c.generateMocks.Execute()
}
