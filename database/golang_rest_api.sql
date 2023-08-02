-- CREATING AND USING DATABASE
CREATE DATABASE IF NOT EXISTS golang_rest_api;
USE golang_rest_api;

-- CREATING TABLES

DROP TABLE IF EXISTS tasks;
CREATE TABLE IF NOT EXISTS tasks (
	task_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	task_name VARCHAR(100) NOT NULL,
	start_date DATE NOT NULL,
	end_date DATE NOT NULL,
	description VARCHAR(200) NOT NULL,
	status ENUM('Completed', 'In Progress', 'Pending') DEFAULT 'Pending'
);


