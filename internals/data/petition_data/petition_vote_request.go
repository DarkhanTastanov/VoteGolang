package petition_data

// PetitionVoteRequest represents the body to petition_data for a petition_data petition_data request.
type PetitionVoteRequest struct {
	UserId     uint `json:"user_id" swaggerignore:"true"`
	PetitionID uint `json:"petition_id"`
	// Enum values: favor, against
	VoteType VoteType `json:"vote_type" enum:"favor,against" example:"favor, against"`
}
