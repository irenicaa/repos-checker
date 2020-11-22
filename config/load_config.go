package config

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/irenicaa/repos-checker/loader"
	"github.com/irenicaa/repos-checker/loader/sources/bitbucket"
	"github.com/irenicaa/repos-checker/loader/sources/external"
	filesystem "github.com/irenicaa/repos-checker/loader/sources/file-system"
	"github.com/irenicaa/repos-checker/loader/sources/github"
	"github.com/irenicaa/repos-checker/loader/sources/gitlab"
)

// Config ...
type Config []SourceConfig

// SourceConfig ...
type SourceConfig struct {
	Name string
}

// LoadConfig ...
func LoadConfig(reader io.Reader) ([]loader.Source, error) {
	configBytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("unable to read a config: %v", err)
	}

	var config Config
	if err := json.Unmarshal(configBytes, &config); err != nil {
		return nil, fmt.Errorf("unable to unmarshal the config: %v", err)
	}

	var sources []loader.Source
	for _, sourceConfig := range config {
		var source loader.Source
		switch sourceConfig.Name {
		case "github":
			source = github.Source{}
		case "bitbucket":
			source = bitbucket.Source{}
		case "gitlab":
			source = gitlab.Source{}
		case "file-system":
			source = filesystem.Source{}
		case "external":
			source = external.Source{}
		default:
			return nil, fmt.Errorf("unknown source %s", sourceConfig.Name)
		}

		sources = append(sources, source)
	}

	return sources, nil
}
