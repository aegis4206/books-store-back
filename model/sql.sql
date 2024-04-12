create table carts (
	id varchar(100) primary key,
	total_count int not null,
	total_amount int not null,
	user_id int not null,
	foreign key (user_id) references users(id)
)
create table cart_items(
	id serial primary key,
	count int not null,
	amount int not null,
	book_id int not null,
	cart_id varchar(100) not null,
	foreign key(book_id) references books(id),
	foreign key(cart_id) references carts(id)
)