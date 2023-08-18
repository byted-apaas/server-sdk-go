package operator

import "github.com/byted-apaas/server-sdk-go/service/data/cond"

var (
	And                = cond.And
	Or                 = cond.Or
	Eq                 = cond.Eq
	Neq                = cond.Neq
	Gt                 = cond.Gt
	Gte                = cond.Gte
	Lt                 = cond.Lt
	Lte                = cond.Lte
	IsOnOrBefore       = cond.IsOnOrBefore
	IsOnOrAfter        = cond.IsOnOrAfter
	Contain            = cond.Contain
	NotContain         = cond.NotContain
	In                 = cond.In
	NotIn              = cond.NotIn
	Empty              = cond.Empty
	NotEmpty           = cond.NotEmpty
	HasAnyOf           = cond.HasAnyOf
	HasAnyOfHierarchy  = cond.HasAnyOfHierarchy
	HasNoneOf          = cond.HasNoneOf
	HasNoneOfHierarchy = cond.HasNoneOfHierarchy
)
