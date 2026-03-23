package pt.atlanse.eduquiz.services

import groovy.util.logging.Slf4j
import io.micronaut.data.model.Page
import io.micronaut.data.model.Pageable
import jakarta.inject.Singleton
import jakarta.persistence.EntityManager
import jakarta.persistence.PersistenceContext
import jakarta.persistence.TypedQuery
import jakarta.persistence.criteria.CriteriaBuilder
import jakarta.persistence.criteria.CriteriaQuery
import jakarta.persistence.criteria.Predicate
import jakarta.persistence.criteria.Root
import jakarta.transaction.Transactional
import pt.atlanse.eduquiz.domain.CategoryEntity
import pt.atlanse.eduquiz.repositories.CategoryRepository

@Slf4j
@Singleton
class CategoryService {
    CategoryRepository categoryRepository

    @PersistenceContext
    EntityManager entityManager

    CategoryService(CategoryRepository categoryRepository) {
        this.categoryRepository = categoryRepository
    }

    /**
     * Finds all {@link CategoryEntity}
     * @return List of {@link CategoryEntity}
     * */
    Page<CategoryEntity> findAll(Map name, Pageable pageable) {
        applyFilters(name, pageable)
    }

    @Transactional
    Page<CategoryEntity> applyFilters(Map map, Pageable pageable) {
        // 1. init criteria searches
        CriteriaBuilder cb = entityManager.getCriteriaBuilder()
        CriteriaQuery<CategoryEntity> query = cb.createQuery(CategoryEntity.class)
        Root<CategoryEntity> root = query.from(CategoryEntity.class)
        CriteriaQuery<CategoryEntity> whereQuery = query.select(root)

        // 2. Create search predicates and group
        List<Predicate> predicates = new ArrayList<>()

        predicates << (map ? cb.like(cb.lower(root.get("name")), "%" + map.name.toLowerCase() + "%") : null)
        predicates.removeIf { !it }

        // 3. Create query using all the not null predicates from above
        whereQuery.where(predicates.toArray() as Predicate[])

        // 4. Create query "pagination"
        TypedQuery<CategoryEntity> typedQuery = entityManager.createQuery(whereQuery)
        typedQuery.setMaxResults(pageable.size)
        typedQuery.setFirstResult(pageable.number > 0 ? pageable.number * pageable.size : 0)

        // 5. Run query and parse results
        List<CategoryEntity> categories = typedQuery.getResultList()

        return Page.of(categories, pageable, categoryRepository.count())
    }
}
