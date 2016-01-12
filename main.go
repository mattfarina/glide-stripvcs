package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/Masterminds/glide/cfg"
	"github.com/Masterminds/vcs"
)

func main() {
	yml, err := ioutil.ReadFile("glide.lock")
	if err != nil && os.IsNotExist(err) {
		fmt.Println("glide.lock file not found. Please run glide up to generate it.")
		os.Exit(1)
	}
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	l, err := cfg.LockfileFromYaml(yml)
	if err != nil {
		fmt.Println("Unable to parse glide.lock file.")
		os.Exit(1)
	}

	for _, d := range l.Imports {
		dir := filepath.Join("vendor", filepath.FromSlash(d.Name))
		t, err := vcs.DetectVcsFromFS(dir)
		if err != nil {
			fmt.Println(err)
			continue
		}

		p := filepath.Join(dir, "."+string(t))
		if _, err = os.Stat(p); os.IsNotExist(err) {
			fmt.Println("VCS data not present for", d.Name)
			continue
		}

		fmt.Println("Removing VCS data for", d.Name)
		err = os.RemoveAll(p)
		if err != nil {
			fmt.Println(err)
		}
	}
}
