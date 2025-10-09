package types

import "github.com/limes-cloud/kratosx/model/hook"

type GetScopeResponse struct {
	AllDept    bool
	DeptScopes []uint32
	Rule       *hook.ConditionGroup
	Fields     []string
}
