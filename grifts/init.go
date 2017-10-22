package grifts

import (
  "github.com/gobuffalo/buffalo"
	"site1/actions"
)

func init() {
  buffalo.Grifts(actions.App())
}
