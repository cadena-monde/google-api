// Copyright 2011 Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func openUrl(url string) {
	var err error
	if runtime.GOOS == "windows" {
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
		if err == nil {
			return
		}
	} else {
		try := []string{"xdg-open", "google-chrome", "open"}
		for _, bin := range try {
			err = exec.Command(bin, url).Run()
			if err == nil {
				return
			}
		}
	}
	log.Printf("Error opening URL in browser: %v", err)
}

func osUserCacheDir() string {
	switch runtime.GOOS {
	case "darwin":
		return filepath.Join(os.Getenv("HOME"), "Library", "Caches")
	case "linux", "freebsd":
		return filepath.Join(os.Getenv("HOME"), ".cache")
	}
	log.Printf("TODO: osUserCacheDir on GOOS %q", runtime.GOOS)
	return "."
}

func condDebugTransport(rt http.RoundTripper) http.RoundTripper {
	//TODO: Implementar log
	// if *debug {
	// 	return &logTransport{rt}
	// }
	return rt
}
