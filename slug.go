package areena

import (
	"github.com/tv42/slug"
)

// Slugify create slug e.g. for filename
func Slugify(s string) (res string) {
	return slug.Slug(s)
}
