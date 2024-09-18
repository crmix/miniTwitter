package constants

type Sentinel string

func (s Sentinel) Error() string {
	return string(s)
}

const (
	// PGForeignKeyViolationCode is used to check foriegn key violation in database
	PGForeignKeyViolationCode = "23503"
	// PGUniqueKeyViolationCode is used to check unique key violation in database
	PGUniqueKeyViolationCode = "23505"
)