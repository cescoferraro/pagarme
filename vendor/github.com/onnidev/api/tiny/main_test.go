package tiny_test

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/onnidev/api/infra"
	"github.com/onnidev/api/tiny"
)

// TestSpec isk kjsdf
func TestSpecReturn(t *testing.T) {
	t.Run("Get Tiny.com contatos", func(t *testing.T) {
		u, _ := url.Parse("https://api.tiny.com.br/api2/contatos.pesquisa.php")
		q := u.Query()
		q.Set("pesquisa", "")
		q.Set("formato", "json")
		q.Set("token", tiny.TOKEN)
		u.RawQuery = q.Encode()
		fmt.Println(u)
		ajax := infra.Ajax{
			Method: "POST",
			Path:   u.String(),
		}
		_, _ = infra.NewTestRequest(t, ajax)
	})
}
