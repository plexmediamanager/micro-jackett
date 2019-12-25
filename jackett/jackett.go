package jackett

import (
    "encoding/json"
    format "fmt"
    "github.com/plexmediamanager/micro-jackett/errors"
    "github.com/plexmediamanager/service/helpers"
    "io/ioutil"
    "net/http"
    "net/http/cookiejar"
    "net/url"
    "time"
)

const (
    MaxRedirects        =   10
)

type Client struct {
    host                    string
    version                 string
    apiKey                  string

    client                  *http.Client
    serverConfiguration     *ServerConfiguration
    indexers                *Indexers
    configuredIndexers      []*Indexer
}

// Initialize Jackett Client
func Initialize() *Client {
    client := &Client {
        host:           helpers.GetEnvironmentVariableAsString("JACKETT_HOST", ""),
        version:        helpers.GetEnvironmentVariableAsString("JACKETT_API_VERSION", ""),
        apiKey:         helpers.GetEnvironmentVariableAsString("JACKETT_API_KEY", ""),
        client:         &http.Client{
            CheckRedirect: func() func(req *http.Request, via []*http.Request) error {
                redirects := 0
                return func(req *http.Request, via []*http.Request) error {
                    if redirects > MaxRedirects {
                        return errors.JackettUnableTooManyRedirects.ToErrorWithArguments(nil, MaxRedirects)
                    }
                    redirects++
                    return nil
                }
            }(),
        },
    }
    client.client.Jar, _ = cookiejar.New(nil)
    return client
}

// Build request URL
func (client *Client) buildRequestURL(path string, query string, trackers []string, categories []string) string {
    requestOptions := url.Values{}
    requestURL := format.Sprintf(
        "%s/api/v%s/%s?apikey=%s&_=%d",
        client.host,
        client.version,
        path,
        client.apiKey,
        client.generateRequestTimestamp(),
    )

    if len(query) > 0 {
        requestOptions.Add("Query", query)
    }

    if len(trackers) > 0 {
        for _, tracker := range trackers {
            requestOptions.Add("Tracker[]", tracker)
        }
    }

    if len(categories) > 0 {
        for _, category := range categories {
            requestOptions.Add("Category[]", category)
        }
    }

    encodedParameters := requestOptions.Encode()
    if len(encodedParameters) > 0 {
        requestURL += "&" + encodedParameters
    }
    return requestURL
}

// Generate request timestamp suitable for Jackett
func (client *Client) generateRequestTimestamp() int64 {
    return time.Now().UnixNano() / 1000000
}

// Send GET request to Jackett server
func (client *Client) sendGetRequest(unmarshal interface{}, path string, query string, trackers []string, categories []string) (interface{}, error) {
    request, err := http.NewRequest("GET", client.buildRequestURL(path, query, trackers, categories), nil)
    if err != nil {
        return nil, errors.JackettUnableToCreateHTTPGetRequest.ToError(err)
    }
    request.Header.Set("Accepts", "application/json")
    request.Header.Set("Content-Type", "application/json")
    request.Header.Set("User-Agent", "FreedomCore Micro Jackett Client")


    response, err := client.client.Do(request)
    if err != nil {
        return nil, errors.JackettUnableToExecuteHTTPGetRequest.ToError(err)
    }

    responseBody, err := ioutil.ReadAll(response.Body)
    if err != nil {
        return nil, errors.JackettUnableToReadResponseBody.ToError(err)
    }
    defer response.Body.Close()

    err = json.Unmarshal(responseBody, &unmarshal)
    if err != nil {
        return nil, errors.JackettUnmarshalError.ToError(err)
    }
    return unmarshal, nil
}