package service

import (
	"errors"
	"fmt"
	goshopify "github.com/bold-commerce/go-shopify/v4"
	"golang.org/x/net/context"
	"ms-shopify/internal/dto/shopify_dtos"
	"ms-shopify/internal/graphql_queries"
	"ms-shopify/internal/models/shopify_models"
	"ms-shopify/pkg/dotenv"
	"ms-shopify/pkg/logster"
)

type shopifyClientStruct struct {
	*goshopify.Client
}

var ShopifyClient *shopifyClientStruct = &shopifyClientStruct{}

func (s *shopifyClientStruct) CreateClient(shop string, token string) {
	logster.StartFuncLogMsg(fmt.Sprintf("Creating client for %s", shop))

	// Create an app somewhere.
	app := goshopify.App{
		ApiKey:      dotenv.GetEnv("SHOPIFY_API_KEY"),
		ApiSecret:   dotenv.GetEnv("SHOPIFY_API_SECRET"),
		RedirectUrl: dotenv.GetEnv("SHOPIFY_REDIRECT_URL"),
		Scope:       dotenv.GetEnv("SHOPIFY_SCOPE"),
	}

	// Create a new API client
	client, err := goshopify.NewClient(app, shop, token, goshopify.WithVersion("2025-04"))

	if err != nil {
		logster.Error(err, "Error creating client")
		logster.EndFuncLog()
	}

	ShopifyClient.Client = client
	logster.EndFuncLogMsg("Shopify Client created")
}

func (s *shopifyClientStruct) GetDiscountCode(code string) (*shopify_models.CodeDiscountNodeByCodeResponse, error) {
	logster.StartFuncLogMsg(fmt.Sprintf("Code: %s", code))

	in := shopify_dtos.CodeDiscountNodeByCode{Code: code}
	out := new(shopify_models.CodeDiscountNodeByCodeResponse)

	logster.Info(fmt.Sprintf("Query: %v", graphql_queries.CodeDiscountNodeByCode))
	err := s.GraphQL.Query(
		context.Background(),
		graphql_queries.CodeDiscountNodeByCode,
		in,
		out)

	if err != nil {
		logster.Error(err, "Error creating discount code")
		logster.EndFuncLog()
		return nil, err
	}

	logster.EndFuncLogMsg(fmt.Sprintf("Code: %s", out.CodeDiscountNodeByCode.ID))
	return out, nil
}

func (s *shopifyClientStruct) GetCollectionByTitle(in shopify_dtos.CollectionsSearch) (*shopify_models.CollectionsResponse, error) {
	logster.StartFuncLog()
	out := new(shopify_models.CollectionsResponse)
	logster.Info(fmt.Sprintf("Query: %v", graphql_queries.GetCollections))
	err := s.GraphQL.Query(
		context.Background(),
		graphql_queries.GetCollections,
		in,
		out)

	if err != nil {
		logster.Error(err, "Error getting collection by handle")
		logster.EndFuncLog()
		return nil, err
	}

	logster.EndFuncLog()
	return out, nil
}

func (s *shopifyClientStruct) CreateCollection(in shopify_dtos.CreateCollection) (*shopify_models.CollectionCreateResponse, error) {
	logster.StartFuncLog()
	out := new(shopify_models.CollectionCreateResponse)
	logster.Info(fmt.Sprintf("Query: %v", graphql_queries.CreateCollection))
	err := s.GraphQL.Query(
		context.Background(),
		graphql_queries.CreateCollection,
		in,
		out)

	if err != nil {
		logster.Error(err, "Error creating collection")
		logster.EndFuncLog()
		return nil, err
	}

	if len(out.CollectionCreate.UserErrors) > 0 {
		logster.Error(nil, fmt.Sprintf("Error creating collection: %v", out.CollectionCreate.UserErrors))
		logster.EndFuncLog()
		return nil, errors.New(fmt.Sprintf("Error creating collection: %v", out.CollectionCreate.UserErrors))
	}

	return out, nil
}

func (s *shopifyClientStruct) CollectionAddProductsV2(in shopify_dtos.CollectionAddProducts) (*shopify_models.CollectionAddProductsV2Response, error) {
	logster.StartFuncLog()
	out := new(shopify_models.CollectionAddProductsV2Response)
	logster.Info(fmt.Sprintf("Query: %v", graphql_queries.CollectionAddProductsV2))
	err := s.GraphQL.Query(
		context.Background(),
		graphql_queries.CollectionAddProductsV2,
		in,
		out)

	if err != nil {
		logster.Error(err, "Error adding products to collection")
		logster.EndFuncLog()
		return nil, err
	}

	if len(out.CollectionAddProductsV2.UserErrors) > 0 {
		logster.Error(nil, fmt.Sprintf("Error adding products to collection: %v", out.CollectionAddProductsV2.UserErrors))
		logster.EndFuncLog()
		return nil, errors.New(fmt.Sprintf("Error adding products to collection: %v", out.CollectionAddProductsV2.UserErrors))
	}

	return out, nil
}

func (s *shopifyClientStruct) GetPublications() (*shopify_models.PublicationResponse, error) {
	logster.StartFuncLog()
	out := new(shopify_models.PublicationResponse)
	logster.Info(fmt.Sprintf("Query: %v", graphql_queries.GetPublications))
	err := s.GraphQL.Query(
		context.Background(),
		graphql_queries.GetPublications,
		nil,
		out,
	)

	if err != nil {
		logster.Error(err, "Error getting publications")
		logster.EndFuncLog()
		return nil, err
	}

	return out, nil
}

func (s *shopifyClientStruct) PublishCollectionToPublication(in shopify_dtos.PublishCollectionToPublication) (*shopify_models.PublishablePublishResponse, error) {
	logster.StartFuncLog()
	out := new(shopify_models.PublishablePublishResponse)
	logster.Info(fmt.Sprintf("Query: %v", graphql_queries.PublishCollectionToPublication))
	err := s.GraphQL.Query(
		context.Background(),
		graphql_queries.PublishCollectionToPublication,
		in,
		out)

	if err != nil {
		logster.Error(err, "Error publishing collection to publication")
		logster.EndFuncLog()
		return nil, err
	}

	if len(out.PublishablePublish.UserErrors) > 0 {
		logster.Error(nil, fmt.Sprintf("Error publishing collection to publication: %v", out.PublishablePublish.UserErrors))
		logster.EndFuncLog()
		return nil, errors.New(fmt.Sprintf("Error publishing collection to publication: %v", out.PublishablePublish.UserErrors))
	}

	return out, nil
}

func (s *shopifyClientStruct) CreateDiscountCode(in shopify_dtos.BasicCodeBxgyDiscount) (*shopify_models.DiscountCodeBasicCreateResponse, error) {
	logster.StartFuncLog()
	out := new(shopify_models.DiscountCodeBasicCreateResponse)
	logster.Info(fmt.Sprintf("Query: %v", graphql_queries.DiscountCodeBxgyCreate))
	err := s.GraphQL.Query(
		context.Background(),
		graphql_queries.DiscountCodeBxgyCreate,
		in,
		out)

	if err != nil {
		logster.Error(err, "Error creating discount code")
		logster.EndFuncLog()
		return nil, err
	}

	if len(out.DiscountCodeBasicCreate.UserErrors) > 0 {
		logster.Error(nil, fmt.Sprintf("Error creating discount code: %v", out.DiscountCodeBasicCreate.UserErrors))
		logster.EndFuncLog()
		return nil, errors.New(fmt.Sprintf("Error creating discount code: %v", out.DiscountCodeBasicCreate.UserErrors))
	}

	logster.EndFuncLog()
	return out, nil
}

func (s *shopifyClientStruct) GetCustomerEmail(in shopify_dtos.GetCustomerEmail) (*shopify_models.CustomerResponse, error) {
	logster.StartFuncLog()
	out := new(shopify_models.CustomerResponse)
	logster.Info(fmt.Sprintf("Query: %v", graphql_queries.GetCustomerEmail))

	err := s.GraphQL.Query(
		context.Background(),
		graphql_queries.GetCustomerEmail,
		in,
		out)

	if err != nil {
		logster.Error(err, "Error getting customer email")
		logster.EndFuncLog()
		return nil, err
	}

	logster.EndFuncLog()
	return out, nil
}
