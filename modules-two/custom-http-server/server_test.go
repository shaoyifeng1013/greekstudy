package custom_http_server

import "testing"

func TestCustomerHttpServer(t *testing.T) {
	server := NewCustomHttpServer()
	server.s.ListenAndServe()
}
