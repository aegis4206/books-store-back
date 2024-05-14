create table carts (
	id uuid primary key,
	-- id varchar(100) primary key,
	total_count int not null,
	total_amount int not null,
	user_id int not null,
	foreign key (user_id) references users(id)
);
create table cart_items(
	id serial primary key,
	count int not null,
	amount int not null,
	book_id int not null,
	cart_id uuid not null,
	-- cart_id varchar(100) not null,
	foreign key(book_id) references books(id),
	foreign key(cart_id) references carts(id)
);

create table orders(
	id uuid primary key,
	create_time timestamp with time zone default current_timestamp,
	total_count int not null,
	total_amount int not null,
	state int not null,
	user_id int ,
	foreign key (user_id) references users(id)
);

create table order_items(
	id serial primary key,
	count int not null,
	amount int not null,
	title varchar(100) not null,
	author varchar(100) not null,
	price int,
	imgPath varchar(100) not null,
	order_Id uuid not null,
	foreign key (order_Id) references orders(id)
);