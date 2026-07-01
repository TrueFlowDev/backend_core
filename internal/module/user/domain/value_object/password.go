package value_object

type HashedPassword struct {
	value string
}

func NewHashedPassword(hash string) HashedPassword {
	return HashedPassword{value: hash}
}

func (p HashedPassword) Value() string { return p.value }
