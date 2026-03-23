#!/bin/bash

set -e

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

ENV=${1:-dev}

case $ENV in
dev)
    MS_EMAIL_HOST="http://localhost:8093"
    MS_TAGPEAK_HOST="http://localhost:8097"
    MS_SHOPIFY_HOST="http://localhost:8297"
    MS_IMAGES_HOST="http://localhost:8099"
    MS_CJ_HOST="http://localhost:8397"
    MS_INTERACTIVE_BROKERS_HOST="http://localhost:8197"
    ;;
qa)
    MS_EMAIL_HOST="https://qa.tagpeak.atlanse.ddns.net/email"
    MS_TAGPEAK_HOST="https://qa.tagpeak.atlanse.ddns.net/tagpeak"
    MS_SHOPIFY_HOST="https://qa.tagpeak.atlanse.ddns.net/shopify" # TODO: Verify shopify path
    MS_IMAGES_HOST="https://qa.tagpeak.atlanse.ddns.net/images"
    MS_CJ_HOST="https://qa.tagpeak.atlanse.ddns.net/cj"
    MS_INTERACTIVE_BROKERS_HOST="https://qa.tagpeak.atlanse.ddns.net/ibkr"
    ;;
pre)
    MS_EMAIL_HOST="https://pre.tagpeak.atlanse-cloud.ddns.net/email"
    MS_TAGPEAK_HOST="https://pre.tagpeak.atlanse-cloud.ddns.net/tagpeak"
    MS_SHOPIFY_HOST="https://pre.tagpeak.atlanse-cloud.ddns.net/shopify" # TODO: Verify shopify path
    MS_IMAGES_HOST="https://pre.tagpeak.atlanse-cloud.ddns.net/images"
    MS_CJ_HOST="https://pre.tagpeak.atlanse-cloud.ddns.net/cj"
    MS_INTERACTIVE_BROKERS_HOST="https://pre.tagpeak.atlanse-cloud.ddns.net/ibkr"
    ;;
prod)
    MS_EMAIL_HOST="https://app.tagpeak.com/email"
    MS_TAGPEAK_HOST="https://app.tagpeak.com/tagpeak"
    MS_SHOPIFY_HOST="https://app.tagpeak.com/shopify"
    MS_IMAGES_HOST="https://app.tagpeak.com/images"
    MS_CJ_HOST="https://app.tagpeak.com/cj"
    MS_INTERACTIVE_BROKERS_HOST="https://app.tagpeak.com/ibkr"
    ;;
*)
    echo -e "${RED}Invalid environment: $ENV${NC}"
    echo "Usage: $0 [dev|qa|pre|prod]"
    echo "  dev  - Local development (localhost)"
    echo "  qa   - QA environment"
    echo "  pre  - Pre-production environment"
    echo "  prod - Production environment"
    exit 1
    ;;
esac

echo "======================================"
echo "  Health Check Test Script"
echo "======================================"
echo -e "Environment: ${BLUE}${ENV}${NC}"
echo ""

test_health_endpoint() {
    local service_name=$1
    local url=$2

    echo -n "Testing ${service_name} [${url}]... "

    response=$(curl -s -o /dev/null -w "%{http_code}" --max-time 5 "${url}" 2>/dev/null)

    if [ $? -ne 0 ]; then
        echo -e "${RED}FAILED${NC} - Connection error"
        return 1
    fi

    if [ "$response" = "200" ]; then
        echo -e "${GREEN}OK${NC} (HTTP $response)"
        return 0
    else
        echo -e "${RED}FAILED${NC} (HTTP $response)"
        return 1
    fi
}

passed=0
failed=0

test_health_endpoint "ms-email              " "${MS_EMAIL_HOST}/health" && ((passed++)) || ((failed++))
test_health_endpoint "ms-tagpeak            " "${MS_TAGPEAK_HOST}/health" && ((passed++)) || ((failed++))
test_health_endpoint "ms-shopify            " "${MS_SHOPIFY_HOST}/health" && ((passed++)) || ((failed++))
test_health_endpoint "ms-images             " "${MS_IMAGES_HOST}/health" && ((passed++)) || ((failed++))
test_health_endpoint "ms-cj                 " "${MS_CJ_HOST}/health" && ((passed++)) || ((failed++))
test_health_endpoint "ms-interactive-brokers" "${MS_INTERACTIVE_BROKERS_HOST}/health" && ((passed++)) || ((failed++))

echo ""
echo "======================================"
echo -e "Results: ${GREEN}${passed} passed${NC}, ${RED}${failed} failed${NC}"
echo "======================================"

if [ $failed -gt 0 ]; then
    exit 1
fi

exit 0
