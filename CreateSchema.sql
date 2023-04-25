CREATE TABLE `Cards` (
  `ID` integer PRIMARY KEY,
  `Owner` varchar(40) COMMENT 'Daimyo.Nickname',
  `BankInfo` varchar(40),
  `LimitInfo` float,
  `Balance` float
);

CREATE TABLE `Samurais` (
  `Nickname` varchar(255) PRIMARY KEY,
  `TurnOver` float,
  `TelegramUsername` varchar(40),
  `Owner` varchar(40) COMMENT 'Daimyo.Nickname'
);

CREATE TABLE `Temp` (
  `Nickname` varchar(255) PRIMARY KEY,
  `TurnOver` float,
  `TelegramUsername` varchar(40),
  `Owner` varchar(40) COMMENT 'Daimyo.Nickname'
);
drop table Temp;

CREATE TABLE `Daimyo` (
  `Nickname` varchar(40) PRIMARY KEY,
  `Owner` varchar(40),
  `TelegramUsername` varchar(40)
);

CREATE TABLE `Shogun` (
  `TelegramUsername` varchar(40),
  `Nickname` varchar(40) PRIMARY KEY
);

CREATE TABLE `Collectors` (
  `TelegramUsername` varchar(40),
  `Nickname` varchar(40) PRIMARY KEY
);

CREATE TABLE `Applications` (
  `Daimyo` varchar(40),
  `Sum` float,
  `ID` integer PRIMARY KEY COMMENT 'card_id'
);

CREATE TABLE `Admins` (
  `TelegramUsername` varchar(40),
  `Nickname` varchar(40) PRIMARY KEY
);

ALTER TABLE `Cards` ADD FOREIGN KEY (`Owner`) REFERENCES `Daimyo` (`Nickname`);

ALTER TABLE `Samurais` ADD FOREIGN KEY (`Owner`) REFERENCES `Daimyo` (`Nickname`);

ALTER TABLE `Daimyo` ADD FOREIGN KEY (`Owner`) REFERENCES `Shogun` (`Nickname`);

ALTER TABLE `Applications` ADD FOREIGN KEY (`Daimyo`) REFERENCES `Daimyo` (`Nickname`);

