// Code generated by FTL. DO NOT EDIT.
package other

import "github.com/TBD54566975/ftl/go-runtime/ftl/reflection"

func init() {
	reflection.Register(
		reflection.SumType[SecondTypeEnum](
			*new(A),
			*new(B),
		),
		reflection.SumType[TypeEnum](
			*new(MyBool),
			*new(MyBytes),
			*new(MyFloat),
			*new(MyInt),
			*new(MyTime),
			*new(MyList),
			*new(MyMap),
			*new(MyString),
			*new(MyStruct),
			*new(MyOption),
			*new(MyUnit),
		),
	)
}