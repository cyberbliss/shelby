package mods

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	"github.com/talal/go-bits/color"
	"gopkg.in/ini.v1"
)

func getActiveConfig() (string, error) {
	var cpath string
	// if running in GCP Cloud shell, the gcloud config path is different
	acFile, isCS := os.LookupEnv("CLOUDSDK_CONFIG")
	if isCS {
		acPath := path.Join(acFile, "active_config")
		ac, err := ioutil.ReadFile(acPath)
		if err != nil {
			return "", nil
		}
		cpath = path.Join(acFile, "configurations", "config_"+string(ac))
	} else {
		acPath := path.Join(homePath(), ".config", "gcloud", "active_config")
		absolutePath, err := filepath.Abs(acPath)
		if err != nil {
			return "", err
		}

		ac, err := ioutil.ReadFile(absolutePath)
		if err != nil {
			return "", err
		}
		cpath = path.Join(homePath(), ".config", "gcloud", "configurations", "config_"+string(ac))
	}

	return cpath, nil

}

func dispActiveGcloudContext(status chan string) {
	const icon = `☁`

	cfPath, err := getActiveConfig()
	if err != nil {
		status <- ""
		return
	}

	cfg, err := ini.Load(cfPath)
	if err != nil {
		status <- ""
		return
	}
	gcloudProject := cfg.Section("core").Key("project").String()

	if gcloudProject != "" {
		status <- color.Sprintf(color.BrightBlue, "%v %s", icon, gcloudProject)
	} else {
		status <- ""
	}

}
