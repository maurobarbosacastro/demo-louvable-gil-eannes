package pagination

import (
	"fmt"
	"gorm.io/gorm"
	"math"
	"strings"
)

// Paginate - A utility function to apply pagination, sorting, and filtering to a GORM query.
//
// Usage:
// - Use this function with a GORM `db` query to paginate results, sort them, and apply filters.
// - Sorting syntax: Use "fieldName asc" or "fieldName desc".
//   Example: `sort=name desc` in a URL query will sort by the "name" field in descending order.
// - Pagination defaults: If no values are provided, the service defaults to a limit of 10 items per page and starts from page 1.
//
// Parameters:
// - `value`: The model or entity type to query.
// - `pagination`: A pointer to a `PaginationResult` struct that contains pagination details like page, limit, sort, etc.
// - `filters`: A dynamic filter (can be any type) used to filter the query results.
// - `db`: The GORM database instance used for querying.
//
// Behavior:
// - If `pagination.Page` is greater than 0, the function calculates the total rows and total pages.
// - The returned GORM scope function applies pagination, sorting, and filtering to the query when executed.
// - If `pagination.Page` is 0, all records are returned without pagination.
//
// Returns:
// - A GORM scope function that can be chained with other query methods.

func Paginate(value interface{}, pagination *PaginationResult, filters interface{}, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	db.Model(value).Where(filters).Count(&totalRows)

	pagination.TotalRows = totalRows
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))
	pagination.TotalPages = totalPages

	originTable := valueTableName(value, db)

	return func(db *gorm.DB) *gorm.DB {

		if strings.Contains(pagination.GetSort(), ".") {
			// Split the sort string into relation and field Example: store.name asc -> store, name asc
			parts := strings.Split(pagination.GetSort(), ".")
			relation := parts[0]                               // Relation name (Table name) Example: store
			field := parts[1]                                  // Field name (Column name) Example: name asc
			relationColumn := fmt.Sprintf("%s_uuid", relation) // Relation column name Example: store_uuid

			//Join to the relation table
			db = db.Joins(fmt.Sprintf("inner join %s on %s.%s = %s.%s", relation, relation, "uuid", originTable, relationColumn))
			//Join example: inner join store on store.uuid = transaction.store_uuid

			//Order by the field in the relation table
			db = db.Order(fmt.Sprintf("%s.%s", relation, field))
			//Order example: order by store.name asc
		}

		if pagination.Page == 0 {
			return db
		} else {
			// Apply pagination
			return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort()).Where(filters)
		}
	}
}

func valueTableName(value interface{}, db *gorm.DB) string {
	stmt := &gorm.Statement{DB: db}
	_ = stmt.Parse(value)
	return stmt.Schema.Table
}
