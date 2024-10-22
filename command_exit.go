package main

import "os"

func commandExit(config *config, args ...string) error {
	os.Exit(0)
	return nil
}
