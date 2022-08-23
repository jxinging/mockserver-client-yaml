package mockyaml

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

func unmarshalApis(filename string) (rs []MockExpectation, err error) {
	bs, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(bs, &rs); err != nil {
		return nil, err
	}
	return rs, nil
}

func WalkYamlFiles(dir string, call func([]MockExpectation) error) error {
	entrys, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, e := range entrys {
		if !strings.HasSuffix(e.Name(), ".yaml") && !strings.HasSuffix(e.Name(), ".yml") {
			continue
		}
		if e.IsDir() {
			continue
		}

		file, err := e.Info()
		if err != nil {
			return err
		}

		apis, err := unmarshalApis(filepath.Join(dir, file.Name()))
		if err != nil {
			return err
		}
		// log.Printf("%v", apis)
		if err := call(apis); err != nil {
			return err
		}
	}
	return nil
}

func CreateMockExpectation(mockserver string, mockReq *MockExpectation) error {
	u, err := url.Parse(mockserver)
	if err != nil {
		return err
	}
	u.Path = "/mockserver/expectation"

	reqContent, err := json.MarshalIndent(mockReq, "", "  ")
	if err != nil {
		return errors.Wrap(err, "marshal mockserver request")
	}

	log.Printf("create mockserver expectation: %s", reqContent)
	httpReq, err := http.NewRequest("PUT", u.String(), bytes.NewBuffer(reqContent))
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return err
	}
	if resp.StatusCode >= 300 {
		return fmt.Errorf(resp.Status)
	}
	return nil
}
