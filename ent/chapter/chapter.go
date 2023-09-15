// Code generated by ent, DO NOT EDIT.

package chapter

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the chapter type in the database.
	Label = "chapter"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldReleasedTime holds the string denoting the released_time field in the database.
	FieldReleasedTime = "released_time"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldSort holds the string denoting the sort field in the database.
	FieldSort = "sort"
	// FieldHasFreePreview holds the string denoting the has_free_preview field in the database.
	FieldHasFreePreview = "has_free_preview"
	// FieldCourseID holds the string denoting the course_id field in the database.
	FieldCourseID = "course_id"
	// EdgeLessons holds the string denoting the lessons edge name in mutations.
	EdgeLessons = "lessons"
	// EdgeCourse holds the string denoting the course edge name in mutations.
	EdgeCourse = "course"
	// Table holds the table name of the chapter in the database.
	Table = "chapters"
	// LessonsTable is the table that holds the lessons relation/edge.
	LessonsTable = "lessons"
	// LessonsInverseTable is the table name for the Lesson entity.
	// It exists in this package in order to avoid circular dependency with the "lesson" package.
	LessonsInverseTable = "lessons"
	// LessonsColumn is the table column denoting the lessons relation/edge.
	LessonsColumn = "chapter_id"
	// CourseTable is the table that holds the course relation/edge.
	CourseTable = "chapters"
	// CourseInverseTable is the table name for the Course entity.
	// It exists in this package in order to avoid circular dependency with the "course" package.
	CourseInverseTable = "courses"
	// CourseColumn is the table column denoting the course relation/edge.
	CourseColumn = "course_id"
)

// Columns holds all SQL columns for chapter fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldReleasedTime,
	FieldDescription,
	FieldSort,
	FieldHasFreePreview,
	FieldCourseID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultReleasedTime holds the default value on creation for the "released_time" field.
	DefaultReleasedTime func() time.Time
	// UpdateDefaultReleasedTime holds the default value on update for the "released_time" field.
	UpdateDefaultReleasedTime func() time.Time
	// DefaultHasFreePreview holds the default value on creation for the "has_free_preview" field.
	DefaultHasFreePreview int
)

// OrderOption defines the ordering options for the Chapter queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByReleasedTime orders the results by the released_time field.
func ByReleasedTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldReleasedTime, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// BySort orders the results by the sort field.
func BySort(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSort, opts...).ToFunc()
}

// ByHasFreePreview orders the results by the has_free_preview field.
func ByHasFreePreview(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHasFreePreview, opts...).ToFunc()
}

// ByCourseID orders the results by the course_id field.
func ByCourseID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCourseID, opts...).ToFunc()
}

// ByLessonsCount orders the results by lessons count.
func ByLessonsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newLessonsStep(), opts...)
	}
}

// ByLessons orders the results by lessons terms.
func ByLessons(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newLessonsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByCourseField orders the results by course field.
func ByCourseField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCourseStep(), sql.OrderByField(field, opts...))
	}
}
func newLessonsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(LessonsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, LessonsTable, LessonsColumn),
	)
}
func newCourseStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CourseInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, CourseTable, CourseColumn),
	)
}
