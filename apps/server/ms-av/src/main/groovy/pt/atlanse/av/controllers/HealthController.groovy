package pt.atlanse.av.controllers

import io.micronaut.http.HttpResponse
import io.micronaut.http.MutableHttpResponse
import io.micronaut.http.annotation.Controller
import io.micronaut.http.annotation.Get

@Controller("/api/health")
class HealthController {
	@Get
    MutableHttpResponse status(){
		HttpResponse.ok()
	}
}
