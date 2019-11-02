package db

import (
	"testing"
)

func TestGetUserSubmitResultCount(t *testing.T) {
	userID := 85
	res, err := GetUserSubmitResultCount(userID)
	if err != nil {
		t.Errorf("error:%s\n", err)
	}
	t.Log(res)
}
