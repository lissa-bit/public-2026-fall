//go:build !solution

package bankaccount

// NewAccount creates an account with the supplied initial balance.
func NewAccount(initial Money) (*Account, error) {
	if initial < 0 {
		return nil, ErrInvalidAmount
	}
	return &Account{balance: initial}, nil
}

// Balance returns the current account balance.
func Balance(account *Account) (Money, error) {
	if account == nil {
		return 0, ErrNilAccount
	}

	if account.balance < 0 {
		return 0, ErrInvalidAmount
	}

	return account.balance, nil
}

// Deposit adds money to the account.
func Deposit(account *Account, amount Money) error {
	if account == nil {
		return ErrNilAccount
	}

	if account.balance < 0 {
		return ErrInvalidAmount
	}

	if amount < 0 {
		return ErrInvalidAmount
	}

	account.balance += amount
	return nil
}

// Withdraw removes money from the account.
func Withdraw(account *Account, amount Money) error {
	if account == nil {
		return ErrNilAccount
	}

	if account.balance < 0 {
		return ErrInvalidAmount
	}

	if (account.balance - amount) < 0 {
		return ErrInsufficientFunds
	}

	if amount < 0 {
		return ErrInvalidAmount
	}

	account.balance -= amount
	return nil
}

// Transfer moves money from one account to another.
func Transfer(from, to *Account, amount Money) error {
	if from == nil || to == nil {
		return ErrNilAccount
	}

	if from.balance < 0 || to.balance < 0 {
		return ErrInvalidAmount
	}

	if from.balance < amount {
		return ErrInsufficientFunds
	}

	if amount < 0 {
		return ErrInvalidAmount
	}

	from.balance -= amount
	to.balance += amount

	return nil
}
