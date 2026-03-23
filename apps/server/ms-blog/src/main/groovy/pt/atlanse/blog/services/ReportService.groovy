package pt.atlanse.blog.services

import groovy.util.logging.Slf4j
import io.micronaut.data.model.Page
import io.micronaut.data.model.Pageable
import io.micronaut.http.HttpStatus
import jakarta.inject.Inject
import jakarta.inject.Singleton
import pt.atlanse.blog.domain.CommentEntity
import pt.atlanse.blog.domain.ReportEntity
import pt.atlanse.blog.models.Comment
import pt.atlanse.blog.models.CustomException
import pt.atlanse.blog.models.Likeable
import pt.atlanse.blog.models.Report
import pt.atlanse.blog.models.Reports
import pt.atlanse.blog.repository.ArticleRepository
import pt.atlanse.blog.repository.CommentRepository
import pt.atlanse.blog.repository.ReportRepository
import pt.atlanse.blog.repository.TranslationRepository

import java.security.Principal
import java.time.LocalDateTime
import java.time.format.DateTimeFormatter

@Slf4j
@Singleton
class ReportService {

    @Inject
    CommentService comments

    CommentRepository comment

    ReportRepository reports

    TranslationRepository translations

    ArticleRepository articles

    ReportService(ReportRepository reports, CommentRepository comment, ArticleRepository articles, TranslationRepository translations) {
        this.reports = reports
        this.comment = comment
        this.articles = articles
        this.translations = translations
    }

    private ReportEntity fillReportEntity(Likeable liked, String who) {
        ReportEntity report = new ReportEntity()
        report.setCreatedBy(who)
        report.setUpdatedBy(who)

        if (liked instanceof CommentEntity) {
            report.setComment(liked)
            report.setArticle((liked as CommentEntity).article)
            return report
        }

        if (reports.findByCommentAndCreatedBy(liked as CommentEntity, who)) {
            throw new CustomException(
                "User already reported this comment",
                "This comment has already 1 report from this user",
                HttpStatus.CONFLICT
            )
        }

        return report
    }

    Report add(Likeable liked, String reason, String status, String who) {
        log.info "Creating report for Likeable object: ${ liked.toString() }"

        ReportEntity report = fillReportEntity(liked, who)
        report.setReason(reason)
        report.setStatus(status)
        try {
            ReportEntity reportEntity = reports.save(report)
            Report reportResponse = new Report()
            reportResponse.id = reportEntity.getId()
            reportResponse.reason = reportEntity.getReason()
            reportResponse.status = reportEntity.getStatus()
            reportResponse.createdAt = reportEntity.getCreatedAt()
            return reportResponse

        } catch (Exception e) {
            log.error "Error while trying to create Report entity; Reason: ${ e.message }"
            throw new CustomException(
                "Error creating Report",
                "An error occured while creating the Report entity; More details about the error: ${ e.message }",
                HttpStatus.INTERNAL_SERVER_ERROR
            )
        }
    }

    Reports getReports(Pageable pages, String status, Principal principal) {
        log.info "Getting reports for comment"
        Page<ReportEntity> reportEntityList = reports.findAllByStatus(status, pages)

        List<Report> reportList = []
        reportEntityList.each {
            Report report = new Report()

            report.id = it.getId()
            report.createdBy = principal.name
            report.comment = buildComment(comment.findById(it.comment.id).get(), principal)
            if (it.comment.article) {
                report.articleName = translations.findByArticleIdAndLang(it.comment.article.id, 'en').get().title
            } else {
                CommentEntity com = comment.findById(it.comment.parent.id).get()
                report.articleName = translations.findByArticleIdAndLang(com.article.id, 'en').get().title
            }

            report.reason = it.getReason()
            report.status = it.getStatus()
            report.author = principal.name
            report.createdAt = it.getCreatedAt().format(DateTimeFormatter.ISO_LOCAL_DATE_TIME)
            report.updatedAt = it.getCreatedBy()
            report.context = getRelatedComments(it.comment.parent, it.comment.getCreatedAt(), principal)
            reportList.add(report)
        }

        return new Reports(
            pageNumber: pages.getNumber() + 1,
            totalPages: reportEntityList.totalPages,
            totalElements: reportEntityList.totalSize,
            reports: reportList,
        )

    }

    Comment buildComment(CommentEntity comment, Principal principal) {
        return new Comment(
            id: comment.id,
            author: principal.name,
            text: comment.text,
            createdAt: comment.createdAt.format(DateTimeFormatter.ISO_LOCAL_DATE_TIME),
            updatedAt: comment.updatedAt.format(DateTimeFormatter.ISO_LOCAL_DATE_TIME),
        )
    }

    List<Comment> getRelatedComments(CommentEntity parent, LocalDateTime date, Principal principal) {

        List<Comment> comments = []
        Comment comm
        if (parent) {
            CommentEntity com = comment.findById(parent.id).get()
            comm = buildComment(com, principal)

            comments.add(comm)

            (comment.findTop3ByParentAndHiddenFalseAndCreatedAtLessThanOrderByCreatedAtDesc(parent, date) as List<CommentEntity>).each {
                comm = buildComment(it, principal)
                comments.add(comm)
            }
            return comments.sort { it.id } as List<Comment>
        }

        (comment.findTop4ByParentIsNullAndCreatedAtLessThanOrderByCreatedAtDesc(date) as List<CommentEntity>).each {
            comm = buildComment(it, principal)
            comments.add(comm)
        }
        return comments.sort { it.id }
    }

    void updateReportStatus(long reportId, String status) {
        try {
            ReportEntity reportEntity = reports.findById(reportId).orElseThrow {
                new CustomException(
                    "report not found",
                    "report with id $reportId was not found",
                    HttpStatus.NOT_FOUND
                )
            }
            reportEntity.status = status

            //Status = Block set comment hidden
            if (status == "Blocked") {
                comments.setCommentHidden(reportEntity.comment.id)
            }
            reports.update(reportEntity)

        } catch (Exception e) {
            throw new CustomException(
                "Error updating Report",
                "An error occured while updating the Report entity; More details about the error: ${ e.message }",
                HttpStatus.INTERNAL_SERVER_ERROR
            )
        }

    }

}
