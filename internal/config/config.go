package config

import (
	"github.com/richardcase/herald/internal/backend"
	"github.com/richardcase/herald/internal/colorscheme"
)

type Config struct {
	Colors  colorscheme.Colorscheme
	Backend backend.Backend
	Filters Filters
}

type Filters []Filter

func (f Filters) GetFilterByName(name string) (bool, *Filter) {
	for _, filter := range f {
		if filter.Name == name {
			return true, &filter
		}
	}

	return false, nil
}

type Filter struct {
	Name        string
	Description string
	RepoFilter  []string
}

func New(colors colorscheme.Colorscheme) (Config, error) {
	// Create a new config
	config := Config{}

	// Set the colorscheme
	config.Colors = colors

	// Get the backend
	backend, err := backend.New(backend.File)
	if err != nil {
		return config, err
	}

	// Set the backend
	config.Backend = backend

	//TODO: get this from file
	config.Filters = Filters{
		{
			Name:        "kubernetes",
			Description: "Kubernetes project notifications",
			RepoFilter:  []string{"kubernetes*"},
		},
		{
			Name:        "rancher",
			Description: "Rancher notifications",
			RepoFilter:  []string{"rancher*"},
		},
	}

	// Return the config
	return config, nil
}

func (c Config) Close() error {
	return c.Backend.Close()
}
