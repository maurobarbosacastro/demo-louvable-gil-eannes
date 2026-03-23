package services

import (
	"fmt"
	"ms-cj/internal/dto"
	"ms-cj/internal/responses"
	"ms-cj/internal/services/redis"
	"ms-cj/pkg/dotenv"
	"ms-cj/pkg/logster"
	"strings"
	"time"

	"github.com/samber/lo"
)

func ProcessTransactions() {
	logster.StartFuncLog()

	startDate := time.Now().Add(-7 * 24 * time.Hour).Format("2006-01-02T15:04:05Z")
	endDate := time.Now().Format("2006-01-02T15:04:05Z")

	dtoParam := dto.CjTransactionDTO{
		PublisherId: dotenv.GetEnv("CJ_TAGPEAK_CID"),
		StartDate:   startDate,
		EndDate:     endDate,
		ApiToken:    dotenv.GetEnv("CJ_TOKEN"),
	}

	cjTransactions, err := GetCjTransaction(dtoParam)
	if err != nil {
		logster.Error(err, "Error getting CJ transactions")
		return
	}

	logster.Info(fmt.Sprintf("Transactions size %d", len(cjTransactions.PublisherCommissions.Records)))

	if cjTransactions.PublisherCommissions.Count == 0 {
		return
	}

	originals := lo.Filter(cjTransactions.PublisherCommissions.Records, func(r responses.CommissionRecord, _ int) bool {
		return r.Original
	})

	notOriginals := lo.Filter(cjTransactions.PublisherCommissions.Records, func(r responses.CommissionRecord, _ int) bool {
		return !r.Original
	})

	records := append(originals, notOriginals...)

	cjRefPrefix := dotenv.GetEnv("CJ_REF_PREFIX")
	records = lo.Filter(records, func(r responses.CommissionRecord, _ int) bool {
		// We split by _ to check from which env it comes. Prod does not have a prefix, other envs have.
		// So, if the prefix is set to empty but the record has a prefix, we filter it out.
		splits := strings.Split(r.ShopperId, "_")

		if cjRefPrefix == "" {
			// Production: Keep only records WITHOUT prefix (single part)
			return len(splits) == 1
		} else {
			// Non-prod: Keep only records WITH the correct prefix
			return len(splits) == 2 && splits[0] == cjRefPrefix
		}
	})

	for _, cjTransaction := range records {
		redis.PushTransactionToRedisQueue(redis.CjTransactionKey, &cjTransaction)
	}

	logster.EndFuncLog()
}
