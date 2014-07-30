package gojenkins

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Auth struct {
	Username string
	ApiToken string
}

type Jenkins struct {
	auth    *Auth
	baseUrl string
}

type Options struct {
	useApiEndPoint bool
}

type Crumb struct {
	Crumb             string `json:"crumb"`
	CrumbRequestField string `json:"crumbRequestField"`
}

var DefaultOptions = Options{}

func init() {
	DefaultOptions.useApiEndPoint = true
}

func NewJenkins(auth *Auth, baseUrl string) *Jenkins {
	return &Jenkins{
		auth:    auth,
		baseUrl: baseUrl,
	}
}

func (jenkins *Jenkins) buildUrl(path string, params url.Values, options Options) (requestUrl string) {
	requestUrl = jenkins.baseUrl + path

	// Not the most efficient, but it is easier to read
	if options.useApiEndPoint {
		requestUrl = requestUrl + "/api/json"
	}

	if params != nil {
		queryString := params.Encode()
		if queryString != "" {
			requestUrl = requestUrl + "?" + queryString
		}
	}

	return
}

func (jenkins *Jenkins) sendRequest(req *http.Request) (*http.Response, error) {
	/* For Crumb on CSRF Protection, we need to obtain the crumb header needed in
	 * a post request.
	 */
	crb := Crumb{}
	jenkins.get("/crumbIssuer", nil, &crb, Options{useApiEndPoint: true})

	if crb.Crumb != "" && crb.CrumbRequestField != "" {
		req.Header.Set(crb.CrumbRequestField, crb.Crumb)
	}

	req.SetBasicAuth(jenkins.auth.Username, jenkins.auth.ApiToken)
	return http.DefaultClient.Do(req)
}

func (jenkins *Jenkins) parseResponse(resp *http.Response, body interface{}) (err error) {
	defer resp.Body.Close()

	if body == nil {
		return
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	return json.Unmarshal(data, body)
}

func (jenkins *Jenkins) get(path string, params url.Values, body interface{}, options Options) (err error) {
	requestUrl := jenkins.buildUrl(path, params, options)
	req, err := http.NewRequest("GET", requestUrl, nil)
	// always required basic auth on both GET/POST
	req.SetBasicAuth(jenkins.auth.Username, jenkins.auth.ApiToken)

	if err != nil {
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	return jenkins.parseResponse(resp, body)
}

func (jenkins *Jenkins) post(path string, params url.Values, body interface{}, options Options) (err error) {
	requestUrl := jenkins.buildUrl(path, params, options)

	req, err := http.NewRequest("POST", requestUrl, nil)
	if err != nil {
		return
	}

	resp, err := jenkins.sendRequest(req)
	if err != nil {
		return
	}

	return jenkins.parseResponse(resp, body)
}
