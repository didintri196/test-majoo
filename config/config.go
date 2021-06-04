package config

import (
	"os"
	"path/filepath"
)

type Config struct{}

func GetDir() string {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return dir
}
