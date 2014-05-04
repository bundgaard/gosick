// Object and ObjectBase is an abstract class for all scheme expressions.
// A return value of a method which returns scheme object is Object.
// And ObjectBase has Object's implementation of String().

package scheme

type Object interface {
	String() string
}

type ObjectBase struct {
}

func (o *ObjectBase) String() string {
	return "This type's Eval() is not implemented yet."
}
