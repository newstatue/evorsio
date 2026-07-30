package s3

import (
	"strings"
	"testing"
)

// 1. 65a8e27d8879283831b664bd8b7f0ad4 Hello, World!
// 2. 65a8e27d8879283831b664bd8b7f0ad4 Hello, World!
// 3. 82bb413746aee42f89dea2b59614f9ef Hello, World
// 4. 65a8e27d8879283831b664bd8b7f0ad4 Hello, World!
func TestS3(t *testing.T) {
	client, err := New("http://localhost:8333", "admin", "secret", "evorsio")
	if err != nil {
		t.Fatal(err)
	}
	context := t.Context()
	upload, err := client.Upload(context, "test.txt", strings.NewReader("Hello, World!"), "text/plain")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Uploaded object: %s", upload)
}
