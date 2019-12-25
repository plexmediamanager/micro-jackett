package jackett

import (
    "github.com/plexmediamanager/micro-jackett/errors"
)

type SearchTorrent struct {
    FirstSeen            string             `json:"FirstSeen"`
    Tracker              string             `json:"Tracker"`
    TrackerID            string             `json:"TrackerId"`
    CategoryDesc         string             `json:"CategoryDesc"`
    BlackHoleLink        interface{}        `json:"BlackholeLink"`
    Title                string             `json:"Title"`
    GUID                 string             `json:"Guid"`
    Link                 string             `json:"Link"`
    Comments             string             `json:"Comments"`
    PublishDate          string             `json:"PublishDate"`
    Category             []int64            `json:"Category"`
    Size                 int64              `json:"Size"`
    Files                interface{}        `json:"Files"`
    Grabs                int64              `json:"Grabs"`
    Description          interface{}        `json:"Description"`
    RageID               interface{}        `json:"RageID"`
    TVDBID               interface{}        `json:"TVDBId"`
    Imdb                 interface{}        `json:"Imdb"`
    TMDb                 interface{}        `json:"TMDb"`
    Seeders              int64              `json:"Seeders"`
    Peers                int64              `json:"Peers"`
    BannerURL            interface{}        `json:"BannerUrl"`
    InfoHash             interface{}        `json:"InfoHash"`
    MagnetURI            interface{}        `json:"MagnetUri"`
    MinimumRatio         float64            `json:"MinimumRatio"`
    MinimumSeedTime      int64              `json:"MinimumSeedTime"`
    DownloadVolumeFactor float64            `json:"DownloadVolumeFactor"`
    UploadVolumeFactor   float64            `json:"UploadVolumeFactor"`
    Gain                 float64            `json:"Gain"`
}

type SearchIndexer struct {
    ID                  string              `json:"ID"`
    Name                string              `json:"Name"`
    Status              uint64              `json:"Status"`
    Results             uint64              `json:"Results"`
    Error               interface{}         `json:"Error"`
}

type SearchResponse struct {
    Results             []SearchTorrent     `json:"Results"`
    Indexers            []SearchIndexer     `json:"Indexers"`
}

// Perform search by query value on specified trackers, with specified categories
func (client *Client) Search(query string, trackers []string, categories []string) (*SearchResponse, error) {
    var result SearchResponse
    newTrackers := make([]string, 0)
    if trackers == nil {
        for _, tracker := range client.configuredIndexers {
            newTrackers = append(newTrackers, tracker.ID)
        }
    } else {
        newTrackers = trackers
    }
    response, err := client.sendGetRequest(&result, "indexers/all/results", query, newTrackers, categories)
    if err != nil {
        return nil, errors.JackettUnableToPerformSearch.ToError(err)
    }
    return response.(*SearchResponse), nil
}