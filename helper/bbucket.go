package helper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// Credential struct for codefreeze app config
type Credential struct {
	name string
	pass string
}

// getBBCredentials reading creadentials from env variables
// and returns structure with user name and user password
func getBBCredentials() *Credential {
	creds := &Credential{}
	creds.name = os.Getenv("BB_USER")
	creds.pass = os.Getenv("BB_PASS")
	return creds
}

// GetRepos returns list of repostories
func GetRepos(bbhost string, repoName string, projectName string) map[string]interface{} {
	// get user name and pass for BitBucket repo
	user := getBBCredentials()

	url := fmt.Sprintf("%s/rest/api/1.0/projects/%s/repos/%s/pull-requests", bbhost, projectName, repoName)
	jsonData := map[string]string{"order": "OLDEST"}
	jsonValue, _ := json.Marshal(jsonData)
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(user.name, user.pass)
	fmt.Println("++", req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	// fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)

	/* declared map of string with empty interface which will hold the value of the parsed json. */
	var result map[string]interface{}

	/* Unmarshal the json string string by converting it to byte into map */
	json.Unmarshal([]byte(body), &result)

	return result
}
