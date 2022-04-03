package main

import (
	"encoding/json"
	"fmt"
	"sort"
)

type deckBuilderDeck struct {
	cards []*Card

	deckType     string
	backSideName string
	backSideURL  string
}
type DeckBuilder struct {
	decks map[string]*deckBuilderDeck
}

func NewDeckBuilder() *DeckBuilder {
	return &DeckBuilder{
		decks: make(map[string]*deckBuilderDeck),
	}
}

// collect
func (b *DeckBuilder) AddCard(deck *Deck, card *Card) {
	if _, ok := b.decks[deck.GetType()]; !ok {
		b.decks[deck.GetType()] = &deckBuilderDeck{
			deckType:     deck.GetType(),
			backSideName: deck.GetBackSideName(),
			backSideURL:  deck.GetBackSideURL(),
		}
	}
	b.decks[deck.GetType()].cards = append(b.decks[deck.GetType()].cards, card)
}
func (b *DeckBuilder) splitCards(deckType string) (cards [][]*Card) {
	for leftBorder := 0; leftBorder < len(b.decks[deckType].cards); leftBorder += MaxCardsOnPage {
		// Calculate right border for current deck
		rightBorder := min(len(b.decks[deckType].cards), leftBorder+MaxCardsOnPage)
		cards = append(cards, b.decks[deckType].cards[leftBorder:rightBorder])
	}
	return
}
func (b *DeckBuilder) getImageSize(count int) (cols, rows int) {
	cols = 10
	rows = 7
	images := cols * rows
	for r := 2; r <= 7; r++ {
		for c := 2; c <= 10; c++ {
			possible := c * r
			if possible < images && possible >= count {
				images = possible
				cols = c
				rows = r
			}
		}
	}
	return
}
func (b *DeckBuilder) GetDecks(deckType string) (decks []*Deck) {
	for index, cards := range b.splitCards(deckType) {
		// Calculate optimal count of columns and rows for result image
		columns, rows := b.getImageSize(len(cards) + 1)
		decks = append(decks, &Deck{
			Cards:   cards,
			Columns: columns,
			Rows:    rows,
			FileName: fmt.Sprintf("%s_%d_%d_%dx%d.png", cleanTitle(b.decks[deckType].deckType), index+1, len(cards),
				columns, rows),
			BackSide: &b.decks[deckType].backSideURL,
			Type:     deckType,
		})
	}
	return
}
func (b *DeckBuilder) GetTypes() (types []string) {
	for deckType := range b.decks {
		types = append(types, deckType)
	}
	sort.SliceStable(types, func(i, j int) bool {
		return types[i] < types[j]
	})
	return
}

// draw
func (b *DeckBuilder) DrawDecks() map[string]string {
	// List of result files
	res := make(map[string]string)
	for _, deckType := range b.GetTypes() {
		decks := b.GetDecks(deckType)
		for _, deck := range decks {
			NewDeckDrawer(deck).Draw()
			// Add current deck title
			res[deck.FileName] = ""
		}
		// Add back side image title
		res[decks[0].GetBackSideName()] = ""
	}
	return res
}

// tts
func (b *DeckBuilder) GenerateTTSDeck() []byte {
	res := TTSSaveObject{}
	for _, deckType := range b.GetTypes() {
		tts := NewTTSBuilder()
		decks := b.GetDecks(deckType)
		for deckId, deck := range decks {
			for j, card := range deck.Cards {
				cardId := (deckId+1)*100 + j
				tts.AddCard(deck, card, deckId+1, cardId)
			}
		}
		res.ObjectStates = append(res.ObjectStates, tts.GetObjects()...)
	}
	data, _ := json.MarshalIndent(res, "", "  ")
	return data
}
