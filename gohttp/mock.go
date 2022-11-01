package gohttp

var (
	mocks map[string]*Mock
)

// mocking key => method + url + body

type Mock struct {
	Method      string
	Url         string
	RequestBody string

	Error              error
	ResponseBody       string
	ResponseStatusCode int
}
