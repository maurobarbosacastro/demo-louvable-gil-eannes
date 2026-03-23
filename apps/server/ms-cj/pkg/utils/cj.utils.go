package utils

import (
	"fmt"
	"ms-cj/internal/dto"
)

func MapCustomCommissionQuery(dtoParam dto.CjTransactionDTO) string {
	return fmt.Sprintf(
		`{
        publisherCommissions(
            forPublishers: ["%s"]
            sincePostingDate: "%s"
            beforePostingDate: "%s"
        ) {
            count
            payloadComplete
            records {
              actionStatus
              actionTrackerName
              actionType
              advertiserId
              advertiserName
              aid
              commissionId
              clickDate
              clickReferringURL
              concludingDeviceName
              correctionReason
              country
              coupon
              eventDate
              items {
                quantity
                perItemSaleAmountAdvCurrency
                perItemSaleAmountPubCurrency
                perItemSaleAmountUsd
                totalCommissionAdvCurrency
                totalCommissionPubCurrency
                totalCommissionUsd
              }
              lockingDate
              orderId
              original
              postingDate
              pubCommissionAmountPubCurrency
              pubCommissionAmountUsd
              saleAmountPubCurrency
              saleAmountUsd
              verticalAttributes {
                  businessUnit
              }
              websiteName
              shopperId
              validationStatus
            }
        }
    }`, dtoParam.PublisherId, dtoParam.StartDate, dtoParam.EndDate)
}

func MapCustomAdIdQuery(dtoParam dto.AdIdDTO) string {
	return fmt.Sprintf(
		`{
          shoppingProductFeeds(
            companyId: %s,
            partnerIds: [%s]
          ) {
            totalCount
            count
            resultList {
              adId
              feedName
              advertiserId
              productCount
              advertiserCountry
              lastUpdated
              advertiserName
              language
              currency
              sourceFeedType
            }
          }
        }`, dtoParam.PublisherId, dtoParam.AdvertiserId)
}
