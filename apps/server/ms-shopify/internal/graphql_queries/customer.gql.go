package graphql_queries

var GetCustomerEmail string = `
query getCustomerEmail($id: ID!) {
  customer(id: $id) {
    email
  }
}
`
