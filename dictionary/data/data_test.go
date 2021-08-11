package data

import (
	"language-structure/dictionary/dto"
	"os"
	"path"
	"runtime"
	"testing"
)

var (
	javaLanguage = "Java"
)

func dir() string {
	_, filename, _, _ := runtime.Caller(0)
	// The ".." may change depending on you folder structure
	dir := path.Join(path.Dir(filename), "../..")
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}
	return dir
}

func TestShouldCreateEmptyData(t *testing.T) {
	data := Languages{}
	if data.Size() != 0 {
		t.Errorf("The data should start empty")
	}
	if data.Contains(javaLanguage) {
		t.Errorf("The data should be empty")
	}

	if data.Size() != 0 {
		t.Errorf("The Should have size 0")
	}
}

func TestShouldCreateRuleManager(t *testing.T) {
	yamlFile := dir() + "/languages/java.yaml"

	if _, err := os.Stat(yamlFile); os.IsNotExist(err) {
		print("file does not exist")
	}
	data := Create()
	language, _ := dto.ReadConf(yamlFile)
	data.Add(language.ToRuleManager())

	if !data.Contains(javaLanguage) {
		t.Errorf("Should have Java language")
	}

	if data.Size() != 1 {
		t.Errorf("The value should have one element")
	}

}
