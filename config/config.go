package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

var (
	Config ConfigFile
)

type ConfigFile struct {
	URL     string `toml:"url"`
	Private bool   `toml:"private"`

	ExtensionFilter     []string `toml:"extension_filter"`
	ExtensionFilterMode string   `toml:"extension_filter_mode"`

	Uploads struct {
		Folder   string `toml:"folder"`
		MaxSize  string `toml:"max_size"`
		Hash     bool   `toml:"hash"`
		IDLength int    `toml:"id_length"`
	} `toml:"uploads"`

	Database string `toml:"database"`

	ServeFiles  bool   `toml:"serve_files"`
	FileBaseURL string `toml:"file_base_url"`
}

func (c *ConfigFile) Defaults() {
	c.URL = "http://localhost:8080/"
	c.FileBaseURL = c.URL
	c.Private = true

	// stolen from lolisafe's default
	c.ExtensionFilter = []string{
		".bash_profile",
		".bash",
		".bashrc",
		".bat",
		".bsh",
		".cmd",
		".com",
		".csh",
		".exe",
		".exec",
		".jar",
		".msi",
		".nt",
		".profile",
		".ps1",
		".psm1",
		".scr",
		".sh",
	}

	c.ExtensionFilterMode = "deny"

	c.Uploads.Folder = "uploads"
	c.Uploads.MaxSize = "128M"
	c.Uploads.IDLength = 10
	c.Uploads.Hash = true

	c.Database = "imagebucket.db"
	c.ServeFiles = true
}

func LoadConfig(file string) error {

	// set defaults before loading the actual config file.
	// any settings explicitly defined will be used instead
	Config.Defaults()

	log.Printf("loading config from %s ...\n", file)

	_, err := toml.DecodeFile(file, &Config)
	if err != nil {
		log.Fatalf("could not read config file %s: %v\n", file, err)
	}

	return nil
}
