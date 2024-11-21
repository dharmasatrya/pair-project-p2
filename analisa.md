## ERD Title: Toko Game Gacor

1. Entities and Their Attributes:
Entity: Table_Name (e.g., Customers)

## Attributes:

- Field_Name (Attribute) : e.g., CustomerId (PK, AI)
- Field (Datatype)
- etc...

A. Entity: users

- Attributes:
- userID: PK AI 
- name: VARCHAR(50) NOT NULL
- email: VARCHAR(50) NOT NULL

B. Entity: product_categories

- Attributes:
- categoryID: PK AI 
- name: VARCHAR(50) NOT NULL

C. Entity: products

- Attributes:
- productID: PK AI 
- name: VARCHAR(50) NOT NULL
- price: INT NOT NULL
- quantity: INT NOT NULL
- categoryID FK product_categories.categoryID NOT NULL

D. Entity: transactions

- Attributes:
- trxID: PK AI 
- userID: FK users.userID NOT NULL
- productID: FK products.productID NOT NULL
- purchasedAt: DATE NOT NULL

## Relationships:
- Table_Name to Table_Name: (e.g., Customers to Orders)

A. Type: One to Many
- Description: One user can have many transactions.
- users to transactions

B. Type: One to Many
- Description: One category product can have many products.
- product_category to products

C. Type: One to Many
- Description: One product can have many transactions.
- products to transactions

## Integrity Constraints:
- self explanatory
- Create trigger update_product_quantity_after_transaction: This trigger decreases the quantity by 1 for the product involved in a new transaction.

## Note
- etc ...