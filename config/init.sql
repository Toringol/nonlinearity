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
VALUES (1, 'История люблю', 'История о влюбленной паре', 'https://toringolimagestorage.s3.eu-north-1.amazonaws.com/storyImage/romance.jpeg', 
'https://toringolimagestorage.s3.eu-north-1.amazonaws.com/stories/test-story.json', 'Барбара Картленд', false, 1, 7.3, 3);

INSERT INTO genres (`id`, `genre`) VALUES (1, 'романтика');

INSERT INTO stories (`id`, `title`, `description`, `image`, `storyPath`, `author`, `editorChoice`, `ratingsNumber`, `rating`, `views`)
VALUES (2, 'Путник', 'Однажды, блуждая по пустуны, путник наткнулся на нечто', 'https://toringolimagestorage.s3.eu-north-1.amazonaws.com/storyImage/drama.jpeg', 
'https://toringolimagestorage.s3.eu-north-1.amazonaws.com/stories/test-story.json', 'Вирджиния Эндрюс', false, 1, 4.5, 4);

INSERT INTO genres (`id`, `genre`) VALUES (2, 'драма');

INSERT INTO stories (`id`, `title`, `description`, `image`, `storyPath`, `author`, `editorChoice`, `ratingsNumber`, `rating`, `views`)
VALUES (3, 'Жизнь', 'Жизнь глазами Боба Дикмана', 'https://toringolimagestorage.s3.eu-north-1.amazonaws.com/storyImage/action-realism.jpg', 
'https://toringolimagestorage.s3.eu-north-1.amazonaws.com/stories/test-story.json', 'Боб Дикман', false, 1, 4.7, 4);

INSERT INTO genres (`id`, `genre`) VALUES (3, 'экшен');
INSERT INTO genres (`id`, `genre`) VALUES (3, 'реализм');

INSERT INTO stories (`id`, `title`, `description`, `image`, `storyPath`, `author`, `editorChoice`, `ratingsNumber`, `rating`, `views`)
VALUES (4, 'История игрушек', 'История о секретной жизни игрушек, когда людей нет рядом', 'https://toringolimagestorage.s3.eu-north-1.amazonaws.com/storyImage/comedy-fantasy.jpg', 
'https://toringolimagestorage.s3.eu-north-1.amazonaws.com/stories/test-story.json', '‎Джон Лессетер', false, 1, 5.2, 4);

INSERT INTO genres (`id`, `genre`) VALUES (4, 'комедия');
INSERT INTO genres (`id`, `genre`) VALUES (4, 'фантастика');

INSERT INTO stories (`id`, `title`, `description`, `image`, `storyPath`, `author`, `editorChoice`, `ratingsNumber`, `rating`, `views`)
VALUES (5, 'Жизнь, полная дыр', 'Короткие истории', 'https://toringolimagestorage.s3.eu-north-1.amazonaws.com/storyImage/drama-horror-action.png', 
'https://toringolimagestorage.s3.eu-north-1.amazonaws.com/stories/test-story.json', 'Мэри Папас', false, 1, 7, 5);

INSERT INTO genres (`id`, `genre`) VALUES (5, 'драма');
INSERT INTO genres (`id`, `genre`) VALUES (5, 'ужасы');
INSERT INTO genres (`id`, `genre`) VALUES (5, 'экшен');

INSERT INTO stories (`id`, `title`, `description`, `image`, `storyPath`, `author`, `editorChoice`, `ratingsNumber`, `rating`, `views`)
VALUES (6, 'Мистическая история', 'Мистическая история убийства, которая повлекла за собой невероятные изменения', 'https://toringolimagestorage.s3.eu-north-1.amazonaws.com/storyImage/detective.jpg', 
'https://toringolimagestorage.s3.eu-north-1.amazonaws.com/stories/test-story.json', '‎Стивен Кинг', false, 1, 7.5, 6);

INSERT INTO genres (`id`, `genre`) VALUES (6, 'детектив');

INSERT INTO stories (`id`, `title`, `description`, `image`, `storyPath`, `author`, `editorChoice`, `ratingsNumber`, `rating`, `views`)
VALUES (7, 'Реальная история', 'История основанная на реальных событиях', 'https://toringolimagestorage.s3.eu-north-1.amazonaws.com/storyImage/realism-romance.jpg', 
'https://toringolimagestorage.s3.eu-north-1.amazonaws.com/stories/test-story.json', '‎Стивен Кинг', true, 1, 8.5, 7);

INSERT INTO genres (`id`, `genre`) VALUES (7, 'реализм');
INSERT INTO genres (`id`, `genre`) VALUES (7, 'романтика');

INSERT INTO stories (`id`, `title`, `description`, `image`, `storyPath`, `author`, `editorChoice`, `ratingsNumber`, `rating`, `views`)
VALUES (8, 'Американская история ужасов', 'Если вы ищете настоящие ужасы, то вы попали в правильное место', 'https://toringolimagestorage.s3.eu-north-1.amazonaws.com/storyImage/horror.jpeg', 
'https://toringolimagestorage.s3.eu-north-1.amazonaws.com/stories/test-story.json', 'Райан Мерфи', false, 1, 9, 10);

INSERT INTO genres (`id`, `genre`) VALUES (8, 'ужасы');


INSERT INTO `users` (`id`, `login`, `password`, `avatar`) VALUES (1, 'testUser', 'Op4wlDuwonp1pf5ubNutOxJkSx9QBWnjYi8bBd/Qn0mXCgvaegRNWA', 'https://toringolimagestorage.s3.eu-north-1.amazonaws.com/avatars/defaultAvatar.png');

INSERT INTO userFavourites (`id`, `drama`, `romance`, `comedy`, `horror`, `detective`, `fantasy`, `action`, `realism`) 
VALUES (1, 5, 1, 3, 10, 7, 2, 6, 8);


