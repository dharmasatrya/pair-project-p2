## ERD Title: Video Game Store System

1. Entities and Their Attributes:
Entity: Table_Name (e.g., Customers)

## Attributes:

- Field_Name (Attribute) : e.g., CustomerId (PK, AI)
- Field (Datatype)
- etc...

A. Entity: developers

- Attributes:
- developerID: PK AI 
- studioName: VARCHAR UNIQUE NOT NULL
- location: VARCHAR NOT NULL

B. Entity: games

- Attributes:
- gameID: PK AI 
- name: VARCHAR UNIQUE NOT NULL
- price: DECIMAL(10,2) NOT NULL
- releaseDate: DATE NOT NULL

C. Entity: players

- Attributes:
- playerID: PK AI 
- name: VARCHAR NOT NULL
- email: VARCHAR UNIQUE NOT NULL

D. Entity: transactions

- Attributes:
- trxID: PK AI 
- gameID: FK games.gameID
- playerID: FK players.playerID
- purchasedAt: DATE

E. Entity: player_interactions

- Attributes:
- interactionID: PK AI 
- gameID: FK games.gameID
- playerID: FK players.playerID
- hoursPlayed: INT

## Relationships:
- Table_Name to Table_Name: (e.g., Customers to Orders)

A. Type: One to Many
- Description: One developer can have many games.
- developers to games

B. Type: One to Many
- Description: One games can have many transactions
- games to transactions

C. Type: One to Many
- Description: One player can have many transactions
- players to transactions

D. Type: One to Many
- Description: One player can have many games interracted
- players to player_interactions

E. Type: One to Many
- Description: One game can have many players interracted
- games to player_interactions

## Integrity Constraints:
- self explanatory

## Note
- hoursPlayed will be calculated later (just an example of 1 interaction)