package data

import (
	"language-structure/dictionary/dto"
	"language-structure/dictionary/engine"
	"os"
	"path"
	"reflect"
	"runtime"
	"sort"
	"testing"
)

const (
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
	_, ok := data.Get(javaLanguage)
	if !ok {
		t.Errorf("The data should be there")
	}

}

func TestShouldMergeFiles(t *testing.T) {
	yamlFile := dir() + "/languages/java.yaml"
	yamlFile2 := dir() + "/languages/java_2.yaml"
	data := Create()
	_, ok := data.Get(javaLanguage)
	if ok {
		t.Errorf("The data should be empty")
	}
	data.Add(createReadRuleManager(yamlFile))
	data.Add(createReadRuleManager(yamlFile2))

	java, ok := data.Get(javaLanguage)
	if !ok {
		t.Errorf("The 'Java' element should be in the structure")
	}
	extensions := []string{".java", ".jsp"}
	if !reflect.DeepEqual(java.Extensions, extensions) {
		t.Errorf("The extensions structure should have '.java' and '.jsp'")
	}
}

func TestShouldNotHaveDuplicateExtensions(t *testing.T) {
	data := Create()
	data.Add(createReadRuleManager(dir() + "/languages/java.yaml"))
	data.Add(createReadRuleManager(dir() + "/languages/java_2.yaml"))
	data.Add(createReadRuleManager(dir() + "/languages/java_3.yaml"))
	java, ok := data.Get(javaLanguage)
	if !ok {
		t.Errorf("The 'Java' element should be in the structure")
	}
	extensions := []string{".jsp", ".java"}
	sort.Strings(java.Extensions)
	sort.Strings(extensions)
	if !reflect.DeepEqual(java.Extensions, extensions) {
		t.Errorf("The extensions structure should have '.java' and '.jsp' %s", java.Extensions)
	}
}

func TestNotHaveDuplicatedRules(t *testing.T)  {
	data := Create()
	data.Add(createReadRuleManager(dir() + "/languages/java.yaml"))
	data.Add(createReadRuleManager(dir() + "/languages/java_2.yaml"))
	data.Add(createReadRuleManager(dir() + "/languages/java_3.yaml"))
	data.Add(createReadRuleManager(dir() + "/languages/java_4.yaml"))
	java, ok := data.Get(javaLanguage)
	if !ok {
		t.Errorf("The 'Java' element should be in the structure")
	}
	if len(java.Rules) != 3 {
		t.Errorf("It should have three rules inside the Java rule, current size %d", len(java.Rules))
	}
}

func TestFindRuleById(t *testing.T) {
	id:= "d34c6b79-4051-4f73-bf8e-37db9becc111"
	data := Create()
	data.Add(createReadRuleManager(dir() + "/languages/java.yaml"))
	data.Add(createReadRuleManager(dir() + "/languages/java_2.yaml"))
	data.Add(createReadRuleManager(dir() + "/languages/java_3.yaml"))
	data.Add(createReadRuleManager(dir() + "/languages/java_4.yaml"))
	_, ok := data.FindRuleById(id)
	if !ok {
		t.Errorf("The {id} element should be in the structure")
	}
}

func TestShouldNotFindRuleById(t *testing.T) {
	id:= "id-not-exist"
	data := Create()
	data.Add(createReadRuleManager(dir() + "/languages/java.yaml"))
	data.Add(createReadRuleManager(dir() + "/languages/java_2.yaml"))
	data.Add(createReadRuleManager(dir() + "/languages/java_3.yaml"))
	data.Add(createReadRuleManager(dir() + "/languages/java_4.yaml"))
	_, ok := data.FindRuleById(id)
	if ok {
		t.Errorf("The id does not exist in the structure")
	}
}

func createReadRuleManager(file string) *engine.RuleManager {
	conf, _ := dto.ReadConf(file)
	return conf.ToRuleManager()
}

