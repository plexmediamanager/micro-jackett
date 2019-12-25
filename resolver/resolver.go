package resolver

import (
    "context"
    "encoding/json"
    "github.com/plexmediamanager/micro-jackett/jackett"
    "github.com/plexmediamanager/micro-jackett/proto"
    "github.com/plexmediamanager/micro-torrent/errors"
)

type JackettService struct {
    Jackett           *jackett.Client
}

// Convert structure to bytes and return error if needed
func structureToBytesWithError(structure interface{}, err error, response *proto.JackettResponse) error {
    if err != nil {
        return err
    }
    bytes, err := json.Marshal(structure)
    if err != nil {
        return errors.UnmarshalError.ToError(err)
    }
    response.Result = bytes
    return nil
}

func (service JackettService) GetServerConfiguration (_ context.Context, parameters *proto.JackettEmpty, response *proto.JackettResponse) error {
    result := service.Jackett.GetServerConfiguration()
    return structureToBytesWithError(result, nil, response)
}

func (service JackettService) GetConfiguredIndexers (_ context.Context, parameters *proto.JackettEmpty, response *proto.JackettResponse) error {
    result := service.Jackett.GetConfiguredIndexers()
    return structureToBytesWithError(result, nil, response)
}

func (service JackettService) Search (_ context.Context, parameters *proto.JackettSearch, response *proto.JackettResponse) error {
    result, err := service.Jackett.Search(parameters.Query, parameters.Trackers, parameters.Categories)
    return structureToBytesWithError(result, err, response)
}
