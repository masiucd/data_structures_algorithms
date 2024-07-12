package gender

type Gender int

const (
	Male Gender = iota
	Female
	Other
)

func (g Gender) String() string {
	return [...]string{"Male", "Female", "Other"}[g]
}
