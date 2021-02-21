package translator

func (t *TranslatorStore) GetTranslation(data string, translatorAPI TranslatorAPI) (translatedData string, err error) {
	translatedData, err = t.Get(data)
	if err == nil {
		return translatedData, err
	}

	translatedData, err = translatorAPI.GetTranslation(data)
	if err != nil {
		return "", err
	}

	t.Save(data, translatedData)
	return translatedData, err
}
