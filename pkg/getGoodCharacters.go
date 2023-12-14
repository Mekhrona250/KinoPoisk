package pkg

import "kinopoisk/pkg/models"

func GetGoodCharacters() []models.Character {
	result := make([] models.Character,0)
	for _, val := range Characters {
		if val.IsEvil == false {
			result = append(result,val)
		}
		
	}
	return result

}
