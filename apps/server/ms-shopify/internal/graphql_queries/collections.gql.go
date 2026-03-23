package graphql_queries

var GetCollections string = `#graphql
    query CustomCollectionList($title: String!) {
      collections(first: 1, query: $title) {
        nodes {
          id
          title
        }
      }
    }`

var CollectionAddProductsV2 string = `#graphql
    mutation CollectionAddProductsV2($id: ID!, $productIds: [ID!]!) {
      collectionAddProductsV2(id: $id, productIds: $productIds) {
        job {
          id
        }
        userErrors {
          field
          message
          code
        }
      }
    }`

var CreateCollection string = `#graphql
    mutation createCollectionMetafields($input: CollectionInput!) {
      collectionCreate(input: $input) {
        collection {
          id
          metafields(first: 3) {
            edges {
              node {
                id
                namespace
                key
                value
              }
            }
          }
        }
        userErrors {
          message
          field
        }
      }
    }`
