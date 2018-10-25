package htmlparse

import (
	"strings"
	"testing"
)

// TestParse tests that <pre> are not lost
func TestParse(t *testing.T) {

	var htm = `<!DOCTYPE html>
<html>
<head>
    <title></title>
</head>
<body>
 body content
 <p>more content</p>
 <pre>gold1</pre>
 <p>more content</p>
 <pre>gold2</pre>
</body>
</html>`
	b, err := Parse(strings.NewReader(htm))
	if err != nil {
		t.Error(err)
	}
	gold := `<pre>gold1</pre><pre>gold2</pre>`
	if string(b) != gold {
		t.Logf("wanted %s, got %q", gold, string(b))
		t.FailNow()
	}
}
