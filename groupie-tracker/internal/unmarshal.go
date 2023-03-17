package internal

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)


func GetArtists() ([]Artist, error) {
	var result DateLocation
	var netClient = http.Client{
        Timeout: time.Second * 10,
    }
	relationsData, err := netClient.Get(ApiRelation)
	if err != nil {
		return []Artist{}, err
	}

	defer relationsData.Body.Close()
	relationsDataBody, err := io.ReadAll(relationsData.Body)
	if err != nil {
		return []Artist{}, err
	}
	jsonErr := json.Unmarshal(relationsDataBody, &result)
	if jsonErr != nil {
		return []Artist{}, jsonErr
	}
	var netClient2 = http.Client{
        Timeout: time.Second * 10,
    }
	artistsData, err := netClient2.Get(ApiArtists)
	if err != nil {
		return []Artist{}, err
	}

	var ArtistArr []Artist

	defer artistsData.Body.Close()

	dataArtistsBody, err := io.ReadAll(artistsData.Body)
	if err != nil {
		return []Artist{}, err
	}

	json.Unmarshal(dataArtistsBody, &ArtistArr)
	var newArtist []Artist
	for i, v := range ArtistArr {
		v.Relations = result.Index[i].DatesLocations
		newArtist = append(newArtist, v)
	}

	return newArtist, nil
}

func GetOneArtist(id string) (Artist, error) {
	var Singer Artist
	var netClient = http.Client{
        Timeout: time.Second * 10,
    }
	data, err := netClient.Get(ApiArtists + "/" + id)
	if err != nil {
		return Artist{}, err
	}

	defer data.Body.Close()
	

	dataBody, err := io.ReadAll(data.Body)
	if err != nil {
		return Artist{}, err
	}
	json.Unmarshal(dataBody, &Singer)

	return Singer, nil
}

func GetArtistLocation(id string) (Location, error) {
	var Relation Location
	var netClient = http.Client{
        Timeout: time.Second * 10,
    }
	data, err := netClient.Get(ApiRelation + "/" + id)
	if err != nil {
		return Location{}, err
	}
	defer data.Body.Close()

	dataBody, err := io.ReadAll(data.Body)
	if err != nil {
		return Location{}, err
	}
	jsonErr := json.Unmarshal(dataBody, &Relation)
	if jsonErr != nil {
		return Location{}, jsonErr
	}
	return Relation, nil
}
