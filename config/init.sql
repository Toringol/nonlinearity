SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `login` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `avatar` varchar(255) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `userFavourites`;
CREATE TABLE `userFavourites` (
  `id` int(11) NOT NULL,
  `drama` int(11) DEFAULT NULL,
  `romance` int(11) DEFAULT NULL,
  `comedy` int(11) DEFAULT NULL,
  `horror` int(11) DEFAULT NULL,
  `detective` int(11) DEFAULT NULL,
  `fantasy` int(11) DEFAULT NULL,
  `action` int(11) DEFAULT NULL,
  `realism` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `stories`;
CREATE TABLE `stories` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL,
  `description` varchar(255) NOT NULL,
  `image` varchar(255) NOT NULL,
  `storyPath` varchar(255) NOT NULL,
  `author` varchar(255) NOT NULL,
  `editorChoice` BOOLEAN DEFAULT NULL,
  `ratingsNumber` int(11) DEFAULT 0,
  `rating` float DEFAULT 0,
  `views` int(11) DEFAULT 0,
  `publicationDate` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


DROP TABLE IF EXISTS `genres`;
CREATE TABLE `genres` (
  `id` int(11) NOT NULL,
  `genre` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


DROP TABLE IF EXISTS `storyRatingViews`;
CREATE TABLE `storyRatingViews` (
  `storyID` int(11) NOT NULL,
  `userID` int(11) NOT NULL,
  `view` BOOLEAN DEFAULT NULL,
  `rating` BOOLEAN DEFAULT NULL,
  `previousRate` float DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- First insert to DB

INSERT INTO stories (`id`, `title`, `description`, `image`, `storyPath`, `author`, `editorChoice`)
VALUES (1, 'LoveStory', 'Story about one man and girl', 'somePic', 'SomeStory', 'Toringol', true);

INSERT INTO genres (`id`, `genre`) VALUES (1, 'someGenre');
INSERT INTO genres (`id`, `genre`) VALUES (1, 'someGenre');

INSERT INTO `users` (`id`, `login`, `password`, `avatar`) VALUES (1, 'test', '123', 'somePic');




