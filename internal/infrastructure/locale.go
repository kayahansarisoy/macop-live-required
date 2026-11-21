package infrastructure

import "golang.org/x/text/language"

// PreferredLanguage picks the language a notification is written in from what the client says it accepts.
// A header that cannot be read, or that names no language, falls back to Turkish.
func PreferredLanguage(accept string) language.Tag {
	tags, _, err := language.ParseAcceptLanguage(accept)
	if err != nil || len(tags) == 0 {
		return language.Turkish
	}
	return tags[0]
}
