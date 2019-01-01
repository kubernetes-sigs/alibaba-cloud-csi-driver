package disk

import (
	"log"
	"testing"
)

func setup() {
	logPath = "tmp/mount-entries"
	path := "tmp"
	_, err := run("mkdir -p " + path)
	if err != nil {
		log.Fatal(err)
	}
}

func teardown() {
	path := "tmp"
	_, err := run("rm -rf " + path)
	if err != nil {
		log.Fatal(err)
	}
}

func TestAttachEntry(t *testing.T) {
	setup()
	defer teardown()

	err := DefaultAttachEntry.Add("aaa", "bbb")
	if err != nil {
		t.Fatal(err)
	}
	err = DefaultAttachEntry.Add("ccc", "ddd")
	if err != nil {
		t.Fatal(err)
	}

	_, err = DefaultAttachEntry.Get("aaa")
	if err != nil {
		t.Fatal(err)
	}

	DefaultAttachEntry.Remove("aaa")
	_, err = DefaultAttachEntry.Get("aaa")
	if err == nil {
		t.Fail()
	}
}
