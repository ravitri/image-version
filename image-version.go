package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/openshift/library-go/pkg/image/dockerv1client"
)

var (
	release_image = "quay.io/openshift-release-dev/ocp-release@sha256:dd75546170e65d7d17130de10a6ffeb425f960399640632cbc8426b9da338458"
)

func main() {

	configDigest := fetchConfigDigest(release_image)

	version := fetchImageVersion(configDigest)

	fmt.Println("Release Image Version: ", version)

}

func fetchConfigDigest(image string) string {

	sha := strings.SplitN(image, ":", 2)
	url := "https://quay.io/v2/openshift-release-dev/ocp-release/manifests/sha256:" + sha[1]

	client := http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Accept", "application/vnd.docker.distribution.manifest.v2+json")
	res, getErr := client.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	manifest := &dockerv1client.DockerImageManifest{}

	jsonErr := json.Unmarshal(body, &manifest)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return manifest.Config.Digest
}

func fetchImageVersion(digest string) string {
	blobURL := "https://quay.io/v2/openshift-release-dev/ocp-release/blobs/" + digest

	client := http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest(http.MethodGet, blobURL, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, getErr := client.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	imageConfig := &dockerv1client.DockerImageConfig{}

	jsonErr := json.Unmarshal(body, &imageConfig)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	version := imageConfig.Config.Labels["io.openshift.release"]

	return version
}
