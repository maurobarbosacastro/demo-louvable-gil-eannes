package pt.atlanse.eduquiz

import io.micronaut.context.event.ApplicationEventListener
import io.micronaut.liquibase.LiquibaseConfigurationProperties
import io.micronaut.liquibase.LiquibaseMigrator
import io.micronaut.runtime.server.event.ServerStartupEvent
import jakarta.inject.Inject
import jakarta.inject.Singleton
import jakarta.transaction.Transactional
import liquibase.exception.LiquibaseException

import javax.sql.DataSource

@Transactional
@Singleton
class Bootstrap implements ApplicationEventListener<ServerStartupEvent> {

    @Inject
    LiquibaseMigrator liquibaseMigrator

    LiquibaseConfigurationProperties configurationProperties

    @Inject
    DataSource dataSource

    Bootstrap(LiquibaseConfigurationProperties configurationProperties) {
        this.configurationProperties = configurationProperties
    }

    /**
     * Handle an application event.
     *
     * @param event the event to respond to
     */
    @Override
    void onApplicationEvent(ServerStartupEvent event) {
        try {

            // Update Liquibase manually
            liquibaseMigrator.run(configurationProperties, dataSource)
        } catch (LiquibaseException e) {
            // Handle Liquibase exception
            e.printStackTrace()
        }
    }
}
