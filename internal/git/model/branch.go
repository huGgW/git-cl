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

type WorktreeBranch struct {
	Name_ string
	Path  string
}

func (w WorktreeBranch) Name() string {
	return w.Name_
}

func (w WorktreeBranch) IsCurrent() bool {
	// since we only going to allow exec program only in real repository, this should be always false.
	return false
}
