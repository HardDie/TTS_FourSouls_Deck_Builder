package cards

import (
	"tts_deck_build/internal/config"
	"tts_deck_build/internal/decks"
	"tts_deck_build/internal/utils"
)

type CardService struct {
	storage *CardStorage
}

func NewService() *CardService {
	return &CardService{
		storage: NewCardStorage(config.GetConfig(), decks.NewService()),
	}
}

func (s *CardService) Create(gameID, collectionID, deckID string, dto *CreateCardDTO) (*CardInfo, error) {
	card := NewCardInfo(dto.Title, dto.Description, dto.Image, dto.Variables)

	// Get all cards in deck
	allCards, err := s.List(gameID, collectionID, deckID, "")
	if err != nil {
		return nil, err
	}

	// Find current biggest index
	var maxID int64
	for _, currentCard := range allCards {
		if currentCard.ID > maxID {
			maxID = currentCard.ID
		}
	}

	// Increase value
	maxID += 1
	// Set new max value to the new card
	card.ID = maxID

	return s.storage.Create(gameID, collectionID, deckID, card)
}

func (s *CardService) Item(gameID, collectionID, deckID string, cardID int64) (*CardInfo, error) {
	return s.storage.GetByID(gameID, collectionID, deckID, cardID)
}

func (s *CardService) List(gameID, collectionID, deckID, sortField string) ([]*CardInfo, error) {
	items, err := s.storage.GetAll(gameID, collectionID, deckID)
	if err != nil {
		return make([]*CardInfo, 0), err
	}
	utils.Sort(&items, sortField)
	return items, nil
}

func (s *CardService) Update(gameID, collectionID, deckID string, cardID int64, dto *UpdateCardDTO) (*CardInfo, error) {
	return s.storage.Update(gameID, collectionID, deckID, cardID, dto)
}

func (s *CardService) Delete(gameID, collectionID, deckID string, cardID int64) error {
	return s.storage.DeleteByID(gameID, collectionID, deckID, cardID)
}
