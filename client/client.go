package client

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"encoding/json"

	"github.com/elebertus/auction-api/models"
)

/*
    This is pretty much a `utils` package. It is going to be refactored into
	`database` and `http` packages, or at least have the functionality logically
	split into packages. The janky chan bullshit is a workaround for just testing
	the flow of all of this from the api.

	General TODOS:
	  - Setting up a logger for nice non-filthy logs
	  - Metrics instrumentation
	  - Better error handling
	  - Establishing API response for error conditions
	  - Persistence
	  - Caching
*/
var client = &http.Client{Timeout: 10 * time.Second}
var fileChan = make(chan string, 1)

func toChan(a *models.AuctionData) {
	fileChan <- a.Files[0].Url
	fmt.Printf("[toChan] len(fileChan): %v\n", len(fileChan))
	close(fileChan)
	fmt.Println("Returning from toChan")
	return
}

func fromChan(ic <-chan string) {
	for s := range ic {
		fileChan <- s
	}
	close(fileChan)
	return
}

// TODO we never want the api to be able to control writing and
// reading of a file. There should be a worker process to grab
// fresh models.AuctionDataFile objects from the bnet api and
// persist them to a datastore.
func writeToDataStore() {
	return
}

// TODO this will replace the GetAuctionData > jsonFromFile flow.
// Will require overall refactoring.
func loadFromDataStore() {
	return
}
func loadFile() (*os.File, error) {
	f, err := os.Open("./auctions.json")
	if err != nil {
		fmt.Printf("[loadFile] Error opening file: %v\n", err)
		return nil, err

	}
	return f, nil
}

// JSONFromFile load json from a file in the format
// of models.GeneratedAuction
func jsonFromFile(r *http.Response, a *models.AuctionContents) (*models.AuctionContents, error) {

	out, err := os.Create("./auctions.json")
	if err != nil {
		fmt.Printf("[jsonFromFile] Error creating file: %v\n", err)
		return a, err
	}
	defer out.Close()

	_, err = io.Copy(out, r.Body)
	if err != nil {
		fmt.Printf("[jsonFromFile] Error copying r.Body to file: %v\n", err)
		return a, err
	}

	f, err := loadFile()
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
	}
	err = json.NewDecoder(f).Decode(a)
	fmt.Printf("[jsonFromFile] a.Auctions: %v\n", len(a.Auctions))

	defer os.Remove("./auction.json")
	return a, err
}

// TODO make the key, locale, and realm configurable
// GetAuctionFile get the auction file for realm
func GetAuctionFile(f *models.AuctionData) error {
	u := "https://us.api.battle.net/wow/auction/data/Hydraxis?locale=en_US&apikey=<somekey>"
	response, err := client.Get(u)
	if err != nil {
		fmt.Printf("Error requesting auction file: %v\n", err)
		return err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(f)
	toChan(f)
	return err
}

// TODO this functionality is mostly to test the idea.
// we will refactor this to get a specific thing, or make
// it a private function not published to the api.
// GetAuctionData get the AuctionFile and parse it
func GetAuctionData(f *models.AuctionContents) error {
	u := <-fileChan
	// TODO get the response and do things with it
	response, err := client.Get(u)
	if err != nil {
		fmt.Printf("Error getting auction file data: %v\n", err)
		return err
	}
	defer response.Body.Close()
	j, err := jsonFromFile(response, f)
	err = json.NewDecoder(response.Body).Decode(j)
	return err

}
