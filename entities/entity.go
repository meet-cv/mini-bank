package entities

import "time"

type User struct {
	tableName           struct{}          `pg:"users"`
	User_Id             int               `pg:"user_id,pk" json:"user_id"`
	User_Name           string            `pg:"user_name,notnull" json:"user_name"`
	User_Address        map[string]string `pg:"user_address,type:json" json:"user_address"`
	User_Contact_Number string            `pg:"user_contact_number" json:"user_contact_number"`
	Created_At          time.Time         `pg:"created_at,notnull,deafult:now()" json:"created_at"`
	Updated_At          time.Time         `pg:"updated_at" json:"updated_at"`
}

type Account struct {
	Account_Id      int       `pg:"acc_id,pk" json:"acc_id"`
	Account_Number  int       `pg:"acc_number,notnull,unique" json:"acc_number"`
	Account_Type    string    `pg:"acc_type,notnull" json:"acc_type"`
	Account_Balance float64   `pg:"acc_balance,notnull" json:"acc_balance"`
	Created_At      time.Time `pg:"created_at,notnull" json:"created_at"`
	Updated_At      time.Time `pg:"updated_at" json:"updated_at"`
}

type AccountHolder struct {
	User_Id    int `pg:"user_id" json:"user_id"`
	Account_Id int `pg:"acc_id" json:"acc_id"`
}

type Transaction struct {
	Transaction_Id     int     `pg:"transac_id,pk" json:"transac_id"`
	Transaction_Type   string  `pg:"transac_type,notnull" json:"transac_type"`
	Transaction_Amount float64 `pg:"transac_amount,notnull" json:"transac_amount"`
	Transaction_Date   string  `pg:"transac_date,notnull" json:"transac_date"`
	User_Id            int     `pg:"user_id,notnull" json:"user_id"`
	Beneficiary_Id     int     `pg:"beneficiary_id,notnull" json:"beneficiary_id"`
}
