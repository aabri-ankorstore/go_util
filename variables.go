package go_util

import (
	"embed"
)

var (
	AppName    = "ankor"
	TemplateFS embed.FS
	PluginFS   embed.FS
	confDir    string
	Force      bool
)

const (
	Homebrew = "homebrew"
)

func init() {
	dirs := NewDirs()
	confDir = dirs.GetConfigDir()
}
