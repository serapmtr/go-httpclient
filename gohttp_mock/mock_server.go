package gohttpmock

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
	"sync"
)

var (
	mockupServer = mockServer{
		mocks: make(map[string]*Mock),
	}
)

type mockServer struct {
	enabled     bool
	serverMutex sync.Mutex

	mocks map[string]*Mock
}

func StartMockServer() {
	mockupServer.serverMutex.Lock()
	defer mockupServer.serverMutex.Unlock()

	mockupServer.enabled = true
}

func StopMockServer() {
	mockupServer.serverMutex.Lock()
	defer mockupServer.serverMutex.Unlock()

	mockupServer.enabled = false
}

func DeleteMocks() {
	mockupServer.serverMutex.Lock()
	defer mockupServer.serverMutex.Unlock()

	mockupServer.mocks = make(map[string]*Mock)
}

func AddMock(mock Mock) {
	mockupServer.serverMutex.Lock()
	defer mockupServer.serverMutex.Unlock()

	key := mockupServer.getMockKey(mock.Method, mock.Url, mock.RequestBody)

	mockupServer.mocks[key] = &mock
}

func (m *mockServer) getMockKey(method, url, body string) string {
	hasher := md5.New()
	hasher.Write([]byte(method + url + m.cleanBody(body)))
	key := hex.EncodeToString(hasher.Sum(nil))
	fmt.Printf("KEY: '%s' \n", key)
	return key
}

func (m *mockServer) cleanBody(body string) string {
	body = strings.TrimSpace(body)
	if body == "" {
		return ""
	}
	body = strings.ReplaceAll(body, "\t", "") // body, old, new

	return body
}

func GetMock(method, url, body string) *Mock {
	// if m.enabled false, then in test we call actual http function
	if !mockupServer.enabled {
		return nil
	}
	if mock := mockupServer.mocks[mockupServer.getMockKey(method, url, body)]; mock != nil {
		return mock
	}

	return &Mock{
		Error: errors.New(fmt.Sprintf("No matching  %s from '%s' with given body", method, url)),
	}
}