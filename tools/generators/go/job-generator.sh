#!/bin/bash

# TagPeak Job Generator Script
# Creates job skeletons for Go microservices
# Usage: ./job-generator.sh [microservice] [job-name]

set -e

# Color codes for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Function to discover Go microservices dynamically
discover_go_microservices() {
    local -a services=()

    # Find directories in apps/server that contain go.mod files
    while IFS= read -r -d '' go_mod_file; do
        local dir
        local service_name
        dir=$(dirname "$go_mod_file")
        service_name=$(basename "$dir")
        services+=("$service_name")
    done < <(find apps/server -name "go.mod" -type f -print0 | sort -z)

    # Set global array
    GO_MICROSERVICES=("${services[@]}")

    if [ ${#GO_MICROSERVICES[@]} -eq 0 ]; then
        log_error "No Go microservices found in apps/server directory"
        exit 1
    fi
}

# Helper functions
log_info() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

log_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

log_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

log_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Function to convert job name to camelCase
to_camel_case() {
    local input="$1"
    # Replace hyphens and underscores with spaces, then convert to camelCase
    echo "$input" | sed 's/[-_]/ /g' | awk '{
        for (i=1; i<=NF; i++) {
            if (i==1) {
                printf "%s", tolower($i)
            } else {
                printf "%s", toupper(substr($i,1,1)) tolower(substr($i,2))
            }
        }
    }'
}

# Function to convert job name to UPPER_SNAKE_CASE
to_upper_snake_case() {
    local input="$1"
    # First handle camelCase by adding underscores before uppercase letters
    # Then replace hyphens and spaces with underscores
    # Finally convert to uppercase and clean up multiple underscores
    echo "$input" | sed 's/\([a-z]\)\([A-Z]\)/\1_\2/g' | sed 's/[-[:space:]]/_/g' | sed 's/_\+/_/g' | tr '[:lower:]' '[:upper:]'
}

# Function to show microservice selection menu
select_microservice() {
    # Discover microservices first
    discover_go_microservices

    echo -e "\n${BLUE}Available Go Microservices:${NC}"
    for i in "${!GO_MICROSERVICES[@]}"; do
        echo "  $((i + 1)). ${GO_MICROSERVICES[i]}"
    done

    while true; do
        echo -n -e "\n${YELLOW}Select a microservice (1-${#GO_MICROSERVICES[@]}):${NC} "
        read -r selection

        if [[ "$selection" =~ ^[0-9]+$ ]] && [ "$selection" -ge 1 ] && [ "$selection" -le "${#GO_MICROSERVICES[@]}" ]; then
            MICROSERVICE="${GO_MICROSERVICES[$((selection - 1))]}"
            break
        else
            log_error "Invalid selection. Please choose a number between 1 and ${#GO_MICROSERVICES[@]}."
        fi
    done
}

# Function to prompt for job name
prompt_job_name() {
    while true; do
        echo -n -e "\n${YELLOW}Enter job name (e.g., sync-orders, cleanup-data):${NC} "
        read -r job_name

        if [[ -n "$job_name" ]] && [[ "$job_name" =~ ^[a-zA-Z][a-zA-Z0-9_-]*$ ]]; then
            JOB_NAME="$job_name"
            break
        else
            log_error "Invalid job name. Use only letters, numbers, hyphens, and underscores. Must start with a letter."
        fi
    done
}

# Function to validate microservice exists
validate_microservice() {
    local ms="$1"
    local ms_path="apps/server/$ms"

    if [[ ! -d "$ms_path" ]]; then
        log_error "Microservice directory '$ms_path' does not exist."
        return 1
    fi

    if [[ ! -f "$ms_path/go.mod" ]]; then
        log_error "Directory '$ms_path' is not a Go microservice (missing go.mod)."
        return 1
    fi

    return 0
}

# Function to extract example function from template
extract_example_function() {
    local template_file="tools/generators/go/dependencies/jobs.go"
    local start_line
    local end_line

    # Find the start and end of the example function
    start_line=$(grep -n "^func example(" "$template_file" | cut -d: -f1)
    end_line=$(tail -n +"$start_line" "$template_file" | grep -n "^}" | head -1 | cut -d: -f1)
    end_line=$((start_line + end_line - 1))

    # Extract the function
    sed -n "${start_line},${end_line}p" "$template_file"
}

# Function to create jobs directory and copy template
create_job_skeleton() {
    local ms="$1"
    local job_name="$2"
    local ms_path="apps/server/$ms"
    local jobs_dir="$ms_path/jobs"
    local jobs_file="$jobs_dir/jobs.go"
    local template_file="tools/generators/go/dependencies/jobs.go"

    # Create jobs directory if it doesn't exist
    if [[ ! -d "$jobs_dir" ]]; then
        log_info "Creating jobs directory: $jobs_dir"
        mkdir -p "$jobs_dir"
    fi

    local camel_case_name
    local upper_snake_case_name
    camel_case_name=$(to_camel_case "$job_name")
    upper_snake_case_name=$(to_upper_snake_case "$job_name")

    log_info "Processing job with name: $job_name"
    log_info "  - Function name: $camel_case_name"
    log_info "  - Environment variable: ${upper_snake_case_name}_JOB"

    # Check if jobs.go already exists
    if [[ -f "$jobs_file" ]]; then
        log_warning "File $jobs_file already exists. Adding new job function."

        # Check if function already exists
        if grep -q "func $camel_case_name(" "$jobs_file"; then
            log_error "Function $camel_case_name already exists in $jobs_file"
            return 1
        fi

        # Extract the example function from template and customize it
        local new_function
        new_function=$(extract_example_function |
            sed "s/func example(/func $camel_case_name(/g" |
            sed "s/\"exampleJob Started\"/\"${camel_case_name}Job Started\"/g" |
            sed "s/\"exampleJob Ended\"/\"${camel_case_name}Job Ended\"/g" |
            sed "s/\"Error adding cron job exampleJob\"/\"Error adding cron job ${camel_case_name}Job\"/g" |
            sed "s/\"EXAMPLE_JOB\"/\"${upper_snake_case_name}_JOB\"/g")

        # Add new function to the end of the file
        echo "" >>"$jobs_file"
        echo "$new_function" >>"$jobs_file"

        # Add job call to StartJobs function
        local start_jobs_line
        start_jobs_line=$(grep -n "^func StartJobs()" "$jobs_file" | cut -d: -f1)
        local closing_brace_line
        closing_brace_line=$(tail -n +$((start_jobs_line + 1)) "$jobs_file" | grep -n "^}" | head -1 | cut -d: -f1)
        closing_brace_line=$((start_jobs_line + closing_brace_line))

        # Insert the new job call before the closing brace
        sed -i.tmp "${closing_brace_line}i\\
\\	go $camel_case_name(loc)" "$jobs_file"

        # Remove temporary file
        rm -f "${jobs_file}.tmp"

        log_success "Added new job function $camel_case_name to existing $jobs_file"
    else
        # Copy template and customize
        log_info "Creating new jobs.go from template"
        cp "$template_file" "$jobs_file"

        # Replace function name and references
        sed -i.tmp "s/func example(/func $camel_case_name(/g" "$jobs_file"
        sed -i.tmp "s/go example(/go $camel_case_name(/g" "$jobs_file"
        sed -i.tmp "s/\"exampleJob Started\"/\"${camel_case_name}Job Started\"/g" "$jobs_file"
        sed -i.tmp "s/\"exampleJob Ended\"/\"${camel_case_name}Job Ended\"/g" "$jobs_file"
        sed -i.tmp "s/\"Error adding cron job exampleJob\"/\"Error adding cron job ${camel_case_name}Job\"/g" "$jobs_file"
        sed -i.tmp "s/\"EXAMPLE_JOB\"/\"${upper_snake_case_name}_JOB\"/g" "$jobs_file"

        # Remove temporary file
        rm -f "${jobs_file}.tmp"

        log_success "Created new job skeleton: $jobs_file"
    fi
}

# Function to update main.go with jobs.StartJobs() call
update_main_go() {
    local ms="$1"
    local ms_path="apps/server/$ms"
    local main_file="$ms_path/main.go"

    if [[ ! -f "$main_file" ]]; then
        log_warning "main.go not found at $main_file. Skipping main.go update."
        return 0
    fi

    # Check if jobs.StartJobs() already exists
    if grep -q "jobs.StartJobs()" "$main_file"; then
        log_info "jobs.StartJobs() already exists in $main_file"
        return 0
    fi

    # Check if config.InitServer() exists
    if ! grep -q "config.InitServer()" "$main_file"; then
        log_warning "config.InitServer() not found in $main_file. Cannot determine where to add jobs.StartJobs()"
        return 0
    fi

    # Check if jobs import exists
    if ! grep -q "\"$ms/.*jobs\"" "$main_file"; then
        log_info "Adding jobs import to $main_file"

        # Find the import block and add jobs import
        local import_end_line
        import_end_line=$(grep -n "^)" "$main_file" | head -1 | cut -d: -f1)

        if [[ -n "$import_end_line" ]]; then
            # Add jobs import and newline before the closing parenthesis
            sed -i.tmp "${import_end_line}i\\
\\	\"$ms/jobs\"\\
" "$main_file"
            rm -f "${main_file}.tmp"
            log_success "Added jobs import to $main_file"
        fi
    fi

    # Find config.InitServer() line and add jobs.StartJobs() before it
    local config_line
    config_line=$(grep -n "config.InitServer()" "$main_file" | cut -d: -f1)

    if [[ -n "$config_line" ]]; then
        # Insert jobs.StartJobs() and two empty lines before config.InitServer()
        sed -i.tmp "${config_line}i\\
\\	jobs.StartJobs()\\
\\
" "$main_file"
        rm -f "${main_file}.tmp"
        log_success "Added jobs.StartJobs() call to $main_file"
    else
        log_error "Could not find config.InitServer() in $main_file"
        return 1
    fi
}

# Function to add cron dependency
add_cron_dependency() {
    local ms="$1"
    local ms_path="apps/server/$ms"

    if [[ ! -d "$ms_path" ]]; then
        log_warning "Microservice directory $ms_path not found. Skipping cron dependency addition."
        return 0
    fi

    log_info "Adding github.com/robfig/cron@v1.2.0 dependency to $ms"

    # Navigate to microservice directory and add dependency
    (
        cd "$ms_path" || {
            log_error "Failed to change to directory $ms_path"
            return 1
        }

        # Add the cron dependency with specific version
        if go get github.com/robfig/cron@v1.2.0; then
            log_success "Added github.com/robfig/cron@v1.2.0 dependency to $ms"
        else
            log_warning "Failed to add cron dependency. You may need to add it manually: go get github.com/robfig/cron@v1.2.0"
        fi
    )
}

# Function to add environment variable
add_env_variable() {
    local job_name="$1"
    local upper_snake_case_name
    upper_snake_case_name=$(to_upper_snake_case "$job_name")
    local env_var="${upper_snake_case_name}_JOB"
    local env_value="0 0 * * *"
    local env_file=".env"

    # Check if environment variable already exists
    if grep -q "^${env_var}=" "$env_file" 2>/dev/null; then
        log_warning "Environment variable $env_var already exists in $env_file"
        return 0
    fi

    # Add environment variable
    log_info "Adding environment variable to $env_file: $env_var=$env_value"
    echo "" >>"$env_file"
    echo "#JOB - $(echo "$job_name" | tr '[:lower:]' '[:upper:]')" >>"$env_file"
    echo "${env_var}=$env_value" >>"$env_file"

    log_success "Added environment variable: $env_var=$env_value"
}

# Main script logic
main() {
    log_info "TagPeak Job Generator"
    echo "======================="

    # Parse arguments
    MICROSERVICE="$1"
    JOB_NAME="$2"

    # Get microservice if not provided
    if [[ -z "$MICROSERVICE" ]]; then
        select_microservice
    else
        # Discover microservices first for validation
        discover_go_microservices

        # Validate provided microservice
        if ! validate_microservice "$MICROSERVICE"; then
            log_error "Available microservices:"
            printf "  %s\n" "${GO_MICROSERVICES[@]}"
            exit 1
        fi
    fi

    # Get job name if not provided
    if [[ -z "$JOB_NAME" ]]; then
        prompt_job_name
    else
        # Validate job name format
        if [[ ! "$JOB_NAME" =~ ^[a-zA-Z][a-zA-Z0-9_-]*$ ]]; then
            log_error "Invalid job name format. Use only letters, numbers, hyphens, and underscores. Must start with a letter."
            exit 1
        fi
    fi

    log_info "Creating job '$JOB_NAME' for microservice '$MICROSERVICE'"

    # Validate microservice directory exists
    if ! validate_microservice "$MICROSERVICE"; then
        exit 1
    fi

    # Create job skeleton
    create_job_skeleton "$MICROSERVICE" "$JOB_NAME"

    # Update main.go with jobs.StartJobs() call
    update_main_go "$MICROSERVICE"

    # Add cron dependency
    add_cron_dependency "$MICROSERVICE"

    # Add environment variable
    add_env_variable "$JOB_NAME"

    echo ""
    log_success "Job generator completed successfully!"
    log_info "Next steps:"
    echo "  1. Review the generated/updated file: apps/server/$MICROSERVICE/jobs/jobs.go"
    echo "  2. Check the updated main.go: apps/server/$MICROSERVICE/main.go"
    echo "  3. Fix any import statements (logster, dotenv) according to your project structure"
    echo "  4. Implement your job logic in the $(to_camel_case "$JOB_NAME") function"
    echo "  5. Update the environment variable $(to_upper_snake_case "$JOB_NAME")_JOB with your desired cron schedule"
    echo "  6. Test your job implementation"
}

# Run main function with all arguments
main "$@"
