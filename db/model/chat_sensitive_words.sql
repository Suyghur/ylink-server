CREATE TABLE `chat_sensitive_words` (
                                        `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
                                        `word` varchar(128) NOT NULL DEFAULT '',
                                        `add_type` tinyint(8) unsigned DEFAULT '0',
                                        PRIMARY KEY (`id`),
                                        UNIQUE KEY `idx_word` (`word`) USING HASH COMMENT '关键字索引'
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;