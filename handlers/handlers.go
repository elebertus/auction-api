package handlers

import (
	"fmt"
	"net/http"

	"encoding/json"

	"github.com/elebertus/auction-api/client"
	"github.com/elebertus/auction-api/models"
	"github.com/gorilla/mux"
)

// Index will be deprecated
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

// GetAuctionDataFile gets the data files for realm
func GetAuctionDataFile(w http.ResponseWriter, r *http.Request) {

	// TODO We need to inspect a.Files[i].LastModified to see if it has
	// changed, else we retrieve a fresh file.
	auctionData := &models.AuctionData{}
	err := client.GetAuctionFile(auctionData)
	if err != nil {
		err = json.NewEncoder(w).Encode(fmt.Errorf("Error getting auction file url: %v\n", err))
	}
	// TODO We should determine if multiple models.AuctionDataFile
	// are ever returned and if we need the models.AuctionDataFiles
	// struct at all.

	// TODO need to setup a error handling function, or handle errors
	// for returns.
	json.NewEncoder(w).Encode(auctionData.Files[0].Url)
}

// GetAuctionData uses GetAuctionDataFiles to retrieve
// models.AuctionData.Files
func GetAuctionData(w http.ResponseWriter, r *http.Request) {
	auctionData := &models.AuctionData{}
	//json.NewEncoder(w).Encode(auctionData.Files[0].Url)
	if len(auctionData.Files) == 0 {
		fmt.Printf("[GetAuctionData] len(auctionData.Files) == %v\n", len(auctionData.Files))
		err := client.GetAuctionFile(auctionData)
		if err != nil {
			err = json.NewEncoder(w).Encode(fmt.Errorf("Error getting auction file url: %v\n", err))
		}
	}
	as := &models.AuctionContents{}
	client.GetAuctionData(as)
	json.NewEncoder(w).Encode(as)
}

// ShowAuctionField shows a field in models.AuctionFileAuction
func ShowAuctionField(w http.ResponseWriter, r *http.Request) {
	//TODO: maybe return specific things?
	vars := mux.Vars(r)
	auctionField := vars["auctionField"]
	fmt.Fprintln(w, "auctionField:", auctionField)
}
