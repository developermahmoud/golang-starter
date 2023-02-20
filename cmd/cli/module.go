package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/gertd/go-pluralize"
	"github.com/iancoleman/strcase"
)

func createModule(module string) {
	createRepository(module)
	createControllerModule(module)
	createModel(module)
}

func createControllerModule(controller string) {
	var controllerName string
	var controllerData string
	data, err := os.ReadFile("template/controller_module.txt")
	if err != nil {
		color.Red("controller template not found")
		os.Exit(0)
	}
	controllerTemplate := string(data)
	pluralize := pluralize.NewClient()
	if pluralize.IsPlural(controller) {
		controllerName = controller
	} else {
		controllerName = pluralize.Plural(controller)
	}
	f, err := os.Create(fmt.Sprintf("src/controllers/%s_controller.go", controllerName))
	if err != nil {
		color.Red("can not create module")
		os.Exit(0)
	}
	defer f.Close()
	controllerData = strings.ReplaceAll(controllerTemplate, "$CONTROLLER_CAMEL$", strcase.ToCamel(controllerName))
	controllerData = strings.ReplaceAll(controllerData, "$CONTROLLER_LOWER$", strings.ToLower(controllerName))
	f.Write([]byte(controllerData))
}

func createModel(model string) {
	var modelName string
	var modelData string
	data, err := os.ReadFile("template/model.txt")
	if err != nil {
		color.Red("model template not found")
		os.Exit(0)
	}
	modelTemplate := string(data)

	pluralize := pluralize.NewClient()
	if pluralize.IsPlural(model) {
		modelName = pluralize.Singular(model)
	} else {
		modelName = model
	}
	f, err := os.Create(fmt.Sprintf("src/models/%s.go", modelName))
	if err != nil {
		color.Red("can not create model")
		os.Exit(0)
	}
	defer f.Close()
	modelData = strings.ReplaceAll(modelTemplate, "$MODEL$", strcase.ToCamel(modelName))
	f.Write([]byte(modelData))
}

func createRepository(repository string) {
	var repositoryName string
	var repositoryData string
	data, err := os.ReadFile("template/repository.txt")
	if err != nil {
		color.Red("repository template not found")
		os.Exit(0)
	}
	repositoryTemplate := string(data)
	pluralize := pluralize.NewClient()
	if pluralize.IsPlural(repository) {
		repositoryName = repository
	} else {
		repositoryName = pluralize.Plural(repository)
	}
	f, err := os.Create(fmt.Sprintf("src/repositories/%s_repository.go", repositoryName))
	if err != nil {
		color.Red("can not create module")
		os.Exit(0)
	}
	defer f.Close()
	repositoryData = strings.ReplaceAll(repositoryTemplate, "$REPOSITORY_CAMEL$", strcase.ToCamel(repositoryName))
	repositoryData = strings.ReplaceAll(repositoryData, "$REPOSITORY_LOWER$", strings.ToLower(repositoryName))
	f.Write([]byte(repositoryData))
}
