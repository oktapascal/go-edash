package utils

import "database/sql"

// CommitRollback attempts to commit the transaction and rolls it back in case of an error.
// It recovers from panics and handles the rollback and commit operations accordingly.
// The function takes a transaction object as input.
func CommitRollback(tx *sql.Tx) {
	// Attempt to recover from panics
	err := recover()
	if err != nil {
		// Rollback the transaction if an error occurred during recovery
		errRollback := tx.Rollback()
		if errRollback != nil {
			// Panic if rolling back the transaction failed
			panic(errRollback)
		}
		// Panic with the original error
		panic(err)
	} else {
		// Commit the transaction
		errCommit := tx.Commit()
		if errCommit != nil {
			// Panic if committing the transaction failed
			panic(errCommit)
		}
	}
}
