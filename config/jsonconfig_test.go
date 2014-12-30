package config

import "testing"

func Test(t *testing.T) {

	config, err := NewConfig("testdata.json")
	if err != nil {
		t.Fatal(err)
	}
	if config.GetString("key1") != "hello" {
		t.Fatal("error")
	}

	if config.GetString("keyDoesNotExist") != "" {
		t.Fatal("error")
	}

	if config.GetString("obj1") != "" {
		t.Fatal("error")
	}

}
