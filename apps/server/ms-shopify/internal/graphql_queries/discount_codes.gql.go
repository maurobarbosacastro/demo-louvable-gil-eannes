package graphql_queries

var CodeDiscountNodeByCode string = `#graphql
  query codeDiscountNodeByCode($code: String!) {
    codeDiscountNodeByCode(code: $code) {
      codeDiscount {
        __typename
        ... on DiscountCodeBasic {
          codesCount {
            count
          }
          shortSummary
        }
      }
      id
    }
  }`

var DiscountCodeBxgyCreate string = `#graphql
    mutation CreateDiscountCode($basicCodeDiscount: DiscountCodeBasicInput!) {
      discountCodeBasicCreate(basicCodeDiscount: $basicCodeDiscount) {
        codeDiscountNode {
          id
          codeDiscount {
            ... on DiscountCodeBasic {
              title
              startsAt
              customerSelection {
                ... on DiscountCustomerAll {
                  allCustomers
                }
              }
              customerGets {
                value {
                  ... on DiscountPercentage {
                    percentage
                  }
                }
                items {
                  ... on DiscountCollections {
                    collections(first: 10) {
                      nodes {
                        id
                        title
                      }
                    }
                  }
                }
              }
            }
          }
        }
        userErrors {
          field
          message
        }
      }
    }`
