
package errors

import "github.com/plexmediamanager/service/errors"

const (
    ServiceID       errors.Service      =   4
)

var (
    // Network errors
    JackettUnableToCreateHTTPGetRequest     =   errors.Error {
        Code:               errors.Code {
            Service:        ServiceID,
            ErrorType:      errors.TypeNetwork,
            ErrorNumber:    1,
        },
        Message:            "Unable to create get request",
    }
    JackettUnableToExecuteHTTPGetRequest    =   errors.Error {
        Code:               errors.Code {
            Service:        ServiceID,
            ErrorType:      errors.TypeNetwork,
            ErrorNumber:    2,
        },
        Message:            "Unable to execute get request",
    }
    JackettUnableTooManyRedirects           =   errors.Error {
        Code:               errors.Code {
            Service:        ServiceID,
            ErrorType:      errors.TypeNetwork,
            ErrorNumber:    3,
        },
        Message:            "Stopped after %d redirects",
    }

    // Library errors
    JackettUnableToReadResponseBody         =   errors.Error {
        Code:               errors.Code {
            Service:        ServiceID,
            ErrorType:      errors.TypeLibrary,
            ErrorNumber:    1,
        },
        Message:            "Unable to read contents of the response body",
    }
    JackettUnmarshalError                   =   errors.Error {
        Code:               errors.Code {
            Service:        ServiceID,
            ErrorType:      errors.TypeLibrary,
            ErrorNumber:    2,
        },
        Message:            "Unable to read contents of the response body",
    }
    JackettUnableToFetchServerInfo          =   errors.Error {
        Code:               errors.Code {
            Service:        ServiceID,
            ErrorType:      errors.TypeLibrary,
            ErrorNumber:    3,
        },
        Message:            "Unable to fetch server information",
    }
    JackettUnableToLoadIndexers             =   errors.Error {
        Code:               errors.Code {
            Service:        ServiceID,
            ErrorType:      errors.TypeLibrary,
            ErrorNumber:    4,
        },
        Message:            "Unable to perform search",
    }
    JackettUnableToPerformSearch            =   errors.Error {
        Code:               errors.Code {
            Service:        ServiceID,
            ErrorType:      errors.TypeLibrary,
            ErrorNumber:    5,
        },
        Message:            "Unable to perform search",
    }
)