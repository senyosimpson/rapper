package spotify

import (
	urllib "net/url"
	"path"
	"strconv"
	"strings"
)

func getURLParams(url string) (string, int, error) {
	u, _ := urllib.Parse(url)
	q, _ := urllib.ParseQuery(u.RawQuery)
	id := strings.Split(u.Path, "/")[3]
	offset, err := strconv.Atoi(q["offset"][0])
	if err != nil {
		return "", 0, err
	}
	return id, offset, nil
}

func constructURL(options *Options, baseURL string, elements ...string) string {
	u, _ := urllib.Parse(baseURL)
	urlElements := append([]string{u.Path}, elements...)
	u.Path = path.Join(urlElements...)

	if options != nil {
		query := urllib.Values{}
		setOptions(&query, options)
		u.RawQuery = query.Encode()
	}
	return u.String()
}

const nullString = ""
const nullInt = 0

func setOptions(query *urllib.Values, options *Options) {
	if options.Limit != nullInt {
		query.Set("limit", strconv.Itoa(options.Limit))
	}

	if options.Offset != nullInt {
		query.Set("offset", strconv.Itoa(options.Offset))
	}

	if options.Country != nullString {
		query.Set("country", options.Country)
	}

	if options.Market != nullString {
		query.Set("market", options.Market)
	}

	if options.Locale != nullString {
		query.Set("locale", options.Locale)
	}

	if options.IncludeGroup != nullString {
		query.Set("include_groups", options.IncludeGroup)
	}

	if options.Type != nullString {
		query.Set("type", options.Type)
	}

	if options.Query != nullString {
		query.Set("q", options.Query)
	}
}
