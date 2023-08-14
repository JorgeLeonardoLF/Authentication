package database

import "github.com/JorgeLeonardoLF/Authentication/models"

func (apiDbCfg *ApiDBConfig) CreateNewAccountRecord(newAccount models.NewAccount) (int, error) {
	/*write our SQL statement, This is Postgres specific fromatting*/
	sqlStatement := `
	INSERT INTO accounts (username,password)
	VALUES ($1, $2)`

	/*Our apiDbCfg.Connection (Aka a *sql.DB pointing to your data base) has:
	- .Exec("sqlStatement", optional parameters...)*/
	_, onInsertErr := apiDbCfg.Connection.Exec(sqlStatement, newAccount.Username, newAccount.Password)
	if onInsertErr != nil {
		return 400, onInsertErr
	}
	return 200, nil
}
