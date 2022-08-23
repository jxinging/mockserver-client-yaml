package mockyaml

type MockExpectation struct {
	HttpRequest struct {
		Method                string      `json:"method" yaml:"method"`
		Path                  string      `json:"path" yaml:"path"`
		Headers               interface{} `json:"headers,omitempty" yaml:"headers"`
		QueryStringParameters interface{} `json:"queryStringParameters,omitempty" yaml:"queryStringParameters"`
		Body                  interface{} `json:"body,omitempty" yaml:"body"`
	} `json:"httpRequest" yaml:"httpRequest"`
	HttpResponse struct {
		StatusCode int         `json:"statusCode,omitempty" yaml:"statusCode"`
		Headers    interface{} `json:"headers,omitempty" yaml:"headers"`
		Body       interface{} `json:"body,omitempty" yaml:"body"`
	} `json:"httpResponse" yaml:"httpResponse"`
}
