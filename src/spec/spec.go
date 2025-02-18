package spec

import (
	genspec "github.com/michaelawyu/cloudevents-generator/src/generator/spec"
	"github.com/michaelawyu/cloudevents-generator/src/logger"
	utils "github.com/michaelawyu/cloudevents-generator/src/utils"
)

// CEGenSpec is
type CEGenSpec struct {
	Events   map[string]Event `json:"events" yaml:"events"`
	Metadata genspec.Metadata `json:"metadata" yaml:"metadata"`
}

// Parse is
func (spec *CEGenSpec) Parse() ([]genspec.Mod, genspec.Metadata) {
	if spec.Events == nil {
		logger.Logger.Fatal("no events are specified")
	}
	checkMetadataValidity(spec)

	var ms []genspec.Mod
	for n, e := range spec.Events {
		ek, dks := e.parse(n)
		m := genspec.Mod{
			ModName:     utils.FormatName(n, "camel"),
			Event:       ek,
			DataClasses: dks,
		}
		ms = append(ms, m)
	}
	return ms, spec.Metadata
}
