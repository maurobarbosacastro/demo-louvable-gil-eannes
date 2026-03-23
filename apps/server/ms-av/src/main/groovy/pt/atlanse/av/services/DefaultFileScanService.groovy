package pt.atlanse.av.services

import groovy.util.logging.Slf4j
import io.micronaut.data.exceptions.EmptyResultException
import jakarta.inject.Singleton
import jakarta.transaction.Transactional
import pt.atlanse.av.domains.FileScan
import pt.atlanse.av.repositories.FileScanRepository


import java.nio.file.Files
import java.nio.file.Paths

@Slf4j
@Singleton
@Transactional
class DefaultFileScanService /*implements FileScanService*/ {
    private final String SCAN_FOLDER = "/home/scan"
    private final FileScanRepository fileScanRepository

    private StringBuilder sout
    private StringBuilder serr

    // Injected thru constructor
    DefaultFileScanService(FileScanRepository fileScanRepository) {
        this.fileScanRepository = fileScanRepository
    }

    /**
     *
     * @return Map with structure Hash -> Analysis
     * */
    Map<String, Map> findByMd5(String... hashes) {
        log.info("List of hashes: $hashes")
        // Start something that looks like a dictionary for HASH - ANALYSIS
        Map<String, String> analysisResults = new HashMap<>()

        // For each hash received, check if the results are available
        hashes.each { hash ->

            try {
                FileScan scan = this.fileScanRepository.findByMd5(hash)
                analysisResults.put(hash, scan as String)
            }
            catch (EmptyResultException ignored) {
            }

        }

        // return the analysis
        return analysisResults
    }

    void save(String fileName, byte[] base64) {
        fileName = fileName.replace(" ", "_")
        System.out.println "Filename : $fileName"

        File newFile = new File("$SCAN_FOLDER/$fileName")
        String command = "clamscan -r $SCAN_FOLDER/$fileName"
        sout = new StringBuilder()
        serr = new StringBuilder()
        FileOutputStream fos
        byte[] decoder

        /////////////////////////////// WRITE TEMP FILE
        // Create directory if not exists
        // Create file section
        try {
            Files.createDirectories(Paths.get(SCAN_FOLDER))
            fos = new FileOutputStream(newFile)
            decoder = Base64.getDecoder().decode(base64)
            fos.write(decoder)
            fos.close()
        } catch (Exception e) {
            throw new Exception("Error while creating temp file; Details:$e.message")
        }

        /////////////////////////////// ATTEMPT TO SCAN
        try {
            log.info("Running command:$command")
            Process proc = command.execute()
            proc.consumeProcessOutput(sout, serr)
            proc.waitFor()

            if (serr) {
                throw new Exception("Running shell command ended with errors; Details:$serr")
            }
        } catch (Exception e) {
            throw new Exception("Error while running ClamAV clamscan process; Reason: $e.message")
        }

        /////////////////////////////// SAVE OUTPUT TO DB
        try {
            log.info("Saving file $fileName")
            FileScan fileScan = new FileScan(
                name: fileName,
                md5: decoder.md5(),
                bytes: base64.size(),
                clamAv: sout.toString()
            )

            this.fileScanRepository.save(fileScan)
        } catch (Exception e) {
            throw new Exception("Error while appending results on database; Reason: $e.message")
        }

    }

}
