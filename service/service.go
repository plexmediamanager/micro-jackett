package service

import (
    "context"
    "encoding/json"
    microClient "github.com/micro/go-micro/client"
    "github.com/plexmediamanager/micro-jackett/jackett"
    "github.com/plexmediamanager/micro-jackett/proto"
    "github.com/plexmediamanager/micro-torrent/errors"
    "github.com/plexmediamanager/service"
)

// Convert response to structure
func protoToStructure(output interface{}, result *proto.JackettResponse, err error) error {
    if err != nil {
        return err
    }
    err = json.Unmarshal(result.Result, &output)
    if err != nil {
        return errors.UnmarshalError.ToError(err)
    }
    return nil
}

func GetJackettService(client microClient.Client) proto.JackettService {
    return proto.NewJackettService(service.GetServiceName(service.JackettServiceName), client)
}

func GetServerConfiguration(client microClient.Client) (*jackett.ServerConfiguration, error) {
    var result *jackett.ServerConfiguration
    service := GetJackettService(client)
    parameters := &proto.JackettEmpty {}
    response, err := service.GetServerConfiguration(context.TODO(), parameters)
    return result, protoToStructure(&result, response, err)
}

func GetConfiguredIndexers(client microClient.Client) ([]jackett.Indexer, error) {
    var result []jackett.Indexer
    service := GetJackettService(client)
    parameters := &proto.JackettEmpty {}
    response, err := service.GetConfiguredIndexers(context.TODO(), parameters)
    return result, protoToStructure(&result, response, err)
}

func Search(client microClient.Client, query string, trackers []string, categories []string) (*jackett.SearchResponse, error) {
    var result *jackett.SearchResponse
    service := GetJackettService(client)
    parameters := &proto.JackettSearch { Query: query, Trackers: trackers, Categories: categories }
    response, err := service.Search(context.TODO(), parameters)
    return result, protoToStructure(&result, response, err)
}