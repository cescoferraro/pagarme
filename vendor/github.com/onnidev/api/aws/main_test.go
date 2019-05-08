package aws_test

import (
	"testing"

	"github.com/onnidev/api/aws"
)

func TestSpec(t *testing.T) {
	t.Run("AWS List", func(t *testing.T) {
		aws.Run()
	})
}
