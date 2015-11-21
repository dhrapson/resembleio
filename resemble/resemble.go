//Copyright (C) 2015 dhrapson

//This program is free software: you can redistribute it and/or modify
//it under the terms of the GNU General Public License as published by
//the Free Software Foundation, either version 3 of the License, or
//(at your option) any later version.

//This program is distributed in the hope that it will be useful,
//but WITHOUT ANY WARRANTY; without even the implied warranty of
//MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//GNU General Public License for more details.

//You should have received a copy of the GNU General Public License
//along with this program.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"errors"
	"fmt"
	"github.com/dhrapson/gopretend/configure"
	"io/ioutil"
	"os"
	"path/filepath"
)

var defaultYamlFileName = "gopretend.yml"

func main() {
	argsWithoutProg := os.Args[1:]
	configYaml, err := GetConfigData(argsWithoutProg)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting GoPretend...")
	serviceType := configure.ConfigureService(configYaml)
	fmt.Println("Configuring GoPretend as", serviceType.Name(), "...")
	serviceType.Configure()
	fmt.Println("Starting GoPretend Service...")
	serviceType.Serve()
	fmt.Println("Stopping GoPretend Service...")
}

func GetConfigData(cmdLineArgs []string) (configYaml string, err error) {
	var yamlFileName string

	if len(cmdLineArgs) > 0 && len(cmdLineArgs[0]) > 0 {
		yamlFileName = cmdLineArgs[0]
		if !fileExists(yamlFileName) {
			return "", errors.New(yamlFileName + " cannot be found")
		}
		fmt.Println("Using provided config file " + yamlFileName)
		configYaml, err = getFileContent(yamlFileName)
	} else if fileExists(defaultYamlFileName) {
		fmt.Println("Using default config file " + defaultYamlFileName)
		configYaml, err = getFileContent(defaultYamlFileName)
	} else {
		fmt.Println("No config file available, initializing empty. You may configure via API")
	}
	return configYaml, err
}

func getFileContent(name string) (content string, err error) {
	filename, _ := filepath.Abs(name)
	configYaml, err := ioutil.ReadFile(filename)
	return string(configYaml), err
}

func fileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
