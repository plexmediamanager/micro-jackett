package main

import (
    "github.com/micro/go-micro/client"
    "github.com/plexmediamanager/micro-jackett/jackett"
    "github.com/plexmediamanager/micro-jackett/proto"
    "github.com/plexmediamanager/micro-jackett/resolver"
    "github.com/plexmediamanager/service"
    "github.com/plexmediamanager/service/log"
    "time"
)

func main() {
    application := service.CreateApplication()

    err := application.InitializeConfiguration()
    if err != nil {
        log.Panic(err)
    }

    jackettClient := jackett.Initialize()
    err = jackettClient.LoadServerConfiguration()
    if err != nil {
        log.Panic(err)
    }
    err = jackettClient.LoadIndexersList()
    if err != nil {
        log.Panic(err)
    }

    err = application.InitializeMicroService()
    if err != nil {
       log.Panic(err)
    }

    err = application.Service().Client().Init(
       client.PoolSize(10),
       client.Retries(30),
       client.RequestTimeout(1 * time.Second),
    )
    if err != nil {
       log.Panic(err)
    }

    err = proto.RegisterJackettServiceHandler(application.Service().Server(), resolver.JackettService{ Jackett: jackettClient })
    if err != nil {
        log.Panic(err)
    }

    go application.StartMicroService()

    service.WaitForOSSignal(1)
}
