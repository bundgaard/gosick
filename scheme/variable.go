// Scheme's identifier is classified to a symbol or a variable.
// And this type owns a role to express a variable.
// Variable itself does not have a value for identifier,
// interpreter searches it from its code block scope by Variable's identifier.

package scheme

type Variable struct {
	ObjectBase
	identifier string
}

func NewVariable(identifier string) *Variable {
	return &Variable{identifier: identifier}
}
