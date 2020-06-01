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

INSERT INTO stories (`id`, `title`, `description`, `image`, `storyPath`, `author`, `editorChoice`, `ratingsNumber`, `rating`, `views`)
VALUES (1, 'LoveStory', 'Story about one man and girl', 'https://toringolimagestorage.s3.eu-north-1.amazonaws.com/storyImage/romance.jpeg', 
'https://toringolimagestorage.s3.eu-north-1.amazonaws.com/stories/test-story.json', 'Toringol', false, 1, 1.2, 3);

INSERT INTO genres (`id`, `genre`) VALUES (1, 'romance');

INSERT INTO stories (`id`, `title`, `description`, `image`, `storyPath`, `author`, `editorChoice`, `ratingsNumber`, `rating`, `views`)
VALUES (2, 'The man', 'Once upon the time one man from desert find sth', 'https://toringolimagestorage.s3.eu-north-1.amazonaws.com/storyImage/drama.jpeg', 
'https://toringolimagestorage.s3.eu-north-1.amazonaws.com/stories/test-story.json', 'Yulia', false, 1, 4.5, 4);

INSERT INTO genres (`id`, `genre`) VALUES (2, 'drama');

INSERT INTO stories (`id`, `title`, `description`, `image`, `storyPath`, `author`, `editorChoice`, `ratingsNumber`, `rating`, `views`)
VALUES (3, 'Life', 'The life through the eyes Bob Dickman', 'https://toringolimagestorage.s3.eu-north-1.amazonaws.com/storyImage/action-realism.jpg', 
'https://toringolimagestorage.s3.eu-north-1.amazonaws.com/stories/test-story.json', 'Bob Dickman', false, 1, 4.7, 4);

INSERT INTO genres (`id`, `genre`) VALUES (3, 'action');
INSERT INTO genres (`id`, `genre`) VALUES (3, 'realism');

INSERT INTO stories (`id`, `title`, `description`, `image`, `storyPath`, `author`, `editorChoice`, `ratingsNumber`, `rating`, `views`)
VALUES (4, 'Toy Story', 'Toy Story is about the secret life of toys when people are not around', 'https://toringolimagestorage.s3.eu-north-1.amazonaws.com/storyImage/comedy-fantasy.jpg', 
'https://toringolimagestorage.s3.eu-north-1.amazonaws.com/stories/test-story.json', '‎John Lasseter', false, 1, 5.2, 4);

INSERT INTO genres (`id`, `genre`) VALUES (4, 'comedy');
INSERT INTO genres (`id`, `genre`) VALUES (4, 'fantasy');

INSERT INTO stories (`id`, `title`, `description`, `image`, `storyPath`, `author`, `editorChoice`, `ratingsNumber`, `rating`, `views`)
VALUES (5, 'A Life full of holes', 'Short stories', 'https://toringolimagestorage.s3.eu-north-1.amazonaws.com/storyImage/drama-horror-action.png', 
'https://toringolimagestorage.s3.eu-north-1.amazonaws.com/stories/test-story.json', '‎Mary Papas', false, 1, 7, 5);

INSERT INTO genres (`id`, `genre`) VALUES (5, 'drama');
INSERT INTO genres (`id`, `genre`) VALUES (5, 'horror');
INSERT INTO genres (`id`, `genre`) VALUES (5, 'action');

INSERT INTO stories (`id`, `title`, `description`, `image`, `storyPath`, `author`, `editorChoice`, `ratingsNumber`, `rating`, `views`)
VALUES (6, 'Strange story', 'Mysterious murder story', 'https://toringolimagestorage.s3.eu-north-1.amazonaws.com/storyImage/detective.jpg', 
'https://toringolimagestorage.s3.eu-north-1.amazonaws.com/stories/test-story.json', '‎Steven King', false, 1, 7.5, 6);

INSERT INTO genres (`id`, `genre`) VALUES (6, 'detective');

INSERT INTO stories (`id`, `title`, `description`, `image`, `storyPath`, `author`, `editorChoice`, `ratingsNumber`, `rating`, `views`)
VALUES (7, 'Real Story', 'Story based on real story', 'https://toringolimagestorage.s3.eu-north-1.amazonaws.com/storyImage/realism-romance.jpg', 
'https://toringolimagestorage.s3.eu-north-1.amazonaws.com/stories/test-story.json', '‎Steven King', true, 1, 8.5, 7);

INSERT INTO genres (`id`, `genre`) VALUES (7, 'realism');
INSERT INTO genres (`id`, `genre`) VALUES (7, 'romance');

INSERT INTO stories (`id`, `title`, `description`, `image`, `storyPath`, `author`, `editorChoice`, `ratingsNumber`, `rating`, `views`)
VALUES (8, 'American horror story', 'If you are looking for horror, you got to the right place', 'https://toringolimagestorage.s3.eu-north-1.amazonaws.com/storyImage/horror.jpeg', 
'https://toringolimagestorage.s3.eu-north-1.amazonaws.com/stories/test-story.json', 'Ryan Murphy', true, 1, 9, 10);

INSERT INTO genres (`id`, `genre`) VALUES (8, 'horror');


INSERT INTO `users` (`id`, `login`, `password`, `avatar`) VALUES (1, 'test', '123', 'somePic');




