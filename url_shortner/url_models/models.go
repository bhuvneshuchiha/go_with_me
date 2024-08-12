package url_models

type URL_STR struct {
	url       string `json:"url_original"`
	id        int
	short_url string `json:"short_url"`
}
