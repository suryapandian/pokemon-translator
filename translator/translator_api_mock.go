package translator

type TranslatorMock struct{}

func (t *TranslatorMock) GetTranslation(data string) (translatedData string, err error) {
	translatedData = data + "translated"
	return
}
