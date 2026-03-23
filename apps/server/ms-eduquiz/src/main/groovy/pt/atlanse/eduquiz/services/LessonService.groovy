package pt.atlanse.eduquiz.services

import io.micronaut.data.model.Page
import io.micronaut.data.model.Pageable
import jakarta.inject.Inject
import pt.atlanse.eduquiz.DTO.LessonDTO
import pt.atlanse.eduquiz.domain.*
import pt.atlanse.eduquiz.models.Lesson
import pt.atlanse.eduquiz.models.Media
import pt.atlanse.eduquiz.utils.ExceptionService
import jakarta.inject.Singleton

@Singleton
class LessonService {

    @Inject
    FileHandler files

    @Inject
    ElearningBeans elearning

    @Inject
    ImagesClientService imagesClientService

    private Lesson parse(LessonEntity lesson) {
        List<Media> parsedContent = new ArrayList<>()

        return new Lesson(
            id: lesson.id,
            image: lesson.image,
            title: lesson.title,
            subtitle: lesson.subtitle,
            type: lesson.type,
            conclusion: lesson.conclusion,
            status: lesson.status,
            createdBy: lesson.createdBy,
            createdAt: lesson.createdAt,
            updatedBy: lesson.updatedBy,
            updatedAt: lesson.updatedAt
        )
    }

    Lesson findById(String id) {
        LessonEntity lesson = elearning.lessons
            .findById(UUID.fromString(id))
            .orElseThrow(ExceptionService::LessonNotFoundException)
        return parse(lesson)
    }

    Page<Lesson> findAll(Pageable pageable) {
        elearning.lessons.findAll(pageable).map(this::parse)
    }

    Page<Lesson> findAllByModule(String moduleId, Pageable pageable) {
        ModulesEntity module = elearning.modules.findById(UUID.fromString(moduleId))
            .orElseThrow(ExceptionService::ModuleNotFoundException)
        Page<ModulesOrderEntity> modulesOrderEntities = elearning.modulesOrders.findAllByModule(module, pageable)
        modulesOrderEntities.map {
            return parse(it.lesson)
        }
    }

    void create(LessonDTO payload, String createdBy) {
        LessonEntity lesson = new LessonEntity(
            title: payload.title,
            subtitle: payload.subtitle,
            type: payload.type,
            status: payload.status,
            conclusion: payload.conclusion,
            createdBy: createdBy,
            updatedBy: createdBy,
            image: imagesClientService.create(payload.content)
        )

        elearning.lessons.save(lesson)
    }

    void addContentById(String lessonId, ContentEntity content, String updatedBy) {
        LessonEntity lesson = elearning.lessons.findById(UUID.fromString(lessonId))
            .orElseThrow(ExceptionService::LessonNotFoundException)

        lesson.updatedBy = updatedBy

        elearning.lessons.update(lesson)

        elearning.lessonContents.save(new LessonContentEntity(
            lesson: lesson,
            content: content,
            updatedBy: updatedBy,
            createdBy: updatedBy
        ))

    }

    // todo [LESSONS] Clean
    LessonEntity update(String id, LessonDTO payload, String updatedBy) {
        LessonEntity lesson = elearning.lessons.findById(UUID.fromString(id))
            .orElseThrow(ExceptionService::LessonNotFoundException)

        if (payload.title) {
            lesson.title = payload.title
        }

        if (payload.content) {
            lesson.image = imagesClientService.create(payload.content)
        }

        if (payload.subtitle) {
            lesson.subtitle = payload.subtitle
        }

        if (payload.conclusion) {
            lesson.conclusion = payload.conclusion
        }

        if (payload.type) {
            lesson.type = payload.type
        }

        if (payload.status) {
            lesson.status = payload.status
        }

        lesson.updatedBy = updatedBy

        elearning.lessons.update(lesson)
    }

    void delete(String id) {
        LessonEntity lesson = elearning.lessons.findById(UUID.fromString(id))
            .orElseThrow(ExceptionService::LessonNotFoundException)
        elearning.lessons.delete(lesson)
    }

}
