# go-graphql-hackernews

## Description

Hackernews GraphQL API clone built with Go, MySQL, JWT authentication, and Schema-Driven Development.
The GraphQL layer is generated with [gqlgen](https://gqlgen.com/).

## Prerequisites

- Go 1.21+
- MySQL running locally with a database named `hackernews`
- Default DB connection: `root:dbpass@tcp(localhost)/hackernews`
  (edit `internal/pkg/db/mysql/mysql.go` to change credentials)

## Setup

1. **Clone the repo and install dependencies**
   ```bash
   git clone https://github.com/chrikar/go-graphql-hackernews.git
   cd go-graphql-hackernews
   go mod download
   ```

2. **Create the MySQL database**

   Connect to MySQL and create the database:
   ```bash
   mysql -u root -p
   ```
   ```sql
   CREATE DATABASE hackernews;
   EXIT;
   ```
   Database migrations (users and links tables) are applied automatically on startup.

3. **Start the server**
   ```bash
   go run server.go
   ```
   The server listens on port `8080` by default. Override with the `PORT` environment variable.

## Usage

Open the GraphQL Playground at http://localhost:8080/.

### Introspect available operations

```graphql
{
  __schema {
    queryType { name fields { name } }
    mutationType { name fields { name } }
  }
}
```

### Register a new user

Returns a JWT token.

```graphql
mutation {
  createUser(input: { username: "alice", password: "secret" })
}
```

### Log in

Returns a JWT token for an existing user.

```graphql
mutation {
  login(input: { username: "alice", password: "secret" })
}
```

### Refresh a token

```graphql
mutation {
  refreshToken(input: { token: "<your-jwt-token>" })
}
```

### Create a link (requires authentication)

Pass the JWT token in the `Authorization` request header, then run:

```graphql
mutation {
  createLink(input: { title: "Google", address: "https://google.com" }) {
    id
    title
    address
    user { id name }
  }
}
```

In the Playground, add the header under **HTTP Headers** (bottom-left panel):

```json
{
  "Authorization": "<your-jwt-token>"
}
```

### List all links

```graphql
{
  links {
    id
    title
    address
    user { id name }
  }
}
```
