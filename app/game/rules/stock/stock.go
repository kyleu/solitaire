package stock

type Stock struct {
	Name                  string `json:"name,omitempty"`
	CardsShown            int    `json:"cardsShown,omitempty"`
	DealTo                DealTo `json:"dealTo,omitempty"`
	MaximumDeals          int    `json:"maximumDeals,omitempty"`
	CardsDealt            string `json:"cardsDealt,omitempty"`
	StopAfterPartialDeal  bool   `json:"stopAfterPartialDeal,omitempty"`
	CreatePocketWhenEmpty bool   `json:"createPocketWhenEmpty,omitempty"`
	GalleryMode           bool   `json:"galleryMode,omitempty"`
}
