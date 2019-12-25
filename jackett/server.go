package jackett

import "github.com/plexmediamanager/micro-jackett/errors"

type ServerConfiguration struct {
    Notices                 []interface{} `json:"notices"`
    Port                    int64         `json:"port"`
    External                bool          `json:"external"`
    APIKey                  string        `json:"api_key"`
    BlackHoleDirectory      string        `json:"blackholedir"`
    UpdateDisabled          bool          `json:"updatedisabled"`
    Prerelease              bool          `json:"prerelease"`
    Password                string        `json:"password"`
    Logging                 bool          `json:"logging"`
    BasePathOverride        string        `json:"basepathoverride"`
    OMDBKey                 string        `json:"omdbkey"`
    OMDBUrl                 string        `json:"omdburl"`
    AppVersion              string        `json:"app_version"`
    CanRunNetCore           bool          `json:"can_run_netcore"`
    ProxyType               int64         `json:"proxy_type"`
    ProxyURL                string        `json:"proxy_url"`
    ProxyPort               int64         `json:"proxy_port"`
    ProxyUsername           string        `json:"proxy_username"`
    ProxyPassword           string        `json:"proxy_password"`
}

// Load Jackett server configuration
func (client *Client) LoadServerConfiguration() error {
    var result ServerConfiguration
    response, err := client.sendGetRequest(&result, "server/config", "", nil, nil)
    if err != nil {
        return errors.JackettUnableToFetchServerInfo.ToError(err)
    }
    client.serverConfiguration = response.(*ServerConfiguration)
    return nil
}

// Get Jackett server configuration
func (client *Client) GetServerConfiguration() *ServerConfiguration {
    return client.serverConfiguration
}