package pt.atlanse.av.services

import pt.atlanse.av.domains.FileScan

interface FileScanService {
	FileScan save(String fileName, byte[] base64)
//	Map<String, String> findByMd5(String ...hashes)
}
