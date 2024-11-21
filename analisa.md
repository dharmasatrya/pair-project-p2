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
- categoryID FK product_categories.categoryID

D. Entity: transactions

- Attributes:
- trxID: PK AI 
- userID: FK users.userID
- productID: FK products.productID
- purchasedAt: DATE

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

## Note
- etc ...