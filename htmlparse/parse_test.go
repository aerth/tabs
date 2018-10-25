package htmlparse

import (
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	var htm = `<!DOCTYPE html>
<html>
<head>
    <title></title>
</head>
<body>
 body content
 <p>more content</p>
 <pre>gold</pre>
</body>
</html>`

	b, err := Parse(strings.NewReader(htm))
	if err != nil {
		t.Error(err)
	}
	if string(b) != "<pre>gold</pre>" {
		t.Logf("wanted <pre>gold</pre>, got %q", string(b))
		t.FailNow()
	}
}
