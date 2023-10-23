CREATE TABLE categories (
    category_id INT AUTO_INCREMENT PRIMARY KEY,
    category_name VARCHAR(255)
);

CREATE TABLE transactions (
    transaction_id INT AUTO_INCREMENT PRIMARY KEY,
    transaction_date DATE,
    transaction_type ENUM('income', 'expense'), 
    amount INT,
    category_id INT,
    description VARCHAR(300),
    FOREIGN KEY (category_id) REFERENCES categories(category_id)
);
