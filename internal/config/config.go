package config

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/go-playground/validator/v10"
	"gopkg.in/yaml.v2"

	"github.com/richardcase/herald/internal/utils"
)

/*
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

*/

const HeraldDir = "herald"
const ConfigYmlFileName = "config.yml"
const ConfigYamlFileName = "config.yaml"
const DEFAULT_XDG_CONFIG_DIRNAME = ".config"

var validate *validator.Validate

type ViewType string

const (
	NotificationsView ViewType = "notifications"
)

type SectionConfig struct {
	Title   string
	Filters string
	Limit   *int `yaml:"limit,omitempty"`
}

type NotificationSectionConfig struct {
	Title   string
	Filters string
	Limit   *int                     `yaml:"limit,omitempty"`
	Layout  NotificationLayoutConfig `yaml:"layout,omitempty"`
}

func (cfg NotificationSectionConfig) ToSectionConfig() SectionConfig {
	return SectionConfig{
		Title:   cfg.Title,
		Filters: cfg.Filters,
		Limit:   cfg.Limit,
	}
}

type ColumnConfig struct {
	Width  *int  `yaml:"width,omitempty"  validate:"omitempty,gt=0"`
	Hidden *bool `yaml:"hidden,omitempty"`
}

type NotificationLayoutConfig struct {
	UpdatedAt    ColumnConfig `yaml:"updatedAt,omitempty"`
	Repo         ColumnConfig `yaml:"repo,omitempty"`
	Author       ColumnConfig `yaml:"author,omitempty"`
	Assignees    ColumnConfig `yaml:"assignees,omitempty"`
	Title        ColumnConfig `yaml:"title,omitempty"`
	Base         ColumnConfig `yaml:"base,omitempty"`
	ReviewStatus ColumnConfig `yaml:"reviewStatus,omitempty"`
	State        ColumnConfig `yaml:"state,omitempty"`
	Ci           ColumnConfig `yaml:"ci,omitempty"`
	Lines        ColumnConfig `yaml:"lines,omitempty"`
}

type LayoutConfig struct {
	Notifications NotificationLayoutConfig `yaml:"notifications,omitempty"`
}

type PreviewConfig struct {
	Open  bool
	Width int
}

type Defaults struct {
	Preview                PreviewConfig `yaml:"preview"`
	NotificationsLimit     int           `yaml:"notificationsLimit"`
	View                   ViewType      `yaml:"view"`
	Layout                 LayoutConfig  `yaml:"layout,omitempty"`
	RefetchIntervalMinutes int           `yaml:"refetchIntervalMinutes,omitempty"`
}

type Keybinding struct {
	Key     string `yaml:"key"`
	Command string `yaml:"command"`
}

type Keybindings struct {
	Notifications []Keybinding `yaml:"notifications"`
}

type Pager struct {
	Diff string `yaml:"diff"`
}

type HexColor string

type ColorThemeText struct {
	Primary   HexColor `yaml:"primary"   validate:"hexcolor"`
	Secondary HexColor `yaml:"secondary" validate:"hexcolor"`
	Inverted  HexColor `yaml:"inverted"  validate:"hexcolor"`
	Faint     HexColor `yaml:"faint"     validate:"hexcolor"`
	Warning   HexColor `yaml:"warning"   validate:"hexcolor"`
	Success   HexColor `yaml:"success"   validate:"hexcolor"`
}

type ColorThemeBorder struct {
	Primary   HexColor `yaml:"primary"   validate:"hexcolor"`
	Secondary HexColor `yaml:"secondary" validate:"hexcolor"`
	Faint     HexColor `yaml:"faint"     validate:"hexcolor"`
}

type ColorThemeBackground struct {
	Selected HexColor `yaml:"selected" validate:"hexcolor"`
}

type ColorTheme struct {
	Text       ColorThemeText       `yaml:"text"       validate:"required,dive"`
	Background ColorThemeBackground `yaml:"background" validate:"required,dive"`
	Border     ColorThemeBorder     `yaml:"border"     validate:"required,dive"`
}

type ColorThemeConfig struct {
	Inline ColorTheme `yaml:",inline" validate:"dive"`
}

type TableUIThemeConfig struct {
	ShowSeparator bool `yaml:"showSeparator" default:"true"`
}

type UIThemeConfig struct {
	Table TableUIThemeConfig `yaml:"table"`
}

type ThemeConfig struct {
	Ui     UIThemeConfig     `yaml:"ui,omitempty"     validate:"dive"`
	Colors *ColorThemeConfig `yaml:"colors,omitempty" validate:"omitempty,dive"`
}

type Config struct {
	NotificationsSections []NotificationSectionConfig `yaml:"notificationsSections"`
	Defaults              Defaults                    `yaml:"defaults"`
	Keybindings           Keybindings                 `yaml:"keybindings"`
	RepoPaths             map[string]string           `yaml:"repoPaths"`
	Theme                 *ThemeConfig                `yaml:"theme,omitempty" validate:"omitempty,dive"`
	Pager                 Pager                       `yaml:"pager"`
}

type configError struct {
	configDir string
	parser    ConfigParser
	err       error
}

type ConfigParser struct{}

func (parser ConfigParser) getDefaultConfig() Config {
	return Config{
		Defaults: Defaults{
			Preview: PreviewConfig{
				Open:  true,
				Width: 50,
			},
			NotificationsLimit:     20,
			View:                   NotificationsView,
			RefetchIntervalMinutes: 30,
			Layout: LayoutConfig{
				Notifications: NotificationLayoutConfig{
					UpdatedAt: ColumnConfig{
						Width: utils.IntPtr(lipgloss.Width("2mo ago")),
					},
					Repo: ColumnConfig{
						Width: utils.IntPtr(15),
					},
					Author: ColumnConfig{
						Width: utils.IntPtr(15),
					},
					Assignees: ColumnConfig{
						Width:  utils.IntPtr(20),
						Hidden: utils.BoolPtr(true),
					},
					Base: ColumnConfig{
						Width:  utils.IntPtr(15),
						Hidden: utils.BoolPtr(true),
					},
					Lines: ColumnConfig{
						Width: utils.IntPtr(lipgloss.Width("123450 / -123450")),
					},
				},
			},
		},
		NotificationsSections: []NotificationSectionConfig{
			{
				Title:   "My Pull Requests",
				Filters: "is:open author:@me",
			},
			{
				Title:   "Needs My Review",
				Filters: "is:open review-requested:@me",
			},
			{
				Title:   "Involved",
				Filters: "is:open involves:@me -author:@me",
			},
		},
		Keybindings: Keybindings{
			Notifications: []Keybinding{},
		},
		RepoPaths: map[string]string{},
		Theme: &ThemeConfig{
			Ui: UIThemeConfig{
				Table: TableUIThemeConfig{
					ShowSeparator: true,
				},
			},
		},
	}
}

func (parser ConfigParser) getDefaultConfigYamlContents() string {
	defaultConfig := parser.getDefaultConfig()
	yaml, _ := yaml.Marshal(defaultConfig)

	return string(yaml)
}

func (e configError) Error() string {
	return fmt.Sprintf(
		`Couldn't find a config.yml or a config.yaml configuration file.
Create one under: %s

Example of a config.yml file:
%s

For more info, go to https://github.com/dlvhdr/gh-dash
press q to exit.

Original error: %v`,
		path.Join(e.configDir, HeraldDir, ConfigYmlFileName),
		string(e.parser.getDefaultConfigYamlContents()),
		e.err,
	)
}

func (parser ConfigParser) writeDefaultConfigContents(
	newConfigFile *os.File,
) error {
	_, err := newConfigFile.WriteString(parser.getDefaultConfigYamlContents())

	if err != nil {
		return err
	}

	return nil
}

func (parser ConfigParser) createConfigFileIfMissing(
	configFilePath string,
) error {
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		newConfigFile, err := os.OpenFile(
			configFilePath,
			os.O_RDWR|os.O_CREATE|os.O_EXCL,
			0666,
		)
		if err != nil {
			return err
		}

		defer newConfigFile.Close()
		return parser.writeDefaultConfigContents(newConfigFile)
	}

	return nil
}

func (parser ConfigParser) getExistingConfigFile() (*string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	xdgConfigDir := os.Getenv("XDG_CONFIG_HOME")
	if xdgConfigDir == "" {
		xdgConfigDir = filepath.Join(homeDir, DEFAULT_XDG_CONFIG_DIRNAME)
	}

	configPaths := []string{
		os.Getenv(
			"HERALD_CONFIG",
		), // If HERALD_CONFIG is empty, the os.Stat call fails
		filepath.Join(xdgConfigDir, HeraldDir, ConfigYmlFileName),
		filepath.Join(xdgConfigDir, HeraldDir, ConfigYamlFileName),
	}

	// Check if each config file exists, return the first one that does
	for _, configPath := range configPaths {
		if configPath == "" {
			continue // Skip checking if path is empty
		}
		if _, err := os.Stat(configPath); err == nil {
			return &configPath, nil
		}
	}

	return nil, nil
}

func (parser ConfigParser) getDefaultConfigFileOrCreateIfMissing() (string, error) {
	var configFilePath string

	heraldConfig := os.Getenv("HERALD_CONFIG")
	if heraldConfig != "" {
		configFilePath = heraldConfig
	} else {
		configDir := os.Getenv("XDG_CONFIG_HOME")
		if configDir == "" {
			homeDir, err := os.UserHomeDir()
			if err != nil {
				return "", err
			}
			configDir = filepath.Join(homeDir, DEFAULT_XDG_CONFIG_DIRNAME)
		}

		dashConfigDir := filepath.Join(configDir, HeraldDir)
		configFilePath = filepath.Join(dashConfigDir, ConfigYmlFileName)
	}

	// Ensure directory exists before attempting to create file
	configDir := filepath.Dir(configFilePath)
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		if err = os.MkdirAll(configDir, os.ModePerm); err != nil {
			return "", configError{
				parser:    parser,
				configDir: configDir,
				err:       err,
			}
		}
	}

	if err := parser.createConfigFileIfMissing(configFilePath); err != nil {
		return "", configError{parser: parser, configDir: configDir, err: err}
	}

	return configFilePath, nil
}

type parsingError struct {
	err error
}

func (e parsingError) Error() string {
	return fmt.Sprintf("failed parsing config.yml: %v", e.err)
}

func (parser ConfigParser) readConfigFile(path string) (Config, error) {
	config := parser.getDefaultConfig()
	data, err := os.ReadFile(path)
	if err != nil {
		return config, configError{parser: parser, configDir: path, err: err}
	}

	err = yaml.Unmarshal([]byte(data), &config)
	if err != nil {
		return config, err
	}

	err = validate.Struct(config)
	return config, err
}

func initParser() ConfigParser {
	validate = validator.New()

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.Split(fld.Tag.Get("yaml"), ",")[0]
		if name == "-" {
			return ""
		}
		return name
	})

	return ConfigParser{}
}

func ParseConfig(path string) (Config, error) {
	parser := initParser()

	var config Config
	var err error
	var configFilePath string

	if path == "" {
		configFilePath, err = parser.getDefaultConfigFileOrCreateIfMissing()
		if err != nil {
			return config, parsingError{err: err}
		}
	} else {
		configFilePath = path
	}

	config, err = parser.readConfigFile(configFilePath)
	if err != nil {
		return config, parsingError{err: err}
	}

	return config, nil
}
