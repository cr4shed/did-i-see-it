START TRANSACTION;

-- Assumes that the database and tables have been created using create-tables.sql.

INSERT INTO `MediaType` (`Type`) VALUES ('Movie'), ('TV Show');

INSERT INTO `Media` (`MediaTypeId`, `Title`, `Year`) VALUES (1, 'Zombieland', 2009), (2, 'Lost', 2004), (1, 'Hackers', 1995);

INSERT INTO `User` (`Id`, `Username`, `Email`, `PassHash`) VALUES (1, 'TESTUSER', 'TESTUSER@didiseeit.com', 'TESTHASH');

INSERT INTO `Collection` (`UserId`, `Name`) VALUES (1, 'My Collection');

INSERT INTO `View` (`CollectionId`, `MediaId`) VALUES (1, 1);

COMMIT;