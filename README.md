## Телефонный список
Телефонный список — это REST API-сервис, написанный на Go, предназначенный для хранения и управления контактами. Проект использует базу данных PostgreSQL, где часть логики реализована на стороне БД через хранимые функции и процедуры для повышения производительности и гибкости.

🚀 Возможности
Добавление, редактирование и удаление контактов

Поиск и фильтрация

Использование REST API

Оптимизация запросов с помощью хранимых процедур и функций в PostgreSQL

🛠️ Технологии
Go

PostgreSQL

⚙️ Установка и запуск
1. Клонируйте репозиторий
bash
Копировать
Редактировать
git clone https://github.com/kitCHERNOV/phone-list.git
cd phone-list
2. Настройка переменных окружения
Создайте .env файл и укажите в нём:
```
DB_HOST=localhost
PASSWORD=1234
USER_NAME=postgres
DB_NAME=dbname
PORT=5432
```
3. Запуск
Через Go
bash
go run main.go

Некоторые вычисления выполняются на стороне PostgreSQL с использованием заранее определённых процедур и функций, что снижает нагрузку на API и ускоряет работу.
Процедура обновления данных:
```
DECLARE
	params integer[8];
BEGIN
	params[1] := insert_or_get_id_of_firstname(input_data[1]);
	params[2] := insert_or_get_id_of_lastname(input_data[2]);
	params[3] := insert_or_get_id_of_middlename(input_data[3]);
	params[4] := insert_or_get_id_of_street(input_data[4]);
	params[5] := insert_or_get_id_of_house(input_data[5]);
	params[6] := insert_or_get_id_of_building(input_data[6]);
	params[7] := insert_or_get_id_of_apartment(input_data[7]);
	params[8] := insert_or_get_id_of_phonenumber(input_data[8]);
	
	INSERT INTO main (first_name, last_name, middle_name, street, house, building, apartment, phone_number)
	VALUES (params[1], params[2], params[3], params[4], params[5], params[6], params[7], params[8]);
END;
```

Одна из используемых функций:
```
DECLARE
	result_id INT;
	exist BOOL;
BEGIN
	-- существует, есть ли такое значение
	SELECT EXISTS(SELECT 1 FROM apartment WHERE apartment_val = apartment_value)
	INTO exist;
	
	IF NOT exist THEN
		-- Если не существует, то делаем новую запись
		INSERT INTO apartment (apartment_val)
		VALUES (apartment_value)
		RETURNING id INTO result_id;
	ELSE 
		-- Если поле существует, получаем текущий id
		SELECT id INTO result_id FROM apartment WHERE apartment_val = apartment_value;
	END IF;
	
	-- Возвращаем id
	RETURN result_id;
END;
```



🧪 Тестирование
В планах — покрытие основных API-методов unit- и integration-тестами.
