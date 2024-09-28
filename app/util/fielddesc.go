package util

import (
	"net/http"

	"github.com/pkg/errors"
	"github.com/samber/lo"
)

type FieldDesc struct {
	Key         string   `json:"key"`
	Title       string   `json:"title"`
	Description string   `json:"description,omitempty"`
	Type        string   `json:"type,omitempty"`
	Default     string   `json:"default,omitempty"`
	Choices     []string `json:"choices,omitempty"`
}

func (d FieldDesc) Parse(q any) (any, error) {
	switch d.Type {
	case "bool", "boolean":
		return ParseBool(q, "", true)
	case "float":
		return ParseFloat(q, "", true)
	case "int", "number":
		return ParseInt(q, "", true)
	case "string", "":
		return ParseString(q, "", true)
	case "[]string":
		return ParseArrayString(q, "", true)
	case "time":
		return ParseTime(q, "", true)
	default:
		return nil, errors.Errorf("unable to parse [%s] value from string [%s]", d.Type, q)
	}
}

type FieldDescs []*FieldDesc

func (d FieldDescs) Get(key string) *FieldDesc {
	return lo.FindOrElse(d, nil, func(x *FieldDesc) bool {
		return x.Key == key
	})
}

func (d FieldDescs) Keys() []string {
	return lo.Map(d, func(x *FieldDesc, _ int) string {
		return x.Key
	})
}

type FieldDescResults struct {
	FieldDescs FieldDescs `json:"fieldDescs"`
	Values     ValueMap   `json:"values"`
	Missing    []string   `json:"missing,omitempty"`
}

func (a *FieldDescResults) HasMissing() bool {
	return len(a.Missing) > 0
}

func FieldDescsCollect(r *http.Request, args FieldDescs) *FieldDescResults {
	ret := make(ValueMap, len(args))
	var missing []string
	qa := r.URL.Query()
	lo.ForEach(args, func(arg *FieldDesc, _ int) {
		if !qa.Has(arg.Key) {
			missing = append(missing, arg.Key)
		}
		ret[arg.Key] = qa.Get(arg.Key)
	})
	return &FieldDescResults{FieldDescs: args, Values: ret, Missing: missing}
}
