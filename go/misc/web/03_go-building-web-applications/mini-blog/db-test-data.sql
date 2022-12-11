CREATE TABLE IF NOT EXISTS `pages` (
     `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
     `page_guid` varchar(256) NOT NULL DEFAULT '',
     `page_title` varchar(256) DEFAULT NULL,
     `page_content` mediumtext,
     `page_date` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON
   UPDATE CURRENT_TIMESTAMP,
     PRIMARY KEY (`id`),
     UNIQUE KEY `page_guid` (`page_guid`)
   ) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1;

INSERT INTO `pages` (`id`, `page_guid`, `page_title`,
   `page_content`, `page_date`) VALUES (NULL, 'hello-world', 'Hello,
   World', 'I\'m so glad you found this page!  It\'s been sitting
   patiently on the Internet for some time, just waiting for a
   visitor.', CURRENT_TIMESTAMP);

INSERT INTO `pages` (`page_guid`, `page_title`, `page_content`, `page_date`) VALUES
	('a-new-blog', 'A New Blog', 'I hope you enjoyed the last
   blog!  Well brace yourself, because my latest blog is even
   <i>better</i> than the last!', '2015-04-29 02:16:19');

INSERT INTO `pages` (`page_guid`, `page_title`, `page_content`, `page_date`) VALUES
	('lorem-ipsum', 'Lorem Ipsum', 'Lorem ipsum dolor sit amet,
   consectetur adipiscing elit. Maecenas sem tortor, lobortis in
   posuere sit amet, ornare non eros. Pellentesque vel lorem sed nisl
   dapibus fringilla. In pretium...', '2015-05-06 04:09:45');

CREATE TABLE `comments` (
   `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
   `page_id` int(11) NOT NULL,
   `comment_guid` varchar(256) DEFAULT NULL,
   `comment_name` varchar(64) DEFAULT NULL,
   `comment_email` varchar(128) DEFAULT NULL,
   `comment_text` mediumtext,
   `comment_date` timestamp NULL DEFAULT NULL,
   PRIMARY KEY (`id`),
   KEY `page_id` (`page_id`)
   ) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE `users` (
     `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
     `user_name` varchar(32) NOT NULL DEFAULT '',
     `user_guid` varchar(256) NOT NULL DEFAULT '',
     `user_email` varchar(128) NOT NULL DEFAULT '',
     `user_password` varchar(128) NOT NULL DEFAULT '',
     `user_salt` varchar(128) NOT NULL DEFAULT '',
     `user_joined_timestamp` timestamp NULL DEFAULT NULL,
     PRIMARY KEY (`id`)
   ) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE `sessions` (
     `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
     `session_id` varchar(256) NOT NULL DEFAULT '',
     `user_id` int(11) DEFAULT NULL,
     `session_start` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
     `session_update` timestamp NOT NULL DEFAULT '2000-01-01 00:00:00',
     `session_active` tinyint(1) NOT NULL,
     PRIMARY KEY (`id`),
     UNIQUE KEY `session_id` (`session_id`)
   ) ENGINE=InnoDB DEFAULT CHARSET=latin1;