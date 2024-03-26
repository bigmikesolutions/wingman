package test

import "testing"

func Test_ApiServer_Should(t *testing.T) {
	s := NewApiStage(t)

	defer s.Close()
}
