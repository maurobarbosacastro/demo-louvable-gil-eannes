# Docker Installation for Remix Shopify App

This guide explains how to set up and run this Remix Shopify app using Docker. It covers the configuration for different environments: `local`, `qa`, `pre`, and `prod`.

## Prerequisites

- [Docker](https://docs.docker.com/get-docker/)
- [Node.js](https://nodejs.org/en/)
- [Shopify CLI](https://shopify.dev/docs/apps/tools/cli)

## 1. Environment Configuration

This application uses environment variables to configure the Shopify connection. The `app/shopify.server.ts` file shows that the configuration is loaded from `process.env`.

Create a `.env` file in the root of the project. You can use the `shopify.app.<environment>.toml` files as a reference.

### Environment-Specific Configurations

**Local (`local`)**

- **`shopify.app.local.toml`**
- **Description**: For local development and testing.
- Managed by the shopify cli
- **`.env` file for `local`**:

```
SHOPIFY_API_KEY="api-key"
BE_URL="your-local-api-secret"
BO_URL="https://wisconsin-variance-ye-ours.trycloudflare.com"
```

**QA (`qa`)**

- **`shopify.app.qa.toml`**
- **Description**: For quality assurance testing.
- **`.env` file for `qa`**:

```
SHOPIFY_API_KEY="qa-api-key"
SHOPIFY_API_SECRET="your-qa-api-secret"
SHOPIFY_APP_URL="https://qa.tagpeak.shopify.atlanse.ddns.net"
```

**Pre-production (`pre`)**

- **`shopify.app.pre.toml`**
- **Description**: For staging and pre-production testing.
- **`.env` file for `pre`**:

```
SHOPIFY_API_KEY="pre-api-key"
SHOPIFY_API_SECRET="your-pre-api-secret"
SHOPIFY_APP_URL="https://pre.tagpeak.shopify.atlanse-cloud.ddns.net"
```

**Production (`prod`)**

- **`shopify.app.toml`**
- **Description**: For the production environment.
- **`.env` file for `prod`**:

```
SHOPIFY_API_KEY="your-production-api-key"
SHOPIFY_API_SECRET="your-production-api-secret"
SHOPIFY_APP_URL="your-production-app-url"
SCOPES="read_discounts,read_orders,read_products,read_publications,write_discounts,write_products,write_publications"
```

**Important**: Replace `your-*-api-secret`, and other placeholder values with the actual credentials for each environment from your Shopify Partner Dashboard and your infrastructure provider.

## 2. Database Configuration --- IGNORAR

The `prisma/schema.prisma` file is configured to use a SQLite database for local development. For `qa`, `pre`, and `prod` environments, you should use a more robust database and update the `DATABASE_URL` accordingly.

## 3. Building the Application

Before deploying or running the application in production, you need to build the Remix app. This process compiles the application code into static assets and a server build.

### Manual Build

To build the application manually, run the following command in your terminal:

```bash
npm run build
```

This will create a `build` directory containing the production-ready application.

### Docker Build

The `Dockerfile` is configured to automatically build the application. The `RUN npm run build` command is included in the Dockerfile, so when you build the Docker image, the Remix application is also built.

## 4. Build and Run with Docker

Once you have created your `.env` file for the desired environment, you can build and run the application using Docker.

1. **Build the Docker image**:

   ```bash
   docker build -t remix-shopify-app .
   ```

2. **Run the Docker container**:

   ```bash
   docker run -p 3000:3000 --env-file .env remix-shopify-app
   ```

   This command maps port 3000 of the container to port 3000 on your host machine and loads the environment variables from the `.env` file.

## 5. Shopify App Setup

To connect your Dockerized app to Shopify, you need to use the Shopify CLI to select the correct configuration:

```bash
# For local development
shopify app config use shopify.app.local.toml

# For QA
shopify app config use shopify.app.qa.toml

# For pre-production
shopify app config use shopify.app.pre.toml

# For production
shopify app config use shopify.app.toml
```

After selecting the configuration, you can deploy your app:

```bash
shopify app deploy
```

This will update the app URLs and other settings in your Shopify Partner Dashboard to match the selected environment.

## 6. GitLab CI/CD Pipeline

Here is an example of a `.gitlab-ci.yml` file for automating the deployment of this Shopify app.

```yaml
stages:

- build
- deploy

variables:
IMAGE_TAG: $CI_REGISTRY_IMAGE:$CI_COMMIT_REF_SLUG

build:
  stage: build
  image: docker:20.10.16
  services:
  - docker:20.10.16-dind
  script:
  - docker login -u $CI_REGISTRY_USER -p $CI_REGISTRY_PASSWORD $CI_REGISTRY
  - docker build -t $IMAGE_TAG .
  - docker push $IMAGE_TAG

deploy_qa:
  stage: deploy
  image: node:18-alpine
  script:
  - npm install -g @shopify/cli@3.78.0
  - export BO_URL=$BO_URL_QA
  - export BE_URL=$BE_URL_QA
  - export SHOPIFY_API_KEY=$SHOPIFY_API_KEY_QA
  - shopify app config use shopify.app.qa.toml
  - shopify app deploy
  rules:
  - if: $CI_COMMIT_BRANCH == 'develop'

deploy_pre:
  stage: deploy
  image: node:18-alpine
  script:
  - npm install -g @shopify/cli
  - export BO_URL=$BO_URL_PRE
  - export BE_URL=$BE_URL_PRE
  - export SHOPIFY_API_KEY=$SHOPIFY_API_KEY_PRE
  - shopify app config use shopify.app.pre.toml
  - shopify app deploy
  rules:
  - if: $CI_COMMIT_BRANCH == 'staging'

deploy_prod:
  stage: deploy
  image: node:18-alpine
  script:
  - npm install -g @shopify/cli
  - export BO_URL=$BO_URL_PROD
  - export BE_URL=$BE_URL_PROD
  - export SHOPIFY_API_KEY=$SHOPIFY_API_KEY_PROD
  - shopify app config use shopify.app.toml
  - shopify app deploy
  rules:
  - if: $CI_COMMIT_BRANCH == 'main'
```

### CI/CD Variables

You will need to configure the following variables in your GitLab project's CI/CD settings:

- `BO_URL_QA`: The Backoffice URL for the QA environment.
- `BE_URL_QA`: The Backend URL for the QA environment.
- `SHOPIFY_API_KEY_QA`: The API key for the QA environment.
- `BO_URL_PRE`: The Backoffice URL for the PRE environment.
- `BE_URL_PRE`: The Backend URL for the PRE environment.
- `SHOPIFY_API_KEY_PRE`: The API key for the pre-production environment.
- `BO_URL_PROD`: The Backoffice URL for the PROD environment.
- `BE_URL_PROD`: The Backend URL for the PROD environment.
- `SHOPIFY_API_KEY_PROD`: The API key for the production environment.

This pipeline will build a Docker image of your application, push it to the GitLab container registry, and then deploy the app to the appropriate environment based on the branch that triggered the pipeline.
