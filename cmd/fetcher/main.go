package main

import (
	"fmt"
	"os"
	"os/signal"
	"path"
	"syscall"

	"github.com/billy-editiano/dsfetch/cmd/fetcher/misc"
	"github.com/billy-editiano/dsfetch/internal/pkg/datasaur"
	"github.com/billy-editiano/dsfetch/internal/pkg/utils"
	"github.com/robfig/cron/v3"
)

func main() {
	fmt.Println("Hello, this is dsfetcher.fetcher")

	hosts := loadHosts()
	hitFn := misc.HitCallback(hosts)

	c := cron.New()
	c.AddFunc("*/1 * * * *", hitFn)
	c.Start()

	// first run
	hitFn()

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	sig := <-sigchan
	fmt.Println("Received signal", sig)
}

func appConfigDirName() string {
	return "dsfetcher"
}

func appConfigFileName() string {
	return "config.json"
}

func loadHosts() map[string]datasaur.Host {
	homeDir := utils.GetCurrentUserHomeDir()
	configPath := path.Join(homeDir, ".config", appConfigDirName(), appConfigFileName())
	ensureDirectory()
	ensureConfigFile()

	hosts := datasaur.GetHostsFromConfig(configPath)

	fmt.Println("Found", len(hosts), "hosts")
	for _, host := range hosts {
		fmt.Println("  ", host.Name, "=>", host.Host)
	}
	return hosts
}

func ensureDirectory() {
	homeDir := utils.GetCurrentUserHomeDir()
	configDir := path.Join(homeDir, ".config", appConfigDirName())
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		os.MkdirAll(configDir, 0755)
	}
}

func ensureConfigFile() {
	homeDir := utils.GetCurrentUserHomeDir()
	configPath := path.Join(homeDir, ".config", appConfigDirName(), appConfigFileName())
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		os.WriteFile(configPath, []byte("{}"), 0644)
	}
}
