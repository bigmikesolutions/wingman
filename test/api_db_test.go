package test

import "testing"

func Test_Api_Database_ShouldGetInfo(t *testing.T) {
	s := NewApiDatabaseStage(t)
	defer s.Close()

	s.Given().
		ServerIsUpAndRunning()

	s.When().
		QueryDatabase("test-env", "test-postgres")

	s.Then().
		NoClientError()
}
