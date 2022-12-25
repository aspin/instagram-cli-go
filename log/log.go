package log

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/urfave/cli/v2"
	"instagram-cli-go/flags"
	"os"
)

type Config struct {
	FileName string
}

func NewConfigFromCLI(c *cli.Context) Config {
	return Config{
		FileName: c.String(flags.LogFile.Name),
	}
}

func Init(cfg Config) (*os.File, error) {
	f, err := tea.LogToFile(cfg.FileName, "debug")
	return f, err
}
