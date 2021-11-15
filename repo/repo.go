package repo

type repo struct{}

type RepoIndex interface {
}

func NewRepoIndex() RepoIndex {
	return &repo{}
}
