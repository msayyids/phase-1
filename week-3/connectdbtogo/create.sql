create table Employees (
    employees_id int primary key auto_increment,
    first_name varchar(50),
    last_name varchar(50),
    position varchar(50)
);

create table MenuItems (
    item_id int primary key auto_increment,
    name varchar(100),
    description varchar(255),
    price decimal(10,2),
    category varchar(50)
);

create table Orders (
    order_id int primary key auto_increment,
    label_number int,
    employees_id int,
    order_date date,
    status varchar(50),
    foreign key(employees_id) references Employees(employees_id)
);

create table OrderItems (
    order_item_id int primary key auto_increment,
    order_id int,
    item_id int,
    quantity int,
    subtotal decimal(10,2),
    foreign key (order_id) references Orders(order_id),
    foreign key(item_id) references MenuItems(item_id)
);

create table Payments(
    payment_id int primary key auto_increment,
    order_id int,
    payment_date date,
    payment_method varchar(50),
    total_amount decimal(10,2),
    foreign key (order_id) references Orders(order_id)

);


