package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type hooksConfig struct {
	PreCommit  string `json:"pre-commit"`
	PostCommit string `json:"post-commit"`
	PrePush    string `json:"pre-push"`
	PostPush   string `json:"post-push"`
}

func loadHooksConfig(filename string) (*hooksConfig, error) {

	data, err := os.ReadFile(filename)

	if err != nil {
		return nil, fmt.Errorf("error reading file %s: %v", filename, err)
	}

	var config hooksConfig

	err = json.Unmarshal(data, &config)

	if err != nil {
		return nil, fmt.Errorf("error parsing JSON: %v", err)
	}

	return &config, nil
}

func updateHookInConfig(configFile, scriptFile string, hookName string) error {

	scriptContent, err := os.ReadFile(scriptFile)

	if err != nil {
		return fmt.Errorf("error reading file %s: %v", scriptFile, err)
	}

	data, err := os.ReadFile(configFile)

	if err != nil {
		return fmt.Errorf("error reading configuration file %s: %v", configFile, err)
	}

	var config hooksConfig

	err = json.Unmarshal(data, &config)

	if err != nil {
		return fmt.Errorf("error parsing configuration JSON: %v", err)
	}

	var updateField *string

	switch hookName {
	case "pre-commit":
		updateField = &config.PreCommit
	case "post-commit":
		updateField = &config.PostCommit
	case "pre-push":
		updateField = &config.PrePush
	case "post-push":
		updateField = &config.PostPush
	default:
		return fmt.Errorf("unrecognized filename for hook: %s", scriptFile)
	}

	*updateField = string(scriptContent)

	updatedData, err := json.MarshalIndent(config, "", "  ")

	if err != nil {
		return fmt.Errorf("error serializing updated JSON: %v", err)
	}

	err = os.WriteFile(configFile, updatedData, 0644)

	if err != nil {
		return fmt.Errorf("error writing configuration file %s: %v", configFile, err)
	}

	fmt.Println("The configuration file has been updated with the content of the script:", scriptFile)

	return nil
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

func checkAndUpdateHooks(configFile string, hookDir string) error {

	hookFiles := []string{
		"pre-commit.sh",
		"post-commit.sh",
		"pre-push.sh",
		"post-push.sh",
	}

	for _, hookFile := range hookFiles {

		hookPath := filepath.Join(hookDir, hookFile)

		if _, err := os.Stat(hookPath); err == nil {

			fmt.Printf("The file %s exists. Updating configuration....\n", hookPath)

			hookName := strings.TrimSuffix(hookFile, ".sh")

			err := updateHookInConfig(configFile, hookPath, hookName)

			if err != nil {
				return fmt.Errorf("error updating configuration for %s: %v", hookPath, err)
			}

		} else if os.IsNotExist(err) {

		} else {
			return fmt.Errorf("error when verifying the existence of %s: %v", hookPath, err)
		}
	}

	return nil
}

func initializeConfigFile(configFile string) error {

	if _, err := os.Stat(configFile); err == nil {
		return nil
	} else if os.IsNotExist(err) {

		config := hooksConfig{}

		data, err := json.MarshalIndent(config, "", "  ")

		if err != nil {
			return fmt.Errorf("error serializing configuration file: %v", err)
		}

		err = os.WriteFile(configFile, data, 0644)

		if err != nil {
			return fmt.Errorf("error writing configuration file %s: %v", configFile, err)
		}

		fmt.Println("Configuration file created successfully:", configFile)

		return nil

	} else {
		return fmt.Errorf("error verifying file existence %s: %v", configFile, err)
	}
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("A directory must be provided for the hook files.")
		return
	}

	hookDir := os.Args[1]

	configFile := "hooks_config.json"

	err := initializeConfigFile(configFile)

	if err != nil {
		fmt.Println("Error initializing configuration file:", err)
		return
	}

	err = checkAndUpdateHooks(configFile, hookDir)

	if err != nil {
		fmt.Printf("Error verifying and updating hooks: %v\n", err)
		return
	}

	config, err := loadHooksConfig(configFile)

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
