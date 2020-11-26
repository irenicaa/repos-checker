package config

import (
	"encoding/json"
	"errors"
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
	Name        string
	IsReference bool
	Options     json.RawMessage
	SubSources  []SourceConfig
}

// LoadConfig ...
func LoadConfig(reader io.Reader, logger loader.Logger) (
	sources []loader.Source,
	referenceName string,
	err error,
) {
	configBytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, "", fmt.Errorf("unable to read a config: %v", err)
	}

	var config Config
	if err := json.Unmarshal(configBytes, &config); err != nil {
		return nil, "", fmt.Errorf("unable to unmarshal the config: %v", err)
	}

	for _, sourceConfig := range config {
		if sourceConfig.IsReference {
			if referenceName != "" {
				return nil, "", errors.New("more than one source is marked as a reference")
			}

			referenceName = sourceConfig.Name
		}

		source, err := LoadSource(sourceConfig, logger)
		if err != nil {
			return nil, "", fmt.Errorf(
				"unable to load the %s source: %v",
				sourceConfig.Name,
				err,
			)
		}

		sources = append(sources, source)
	}

	return sources, referenceName, nil
}

// LoadSource ...
func LoadSource(sourceConfig SourceConfig, logger loader.Logger) (
	loader.Source,
	error,
) {
	var source loader.Source
	switch sourceConfig.Name {
	case "github":
		source = &github.Source{Logger: logger}
	case "bitbucket":
		source = &bitbucket.Source{Logger: logger}
	case "gitlab":
		source = &gitlab.Source{Logger: logger}
	case "file-system":
		source = &filesystem.Source{Logger: logger}
	case "external":
		source = &external.Source{}
	case "multi-source":
		var subSources []loader.Source
		for _, subSourceConfig := range sourceConfig.SubSources {
			subSource, err := LoadSource(subSourceConfig, logger)
			if err != nil {
				return nil, fmt.Errorf(
					"unable to load the %s sub-source: %v",
					subSourceConfig.Name,
					err,
				)
			}

			subSources = append(subSources, subSource)
		}

		source = loader.MultiSource(subSources)
	default:
		return nil, fmt.Errorf("unknown source %s", sourceConfig.Name)
	}

	if sourceConfig.Options != nil {
		if err := json.Unmarshal(sourceConfig.Options, source); err != nil {
			return nil, fmt.Errorf("unable to unmarshal source options: %v", err)
		}
	}

	return source, nil
}
