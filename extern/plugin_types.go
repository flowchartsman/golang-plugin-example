package extern

type NamePlugin func(Person) string
type ActionPlugin func(Person) (string, error)
