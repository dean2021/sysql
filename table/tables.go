package table

func ColumnDefinition(columns TableColumns) (statement string) {
	var kDisableRowId = "WITHOUT ROWID"
	var indexed bool
	var epilog = make(map[string]bool)
	var pKeys []string
	statement = "("

	var i = 0
	for ; i < len(columns); i++ {
		column := columns[i]
		statement +=
			"`" + column.Name + "` " + ColumnTypeNames[column.Type]
		if column.Options == INDEX {
			indexed = true
		}
		if (column.Options & (INDEX | ADDITIONAL)) == 1 {
			pKeys = append(pKeys, column.Name)
			epilog[kDisableRowId] = true
		}
		if column.Options == HIDDEN {
			statement += " HIDDEN"
		}
		if i < len(columns)-1 {
			statement += ", "
		}
	}

	// If there are only 'additional' columns (rare), pkey is the 'unique row'.
	// Otherwise an additional constraint will create duplicate rowids.
	if !indexed && epilog[kDisableRowId] {
		for _, column := range columns {
			pKeys = append(pKeys, column.Name)
		}
	}

	// Append the primary keys, if any were defined.
	if len(pKeys) > 0 {

		statement += ", PRIMARY KEY ("
		for i, pkey := range pKeys {
			statement += "`" + pkey + "`"
			if i < len(pKeys)-1 {
				statement += ", "
			}
		}
		statement += ")"
	}

	statement += ")"
	//for (auto& ei : epilog) {
	//	if (ei.second) {
	//		statement += ' ' + std::move(ei.first);
	//	}
	//}

	return statement
}
