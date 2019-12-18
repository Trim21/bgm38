package model

//VoteOption table for vote options
type VoteOption struct {
	base
	Text   string
	Count  int
	VoteID uint `sql:"index"`
}

//Vote table for vote
type Vote struct {
	base
	Title   string
	Creator int
}

//VoteFull a model with vote title and all options
type VoteFull struct {
	Title   string
	Options []VoteOption
}

//GetVoteFull get a VoteFull from id
func GetVoteFull(id uint) (VoteFull, error) {
	var options []VoteOption
	var vote Vote
	if err := DB.First(&vote, "ID = ?", id).Error; err != nil {
		return VoteFull{}, err
	}

	DB.Where(&VoteOption{VoteID: uint(id)}).Find(&options)

	return VoteFull{Title: vote.Title, Options: options}, nil
}
