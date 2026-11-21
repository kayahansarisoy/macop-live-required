package infrastructure

import "gopkg.in/yaml.v2"

// NotificationSettings is what the notification sender reads from its settings file.
type NotificationSettings struct {
	Sender   string `yaml:"sender"`
	Language string `yaml:"language"`
}

// ReadNotificationSettings reads the notification settings from their YAML form.
func ReadNotificationSettings(document []byte) (NotificationSettings, error) {
	var settings NotificationSettings
	err := yaml.Unmarshal(document, &settings)
	return settings, err
}
