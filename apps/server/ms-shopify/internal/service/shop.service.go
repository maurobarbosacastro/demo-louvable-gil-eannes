package service

import (
	"fmt"
	"github.com/google/uuid"
	"ms-shopify/internal/dto"
	"ms-shopify/internal/dto/shopify_dtos"
	"ms-shopify/internal/models"
	"ms-shopify/internal/models/shopify_models"
	"ms-shopify/internal/repository"
	"ms-shopify/pkg/dotenv"
	"ms-shopify/pkg/http_client"
	"ms-shopify/pkg/logster"
	"net/http"
	"time"
)

func GetShopByUrl(url string) (*models.Shop, error) {
	return repository.GetShopByUrl(url)
}

func CreateShop(body dto.CreateShopifyShopDTO) (*models.Shop, error) {
	return repository.CreateShop(body.ToModel())
}

func UpdateShop(uuid uuid.UUID, body dto.UpdateShopifyShopDTO) (*models.Shop, error) {
	logster.StartFuncLog()

	shop, err := repository.GetByUuid(uuid)
	if err != nil {
		logster.Error(err, "Error getting shop")
		logster.EndFuncLog()
		return nil, err
	}

	if body.AccessToken != nil {
		shop.AccessToken = body.AccessToken
	}

	if body.State != nil {
		shop.State = *body.State
	}

	if body.InstallationDone != nil {
		shop.InstallationDone = *body.InstallationDone
	}

	logster.EndFuncLog()
	return repository.UpdateShop(*shop)
}

func UpdateTokenByShopUrl(shopUrl string, token string) {
	logster.StartFuncLog()

	shop, err := repository.GetShopByUrl(shopUrl)
	if err != nil {
		logster.Error(err, "Error getting shop")
	}

	shop.AccessToken = &token
	_, err = repository.UpdateShop(*shop)
	if err != nil {
		logster.Error(err, "Error updating shop")
	}

	logster.EndFuncLog()
}

func GetCollectionId() (*string, error) {
	logster.StartFuncLog()
	collectionSearchDTO := shopify_dtos.CollectionsSearch{
		Title: "title:" + dotenv.GetEnv("SHOPIFY_COLLECTION_TITLE"),
	}
	//Check if collection exists, if not create collection
	collection, err := ShopifyClient.GetCollectionByTitle(collectionSearchDTO)
	if err != nil {
		logster.Error(err, "Error getting collection by handle")
		logster.EndFuncLog()
		return nil, err
	}

	if len(collection.Collections.Nodes) == 0 {
		logster.Info("No collection found, creating new collection")
		//Create collection
		collectionCreateDTO := shopify_dtos.CreateCollection{
			Input: shopify_dtos.CreateCollectionInput{
				Title: dotenv.GetEnv("SHOPIFY_COLLECTION_TITLE"),
			},
		}

		createdCollection, err := ShopifyClient.CreateCollection(collectionCreateDTO)
		if err != nil {
			logster.Error(err, "Error creating collection")
			logster.EndFuncLog()
			return nil, err
		}

		logster.EndFuncLogMsg(fmt.Sprintf("Collection created: %s", createdCollection.CollectionCreate.Collection.ID))
		return &createdCollection.CollectionCreate.Collection.ID, nil
	} else {
		//Use existing collection
		logster.EndFuncLogMsg(fmt.Sprintf("Collection found: %s", collection.Collections.Nodes[0].ID))
		return &collection.Collections.Nodes[0].ID, nil
	}
}

func PublishCollection(collectionId string) error {
	logster.StartFuncLog()
	publishableChannels, err := ShopifyClient.GetPublications()
	if err != nil {
		logster.Error(err, "Error getting publications")
		logster.EndFuncLog()
		return err
	}

	for _, channel := range publishableChannels.Publications.Nodes {
		logster.Debug(fmt.Sprintf("Channel: %s", channel.Name))

		publishableChannel := shopify_dtos.PublishCollectionToPublication{
			CollectionId:  collectionId,
			PublicationId: channel.ID,
		}

		publishable, err := ShopifyClient.PublishCollectionToPublication(publishableChannel)
		if err != nil {
			logster.Error(err, "Error publishing collection to publication")
			logster.EndFuncLog()
			return err
		}
		logster.Info(fmt.Sprintf("Published: %v", publishable.PublishablePublish.Publishable.PublishedOnPublication))
	}

	logster.EndFuncLog()
	return nil
}

func HandleDiscountCodeCreation(collectionId string) (*shopify_models.DiscountCodeBasicCreateResponse, error) {
	logster.StartFuncLog()

	eastern, err := time.LoadLocation("America/New_York")

	if err != nil {
		logster.Panic(err, "Error loading location: America/New_York")
	}

	discountDto := shopify_dtos.BasicCodeBxgyDiscount{
		BasicCodeDiscount: shopify_dtos.BasicCodeDiscount{
			Title:    "Tagpeak discount on collection Tagpeak",
			Code:     dotenv.GetEnv("SHOPIFY_TAGPEAK_CODE"),
			StartsAt: time.Now().In(eastern).Format("2006-01-02T15:04:05Z"),
			CustomerSelection: shopify_dtos.CustomerSelection{
				All: true,
			},
			CustomerGets: shopify_dtos.CustomerGets{
				Value: shopify_dtos.Value{
					Percentage: 0,
				},
				Items: shopify_dtos.Items{
					Collections: shopify_dtos.Collections{
						Add: collectionId,
					},
				},
			},
		},
	}

	codeCreated, errCodeCreation := ShopifyClient.CreateDiscountCode(discountDto)

	if errCodeCreation != nil {
		logster.Error(errCodeCreation, "Error creating discount code")
		logster.EndFuncLog()
		return nil, errCodeCreation
	}

	return codeCreated, nil
}

func GetCustomerEmail(customerId string) (string, error) {
	logster.StartFuncLogMsg(customerId)

	customerEmail, err := ShopifyClient.GetCustomerEmail(shopify_dtos.GetCustomerEmail{ID: customerId})
	if err != nil {
		logster.Error(err, "Error getting customer email")
		logster.EndFuncLog()
		return "", err
	}

	logster.EndFuncLog()
	return customerEmail.Customer.Email, nil
}

func AddProductsToCollection(collectionId string, productIds []string) error {
	logster.StartFuncLogMsg(fmt.Sprintf("CollectionId: %s, ProductIds: %v", collectionId, productIds))

	in := shopify_dtos.CollectionAddProducts{
		ID:         collectionId,
		ProductIds: productIds,
	}

	result, err := ShopifyClient.CollectionAddProductsV2(in)
	if err != nil {
		logster.Error(err, "Error adding products to collection")
		logster.EndFuncLog()
		return err
	}

	if result.CollectionAddProductsV2.Job != nil {
		logster.Info(fmt.Sprintf("Job ID: %s", result.CollectionAddProductsV2.Job.ID))
	}

	logster.EndFuncLog()
	return nil
}

func GetShopOfflineToken(shop string) *string {
	logster.StartFuncLogMsg(fmt.Sprintf("Shop: %s", shop))

	httpClient := &http_client.HttpClient{HttpClient: &http.Client{}}
	shopifyPluginUrl := dotenv.GetEnv("MS_SHOPIFY_PLUGIN_URL")

	headers := map[string]string{"x-shopify-shop-domain": shop}
	out := &models.OfflineTokenShopStruct{}
	url := fmt.Sprintf("%sapi/session", shopifyPluginUrl)
	_, err := httpClient.Get(url, headers, out)

	if err != nil {
		logster.Error(err, "Error getting offline token")
		logster.EndFuncLog()
		return nil
	}

	logster.EndFuncLog()
	return &out.OfflineTokenShop
}

