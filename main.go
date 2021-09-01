package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please, enter the path to the desired file: ")

	filePath, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error ocurred: ", err)
		return
	}

	filePath = strings.Replace(filePath, "\n", "", -1)

	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("An error ocurred: ", err)
		return
	}

	var scopes Scopes
	err = json.Unmarshal(file, &scopes)
	if err != nil {
		fmt.Println("An error ocurred: ", err)
		return
	}

	for scopeName, properties := range scopes {
		output, err := exec.Command("bash", "-c", fmt.Sprintf("databricks secrets create-scope --scope %s", scopeName)).Output()
		if err != nil {
			if strings.Contains(string(output), "RESOURCE_ALREADY_EXISTS") {
				fmt.Println("Scope already exists. Skipping creation")
			} else {
				fmt.Println("An error ocurred: ", err)
				return
			}
		}

		for key, secret := range properties.Secrets {
			output, err = exec.Command("bash", "-c", fmt.Sprintf("databricks secrets put --scope %s --key %s --string-value %s", scopeName, key, secret)).Output()
			if err != nil {
				fmt.Println("An error ocurred: ", string(output))
				return
			}
		}

		for principal, permission := range properties.ACLs {
			output, err = exec.Command("bash", "-c", fmt.Sprintf("databricks secrets put-acl --scope %s --principal %s --permission %s", scopeName, principal, permission)).Output()
			if err != nil {
				if err != nil {
					fmt.Println("An error ocurred: ", string(output))
					return
				}
			}
		}
	}
}
