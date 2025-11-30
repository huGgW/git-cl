package model

type LocalBranches struct {
	Branches []Branch
}

type Branch interface {
	Name() string
	IsCurrent() bool
}

type LocalBranch struct {
	Name_      string
	IsCurrent_ bool
}

func (l LocalBranch) Name() string {
	return l.Name_
}

func (l LocalBranch) IsCurrent() bool {
	return l.IsCurrent_
}
