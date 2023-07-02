package docs

import (
	"log"
	env "marketplace/security-api/src/shared"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

func LoadDoc() string {
	dir, _ := os.Getwd()
	docsPath := filepath.Join(dir,"/src/docs/swagger.yaml")
	swaggerFile, err := os.ReadFile(docsPath)
	if err != nil {
		log.Panicf("Error on load swagger doc: %v",err)
	}
	return string(setProperties(&swaggerFile))
}

func setProperties(swaggerFile *[]byte) []byte{
	var config map[string]interface{}
	apiURL := env.GetConfig().UrlApi
	err := yaml.Unmarshal(*swaggerFile, &config)
	if err != nil {
		log.Fatalf("Error on parsing yaml file: %v", err)
	}
	servers := config["servers"].([]interface{})
	if len(servers) > 0 {
		server := servers[0].(map[interface {}]interface{})
		server["url"] = apiURL
	}
	updatedYAML, err := yaml.Marshal(&config)
	if err != nil {
		log.Fatalf("Error on update yaml: %v", err)
	}
	return updatedYAML
}