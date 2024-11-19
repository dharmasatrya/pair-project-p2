CREATE DATABASE game

use game

CREATE TABLE IF NOT EXISTS developers(
	developerID INT AUTO_INCREMENT PRIMARY KEY,
	studioName VARCHAR (50) NOT NULL,
	location VARCHAR (50) NOT NULL
)

CREATE TABLE IF NOT EXISTS games(
	gameID INT AUTO_INCREMENT PRIMARY KEY,
	name VARCHAR (50) UNIQUE NOT NULL,
	price DECIMAL(10,2) NOT NULL,
	releaseDate DATE NOT NULL,
	developerID INT,
	FOREIGN KEY (developerID) REFERENCES developers(developerID)
)

CREATE TABLE IF NOT EXISTS players(
	playerID INT AUTO_INCREMENT PRIMARY KEY,
	name VARCHAR (50) UNIQUE NOT NULL,
	email VARCHAR (50) UNIQUE NOT NULL
)

CREATE TABLE IF NOT EXISTS transactions(
	trxID INT AUTO_INCREMENT PRIMARY KEY,
	gameID INT NOT NULL,
	playerID INT NOT NULL,
	purchasedAt DATE NOT NULL,
	FOREIGN KEY (gameID) REFERENCES games(gameID),
	FOREIGN KEY (playerID) REFERENCES players(playerID)
)

CREATE TABLE IF NOT EXISTS player_interactions(
	interactionID INT AUTO_INCREMENT PRIMARY KEY,
	gameID INT NOT NULL,
	playerID INT NOT NULL,
	hoursPlayed INT NOT NULL,
	FOREIGN KEY (gameID) REFERENCES games(gameID),
	FOREIGN KEY (playerID) REFERENCES players(playerID)
)

INSERT INTO Developers (StudioName, Location) VALUES 
('Creative Visions', 'New York'),
('Dream Games', 'San Francisco'),
('Pixel Studios', 'London');

INSERT INTO games (Name, Price, ReleaseDate, developerID) VALUES 
('Epic Adventure', 59.99, '2021-03-15', 1),
('Space Explorers', 49.99, '2022-01-25', 2),
('Fantasy Quest', 39.99, '2022-05-20', 3);

INSERT INTO Players (Name, Email) VALUES 
('John Doe', 'john.doe@example.com'),
('Alice Smith', 'alice.smith@example.com'),
('Bob Johnson', 'bob.johnson@example.com');

INSERT INTO transactions (gameID, playerID, purchasedAt) VALUES 
(1, 1, '2023-08-10'),
(2, 2, '2023-08-11'),
(3, 3, '2023-08-12'),
(3, 1, '2023-08-12');

INSERT INTO player_interactions  (gameID, playerID, hoursPlayed) VALUES 
(2, 1, '10'),
(1, 3, '1200');