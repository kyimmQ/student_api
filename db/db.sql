-- CREATE NEW DATABASE
DROP DATABASE IF EXISTS student_api_db;

CREATE DATABASE student_api_db;

USE student_api_db;

-- CREATE SUDENT TABLE
CREATE TABLE student (
    std_id int primary key,
    std_name varchar(255) not null,
    academic_year int not null
);

-- CREATE STORED PROCEDURES FOR QUERYING

-- CREATE NEW STUDENT
DELIMITER //
CREATE PROCEDURE create_new_student (
    IN p_std_id INT,
    IN p_std_name VARCHAR(255),
    IN p_academic_year INT
)
BEGIN
    DECLARE std_id_exists INT;

    -- Check if student ID already exists
    SELECT COUNT(*) INTO std_id_exists FROM student WHERE std_id = p_std_id;

    -- If student ID exists, raise an error and exit
    IF std_id_exists > 0 THEN
        SELECT 0 as success;
    ELSE
        -- Insert new student
        INSERT INTO student (std_id, std_name, academic_year) VALUES (p_std_id, p_std_name, p_academic_year);
        SELECT 1 as success;
    END IF;
END //

-- SEARCH SUDENT BY ID
CREATE PROCEDURE search_student_by_id (
    IN p_std_id INT
)
BEGIN
    SELECT * FROM student  WHERE  std_id = p_std_id;
END //


-- SEARCH SUDENT BY NAME
CREATE PROCEDURE search_student_by_name (
    IN p_std_name VARCHAR(255)
)
BEGIN
    IF p_std_name IS NULL THEN
        SELECT * FROM student;
    ELSE
        SELECT * FROM student WHERE std_name LIKE CONCAT("%", p_std_name, "%");
    END IF;
END //
DELIMITER ;

INSERT INTO student (std_id, std_name, academic_year) VALUES
(2111234, 'John', 2021),
(2112345, 'Alice', 2021),
(2113456, 'David', 2021),
(2114567, 'Emma', 2021),
(2115678, 'Michael', 2021),
(2116789, 'Sophia', 2021),
(2117890, 'William', 2021),
(2118901, 'Olivia', 2021),
(2119012, 'James', 2021),
(2110123, 'Alice', 2021),
(2211234, 'Ethan', 2022),
(2212345, 'Emily', 2022),
(2213456, 'Alexander', 2022),
(2214567, 'Mia', 2022),
(2215678, 'Daniel', 2022),
(2216789, 'Abigail', 2022),
(2217890, 'Matthew', 2022),
(2218901, 'Ava', 2022),
(2219012, 'Ryan', 2022),
(2210123, 'Elizabeth', 2022),
(2311234, 'Noah', 2023),
(2312345, 'Samantha', 2023),
(2313456, 'Christopher', 2023),
(2314567, 'Madison', 2023),
(2315678, 'Jacob', 2023),
(2316789, 'James', 2023),
(2317890, 'Madison', 2023),
(2318901, 'Sophie', 2023),
(2319012, 'Benjamin', 2023),
(2310123, 'Amelia', 2023);