## DB Structure

### users 
    (
    user_id integer primary key, 
    user_name string not null, 
    user_address json, 
    user_contact_number string,
    created_at timestampz not null default current_timestamp,
    updated_at timestampz
    )

### accounts 
    (
    acc_id serial primary_key,
    acc_number integer unique, 
    acc_type enum not null, 
    acc_balance real not null check(acc_balance>=0) default 0, 
    created_at timestampz not null default current_timestamp,
    updated_at timestampz
    )

### acc_holder 
    (
    user_id integer not null,
    acc_id interger not null,
    primary key(user_id, acc_number),
    foreign key(user_id) references users on delete cascade, 
    foreign key(acc_id) references accounts on delete cascade, 
    )

### transactions 
    (
    transac_id integer primary_key, 
    user_id integer not null,
    beneficiary_id integer not null, 
    transac_type enum not null, 
    transac_amount real not null check(trasac_amount>=0), 
    transac_datetime timestampz not null default current_timestamp,
    foreign key(user_id) references users,
    foreign key(beneficiary_id) references users,
    )

## Domain
create type acc_type as ENUM('saving','joint','salary');
create type transac_type as ENUM('debit','credit');
