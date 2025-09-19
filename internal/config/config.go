package config

import (
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

type Rule struct {
	Query       string `yaml:"query"`
	URL         string `yaml:"url"`
	Passthrough bool   `yaml:"passthrough"`
}

type Route struct {
	Default     string `yaml:"default"`
	Passthrough bool   `yaml:"passthrough"`
	Rules       []Rule `yaml:"rules"`
}

type Config struct {
	Routes map[string]Route `yaml:"routes"`
}

func LoadConfig(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

func (c *Config) FindRoute(shortcode, queryString string) (string, bool) {
	if route, exists := c.Routes[shortcode]; exists {
		// Check rules in order
		for _, rule := range route.Rules {
			if matchesQuery(rule.Query, queryString) {
				if rule.Passthrough && queryString != "" {
					return rule.URL + "?" + queryString, true
				}
				return rule.URL, true
			}
		}
		// Fall back to default if no rules match
		if route.Default != "" {
			if route.Passthrough && queryString != "" {
				return route.Default + "?" + queryString, true
			}
			return route.Default, true
		}
	}

	return "", false
}

func matchesQuery(ruleQuery, actualQuery string) bool {
	if ruleQuery == "" && actualQuery == "" {
		return true
	}
	if ruleQuery == "" {
		return false
	}

	// Simple exact match for now
	return strings.Contains(actualQuery, ruleQuery)
}