package registry

import (
    "encoding/json"
    "errors"
    "os"
    "path/filepath"
)

type Tool struct {
    Name    string `json:"name"`
    Path    string `json:"path"`
    Command string `json:"command"`
}

var configFile = filepath.Join(os.Getenv("HOME"), ".toolr", "config.json")

func ensureConfig() {
    dir := filepath.Dir(configFile)
    os.MkdirAll(dir, 0755)
    if _, err := os.Stat(configFile); os.IsNotExist(err) {
        os.WriteFile(configFile, []byte("[]"), 0644)
    }
}

func RegisterTool(name, path, command string) error {
    ensureConfig()
    tools := ListTools()
    for _, t := range tools {
        if t.Name == name {
            return errors.New("tool already exists")
        }
    }
    tools = append(tools, Tool{name, path, command})
    data, _ := json.MarshalIndent(tools, "", "  ")
    return os.WriteFile(configFile, data, 0644)
}

func ListTools() []Tool {
    ensureConfig()
    data, err := os.ReadFile(configFile)
    if err != nil {
        return []Tool{}
    }
    var tools []Tool
    json.Unmarshal(data, &tools)
    return tools
}

func GetTool(name string) (Tool, error) {
    tools := ListTools()
    for _, t := range tools {
        if t.Name == name {
            return t, nil
        }
    }
    return Tool{}, errors.New("tool not found")
}
