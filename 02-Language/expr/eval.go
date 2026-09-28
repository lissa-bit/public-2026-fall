//go:build !solution

package expr

// Eval evaluates an arithmetic expression tree.
func Eval(expression *Expression) (int64, error) {
	if expression == nil {
		return 0, ErrInvalidExpression
	}

	if expression.Operation == OperationValue {
		return expression.Value, nil
	}

	if expression.Operation != OperationAdd &&
		expression.Operation != OperationDivide &&
		expression.Operation != OperationMultiply &&
		expression.Operation != OperationSubtract {
		return 0, ErrInvalidExpression
	}

	if expression.Left == nil || expression.Right == nil {
		return 0, ErrInvalidExpression
	}

	left, err := Eval(expression.Left)
	if err != nil {
		return 0, err
	}

	right, err := Eval(expression.Right)
	if err != nil {
		return 0, err
	}

	if expression.Operation == OperationAdd {
		return left + right, nil
	}

	if expression.Operation == OperationSubtract {
		return left - right, nil
	}

	if expression.Operation == OperationMultiply {
		return left * right, nil
	}

	if right == 0 {
		return 0, ErrDivisionByZero
	}

	return left / right, nil
}
