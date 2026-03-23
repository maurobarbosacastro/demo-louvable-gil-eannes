package pt.atlanse.eduquiz.services

import pt.atlanse.eduquiz.repositories.CourseOrderRepository
import pt.atlanse.eduquiz.repositories.CourseRepository

import pt.atlanse.eduquiz.repositories.LessonContentRepository
import pt.atlanse.eduquiz.repositories.LessonRepository
import pt.atlanse.eduquiz.repositories.ModulesOrderRepository
import pt.atlanse.eduquiz.repositories.ModulesRepository
import jakarta.inject.Singleton

/**
 * @deprecated
 * */
@Singleton
class ElearningBeans {

    CourseRepository courses
    CourseOrderRepository courseOrders

    ModulesRepository modules
    ModulesOrderRepository modulesOrders

    LessonRepository lessons

    LessonContentRepository lessonContents

    ElearningBeans(CourseRepository courses,
                   CourseOrderRepository courseOrders,
                   ModulesRepository modules,
                   ModulesOrderRepository modulesOrders,
                   LessonContentRepository lessonContents,
                   LessonRepository lessons) {
        this.courses = courses
        this.courseOrders = courseOrders
        this.modules = modules
        this.modulesOrders = modulesOrders
        this.courseOrders = courseOrders
        this.lessons = lessons
        this.lessonContents = lessonContents
    }
}
