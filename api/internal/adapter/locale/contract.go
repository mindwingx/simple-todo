package locale

type ILocale interface {
	Init()
	Get(key string) string
	Plural(key string, params map[string]string) string
}
