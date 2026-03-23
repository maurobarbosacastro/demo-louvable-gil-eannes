package pt.atlanse.eduquiz.services

import groovy.util.logging.Slf4j
import jakarta.inject.Singleton
import pt.atlanse.eduquiz.domain.ModulesOrderEntity
import pt.atlanse.eduquiz.repositories.ModulesOrderRepository
import pt.atlanse.eduquiz.utils.ExceptionService

@Slf4j
@Singleton
class ModulesOrderService {

    ModulesOrderRepository modulesOrderRepository

    ModulesOrderService(ModulesOrderRepository modulesOrderRepository) {
        this.modulesOrderRepository = modulesOrderRepository
    }

    ModulesOrderEntity findById(String id) {
        modulesOrderRepository.findById(UUID.fromString(id))
            .orElseThrow(ExceptionService::LessonNotFoundException)
    }

    /**
     * Deleting a module order by the id.
     * @param Long id of the module order to delete
     * */
    void delete(String id) {
        ModulesOrderEntity modulesOrder = findById(id)
        modulesOrderRepository.delete(modulesOrder)
    }
}
