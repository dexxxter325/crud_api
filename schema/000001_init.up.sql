CREATE TABLE products (
                           id serial NOT NULL UNIQUE,
                           name varchar(255) NOT NULL UNIQUE,
                           description varchar(255) NOT NULL UNIQUE

);

/*serial-уникальный id к каждому запросу
  varchar-строка,не превышающая 255 символов
  unique-знач.в столбцах уникальные(без повторов)*/
