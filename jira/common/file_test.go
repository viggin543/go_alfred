package common

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Remove("/tmp/da")
	run := m.Run()
	os.Remove("/tmp/da")
	os.Exit(run)
}

func TestName(t *testing.T) {
	AppendToFile("/tmp/da","da")
	AppendToFile("/tmp/da","da")


	content, err := ioutil.ReadFile("/tmp/da")
	if err != nil {
		t.Fatal(err)
	}

	if string(content) != "da\nda\n" {
		t.Fatal("file should contain dada and not " + string(content))
	}



}

//da

//da
