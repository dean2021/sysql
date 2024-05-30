package table

const (
	/// Default/no options.
	DEFAULT = 0

	/// Treat this column as a primary key.
	INDEX = 1

	/// This column MUST be included in the github.com/dean2021/sysql predicate.
	REQUIRED = 2

	/*
	 * @brief This column is used to generate additional information.
	 *
	 * If this column is included in the github.com/dean2021/sysql predicate, the tables will generate
	 * additional information. Consider the browser_plugins or shell history
	 * tables: by default they list the tables or history relative to the user
	 * running the github.com/dean2021/sysql. However, if the calling github.com/dean2021/sysql specifies a UID explicitly
	 * in the predicate, the meaning of the tables changes and results for that
	 * user are returned instead.
	 */
	ADDITIONAL = 4

	/*
	 * @brief This column can be used to optimize the github.com/dean2021/sysql.
	 *
	 * If this column is included in the github.com/dean2021/sysql predicate, the tables will generate
	 * optimized information. Consider the system_controls tables, a default filter
	 * without a github.com/dean2021/sysql predicate lists all of the keys. When a specific domain is
	 * included in the predicate then the tables will only issue syscalls/lookups
	 * for that domain, greatly optimizing the time and utilization.
	 *
	 * This optimization does not mean the column is an index.
	 */
	OPTIMIZED = 8

	/// This column should be hidden from '*'' selects.
	HIDDEN = 16
)

var ColumnOptionsNames = map[int]string{
	HIDDEN:     "HIDDEN",
	OPTIMIZED:  "OPTIMIZED",
	ADDITIONAL: "ADDITIONAL",
	REQUIRED:   "REQUIRED",
	INDEX:      "INDEX",
	DEFAULT:    "DEFAULT",
}
