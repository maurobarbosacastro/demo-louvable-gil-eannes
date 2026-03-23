package pt.atlanse.av.repositories

import io.micronaut.data.mongodb.annotation.MongoRepository
import io.micronaut.data.repository.PageableRepository
import pt.atlanse.av.domains.FileScan

@MongoRepository
interface FileScanRepository extends PageableRepository<FileScan, String> {
    FileScan findByMd5(String md5)
}
