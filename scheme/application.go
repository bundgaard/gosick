// Application is a type to express scheme procedure application.
// Application has a procedure and its argument as list which consists
// of Pair.

package scheme

type Application struct {
	ObjectBase
	procedure Object
	arguments Object
}

func NewApplication(parent Object) *Application {
	return &Application{
		ObjectBase: ObjectBase{parent: parent},
	}
}

func (a *Application) Eval() Object {
	return a.applyProcedure()
}

func (a *Application) String() string {
	pair := NewPair(nil)
	pair.Car = a.procedure
	pair.Cdr = a.arguments
	return pair.String()
}

func (a *Application) applyProcedure() Object {
	evaledObject := a.procedure.Eval()

	if evaledObject.isProcedure() {
		procedure := evaledObject.(*Procedure)
		return procedure.Invoke(a.arguments)
	} else if evaledObject.isSyntax() {
		syntax := evaledObject.(*Syntax)
		return syntax.Invoke(a.arguments)
	} else {
		runtimeError("invalid application")
		return nil
	}
}

func (a *Application) isApplication() bool {
	return true
}
