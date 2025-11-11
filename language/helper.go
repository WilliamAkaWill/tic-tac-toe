package language

import "github.com/WilliamAkaWill/tic-tac-toe/shared"

func GetLanguageService(lang shared.Language) shared.LanguageService {
	switch lang {
	case shared.German:
		return NewGerman()
	case shared.English:
		return NewEnglish()
	}
	return nil
}