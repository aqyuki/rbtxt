package info

import (
	_ "embed"
	"sync"

	"gopkg.in/yaml.v3"
)

// ApplicationInformation is a structure to hold application static information.
type ApplicationInformation struct {
	Version string `yaml:"version"`
}

//go:embed info.yml
var raw []byte

var (
	applicationInfo      *ApplicationInformation = nil
	informationLoadError error                   = nil
	decodeOnce           sync.Once
)

// decode decodes the embed yaml file into ApplicationInformation.
func decode(b []byte) (info *ApplicationInformation, err error) {
	info = &ApplicationInformation{}
	if err = yaml.Unmarshal(b, info); err != nil {
		return nil, err
	}
	return info, nil
}

// GetApplicationInformation returns application information.
func GetApplicationInformation() (*ApplicationInformation, error) {
	decodeOnce.Do(func() {
		info, err := decode(raw)
		if err != nil {
			informationLoadError = err
			return
		}
		applicationInfo = info
	})
	return applicationInfo, informationLoadError
}
