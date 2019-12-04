package model

type Product struct {
	Model
	Code  string
	Price uint
}

type VoteOption struct {
	Model
	Text   string
	Count  int
	VoteID uint `sql:"index"`
}

type Vote struct {
	Model
	Title   string
	Creator int
}

type VoteFull struct {
	Title   string
	Options []VoteOption
}

func GetVoteFull(id uint) (VoteFull, error) {
	var options []VoteOption
	var vote Vote

	if err := DB.First(&vote, id).Error; err != nil {
		return VoteFull{}, err
	}

	DB.Where(&VoteOption{VoteID: uint(id)}).Find(&options)

	return VoteFull{Title: vote.Title, Options: options}, nil
}
