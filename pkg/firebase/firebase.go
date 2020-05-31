package firebase

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path"
)

const firebaseAddressEnvVar = "FIREBASE_ADDRESS"

var endpoint *url.URL

func init() {
	endpoint, _ = url.Parse(os.Getenv(firebaseAddressEnvVar))
}

type Statuses []string

func (s Statuses) Exists(status string) bool {
	for _, existingStatus := range s {
		if existingStatus == status {
			return true
		}
	}

	return false
}

func FetchStatuses() (Statuses, error) {
	u, _ := url.Parse(endpoint.String())
	u.Path = path.Join(u.Path, "statuses.json")

	req, err := http.NewRequest(http.MethodGet, u.String(), http.NoBody)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var allStatuses Statuses
	if err := json.NewDecoder(resp.Body).Decode(&allStatuses); err != nil {
		return nil, err
	}

	return allStatuses, nil
}

type Person struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

type People map[string]Person

func (p People) Exists(person string) bool {
	_, exists := p[person]
	return exists
}

func FetchPeople() (People, error) {
	u, _ := url.Parse(endpoint.String())
	u.Path = path.Join(u.Path, "people.json")

	req, err := http.NewRequest(http.MethodGet, u.String(), http.NoBody)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var allPeople People
	if err := json.NewDecoder(resp.Body).Decode(&allPeople); err != nil {
		return nil, err
	}

	return allPeople, nil
}

func UpdateLocation(name, status string) error {
	u, _ := url.Parse(endpoint.String())
	u.Path = path.Join(u.Path, "people", fmt.Sprintf("%s.json", name))

	payload := map[string]string{"status": status}
	js, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPatch, u.String(), bytes.NewBuffer(js))
	if err != nil {
		return err
	}
	req.Header.Add("content-type", "application/json")

	client := &http.Client{}
	_, err = client.Do(req)

	return err
}
