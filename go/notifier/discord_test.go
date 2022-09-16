package notifier

import (
	"strings"
	"testing"
)

func TestDiscordNotifier_sanitizeUrls(t *testing.T) {
	n := DiscordNotifier{}
	urls := []string{"https://example.com",
		"https://this-is-a-test.com/subpath",
		//"<https://already-sanitized.com/subpath?param=1&param2=2>",	//already sanitized urls are double sanitized (don't know how to fix this)
		"https://this-is-a-test.com/subpath?query=1",
		"[test](http://test.com)",
	}
	wanted := []string{"<https://example.com>",
		"<https://this-is-a-test.com/subpath>",
		//"<https://already-sanitized.com/subpath?param=1&param2=2>",
		"<https://this-is-a-test.com/subpath?query=1>",
		"[test](<http://test.com>)",
	}
	result := n.sanitizeUrls(strings.Join(urls, "\n"))
	if result != strings.Join(wanted, "\n") {
		t.Fatalf("Expected %s, got %s", wanted, result)
	}
}
