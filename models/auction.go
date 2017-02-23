package models

// TODO refactor and remove this, it's dumb
type AuctionData struct {
	Files AuctionDataFiles `json:"files"`
}
type AuctionDataFile struct {
	Url          string `json:"url"`
	LastModified int    `json:"lastModified"`
}

type AuctionDataFiles []AuctionDataFile

type AuctionContents struct {
	Auctions []struct {
		Auc        int64 `json:"auc"`
		Bid        int64 `json:"bid"`
		BonusLists []struct {
			BonusListID int64 `json:"bonusListId"`
		} `json:"bonusLists"`
		Buyout    int64 `json:"buyout"`
		Context   int64 `json:"context"`
		Item      int64 `json:"item"`
		Modifiers []struct {
			Type  int64 `json:"type"`
			Value int64 `json:"value"`
		} `json:"modifiers"`
		Owner        string `json:"owner"`
		OwnerRealm   string `json:"ownerRealm"`
		PetBreedID   int64  `json:"petBreedId"`
		PetLevel     int64  `json:"petLevel"`
		PetQualityID int64  `json:"petQualityId"`
		PetSpeciesID int64  `json:"petSpeciesId"`
		Quantity     int64  `json:"quantity"`
		Rand         int64  `json:"rand"`
		Seed         int64  `json:"seed"`
		TimeLeft     string `json:"timeLeft"`
	} `json:"auctions"`
	Realms []struct {
		Name string `json:"name"`
		Slug string `json:"slug"`
	} `json:"realms"`
}
