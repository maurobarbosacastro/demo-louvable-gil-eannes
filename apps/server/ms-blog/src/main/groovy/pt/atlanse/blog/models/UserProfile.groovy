package pt.atlanse.blog.models

import groovy.transform.TupleConstructor
import pt.atlanse.blog.DTO.ImageDTO

import java.time.LocalDate

@TupleConstructor
class UserProfile {

    // KEYCLOAK ACCOUNT DETAILS
    String username

    String password

    // If username is null, use the email instead
    String email

    // CLIENT ACCOUNT DETAILS
    String firstName

    String lastName

    ImageDTO profilePicture

    String country

    // NON MANDATORY CLIENT FIELDS
    LocalDate birthDate

    String genre

    String relationShip

    String children

    String education

    String career

    String income

    String financialGoals

    String isBlocked
}
