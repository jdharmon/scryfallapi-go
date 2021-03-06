package scryfall

// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
    "github.com/Azure/go-autorest/autorest"
    "github.com/Azure/go-autorest/autorest/azure"
    "net/http"
    "context"
    "github.com/Azure/go-autorest/tracing"
)

// CatalogClient is the client for the Catalog methods of the Scryfall service.
type CatalogClient struct {
    BaseClient
}
// NewCatalogClient creates an instance of the CatalogClient client.
func NewCatalogClient() CatalogClient {
    return NewCatalogClientWithBaseURI(DefaultBaseURI, )
}

// NewCatalogClientWithBaseURI creates an instance of the CatalogClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
    func NewCatalogClientWithBaseURI(baseURI string, ) CatalogClient {
        return CatalogClient{ NewWithBaseURI(baseURI, )}
    }

// GetArtifactTypes sends the get artifact types request.
func (client CatalogClient) GetArtifactTypes(ctx context.Context) (result Catalog, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/CatalogClient.GetArtifactTypes")
        defer func() {
            sc := -1
        if result.Response.Response != nil {
        sc = result.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
    req, err := client.GetArtifactTypesPreparer(ctx)
    if err != nil {
    err = autorest.NewErrorWithError(err, "scryfall.CatalogClient", "GetArtifactTypes", nil , "Failure preparing request")
    return
    }

        resp, err := client.GetArtifactTypesSender(req)
        if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "scryfall.CatalogClient", "GetArtifactTypes", resp, "Failure sending request")
        return
        }

        result, err = client.GetArtifactTypesResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "scryfall.CatalogClient", "GetArtifactTypes", resp, "Failure responding to request")
        }

    return
}

    // GetArtifactTypesPreparer prepares the GetArtifactTypes request.
    func (client CatalogClient) GetArtifactTypesPreparer(ctx context.Context) (*http.Request, error) {
    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPath("/catalog/artifact-types"))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetArtifactTypesSender sends the GetArtifactTypes request. The method will close the
    // http.Response Body if it receives an error.
    func (client CatalogClient) GetArtifactTypesSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            }

    // GetArtifactTypesResponder handles the response to the GetArtifactTypes request. The method always
    // closes the http.Response Body.
    func (client CatalogClient) GetArtifactTypesResponder(resp *http.Response) (result Catalog, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

// GetCardNames sends the get card names request.
func (client CatalogClient) GetCardNames(ctx context.Context) (result Catalog, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/CatalogClient.GetCardNames")
        defer func() {
            sc := -1
        if result.Response.Response != nil {
        sc = result.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
    req, err := client.GetCardNamesPreparer(ctx)
    if err != nil {
    err = autorest.NewErrorWithError(err, "scryfall.CatalogClient", "GetCardNames", nil , "Failure preparing request")
    return
    }

        resp, err := client.GetCardNamesSender(req)
        if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "scryfall.CatalogClient", "GetCardNames", resp, "Failure sending request")
        return
        }

        result, err = client.GetCardNamesResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "scryfall.CatalogClient", "GetCardNames", resp, "Failure responding to request")
        }

    return
}

    // GetCardNamesPreparer prepares the GetCardNames request.
    func (client CatalogClient) GetCardNamesPreparer(ctx context.Context) (*http.Request, error) {
    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPath("/catalog/card-names"))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetCardNamesSender sends the GetCardNames request. The method will close the
    // http.Response Body if it receives an error.
    func (client CatalogClient) GetCardNamesSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            }

    // GetCardNamesResponder handles the response to the GetCardNames request. The method always
    // closes the http.Response Body.
    func (client CatalogClient) GetCardNamesResponder(resp *http.Response) (result Catalog, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

// GetCreatureTypes sends the get creature types request.
func (client CatalogClient) GetCreatureTypes(ctx context.Context) (result Catalog, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/CatalogClient.GetCreatureTypes")
        defer func() {
            sc := -1
        if result.Response.Response != nil {
        sc = result.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
    req, err := client.GetCreatureTypesPreparer(ctx)
    if err != nil {
    err = autorest.NewErrorWithError(err, "scryfall.CatalogClient", "GetCreatureTypes", nil , "Failure preparing request")
    return
    }

        resp, err := client.GetCreatureTypesSender(req)
        if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "scryfall.CatalogClient", "GetCreatureTypes", resp, "Failure sending request")
        return
        }

        result, err = client.GetCreatureTypesResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "scryfall.CatalogClient", "GetCreatureTypes", resp, "Failure responding to request")
        }

    return
}

    // GetCreatureTypesPreparer prepares the GetCreatureTypes request.
    func (client CatalogClient) GetCreatureTypesPreparer(ctx context.Context) (*http.Request, error) {
    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPath("/catalog/creature-types"))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetCreatureTypesSender sends the GetCreatureTypes request. The method will close the
    // http.Response Body if it receives an error.
    func (client CatalogClient) GetCreatureTypesSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            }

    // GetCreatureTypesResponder handles the response to the GetCreatureTypes request. The method always
    // closes the http.Response Body.
    func (client CatalogClient) GetCreatureTypesResponder(resp *http.Response) (result Catalog, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

// GetEnchantmentTypes sends the get enchantment types request.
func (client CatalogClient) GetEnchantmentTypes(ctx context.Context) (result Catalog, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/CatalogClient.GetEnchantmentTypes")
        defer func() {
            sc := -1
        if result.Response.Response != nil {
        sc = result.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
    req, err := client.GetEnchantmentTypesPreparer(ctx)
    if err != nil {
    err = autorest.NewErrorWithError(err, "scryfall.CatalogClient", "GetEnchantmentTypes", nil , "Failure preparing request")
    return
    }

        resp, err := client.GetEnchantmentTypesSender(req)
        if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "scryfall.CatalogClient", "GetEnchantmentTypes", resp, "Failure sending request")
        return
        }

        result, err = client.GetEnchantmentTypesResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "scryfall.CatalogClient", "GetEnchantmentTypes", resp, "Failure responding to request")
        }

    return
}

    // GetEnchantmentTypesPreparer prepares the GetEnchantmentTypes request.
    func (client CatalogClient) GetEnchantmentTypesPreparer(ctx context.Context) (*http.Request, error) {
    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPath("/catalog/enchantment-types"))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetEnchantmentTypesSender sends the GetEnchantmentTypes request. The method will close the
    // http.Response Body if it receives an error.
    func (client CatalogClient) GetEnchantmentTypesSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            }

    // GetEnchantmentTypesResponder handles the response to the GetEnchantmentTypes request. The method always
    // closes the http.Response Body.
    func (client CatalogClient) GetEnchantmentTypesResponder(resp *http.Response) (result Catalog, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

// GetLandTypes sends the get land types request.
func (client CatalogClient) GetLandTypes(ctx context.Context) (result Catalog, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/CatalogClient.GetLandTypes")
        defer func() {
            sc := -1
        if result.Response.Response != nil {
        sc = result.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
    req, err := client.GetLandTypesPreparer(ctx)
    if err != nil {
    err = autorest.NewErrorWithError(err, "scryfall.CatalogClient", "GetLandTypes", nil , "Failure preparing request")
    return
    }

        resp, err := client.GetLandTypesSender(req)
        if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "scryfall.CatalogClient", "GetLandTypes", resp, "Failure sending request")
        return
        }

        result, err = client.GetLandTypesResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "scryfall.CatalogClient", "GetLandTypes", resp, "Failure responding to request")
        }

    return
}

    // GetLandTypesPreparer prepares the GetLandTypes request.
    func (client CatalogClient) GetLandTypesPreparer(ctx context.Context) (*http.Request, error) {
    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPath("/catalog/land-types"))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetLandTypesSender sends the GetLandTypes request. The method will close the
    // http.Response Body if it receives an error.
    func (client CatalogClient) GetLandTypesSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            }

    // GetLandTypesResponder handles the response to the GetLandTypes request. The method always
    // closes the http.Response Body.
    func (client CatalogClient) GetLandTypesResponder(resp *http.Response) (result Catalog, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

// GetLoyalties sends the get loyalties request.
func (client CatalogClient) GetLoyalties(ctx context.Context) (result Catalog, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/CatalogClient.GetLoyalties")
        defer func() {
            sc := -1
        if result.Response.Response != nil {
        sc = result.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
    req, err := client.GetLoyaltiesPreparer(ctx)
    if err != nil {
    err = autorest.NewErrorWithError(err, "scryfall.CatalogClient", "GetLoyalties", nil , "Failure preparing request")
    return
    }

        resp, err := client.GetLoyaltiesSender(req)
        if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "scryfall.CatalogClient", "GetLoyalties", resp, "Failure sending request")
        return
        }

        result, err = client.GetLoyaltiesResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "scryfall.CatalogClient", "GetLoyalties", resp, "Failure responding to request")
        }

    return
}

    // GetLoyaltiesPreparer prepares the GetLoyalties request.
    func (client CatalogClient) GetLoyaltiesPreparer(ctx context.Context) (*http.Request, error) {
    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPath("/catalog/loyalties"))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetLoyaltiesSender sends the GetLoyalties request. The method will close the
    // http.Response Body if it receives an error.
    func (client CatalogClient) GetLoyaltiesSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            }

    // GetLoyaltiesResponder handles the response to the GetLoyalties request. The method always
    // closes the http.Response Body.
    func (client CatalogClient) GetLoyaltiesResponder(resp *http.Response) (result Catalog, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

// GetPlaneswalkerTypes sends the get planeswalker types request.
func (client CatalogClient) GetPlaneswalkerTypes(ctx context.Context) (result Catalog, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/CatalogClient.GetPlaneswalkerTypes")
        defer func() {
            sc := -1
        if result.Response.Response != nil {
        sc = result.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
    req, err := client.GetPlaneswalkerTypesPreparer(ctx)
    if err != nil {
    err = autorest.NewErrorWithError(err, "scryfall.CatalogClient", "GetPlaneswalkerTypes", nil , "Failure preparing request")
    return
    }

        resp, err := client.GetPlaneswalkerTypesSender(req)
        if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "scryfall.CatalogClient", "GetPlaneswalkerTypes", resp, "Failure sending request")
        return
        }

        result, err = client.GetPlaneswalkerTypesResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "scryfall.CatalogClient", "GetPlaneswalkerTypes", resp, "Failure responding to request")
        }

    return
}

    // GetPlaneswalkerTypesPreparer prepares the GetPlaneswalkerTypes request.
    func (client CatalogClient) GetPlaneswalkerTypesPreparer(ctx context.Context) (*http.Request, error) {
    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPath("/catalog/planeswalker-types"))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetPlaneswalkerTypesSender sends the GetPlaneswalkerTypes request. The method will close the
    // http.Response Body if it receives an error.
    func (client CatalogClient) GetPlaneswalkerTypesSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            }

    // GetPlaneswalkerTypesResponder handles the response to the GetPlaneswalkerTypes request. The method always
    // closes the http.Response Body.
    func (client CatalogClient) GetPlaneswalkerTypesResponder(resp *http.Response) (result Catalog, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

// GetPowers sends the get powers request.
func (client CatalogClient) GetPowers(ctx context.Context) (result Catalog, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/CatalogClient.GetPowers")
        defer func() {
            sc := -1
        if result.Response.Response != nil {
        sc = result.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
    req, err := client.GetPowersPreparer(ctx)
    if err != nil {
    err = autorest.NewErrorWithError(err, "scryfall.CatalogClient", "GetPowers", nil , "Failure preparing request")
    return
    }

        resp, err := client.GetPowersSender(req)
        if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "scryfall.CatalogClient", "GetPowers", resp, "Failure sending request")
        return
        }

        result, err = client.GetPowersResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "scryfall.CatalogClient", "GetPowers", resp, "Failure responding to request")
        }

    return
}

    // GetPowersPreparer prepares the GetPowers request.
    func (client CatalogClient) GetPowersPreparer(ctx context.Context) (*http.Request, error) {
    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPath("/catalog/powers"))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetPowersSender sends the GetPowers request. The method will close the
    // http.Response Body if it receives an error.
    func (client CatalogClient) GetPowersSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            }

    // GetPowersResponder handles the response to the GetPowers request. The method always
    // closes the http.Response Body.
    func (client CatalogClient) GetPowersResponder(resp *http.Response) (result Catalog, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

// GetSpellTypes sends the get spell types request.
func (client CatalogClient) GetSpellTypes(ctx context.Context) (result Catalog, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/CatalogClient.GetSpellTypes")
        defer func() {
            sc := -1
        if result.Response.Response != nil {
        sc = result.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
    req, err := client.GetSpellTypesPreparer(ctx)
    if err != nil {
    err = autorest.NewErrorWithError(err, "scryfall.CatalogClient", "GetSpellTypes", nil , "Failure preparing request")
    return
    }

        resp, err := client.GetSpellTypesSender(req)
        if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "scryfall.CatalogClient", "GetSpellTypes", resp, "Failure sending request")
        return
        }

        result, err = client.GetSpellTypesResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "scryfall.CatalogClient", "GetSpellTypes", resp, "Failure responding to request")
        }

    return
}

    // GetSpellTypesPreparer prepares the GetSpellTypes request.
    func (client CatalogClient) GetSpellTypesPreparer(ctx context.Context) (*http.Request, error) {
    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPath("/catalog/spell-types"))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSpellTypesSender sends the GetSpellTypes request. The method will close the
    // http.Response Body if it receives an error.
    func (client CatalogClient) GetSpellTypesSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            }

    // GetSpellTypesResponder handles the response to the GetSpellTypes request. The method always
    // closes the http.Response Body.
    func (client CatalogClient) GetSpellTypesResponder(resp *http.Response) (result Catalog, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

// GetToughnesses sends the get toughnesses request.
func (client CatalogClient) GetToughnesses(ctx context.Context) (result Catalog, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/CatalogClient.GetToughnesses")
        defer func() {
            sc := -1
        if result.Response.Response != nil {
        sc = result.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
    req, err := client.GetToughnessesPreparer(ctx)
    if err != nil {
    err = autorest.NewErrorWithError(err, "scryfall.CatalogClient", "GetToughnesses", nil , "Failure preparing request")
    return
    }

        resp, err := client.GetToughnessesSender(req)
        if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "scryfall.CatalogClient", "GetToughnesses", resp, "Failure sending request")
        return
        }

        result, err = client.GetToughnessesResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "scryfall.CatalogClient", "GetToughnesses", resp, "Failure responding to request")
        }

    return
}

    // GetToughnessesPreparer prepares the GetToughnesses request.
    func (client CatalogClient) GetToughnessesPreparer(ctx context.Context) (*http.Request, error) {
    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPath("/catalog/thoughnesses"))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetToughnessesSender sends the GetToughnesses request. The method will close the
    // http.Response Body if it receives an error.
    func (client CatalogClient) GetToughnessesSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            }

    // GetToughnessesResponder handles the response to the GetToughnesses request. The method always
    // closes the http.Response Body.
    func (client CatalogClient) GetToughnessesResponder(resp *http.Response) (result Catalog, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

// GetWatermarks sends the get watermarks request.
func (client CatalogClient) GetWatermarks(ctx context.Context) (result Catalog, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/CatalogClient.GetWatermarks")
        defer func() {
            sc := -1
        if result.Response.Response != nil {
        sc = result.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
    req, err := client.GetWatermarksPreparer(ctx)
    if err != nil {
    err = autorest.NewErrorWithError(err, "scryfall.CatalogClient", "GetWatermarks", nil , "Failure preparing request")
    return
    }

        resp, err := client.GetWatermarksSender(req)
        if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "scryfall.CatalogClient", "GetWatermarks", resp, "Failure sending request")
        return
        }

        result, err = client.GetWatermarksResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "scryfall.CatalogClient", "GetWatermarks", resp, "Failure responding to request")
        }

    return
}

    // GetWatermarksPreparer prepares the GetWatermarks request.
    func (client CatalogClient) GetWatermarksPreparer(ctx context.Context) (*http.Request, error) {
    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPath("/catalog/watermarks"))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetWatermarksSender sends the GetWatermarks request. The method will close the
    // http.Response Body if it receives an error.
    func (client CatalogClient) GetWatermarksSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            }

    // GetWatermarksResponder handles the response to the GetWatermarks request. The method always
    // closes the http.Response Body.
    func (client CatalogClient) GetWatermarksResponder(resp *http.Response) (result Catalog, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

// GetWordBank sends the get word bank request.
func (client CatalogClient) GetWordBank(ctx context.Context) (result Catalog, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/CatalogClient.GetWordBank")
        defer func() {
            sc := -1
        if result.Response.Response != nil {
        sc = result.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
    req, err := client.GetWordBankPreparer(ctx)
    if err != nil {
    err = autorest.NewErrorWithError(err, "scryfall.CatalogClient", "GetWordBank", nil , "Failure preparing request")
    return
    }

        resp, err := client.GetWordBankSender(req)
        if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "scryfall.CatalogClient", "GetWordBank", resp, "Failure sending request")
        return
        }

        result, err = client.GetWordBankResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "scryfall.CatalogClient", "GetWordBank", resp, "Failure responding to request")
        }

    return
}

    // GetWordBankPreparer prepares the GetWordBank request.
    func (client CatalogClient) GetWordBankPreparer(ctx context.Context) (*http.Request, error) {
    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPath("/catalog/word-bank"))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetWordBankSender sends the GetWordBank request. The method will close the
    // http.Response Body if it receives an error.
    func (client CatalogClient) GetWordBankSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            }

    // GetWordBankResponder handles the response to the GetWordBank request. The method always
    // closes the http.Response Body.
    func (client CatalogClient) GetWordBankResponder(resp *http.Response) (result Catalog, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

