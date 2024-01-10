package schema

import (
	"google.golang.org/protobuf/proto"

	schemapb "github.com/TBD54566975/ftl/protos/xyz/block/ftl/v1/schema"
)

// VerbRef is a reference to a Verb.
type VerbRef Ref

var _ Type = (*VerbRef)(nil)

func (*VerbRef) schemaChildren() []Node { return nil }
func (*VerbRef) schemaType()            {}
func (v VerbRef) String() string        { return makeRef(v.Module, v.Name) }

func (v *VerbRef) ToProto() proto.Message {
	return &schemapb.VerbRef{
		Pos:    posToProto(v.Pos),
		Name:   v.Name,
		Module: v.Module,
	}
}

func VerbRefFromProto(s *schemapb.VerbRef) *VerbRef {
	return &VerbRef{
		Pos:    posFromProto(s.Pos),
		Name:   s.Name,
		Module: s.Module,
	}
}

func verbRefListToSchema(s []*schemapb.VerbRef) []*VerbRef {
	var out []*VerbRef
	for _, n := range s {
		out = append(out, VerbRefFromProto(n))
	}
	return out
}