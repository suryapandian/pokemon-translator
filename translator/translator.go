package translator

type TranslatorService struct {
	ApiClient TranslatorAPI
	Cache     TranslatorCache
}

func NewTranslatorService(apiClient TranslatorAPI, cache TranslatorCache) *TranslatorService {
	return &TranslatorService{ApiClient: apiClient, Cache: cache}
}

func (t *TranslatorService) GetTranslation(data string) (translatedData string, err error) {
	translatedData, err = t.Cache.Get(data)
	if err == nil {
		return translatedData, err
	}

	translatedData, err = t.ApiClient.GetTranslation(data)
	if err != nil {
		return "", err
	}

	t.Cache.Save(data, translatedData)
	return translatedData, err
}
