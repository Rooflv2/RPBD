<h3> 1. Выведите на экран любое сообщение </h3>

```plpgsql
SELECT 'LOL'
```
<h3> 2. Выведите на экран текущую дату </h3>

```plpgsql
SELECT NOW()
```

Или

```plpgsql
SELECT CURRENT_TIME
```

<h3> 3. Создайте две числовые переменные и присвойте им значение. Выполните математические действия с этими числами и выведите результат на экран. </h3>

```plpgsql
CREATE OR REPLACE FUNCTION math(x int, y int) RETURNS int AS $$
BEGIN
    RETURN x + y;
END
$$ LANGUAGE plpgsql
```

<h3>4. Написать программу двумя способами 1 - использование IF, 2 - использование CASE. Объявите числовую переменную и присвоейте ей значение. Если число равно 5 - выведите на экран "Отлично". 4 - "Хорошо". 3 - Удовлетворительно". 2 - "Неуд". В остальных случаях выведите на экран сообщение, что введённая оценка не верна.</h3>  

Вариант решения через конструкцию if:
```plpgsql
CREATE OR REPLACE FUNCTION num_if(x int) RETURNS char AS $$
BEGIN
    IF (x = 5) THEN RETURN 'Отлично';
	  ELSIF (x = 4) THEN RETURN 'Хорошо';
	  ELSIF (x = 3) THEN RETURN 'Удовлетворительно';
	  ELSIF (x = 2) THEN RETURN 'Не удовлетворительно';
	  ELSE RETURN 'Введите число от 2 до 5';
	  END IF;
END
$$ LANGUAGE plpgsql
```

Вариант решения через конструкцию switch case:

```plpgsql
CREATE OR REPLACE FUNCTION num_if(x int) RETURNS char AS $$
BEGIN
    CASE WHEN x = 5 THEN RETURN 'Отлично';
	WHEN x = 4 THEN RETURN 'Хорошо';
	WHEN x = 3 THEN RETURN 'Удовлетворительно';
	WHEN x = 2 THEN RETURN 'Не удовлетворительно';
	ELSE RETURN 'Введите число от 2 до 5';
	END CASE;
END
$$ LANGUAGE plpgsql
```

Вызов функции осуществялется при помощи следующего запроса

```plpgsql
SELECT num_if(5);
```

<h3>5. Выведите все квадраты чисел от 20 до 30 3-мя разными способами (LOOP, WHILE, FOR).</h3>  

Вариант решения через конструкцию LOOP:

```plpgsql
CREATE OR REPLACE PROCEDURE loop() AS $$
DECLARE
	x int := 20;
BEGIN
LOOP
	RAISE NOTICE '%^2:%', x, x*x;
	x := x + 1;
	EXIT WHEN x > 30;		
END LOOP;
END
$$ LANGUAGE plpgsql;
```

Вариант решения через конструкцию WHILE:

```plpgsql
CREATE OR REPLACE PROCEDURE loop() AS $$
DECLARE
	x int := 20;
BEGIN
WHILE x <= 30 LOOP
	RAISE NOTICE '% ^ 2 : %', x, x*x;
	x := x + 1;	
END LOOP;
END
$$ LANGUAGE plpgsql;
```

Вариант решения через конструкцию FOR:

```plpgsql
CREATE OR REPLACE PROCEDURE loop() AS $$
BEGIN
FOR x IN 20..30 LOOP
	RAISE NOTICE '%^2:%', x, x*x;	
END LOOP;
END
$$ LANGUAGE plpgsql;
```

Мне очень импонирует краткость цикла FOR, поэтому я предпочту чаще работать с ним, чем с другими циклами  

Вызов процедуры отличается от функции и осуществялется при помощи следующего запроса

```plpgsql
CALL loop();
```

<h3>6. Последовательность Коллатца. Берётся любое натуральное число. Если чётное - делим его на 2, если нечётное, то умножаем его на 3 и прибавляем 1. Такие действия выполняются до тех пор, пока не будет получена единица. Гипотеза заключается в том, что какое бы начальное число n не было выбрано, всегда получится 1 на каком-то шаге. Задания: написать функцию, входной параметр - начальное число, на выходе - количество чисел, пока не получим 1; написать процедуру, которая выводит все числа последовательности. Входной параметр - начальное число.</h3>

Функция:  

```plpgsql
CREATE OR REPLACE FUNCTION kollatc(x int) RETURNS int AS $$
DECLARE
	amount int := 0;
BEGIN
	WHILE x != 1 LOOP
		IF mod(x, 2) = 0 THEN
			x := x / 2;
		ELSE
			x := x * 3 + 1;
		END IF;
		amount := amount + 1;
	END LOOP;
RETURN amount;
END
$$ LANGUAGE plpgsql;
```

Процедура:  

```plphsql
CREATE OR REPLACE PROCEDURE kollatc_p(x int) AS $$
BEGIN
	WHILE x != 1 LOOP
		RAISE NOTICE '%', x;
		IF mod(x, 2) = 0 THEN
			x := x / 2;
		ELSE
			x := x * 3 + 1;
		END IF;
	END LOOP;
END
$$ LANGUAGE plpgsql;
```

<h3>7. Числа Люка. Объявляем и присваиваем значение переменной - количество числе Люка. Вывести на экран последовательность чисел. Где L0 = 2, L1 = 1 ; Ln=Ln-1 + Ln-2 (сумма двух предыдущих чисел). Задания: написать фунцию, входной параметр - количество чисел, на выходе - последнее число (Например: входной 5, 2 1 3 4 7 - на выходе число 7); написать процедуру, которая выводит все числа последовательности. Входной параметр - количество чисел.</h3>  

Функция:

```plpgsql
CREATE OR REPLACE FUNCTION luka_f(x int) RETURNS int AS $$
DECLARE
L0 int := 2;
L1 int := 1;
L2 int;
BEGIN
	FOR x IN 3..x LOOP
		L2 := L1 + L0;
		L0 := L1;
		L1 := L2;
	END LOOP;
	RETURN L2;
END
$$ LANGUAGE plpgsql;
```

Процедура:

```plpgsql
CREATE OR REPLACE PROCEDURE luka_p(x int) AS $$
DECLARE
L0 int := 2;
L1 int := 1;
L2 int;
BEGIN
	RAISE NOTICE '%', L0;
	RAISE NOTICE '%', L1;
	FOR x IN 3..x LOOP
		L2 := L1 + L0;
		L0 := L1;
		L1 := L2;
		RAISE NOTICE '%', L2;
	END LOOP;
END
$$ LANGUAGE plpgsql;
```

<h3>8. Напишите функцию, которая возвращает количество человек родившихся в заданном году.</h3>

```plpgsql
CREATE OR REPLACE FUNCTION people_by_year(year_p int) RETURNS int AS $$
DECLARE
	amount int;
BEGIN
	SELECT count(*) INTO amount
	FROM people
	WHERE EXTRACT(year FROM people.birth_date) = year_p;
	RETURN amount;
END
$$ LANGUAGE plpgsql;
```

<h3>9. Напишите функцию, которая возвращает количество человек с заданным цветом глаз.</h3>

```plpgsql
CREATE OR REPLACE FUNCTION people_by_eyes(color varchar) RETURNS int AS $$
DECLARE
	amount int;
BEGIN
	SELECT count(*) INTO amount
	FROM people
	WHERE people.eyes = color;
	RETURN amount;
END
$$ LANGUAGE plpgsql;

SELECT people_by_eyes('brown');
```

<h3>10. Напишите функцию, которая возвращает ID самого молодого человека в таблице.</h3>

```plpgsql
CREATE OR REPLACE FUNCTION younger_people_id() RETURNS int AS $$
DECLARE
	y_p int;
BEGIN
	SELECT id INTO y_p
	FROM people
	WHERE birth_date = (SELECT max(birth_date) FROM people);
	RETURN y_p;
END
$$ LANGUAGE plpgsql;
```

<h3>11. Напишите процедуру, которая возвращает людей с индексом массы тела больше заданного. ИМТ = масса в кг / (рост в м)^2.</h3>

```plpgsql
CREATE OR REPLACE PROCEDURE people_by_imt(imt int) AS $$
DECLARE
	p people%ROWTYPE;
BEGIN
	FOR p IN
		SELECT * FROM people
	LOOP
		IF p.weight / ((p.growth/100)^2 > imt THEN
			RAISE NOTICE 'name: %, surname: %', p.name, p.surname;
		END IF;
	END LOOP;
END
$$ LANGUAGE plpgsql
```

<h3>12. Измените схему БД так, чтобы в БД можно было хранить родственные связи между людьми. Код должен быть представлен в виде транзакции (Например (добавление атрибута): BEGIN; ALTER TABLE people ADD COLUMN leg_size REAL; COMMIT;). Дополните БД данными.</h3>  

Создаём таблицу  

```plpgsql
CREATE TABLE "links" (
	"id" integer primary key,
	"people_id" INTEGER NOT NULL,
	"people_id_2" INTEGER NOT NULL, 
	"link" VARCHAR,
	FOREIGN KEY ("people_id") REFERENCES "people" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
	FOREIGN KEY ("people_id_2") REFERENCES "people" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
	);
```
И добавляем данные при помощи транзакции  

```plpgsql
BEGIN;
	INSERT INTO links (id, people_id, people_id_2, link)
	VALUES  (1, 2, 6, 'брат'),
			(2, 4, 5, 'сестра');
COMMIT;
```

<h3>13. Напишите процедуру, которая позволяет создать в БД нового человека с указанным родством.</h3>

```plpgsql
CREATE OR REPLACE PROCEDURE add_people_by_link(id_people int, name_people varchar, id_people_2 int, name_people_2 varchar, name_link varchar)
LANGUAGE plpgsql
AS $$
DECLARE
BEGIN
INSERT INTO people (id, name)
VALUES (id_people, name_people);
IF NOT EXISTS (SELECT * FROM people WHERE id = id_people_2) THEN
INSERT INTO people (id, name)
VALUES (id_people_2, name_people_2);
END IF;
INSERT INTO links (people_id, people_id_2, link)
VALUES (id_people, id_people_2, name_link); 
END;
$$
```

И вызываем функцию при помощи оператора CALL

```plpgsql
CALL add_people_by_link(9, 'makar', 7, 'roma', 'брат')
```

<h3>14. Измените схему БД так, чтобы в БД можно было хранить время актуальности данных человека (выполнить также, как п.12).</h3>

```plpgsql

```

<h3>15. Напишите процедуру, которая позволяет актуализировать рост и вес человека.</h3>

```plpgsql

```
