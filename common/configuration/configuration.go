package configuration

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"reflect"
	"strings"

	"github.com/volcano822/journey/common/filenames"
)

// Configuration: settings that are neccesary for server configuration
type Configuration struct {
	HttpHostAndPort string
	Url             string
	UseLetsEncrypt  bool
}

func NewConfiguration() *Configuration {
	var config Configuration
	err := config.load()
	if err != nil {
		log.Println("Warning: couldn't load " + filenames.ConfigFilename + ", creating new config file.")
		err = config.create()
		if err != nil {
			log.Fatal("Fatal error: Couldn't create configuration.")
			return nil
		}
		err = config.load()
		if err != nil {
			log.Fatal("Fatal error: Couldn't load configuration.")
			return nil
		}
	}
	return &config
}

// Global config - thread safe and accessible from all packages
var Config = NewConfiguration()

func (c *Configuration) save() error {
	data, err := json.Marshal(c)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filenames.ConfigFilename, data, 0600)
}

func (c *Configuration) load() error {
	configWasChanged := false
	data, err := ioutil.ReadFile(filenames.ConfigFilename)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, c)
	if err != nil {
		return err
	}
	// Make sure the url is in the right format
	// Make sure there is no trailing slash at the end of the url
	if strings.HasSuffix(c.Url, "/") {
		c.Url = c.Url[0 : len(c.Url)-1]
		configWasChanged = true
	}
	if !strings.HasPrefix(c.Url, "http://") && !strings.HasPrefix(c.Url, "https://") {
		c.Url = "http://" + c.Url
		configWasChanged = true
	}
	// Check if all fields are filled out
	cReflected := reflect.ValueOf(*c)
	for i := 0; i < cReflected.NumField(); i++ {
		if cReflected.Field(i).Interface() == "" {
			log.Println("Error: " + filenames.ConfigFilename + " is corrupted. Did you fill out all of the fields?")
			return errors.New("Error: Configuration corrupted.")
		}
	}
	// Save the changed config
	if configWasChanged {
		err = c.save()
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Configuration) create() error {
	// TODO: Change default port
	c = &Configuration{HttpHostAndPort: ":8084", Url: "127.0.0.1:8084"}
	err := c.save()
	if err != nil {
		log.Println("Error: couldn't create " + filenames.ConfigFilename)
		return err
	}

	return nil
}
