package scryfall

// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
    "github.com/Azure/go-autorest/autorest"
    "github.com/Azure/go-autorest/autorest/azure"
    "net/http"
    "context"
    "github.com/Azure/go-autorest/tracing"
    "github.com/satori/go.uuid"
)

// RulingsClient is the client for the Rulings methods of the Scryfall service.
type RulingsClient struct {
    BaseClient
}
// NewRulingsClient creates an instance of the RulingsClient client.
func NewRulingsClient() RulingsClient {
    return NewRulingsClientWithBaseURI(DefaultBaseURI, )
}

// NewRulingsClientWithBaseURI creates an instance of the RulingsClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
    func NewRulingsClientWithBaseURI(baseURI string, ) RulingsClient {
        return RulingsClient{ NewWithBaseURI(baseURI, )}
    }

// GetByCodeByNumberID sends the get by code by number id request.
func (client RulingsClient) GetByCodeByNumberID(ctx context.Context, code string, number int32) (result RulingList, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/RulingsClient.GetByCodeByNumberID")
        defer func() {
            sc := -1
        if result.Response.Response != nil {
        sc = result.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
    req, err := client.GetByCodeByNumberIDPreparer(ctx, code, number)
    if err != nil {
    err = autorest.NewErrorWithError(err, "scryfall.RulingsClient", "GetByCodeByNumberID", nil , "Failure preparing request")
    return
    }

        resp, err := client.GetByCodeByNumberIDSender(req)
        if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "scryfall.RulingsClient", "GetByCodeByNumberID", resp, "Failure sending request")
        return
        }

        result, err = client.GetByCodeByNumberIDResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "scryfall.RulingsClient", "GetByCodeByNumberID", resp, "Failure responding to request")
        }

    return
}

    // GetByCodeByNumberIDPreparer prepares the GetByCodeByNumberID request.
    func (client RulingsClient) GetByCodeByNumberIDPreparer(ctx context.Context, code string, number int32) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "code": autorest.Encode("path",code),
        "number": autorest.Encode("path",number),
        }

    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/cards/{code}/{number}/rulings",pathParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetByCodeByNumberIDSender sends the GetByCodeByNumberID request. The method will close the
    // http.Response Body if it receives an error.
    func (client RulingsClient) GetByCodeByNumberIDSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            }

    // GetByCodeByNumberIDResponder handles the response to the GetByCodeByNumberID request. The method always
    // closes the http.Response Body.
    func (client RulingsClient) GetByCodeByNumberIDResponder(resp *http.Response) (result RulingList, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

// GetByID sends the get by id request.
func (client RulingsClient) GetByID(ctx context.Context, ID uuid.UUID) (result RulingList, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/RulingsClient.GetByID")
        defer func() {
            sc := -1
        if result.Response.Response != nil {
        sc = result.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
    req, err := client.GetByIDPreparer(ctx, ID)
    if err != nil {
    err = autorest.NewErrorWithError(err, "scryfall.RulingsClient", "GetByID", nil , "Failure preparing request")
    return
    }

        resp, err := client.GetByIDSender(req)
        if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "scryfall.RulingsClient", "GetByID", resp, "Failure sending request")
        return
        }

        result, err = client.GetByIDResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "scryfall.RulingsClient", "GetByID", resp, "Failure responding to request")
        }

    return
}

    // GetByIDPreparer prepares the GetByID request.
    func (client RulingsClient) GetByIDPreparer(ctx context.Context, ID uuid.UUID) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "id": autorest.Encode("path",ID),
        }

    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/cards/{id}/rulings",pathParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetByIDSender sends the GetByID request. The method will close the
    // http.Response Body if it receives an error.
    func (client RulingsClient) GetByIDSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            }

    // GetByIDResponder handles the response to the GetByID request. The method always
    // closes the http.Response Body.
    func (client RulingsClient) GetByIDResponder(resp *http.Response) (result RulingList, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

// GetByMtgoID sends the get by mtgo id request.
func (client RulingsClient) GetByMtgoID(ctx context.Context, ID int32) (result RulingList, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/RulingsClient.GetByMtgoID")
        defer func() {
            sc := -1
        if result.Response.Response != nil {
        sc = result.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
    req, err := client.GetByMtgoIDPreparer(ctx, ID)
    if err != nil {
    err = autorest.NewErrorWithError(err, "scryfall.RulingsClient", "GetByMtgoID", nil , "Failure preparing request")
    return
    }

        resp, err := client.GetByMtgoIDSender(req)
        if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "scryfall.RulingsClient", "GetByMtgoID", resp, "Failure sending request")
        return
        }

        result, err = client.GetByMtgoIDResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "scryfall.RulingsClient", "GetByMtgoID", resp, "Failure responding to request")
        }

    return
}

    // GetByMtgoIDPreparer prepares the GetByMtgoID request.
    func (client RulingsClient) GetByMtgoIDPreparer(ctx context.Context, ID int32) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "id": autorest.Encode("path",ID),
        }

    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/cards/mtgo/{id}/rulings",pathParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetByMtgoIDSender sends the GetByMtgoID request. The method will close the
    // http.Response Body if it receives an error.
    func (client RulingsClient) GetByMtgoIDSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            }

    // GetByMtgoIDResponder handles the response to the GetByMtgoID request. The method always
    // closes the http.Response Body.
    func (client RulingsClient) GetByMtgoIDResponder(resp *http.Response) (result RulingList, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

// GetByMultiverseID sends the get by multiverse id request.
func (client RulingsClient) GetByMultiverseID(ctx context.Context, ID int32) (result RulingList, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/RulingsClient.GetByMultiverseID")
        defer func() {
            sc := -1
        if result.Response.Response != nil {
        sc = result.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
    req, err := client.GetByMultiverseIDPreparer(ctx, ID)
    if err != nil {
    err = autorest.NewErrorWithError(err, "scryfall.RulingsClient", "GetByMultiverseID", nil , "Failure preparing request")
    return
    }

        resp, err := client.GetByMultiverseIDSender(req)
        if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "scryfall.RulingsClient", "GetByMultiverseID", resp, "Failure sending request")
        return
        }

        result, err = client.GetByMultiverseIDResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "scryfall.RulingsClient", "GetByMultiverseID", resp, "Failure responding to request")
        }

    return
}

    // GetByMultiverseIDPreparer prepares the GetByMultiverseID request.
    func (client RulingsClient) GetByMultiverseIDPreparer(ctx context.Context, ID int32) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "id": autorest.Encode("path",ID),
        }

    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/cards/multiverse/{id}/rulings",pathParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetByMultiverseIDSender sends the GetByMultiverseID request. The method will close the
    // http.Response Body if it receives an error.
    func (client RulingsClient) GetByMultiverseIDSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            }

    // GetByMultiverseIDResponder handles the response to the GetByMultiverseID request. The method always
    // closes the http.Response Body.
    func (client RulingsClient) GetByMultiverseIDResponder(resp *http.Response) (result RulingList, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

