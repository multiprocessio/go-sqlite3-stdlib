package stdlib

import "net/url"

func urlHost(u string) string {
	p, err := url.Parse(u)
	if err != nil {
		return ""
	}

	return p.Host
}

func urlPort(u string) string {
	p, err := url.Parse(u)
	if err != nil {
		return ""
	}

	return p.Port()
}

func urlScheme(u string) string {
	p, err := url.Parse(u)
	if err != nil {
		return ""
	}

	return p.Scheme
}

func urlPath(u string) string {
	p, err := url.Parse(u)
	if err != nil {
		return ""
	}

	return p.Path
}

func urlParam(u, k string) string {
	p, err := url.Parse(u)
	if err != nil {
		return ""
	}

	return p.Query().Get(k)
}

func urlFragment(u string) string {
	p, err := url.Parse(u)
	if err != nil {
		return ""
	}

	return p.Fragment
}

var netFunctions = map[string]any{
	"url_host": urlHost,
	"url_port": urlPort,
	"url_scheme": urlScheme,
	"url_path": urlPath,
	"url_param": urlParam,
	"url_fragment": urlFragment,
}
