package pt.atlanse.eduquiz.services

import groovy.util.logging.Slf4j
import io.micronaut.data.model.Page
import io.micronaut.data.model.Pageable
import jakarta.inject.Inject
import jakarta.persistence.criteria.Join
import jakarta.persistence.criteria.Path
import pt.atlanse.eduquiz.DTO.ModulesDTO
import pt.atlanse.eduquiz.DTO.ModulesParams
import pt.atlanse.eduquiz.domain.*
import pt.atlanse.eduquiz.domain.compositeIds.ModuleCategoryEntity
import pt.atlanse.eduquiz.repositories.*
import pt.atlanse.eduquiz.utils.ExceptionService

import jakarta.persistence.EntityManager
import jakarta.persistence.PersistenceContext
import jakarta.persistence.TypedQuery
import jakarta.persistence.criteria.CriteriaBuilder
import jakarta.persistence.criteria.CriteriaQuery
import jakarta.persistence.criteria.Predicate
import jakarta.persistence.criteria.Root
import jakarta.transaction.Transactional
import jakarta.inject.Singleton

@Slf4j
@Singleton
@Transactional
class ModulesService {

    ModulesRepository modulesRepository
    ModulesOrderRepository modulesOrderRepository
    CourseOrderRepository courseOrderRepository
    CategoryRepository categoryRepository
    ModuleCategoryRepository moduleCategoryRepository

    @PersistenceContext
    EntityManager entityManager

    @Inject
    ImagesClientService imagesClientService

    ModulesService(ModulesRepository modulesRepository,
                   ModulesOrderRepository modulesOrderRepository,
                   CourseOrderRepository courseOrderRepository,
                   CategoryRepository categoryRepository,
                   ModuleCategoryRepository moduleCategoryRepository

    ) {
        this.modulesRepository = modulesRepository
        this.modulesOrderRepository = modulesOrderRepository
        this.courseOrderRepository = courseOrderRepository
        this.categoryRepository = categoryRepository
        this.moduleCategoryRepository = moduleCategoryRepository
    }

    /**
     * Find a single {@link ModulesEntity} by Id
     * @param Long id Allows usage of filters
     * @return An entire {@link ModulesEntity}
     * */
    ModulesEntity findById(String id) {
        modulesRepository.findById(UUID.fromString(id)).orElseThrow {
            ExceptionService::ModuleNotFoundException()
        }
    }

    /**
     * Finds all {@link ModulesEntity}
     * @return List of {@link ModulesEntity}
     * */
    Page<ModulesEntity> findAll(ModulesParams params, Pageable pageable) {
        log.info "Using pageable arguments: Page_number: ${pageable.offset}; Amount_of_articles: ${pageable.size}"
        return applyFilters(params, pageable)
    }

    /**
     * Creating a new module.
     * @param {@link ModulesDTO}, String author that changed
     * */
    ModulesEntity create(ModulesDTO dto, String author) {
        log.info "Creating a module"
        // 1. Build and save an module

        log.debug "Creating module entity"
        ModulesEntity entity = new ModulesEntity(title: dto.title,
            status: dto.status,
            createdBy: author,
            updatedBy: author
        )

        if (dto.description) {
            entity.description = dto.description
        }

        log.debug "Creating image using ClientService"
        // 2. Create content object
        if (dto.image) {
            entity.image = imagesClientService.create(dto.image)
        }

        // 3. Manage Categories
        dto.categories.each {
            categoryRepository.findById(it.id)
                .ifPresent {
                    entity.categories << it
                }
        }

        // 4. Save entity
        log.debug "Saving Module to database with payload: ${entity.toString()}"
        modulesRepository.save(entity)
    }

    /**
     * Add a lesson to a module
     * @param Long moduleId, {@link LessonEntity}, String author that changed
     * */
    void addLessonToModule(String moduleId, LessonEntity lesson, String author) {
        log.info "Add a new lesson to a module"
        // 1. Find the module we are trying to add
        ModulesEntity module = modulesRepository.findById(UUID.fromString(moduleId))
            .orElseThrow(ExceptionService::ModuleNotFoundException)

        // 2. Create a new module order
        ModulesOrderEntity modulesOrder = new ModulesOrderEntity()
        modulesOrder.module = module
        modulesOrder.lesson = lesson
        modulesOrder.updatedBy = author
        modulesOrder.createdBy = author

        // 3. Count the module-order by lesson and module and add 1
        modulesOrder.position = modulesOrderRepository.countByLessonAndModule(lesson, module) + 1
        modulesOrderRepository.save(modulesOrder)
    }

    /**
     * Update a {@link ModulesEntity}.
     * @param {@link ModulesDTO}, String author that changed
     * @param String author that changed
     * @param Long id of the module to update
     * */
    ModulesEntity update(String id, ModulesDTO dto, String author) {
        // Find the module entity with the given ID
        ModulesEntity module = findById(id)
        // Update the module's image if it is present in the DTO
        if (dto.image) {
            // Create a new content entity for the image and associate it with the module
            module.image = imagesClientService.create(dto.image)
        }
        // Update the module's description if it is present in the DTO
        if (dto.description) {
            module.description = dto.description
        }
        // Update the module's title if it is present in the DTO
        if (dto.title) {
            module.title = dto.title
        }

        if (dto.status) {
            module.status = dto.status
        }

        List<CategoryEntity> categoriesToAdd = new ArrayList<>()
        Set<CategoryEntity> categoriesToRemove = new HashSet<>()
        // Process category changes if they are present in the DTO
        if (dto.categories) {
            // Get the current set of category IDs for the module
            Set<String> currentCategoryIds = module.categories*.id.toSet()
            // Loop through each category in the DTO
            dto.categories.each {
                CategoryEntity categoryEntity = categoryRepository.findById(it.id).orElseThrow(ExceptionService::CategoryNotFoundException)
                if ((it.action == 'add') && !currentCategoryIds.contains(categoryEntity.id)) {
                    categoriesToAdd.add(categoryEntity)
                } else if (it.action == 'delete') {
                    categoriesToRemove.add(categoryEntity)
                }
            }

            if (categoriesToAdd) {
                module.categories = categoriesToAdd
            }

            categoriesToRemove.each {
                module.categories.removeIf(cat -> it.id == cat.id)
                moduleCategoryRepository.delete(moduleCategoryRepository.findByModuleAndCategory(module, it))
            }
        }
        // Set the updatedBy field to the author's name
        module.updatedBy = author
        // Update the module entity in the database
        modulesRepository.update(module.with {
            if (categoriesToAdd.size()) {
                return it
            }
            it.categories = []
            return it
        })
        // Return the updated module entity
        return module
    }

    /**
     * Deleting the module by the id.
     * @param Long id of the module to delete
     * */
    void delete(String id) {
        ModulesEntity module = findById(id)

        if (module.categories) {
            // Delete associated rows in module_category table
            List<ModuleCategoryEntity> mce = moduleCategoryRepository.findAllByModule(module)
            moduleCategoryRepository.deleteAll(mce)
        }

        // Delete the ModulesEntity instance
        modulesRepository.delete(module)
    }

    /**
     * Build the pages of type {@link ModulesEntity}
     * @param params Other parameters. Class type {@link ModulesParams}. Allows usage of filters
     * @param pageable ({@link Pageable}) obtained from the parameters and used for the pagination (e.g., size and page)
     * @return An entire {@link Page} book made of type {@link ModulesEntity}
     * */
    @Transactional
    Page<ModulesEntity> applyFilters(ModulesParams params, Pageable pageable) {
        // 1. Init criteria searches
        CriteriaBuilder cb = entityManager.getCriteriaBuilder()
        CriteriaQuery<ModulesEntity> query = cb.createQuery(ModulesEntity.class)
        Root<ModulesEntity> root = query.from(ModulesEntity.class)
        CriteriaQuery<ModulesEntity> whereQuery = query.select(root)
        // 2. Create search predicates and group
        List<Predicate> predicates = new ArrayList<>()

        if (params.category) {
            Join<ModulesEntity, CategoryEntity> categoryJoin = root.joinList("categories", JoinType.LEFT)
            categoryRepository.findById(params.category).ifPresent(category -> {
                Path<UUID> categoryIdPath = categoryJoin.get("id")
                predicates.add(cb.equal(categoryIdPath, category.getId()))
            })
        }

        predicates << (params.description ? cb.like(cb.lower(root.get("description")), "%" + params.description.toLowerCase() + "%") : null)
        predicates << (params.title ? cb.like(cb.lower(root.get("title")), "%" + params.title.toLowerCase() + "%") : null)
        predicates << (params.status ? cb.like(cb.lower(root.get("status")), "%" + params.status.toLowerCase() + "%") : null)
        predicates.removeIf { !it }

        // 3. Create query using all the not null predicates from above
        whereQuery.where(predicates.toArray() as Predicate[])

        // 4. Create Query "Pagination"
        TypedQuery<ModulesEntity> typedQuery = entityManager.createQuery(whereQuery)
        typedQuery.setMaxResults(pageable.size)
        typedQuery.setFirstResult(pageable.number > 0 ? pageable.number * pageable.size : 0)

        // 5. Run query and parse results
        List<ModulesEntity> modulesEntities = typedQuery.getResultList()

        // 6. Return as Page
        return Page.of(modulesEntities, pageable, modulesRepository.count())
    }
}
