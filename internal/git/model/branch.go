package model

type LocalBranches struct {
	Branches []Branch
}

type Branch struct {
	Name      string
	IsCurrent bool
}
