package jackett

import (
    "github.com/plexmediamanager/micro-jackett/errors"
)

type IndexerCaps struct {
    ID                      string          `json:"id"`
    Name                    string          `json:"name"`
}

type Indexer struct {
    ID                      string          `json:"id"`
    Name                    string          `json:"name"`
    Description             string          `json:"description"`
    Type                    string          `json:"type"`
    Configured              bool            `json:"configured"`
    SiteLink                string          `json:"site_link"`
    AlternativeSiteLinks    []string        `json:"alternativesitelinks"`
    Language                string          `json:"language"`
    LastError               string          `json:"last_error"`
    PotatoEnabled           bool            `json:"potatoenabled"`
    Categories              []IndexerCaps   `json:"caps"`
}

type Indexers []Indexer

// Load Jackett indexers list
func (client *Client) LoadIndexersList() error {
    var result Indexers
    response, err := client.sendGetRequest(&result, "indexers", "", nil, nil)
    if err != nil {
        return errors.JackettUnableToLoadIndexers.ToError(err)
    }
    client.indexers = response.(*Indexers)
    configuredIndexers := make([]*Indexer, 0)
    for _, indexer := range result {
        if indexer.Configured {
            configuredIndexers = append(configuredIndexers, &indexer)
        }
    }
    client.configuredIndexers = configuredIndexers
    return nil
}

// Get list of configured Jackett indexers
func (client *Client) GetConfiguredIndexers() []*Indexer {
    return client.configuredIndexers
}