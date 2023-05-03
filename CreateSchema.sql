CREATE TABLE `Cards`
(
  `ID`        varchar(40) PRIMARY KEY,
  `Owner`     varchar(40) COMMENT 'Daimyo.Nickname',
  `BankInfo`  varchar(40),
  `LimitInfo` float,
  `Balance`   float
);

CREATE TABLE `Samurais`
(
  `Nickname`         varchar(255),
  `TelegramUsername` varchar(40) PRIMARY KEY,
  `Owner`            varchar(40) COMMENT 'Daimyo.Nickname'
);

CREATE TABLE `Daimyo`
(
  `Nickname`         varchar(40),
  `Owner`            varchar(40),
  `TelegramUsername` varchar(40) PRIMARY KEY
);

CREATE TABLE `Shogun`
(
  `TelegramUsername` varchar(40) PRIMARY KEY,
  `Nickname`         varchar(40)
);

CREATE TABLE `Collectors`
(
  `TelegramUsername` varchar(40) PRIMARY KEY,
  `Nickname`         varchar(40)
);

CREATE TABLE `Applications`
(
  `Daimyo` varchar(40),
  `Sum`    float,
  `ID`     varchar(40) PRIMARY KEY COMMENT 'card_id'
);
CREATE TABLE `Admins`
(
  `TelegramUsername` varchar(40) PRIMARY KEY,
  `Nickname`         varchar(40)
);

CREATE TABLE `Transactions`
(
  `OperationType`   bool,
  `Amount`          float,
  `Date`            DATETIME,
  `SamuraiUsername` varchar(40),
  `CardID`          varchar(40),
  FOREIGN KEY (`SamuraiUsername`) REFERENCES Samurais (`TelegramUsername`),
  PRIMARY KEY (`SamuraiUsername`, `Date`, `CardID`)
);

CREATE TABLE `Turnovers`
(
  `Amount`          float,
  `Date`            DATE,
  `SamuraiUsername` varchar(40),
  FOREIGN KEY (`SamuraiUsername`) REFERENCES Samurais (`TelegramUsername`),
  PRIMARY KEY (`SamuraiUsername`, `Date`)
);

ALTER TABLE `Cards`
  ADD FOREIGN KEY (`Owner`) REFERENCES `Daimyo` (`TelegramUsername`);

ALTER TABLE `Samurais`
  ADD FOREIGN KEY (`Owner`) REFERENCES `Daimyo` (`TelegramUsername`);

ALTER TABLE `Daimyo`
  ADD FOREIGN KEY (`Owner`) REFERENCES `Shogun` (`TelegramUsername`);

ALTER TABLE `Applications`
  ADD FOREIGN KEY (`Daimyo`) REFERENCES `Daimyo` (`TelegramUsername`);

ALTER TABLE `Transactions`
  ADD FOREIGN KEY (`SamuraiUsername`) REFERENCES `Samurais` (`TelegramUsername`);
