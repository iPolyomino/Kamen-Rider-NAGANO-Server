use imageboard;

SET NAMES utf8mb4;

INSERT INTO `room` (`id`,`name`) VALUES (1, 'ポケモン好き集まれ!');

INSERT INTO `comment` (`sender`,`text`,`room_id`) VALUES ('サトシ','初めまして！よろしくお願いします。', 1);
INSERT INTO `comment` (`sender`,`text`,`room_id`) VALUES ('タケシ','よろしく〜', 1);
INSERT INTO `comment` (`sender`,`text`,`room_id`) VALUES ('カスミ','わたしのポリシーはね… みずタイプポケモンで せめてせめて …せめまくることよ！', 1);
