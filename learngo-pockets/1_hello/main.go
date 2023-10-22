package main

import (
	"flag"
	"fmt"
)

func main() {
	var lang string
	flag.StringVar(&lang, "lang",
		"en", "the required language")
	flag.Parse()

	greeting := greet(language(lang))
	fmt.Println(greeting)
}

// language represents the language's code
type language string

// phrasebook holds greeting for each supported language
var phrasebook = map[language]string{
	"el": "Χαίρετε Κόσμε",     // Greek
	"en": "Hello world",       // English
	"fr": "Bonjour le monde",  // French
	"he": "םולש םלוע",         // Hebrew
	"ur": "ﺎﯿﻧد ﻮﻠﯿﮨ",         // Urdu
	"vi": "Xin chào Thế Giới", // Vietnamese
	"bn": "শুভেচ্ছা পৃথিবী",   // Bengali
}

// greet says hello to world in a specific language
func greet(l language) string {
	greeting, exists := phrasebook[l]
	if !exists {
		return fmt.Sprintf("unsupported language %q", l)
	}
	return greeting
}
