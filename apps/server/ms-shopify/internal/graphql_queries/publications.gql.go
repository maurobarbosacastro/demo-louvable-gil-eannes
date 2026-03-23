package graphql_queries

var GetPublications string = `#graphql
    query Publications{
      publications(first: 10) {
        nodes{
          id,
            name,
          catalog{
            id,
            title
          }
        }
      }
    }`

var PublishCollectionToPublication string = `#graphql
    mutation PublishablePublish($collectionId: ID!, $publicationId: ID!) {
      publishablePublish(id: $collectionId, input: {publicationId: $publicationId}) {
        publishable {
          publishedOnPublication(publicationId: $publicationId)
        }
        userErrors {
          field
          message
        }
      }
    }`
