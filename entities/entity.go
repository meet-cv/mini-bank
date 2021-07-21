package entities

import "time"

type User struct {
	tableName           struct{}               `pg:"users"`
	User_Id             int                    `pg:"user_id, type:serial, pk" json:"user_id" validate:"number"`
	User_Name           string                 `pg:"user_name, notnull, type:varchar(255)" json:"user_name" validate:"alpha"`
	User_Address        map[string]interface{} `pg:"user_address, type:json" json:"user_address" validate:"json"`
	User_Contact_Number string                 `pg:"user_contact_number,type:varchar(255)" json:"user_contact_number" validate:"alphanum"`
	Created_At          time.Time              `pg:"created_at, notnull, type:timestamp, default:now()" json:"created_at" validate:"datetime"`
	Updated_At          time.Time              `pg:"updated_at, type:timestamp" json:"updated_at" validate:"datetime"`
}

type Account struct {
	tableName       struct{}  `pg:"accounts"`
	Account_Id      int       `pg:"acc_id, pk, type:serial" json:"acc_id" validate:"number"`
	Account_Number  string    `pg:"acc_number, notnull, unique, type:varchar(255)" json:"acc_number" validate:"alphanum"`
	Account_Type    string    `pg:"acc_type, notnull, type:acc_type" json:"acc_type" validate:"alpha"`
	Account_Balance float64   `pg:"acc_balance, notnull, type:numeric(20,2), default:0" json:"acc_balance" validate:"numeric"`
	Created_At      time.Time `pg:"created_at, notnull, type:timestamp, default:now()" json:"created_at" validate:"datetime"`
	Updated_At      time.Time `pg:"updated_at, type:timestamp" json:"updated_at" validate:"datetime"`
}

type AccountHolder struct {
	tableName  struct{} `pg:"account_holders"`
	User_Id    int      `pg:"user_id, notnull, type:int" json:"user_id" validate:"number"`
	Account_Id int      `pg:"acc_id, notnull, type:int" json:"acc_id" validate:"number"`
}

type Transaction struct {
	tableName          struct{}  `pg:"transactions"`
	Transaction_Id     int       `pg:"transac_id, pk, type:serial" json:"transac_id" validate:"number"`
	Transaction_Type   string    `pg:"transac_type, notnull, type:transac_type" json:"transac_type" validate:"alpha"`
	Transaction_Amount float64   `pg:"transac_amount, notnull, type:numeric(20,2)" json:"transac_amount" validate:"numeric"`
	Transaction_Date   time.Time `pg:"transac_date, notnull, type:timestamp, default:now()" json:"transac_date" validate:"datetime"`
	User_Id            int       `pg:"user_id, notnull, type:int" json:"user_id" validate:"number"`
	Beneficiary_Id     int       `pg:"beneficiary_id, notnull, type:int" json:"beneficiary_id" validate:"number"`
}

type UserAccount struct {
	UserInfo    User    `json:"user_info"`
	AccountInfo Account `json:"account_info"`
}
