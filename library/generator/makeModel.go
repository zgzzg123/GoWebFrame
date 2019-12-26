package generator

import (
	"strings"
	"os"
	"fmt"
	"io/ioutil"
	"goDoc/library/helper"
)

/**
预处理名称为驼峰名称
 */
func preNameToCamelString(foldName string) string {
	foldNameString := ""
	foldNameMap := strings.Split(foldName, "_")
	for k, v := range foldNameMap {
		if k > 0 {
			foldNameString += helper.Ucfirst(v)
		} else {
			foldNameString += v
		}
	}

	return foldNameString
}
func MakeModel(foldName string) {
	modelPath := "./models/"

	foldNameString := preNameToCamelString(foldName)

	fullPath := modelPath + foldNameString + "Model"
	// check
	if _, err := os.Stat(fullPath); err == nil {
		fmt.Println("path exists 1", fullPath)
	} else {
		fmt.Println("path not exists ", fullPath)
		err := os.Mkdir(fullPath, os.ModePerm)
		if err != nil {
			fmt.Printf("mkdir failed![%v]\n", err)
		} else {
			fmt.Printf("mkdir success!\n")
		}
	}

	contentModelTpl, err := ioutil.ReadFile("library/generator/templates/model.tpl")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(contentModelTpl)
	modelTemplateContent := string(contentModelTpl)

	modelTemplateContent = strings.Replace(modelTemplateContent,"{{tbName}}",TbName,1)

	dataProfileString := makeDataProfiles(true)
	modelTemplateContent = strings.Replace(modelTemplateContent,"{{dataProfilesContent}}",dataProfileString,1)

	searchParamsContent := makeDataProfiles(false)
	modelTemplateContent = strings.Replace(modelTemplateContent,"{{searchParamsContent}}",searchParamsContent,1)

	modelTemplateContent = strings.Replace(modelTemplateContent,"{{scopeFilterContent}}","",1)

	fileName := fullPath + "/main.go"
	helper.WriteFile(fileName, modelTemplateContent)
}