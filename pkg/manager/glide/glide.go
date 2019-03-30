package glide

import (
	"path/filepath"
	"github.com/BurntSushi/toml"
	"strings"
	"github.com/sanjid133/gopher-love/util"
	. "github.com/sanjid133/gopher-love/pkg"
	"context"
)

const (
	Manager  = "glide"
	FileName = "glide.yaml"
)

type Glide struct {
	ctx       context.Context
	directory string
}

var _ LoveBag = &Glide{}

func init() {
	RegistarManager(Manager, func(ctx context.Context) LoveBag { return New(ctx) })
}

func New(ctx context.Context) LoveBag {
	return &Glide{ctx: ctx}
}

type GlideFile struct {
	Import []struct {
		Package string `yaml:"package"`
	} `yaml:"import"`
}

func (d *Glide) Initialize(directory string) LoveBag {
	d.directory = directory
	return d
}

func (d *Glide) File() string {
	return FileName
}

func (d *Glide) Read() ([]*Repository, error) {
	file := filepath.Join(d.directory, FileName)
	var config GopkgConfig
	_, err := toml.DecodeFile(file, &config)
	if err != nil {
		return nil, err
	}
	repos := make([]*Repository, 0)
	for _, c := range config.Constraint {
		repo := &Repository{
			Url: c.Name,
		}
		parts := strings.Split(c.Name, "/")
		if len(parts) > 0 {
			repo.Platform = util.GetPlatform(parts[0])
		}
		if len(parts) > 1 {
			repo.Owner = parts[1]
		}
		if len(parts) > 2 {
			repo.Name = parts[2]
		}
		repos = append(repos, repo)

	}
	return repos, nil
}

