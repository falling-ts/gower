package services

type TransMap map[string]string
type TransCategory map[string]TransMap
type TransAll map[string]any

type TranslateService interface {
	Service

	DBError(err error) error
}
