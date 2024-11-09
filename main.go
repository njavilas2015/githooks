package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type HooksConfig struct {
	PreCommit  string `json:"pre-commit"`
	PostCommit string `json:"post-commit"`
	PrePush    string `json:"pre-push"`
	PostPush   string `json:"post-push"`
}

func loadHooksConfig(filename string) (*HooksConfig, error) {

	data, err := os.ReadFile(filename)

	if err != nil {
		return nil, fmt.Errorf("error reading file %s: %v", filename, err)
	}

	var config HooksConfig

	err = json.Unmarshal(data, &config)

	if err != nil {
		return nil, fmt.Errorf("error parsing JSON: %v", err)
	}

	return &config, nil
}

func createGitHook(hookName string, hookContent string) error {

	hookPath := fmt.Sprintf(".git/hooks/%s", hookName)

	hookFile, err := os.OpenFile(hookPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)

	if err != nil {
		return fmt.Errorf("error creating the hook %s: %v", hookName, err)
	}

	defer hookFile.Close()

	_, err = hookFile.WriteString(hookContent)

	if err != nil {
		return fmt.Errorf("error writing to hook %s: %v", hookName, err)
	}

	err = os.Chmod(hookPath, 0755)

	if err != nil {
		return fmt.Errorf("error when changing hook permissions %s: %v", hookName, err)
	}

	return nil
}

func main() {

	config, err := loadHooksConfig("hooks_config.json")

	if err != nil {
		fmt.Printf("Error loading configuration: %v\n", err)
		return
	}

	hooks := map[string]string{
		"pre-commit":  config.PreCommit,
		"post-commit": config.PostCommit,
		"pre-push":    config.PrePush,
		"post-push":   config.PostPush,
	}

	for hookName, hookContent := range hooks {

		err := createGitHook(hookName, hookContent)

		if err != nil {
			fmt.Printf("Error creating the hook %s: %v\n", hookName, err)
			return
		}
	}

	fmt.Println("Hooks configured correctly.")
}
