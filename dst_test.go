package dst

import (
	"os"
	"testing"
)

func Test(t *testing.T) {
	p := NewTemplate("hello, {{.Name}}\n")

	p.Prepare(map[string]string{
		"Name": "{{.Username}}",
	})

	p.Execute(os.Stdout, map[string]string{
		"Username": "kaaaaaaarl!",
	})
	p.Execute(os.Stdout, map[string]string{
		"Username": "another!",
	})
	p.Execute(os.Stdout, map[string]string{
		"Username": "wowowowowow!",
	})
}
