package pt.atlanse.av.controllers

import groovy.util.logging.Slf4j
import io.micronaut.data.exceptions.EmptyResultException
import io.micronaut.http.HttpRequest
import io.micronaut.http.HttpResponse
import io.micronaut.http.MediaType
import io.micronaut.http.MutableHttpResponse
import io.micronaut.http.annotation.Body
import io.micronaut.http.annotation.Consumes
import io.micronaut.http.annotation.Controller
import io.micronaut.http.annotation.Get
import io.micronaut.http.annotation.Post
import io.micronaut.http.multipart.CompletedFileUpload
import jakarta.inject.Inject
import pt.atlanse.av.domains.FileScan
import pt.atlanse.av.models.FilePayload
import pt.atlanse.av.services.DefaultFileScanService

import java.nio.file.Files
import java.nio.file.Paths

@Slf4j
@Controller("/api/")
class ClamAVController {
    final String SCAN_FOLDER = "/home/scan"
    final String LOG_FOLDER = "/home/log"
    final Integer SCAN_TIMEOUT = 10000

    @Inject
    DefaultFileScanService fileScanService

    StringBuilder sout
    StringBuilder serr

    /**
     * @deprecated To be removed; Use clampScanV2 instead...
     */
    @Post("/scan")
    @Consumes(MediaType.MULTIPART_FORM_DATA)
    MutableHttpResponse clampScan(CompletedFileUpload f) {
        log.info("Preparing to scan file")
        /////////////////////////////// WRITE TEMP FILE
        // Create directory if not exists
        Files.createDirectories(Paths.get(SCAN_FOLDER))

        // Create file section
        File newFile = new File(SCAN_FOLDER + "/" + f.filename)
        FileOutputStream fos = new FileOutputStream(newFile)
        byte[] decoder = f.getBytes()
        fos.write(decoder)

        // TODO unused variable declaration
        String nameWithoutExtension = f.filename.replaceFirst("[.][^.]+\$", "")

        // Close/save new file
        fos.close()

        /////////////////////////////// ATTEMPT TO SCAN
        sout = new StringBuilder()
        serr = new StringBuilder()

        String command = "clamscan -r $SCAN_FOLDER/$f.filename"
        Process proc = command.execute()
        log.info("Running command:$command")

        proc.consumeProcessOutput(sout, serr)
//		proc.waitForOrKill(SCAN_TIMEOUT)
        proc.waitFor()

        if (serr) {
            log.error "Process ended with errors; Details:$serr"
        }

        HttpResponse.ok().body(message: sout)
    }

    @Post("/v2/scan")
    @Consumes(MediaType.APPLICATION_JSON)
    MutableHttpResponse clampScanV2(@Body FilePayload fileUpload) {
        log.info("Preparing to scan file")
        FileScan results
        try {
            new Thread().start {
                log.info("Preparing to scan file")
                fileScanService.save(fileUpload.fileName, fileUpload.base64)
            }

            return HttpResponse.ok()
        } catch (Exception e) {
            log.error "Scan ended with errors; Details:$e.message"
            return HttpResponse.badRequest().body(
                message: "System melting",
                reason: e.message
            )
        }
    }

    @Post("/analysis")
    @Consumes(MediaType.APPLICATION_JSON)
    MutableHttpResponse analysis(String... hashes) {
        try {
            def results = fileScanService.findByMd5(hashes)
            return HttpResponse.ok().body(results)
        } catch (EmptyResultException ignored) {
            log.error "Query came empty"
            return HttpResponse.ok()
        } catch (Exception e) {
            log.error "Scan ended with errors; Details:$e.message"
            return HttpResponse.badRequest().body(
                message: "System melting",
                reason: e.message
            )
        }
    }

    /**
     * @deprecated To be removed in the future
     * */
    @Get
    MutableHttpResponse clampHelp(HttpRequest request) {
        sout = new StringBuilder()
        serr = new StringBuilder()

        Process proc = 'clamscan -h'.execute()
        proc.consumeProcessOutput(sout, serr)
        proc.waitForOrKill(1000)

        log.info("out> $sout\\nerr> $serr")

        HttpResponse.ok().body(output: "$sout", err: "$serr")
    }
}
