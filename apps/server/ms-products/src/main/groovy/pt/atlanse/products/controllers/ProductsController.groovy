package pt.atlanse.products.controllers


import groovy.util.logging.Slf4j
import io.micronaut.context.annotation.Parameter
import io.micronaut.core.annotation.NonNull
import io.micronaut.data.model.Pageable
import io.micronaut.http.HttpRequest
import io.micronaut.http.HttpResponse
import io.micronaut.http.HttpStatus
import io.micronaut.http.MutableHttpResponse
import io.micronaut.http.annotation.Body
import io.micronaut.http.annotation.Controller
import io.micronaut.http.annotation.Delete
import io.micronaut.http.annotation.Error
import io.micronaut.http.annotation.Get
import io.micronaut.http.annotation.Patch
import io.micronaut.http.annotation.Post
import io.micronaut.http.annotation.Put
import io.micronaut.http.annotation.QueryValue
import io.micronaut.scheduling.TaskExecutors
import io.micronaut.scheduling.annotation.ExecuteOn
import jakarta.inject.Inject
import jakarta.validation.ConstraintViolationException
import jakarta.validation.Valid
import jakarta.validation.constraints.NotBlank
import pt.atlanse.products.domains.BrandEntity
import pt.atlanse.products.domains.CategoryEntity

import pt.atlanse.products.domains.DiscountEntity
import pt.atlanse.products.domains.IngredientEntity
import pt.atlanse.products.domains.ProductEntity
import pt.atlanse.products.domains.RatingEntity
import pt.atlanse.products.domains.SellingStoreEntity
import pt.atlanse.products.dtos.Product
import pt.atlanse.products.dtos.ProductExtrasDTO
import pt.atlanse.products.dtos.ProductParams
import pt.atlanse.products.dtos.ProductSearchDTO
import pt.atlanse.products.dtos.Review
import pt.atlanse.products.models.CustomException
import pt.atlanse.products.repositories.ProductRepository
import pt.atlanse.products.repositories.RatingRepository
import pt.atlanse.products.repositories.SellingStoreRepository
import pt.atlanse.products.services.BrandService
import pt.atlanse.products.services.CategoryService

import pt.atlanse.products.services.DiscountService
import pt.atlanse.products.services.ImagesClientService
import pt.atlanse.products.services.IngredientService
import pt.atlanse.products.services.ProductExtrasService
import pt.atlanse.products.services.ProductService
import pt.atlanse.products.services.RatingService
import pt.atlanse.products.services.ReviewService
import pt.atlanse.products.services.SellingStoreService


@Slf4j
@ExecuteOn(TaskExecutors.IO)
@Controller("/api/products")
class ProductsController {

    @Inject
    ProductService products

    @Inject
    ProductExtrasService productExtrasService

    @Inject
    ReviewService reviews

    @Inject
    RatingService ratings

    @Inject
    BrandService brands

    @Inject
    DiscountService discounts

    @Inject
    CategoryService categories

    @Inject
    SellingStoreService sellingStores

    @Inject
    IngredientService ingredientService

    @Inject
    ImagesClientService imagesClientService

    @Inject
    ProductRepository productRepository

    @Inject
    RatingRepository ratingRepository

    @Inject
    SellingStoreRepository sellingStoreRepository

    @Error(exception = CustomException.class)
    MutableHttpResponse exceptionHandle(HttpRequest request, CustomException ex) {
        log.error "The exception: ${ ex.toString() }"
        return HttpResponse.status(ex.status).body([message: ex.title,
                                                    details: ex.details,
                                                    link   : request.path])
    }

    @Error(exception = ConstraintViolationException.class)
    MutableHttpResponse randomExceptionHandle(HttpRequest request, ConstraintViolationException ex) {
        log.error "The exception: ${ ex.toString() }"
        return HttpResponse.badRequest([message: "The payload is incorrect",
                                        details: "The payload for the specified action is invalid; Reason: ${ ex.message }",
                                        link   : request.path])
    }

    @Get("{?params*}")
    MutableHttpResponse findAll(ProductParams params, Pageable pageable) {
        // Find all brands using pagination
        HttpResponse.ok(products.findAll(params, pageable))
    }

    @Get("/{id}")
    MutableHttpResponse find(@NonNull @NotBlank UUID id) {
        // 1. Find brand by id and return with OK status
        HttpResponse.ok(products.findById(id))
    }

    @Post
    MutableHttpResponse add(@Body @Valid Product product) {

        // 1. Brand and Stores are mandatory values
        BrandEntity brand = brands.getById(product.brand)

        // Create empty stores list
        List<SellingStoreEntity> stores = new ArrayList<>()

        List<IngredientEntity> ingredientList = ingredientService.manageIngredients(product.ingredients)

        // 2. Verify if product starts with a discount
        DiscountEntity discount = product.discount ? discounts.findById(product.discount) : null
        CategoryEntity category = product.category ? categories.getById(product.category) : null
        String productImage = product.image ? imagesClientService.create(product.image) : null
        // 6. Create brand using the request's body and content created
        ProductEntity productCreated = products.create(product, productImage, brand, stores, category, discount, ingredientList)

        // For each selling store, create them
        product.sellingStores.each {
            stores << sellingStores.create(it, productCreated)
        }

        ingredientList.each {
            ingredientService.addProductToIngredient(productCreated, it)
        }

        // 7. Return final response
        HttpResponse.status(HttpStatus.CREATED).body([id: productCreated.id])
    }

    @Put("/{id}")
    MutableHttpResponse overwrite(@NonNull @NotBlank UUID id, @Body @Valid Product product) {
        HttpResponse.status(HttpStatus.SERVICE_UNAVAILABLE)
    }

    @Patch("/{id}")
    MutableHttpResponse partialUpdate(@NonNull @NotBlank UUID id, @Body @Valid Product product) {

        // Get current product
        ProductEntity currentProduct = products.getById(id)

        // Create empty stores list
        List<SellingStoreEntity> stores = new ArrayList<>()

        // If sellingStores is in the payload, delete all and create new ones (if is not an empty list)
        if (product.sellingStores) {
            List<SellingStoreEntity> oldStores = sellingStoreRepository.findByProduct(currentProduct)

            oldStores.each {
                sellingStores.delete(it.id)
            }

            // For each new selling store, create them
            if (product.sellingStores != []) {
                product.sellingStores.each {
                    stores << sellingStores.create(it, currentProduct)
                }
            }
        }

        BrandEntity brand = product.brand ? brands.getById(product.brand) : null
        CategoryEntity category = product.category ? categories.getById(product.category) : null

        List<IngredientEntity> productIngredientsList = []
        currentProduct.ingredients.each {
            if (product.ingredients.find { iti -> iti.action != "delete" && iti.action != "create" && UUID.fromString(iti.content) == it.id }) {
                productIngredientsList.add(it)
            }
        }

        List<IngredientEntity> ingredientsUpdatedList = ingredientService.manageIngredients(product.ingredients, currentProduct)

        if (ingredientsUpdatedList.size() > 0) {
            ingredientsUpdatedList.addAll(productIngredientsList)
            ingredientsUpdatedList.unique()
        }

        ingredientsUpdatedList.each {
            ingredientService.addProductToIngredient(currentProduct, it)
        }

        // 1. Verify if the image is new
        String productImage = null

        if (product.image) {
            productImage = imagesClientService.create(product.image)
        }

        HttpResponse.ok(products.partialUpdate(id, product, productImage, stores, ingredientsUpdatedList, brand, category))
    }

    @Delete("/{id}")
    MutableHttpResponse delete(@NonNull @NotBlank UUID id) {
        HttpResponse.ok(products.delete(id))
    }

    @Get("/{id}/reviews")
    MutableHttpResponse getReviews(@NonNull @NotBlank UUID id, Pageable pageable) {
        ProductEntity product = products.getById(id)
        HttpResponse.ok(reviews.findAll(product, pageable))
    }

    @Post("/{id}/reviews")
    MutableHttpResponse addReview(@NonNull @NotBlank UUID id, @Body @Valid Review payload) {

        // Create the rating with the value received
        RatingEntity rating = ratings.create(payload.amount)

        // 2. Create a new review
        ProductEntity product = products.getById(id)
        reviews.create(payload, rating, product)

        // Update the rating in the product
        Long currentRating = product.rating ? product.rating.amount : 0
        Long newRating = (Long) (((currentRating * product.reviews.size()) + rating.amount) / (product.reviews.size() + 1))

        if (product.rating == null) {
            RatingEntity globalRating = ratings.create(newRating)
            product.rating = globalRating
            productRepository.update(product)
        } else {
            product.rating.amount = newRating
            ratingRepository.update(product.rating)
        }

        // 3. Return final response
        HttpResponse.status(HttpStatus.CREATED)
    }

    @Get("/{id}/ingredients")
    MutableHttpResponse getIngredients(@NonNull @NotBlank UUID id, Pageable pageable) {
        ProductEntity product = products.getById(id)
        HttpResponse.ok(ingredientService.findAll(product, pageable))
    }

    @Post("/{id}/ingredients")
    MutableHttpResponse addIngredient(@NonNull @NotBlank UUID id, @NonNull @NotBlank(message = "Ingredient must have a name") UUID ingredient) {
        // 1. Create new ingredient
        IngredientEntity entity = ingredientService.findById(ingredient)

        // 3. Create new relation
        products.addIngredient(id, entity)

        // 3. Return final response
        HttpResponse.status(HttpStatus.CREATED)
    }

    @Delete("/{id}/ingredients/{ingredientId}")
    MutableHttpResponse deleteIngredient(@NonNull @NotBlank UUID id, @NonNull @NotBlank UUID ingredientId) {
        IngredientEntity ingredient = ingredientService.findById(ingredientId)

        HttpResponse.ok(products.deleteIngredient(id, ingredient))
    }

    @Post("searches")
    MutableHttpResponse productSearch(@Body @Valid ProductSearchDTO payload, Pageable pageable) {

        List<IngredientEntity> ingredients = []
        if (payload.popularIngredients) {
            payload.popularIngredients.each {
                ingredients.add(ingredientService.findById(it))
            }
        }

        HttpResponse.ok(products.productSearch(payload, ingredients, pageable))
    }

    @Post("/{id}/extras")
    MutableHttpResponse addExtras(@NotBlank UUID id, @Body @Valid ProductExtrasDTO extras ) {
        HttpResponse.ok(productExtrasService.create(id, extras.extras))
    }

    @Delete("/{id}/extras")
    MutableHttpResponse deleteExtraRelation(@NotBlank UUID id, @Body @Valid ProductExtrasDTO extras ) {
        HttpResponse.ok(productExtrasService.delete(id, extras.extras))
    }

}


