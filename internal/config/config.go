package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type ChainConfig struct {
	Name     string `yaml:"name"`
	RpcURL   string `yaml:"rpc_url"`
	ChainID  int64  `yaml:"chain_id"`
	Symbol   string `yaml:"symbol"`
	Explorer string `yaml:"explorer"`
}

func LoadChainConfig(path string) (*ChainConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg ChainConfig
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

func LoadAllChains(dir string) ([]ChainConfig, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var chains []ChainConfig
	for _, f := range files {
		if filepath.Ext(f.Name()) == ".yaml" || filepath.Ext(f.Name()) == ".yml" {
			cfg, err := LoadChainConfig(filepath.Join(dir, f.Name()))
			if err != nil {
				return nil, fmt.Errorf("failed to load %s: %w", f.Name(), err)
			}
			chains = append(chains, *cfg)
		}
	}
	return chains, nil
}