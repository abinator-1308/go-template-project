package config

import _ "embed"

// NOTE - THIS IS USED IN INTEGRATION TESTS. In Main application raed from the file (Avoid using this variable)

//go:embed app.yaml
var ApplicationConfigString string
