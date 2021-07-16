## DB Structure

### users 
    (
    user_id integer primary key, 
    user_name string not null, 
    user_address string, 
    user_contact_number integer
    )

### accounts 
    (
    acc_number integer primary_key, 
    acc_type string not null, 
    acc_balance real not null check(>=0), 
    )

### acc_holder (
    user_id integer,
    acc_number interger,
    primary key(user_id, acc_number) 
    )

### loan 
    (
    loan_number integer primary_key, 
    load_type string not null,
    loan_amount real not null check(>=0), 
    loan_balance real not null check(>=0), 
    user_id integer not null references users(user_id)
    )

### transactions 
    (
    transac_id integer primary_key, 
    transac_type string not null, 
    transac_amount real not null check(>=0), 
    transac_date date not null,
    user_id integer not null references users(user_id),
    beneficiary_id integer not null, 
    )

## Domain
**acc_type** 
    - ("saving","salary,","joint")
**transac_type**
    - ("debit","credit")
**loan_type domain** 
    - ("home","education","car")