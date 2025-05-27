package learnfunctiondepth

import "fmt"

type fetchLanguage func(string) string

type details struct {
	languageDetails map[string]string
}

func LearnFunctionAsParam() {
	userLanguage := make(map[string]string, 3)
	userLanguage["Suraj"] = "Hindi"
	userLanguage["Pankaj"] = "Marathu"
	userLanguage["Naman"] = "Gujarati"

	langDetails := details{
		userLanguage,
	}
	langDetails.languagePopulation("Pankaj", langDetails.getLanguageDetails)

}

func (d details) languagePopulation(user string, getLanguage fetchLanguage) {

	userLangDetails := getLanguage(user)

	fmt.Println(userLangDetails)
}

func (d details) getLanguageDetails(userName string) string {
	return d.languageDetails[userName]
}
