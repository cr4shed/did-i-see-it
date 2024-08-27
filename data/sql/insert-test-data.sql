START TRANSACTION;

-- Assumes that the database and tables have been created using create-tables.sql.

INSERT INTO `MediaType` (`Type`) VALUES ('Movie'), ('TV Show');

INSERT INTO `Media` (`MediaTypeId`, `Title`) VALUES (1, 'Zombieland'), (2, 'Lost');

INSERT INTO `User` (`Id`) VALUES (1);

INSERT INTO `Collection` (`UserId`, `Name`) VALUES (1, 'My Collection');

INSERT INTO `View` (`CollectionId`, `MediaId`) VALUES (1, 1);

COMMIT;