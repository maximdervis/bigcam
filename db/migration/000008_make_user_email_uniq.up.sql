alter table users
add constraint email_uniq 
unique (email);
