package regkeys

import (
	"fmt"
	"strings"

	"golang.org/x/sys/windows/registry"
)

var keys = []string{
	"SYSTEM\\ControlSet001\\Control\\Class\\{4D36E968-E325-11CE-BFC1-08002BE10318}\\0000\\DriverDesc",
	"SYSTEM\\ControlSet001\\Control\\Class\\{4D36E968-E325-11CE-BFC1-08002BE10318}\\0000\\ProviderName",
}

func getRegKeyValues() map[string]string {
	keyValues := make(map[string]string)
	for _, fullKey := range keys {
		parts := strings.Split(fullKey, "\\")
		key := strings.Join(parts[:len(parts)-1], "\\")
		valueName := parts[len(parts)-1]
		regKey, err := registry.OpenKey(registry.LOCAL_MACHINE, key, registry.QUERY_VALUE)
		if err != nil {
			fmt.Printf("Failed to open registry key: %s, error: %v\n", key, err)
			continue
		}
		value, _, err := regKey.GetStringValue(valueName)
		if err != nil {
			fmt.Printf("Failed to get value for key: %s, error: %v\n", fullKey, err)
			regKey.Close()
			continue
		}
		keyValues[fullKey] = value
		regKey.Close()
	}
	return keyValues
}

func BackupRegKeys() {
	keyValues := getRegKeyValues()
	for fullKey, value := range keyValues {
		parts := strings.Split(fullKey, "\\")
		key := strings.Join(parts[:len(parts)-1], "\\")
		backupValueName := parts[len(parts)-1] + "_backup"
		regKey, err := registry.OpenKey(registry.LOCAL_MACHINE, key, registry.SET_VALUE)
		if err != nil {
			fmt.Printf("Failed to open registry key: %s, error: %v\n", key, err)
			continue
		}
		err = regKey.SetStringValue(backupValueName, value)
		if err != nil {
			fmt.Printf("Failed to set backup value for key: %s, error: %v\n", fullKey, err)
		}
		regKey.Close()
	}
}

func RestoreRegKeys() {
	for _, fullKey := range keys {
		parts := strings.Split(fullKey, "\\")
		key := strings.Join(parts[:len(parts)-1], "\\")
		valueName := parts[len(parts)-1]
		backupValueName := valueName + "_backup"
		regKey, err := registry.OpenKey(registry.LOCAL_MACHINE, key, registry.QUERY_VALUE|registry.SET_VALUE)
		if err != nil {
			fmt.Printf("Failed to open registry key: %s, error: %v\n", key, err)
			continue
		}
		backupValue, _, err := regKey.GetStringValue(backupValueName)
		if err != nil {
			fmt.Printf("Failed to get backup value for key: %s, error: %v\n", fullKey, err)
			regKey.Close()
			continue
		}
		err = regKey.SetStringValue(valueName, backupValue)
		if err != nil {
			fmt.Printf("Failed to restore value for key: %s, error: %v\n", fullKey, err)
			regKey.Close()
			continue
		}
		err = regKey.DeleteValue(backupValueName)
		if err != nil {
			fmt.Printf("Failed to delete backup value for key: %s, error: %v\n", fullKey, err)
		}
		regKey.Close()
	}
}

func BackupExists() bool {
	for _, fullKey := range keys {
		parts := strings.Split(fullKey, "\\")
		key := strings.Join(parts[:len(parts)-1], "\\")
		backupValueName := parts[len(parts)-1] + "_backup"
		regKey, err := registry.OpenKey(registry.LOCAL_MACHINE, key, registry.QUERY_VALUE)
		if err != nil {
			fmt.Printf("Failed to open registry key to check backup: %s, error: %v\n", key, err)
			continue
		}
		_, _, err = regKey.GetStringValue(backupValueName)
		if err == nil {
			regKey.Close()
			return true
		}
		regKey.Close()
	}
	return false
}

func ChangeRegKeys() {
	if !BackupExists() {
		BackupRegKeys()
	}

	keyValues := getRegKeyValues()
	for fullKey := range keyValues {
		parts := strings.Split(fullKey, "\\")
		key := strings.Join(parts[:len(parts)-1], "\\")
		valueName := parts[len(parts)-1]
		regKey, err := registry.OpenKey(registry.LOCAL_MACHINE, key, registry.SET_VALUE)
		if err != nil {
			fmt.Printf("Failed to open registry key: %s, error: %v\n", key, err)
			continue
		}
		err = regKey.SetStringValue(valueName, "VBox")
		if err != nil {
			fmt.Printf("Failed to set changed value for key: %s, error: %v\n", fullKey, err)
		}
		regKey.Close()
	}
}
