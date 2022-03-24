CREATE TABLE `chat_configs` (
                           `id` int unsigned NOT NULL AUTO_INCREMENT,
                           `conf_name` varchar(255) NOT NULL DEFAULT '',
                           `conf_key` varchar(255) NOT NULL DEFAULT '',
                           `conf_value` varchar(1024) NOT NULL DEFAULT '',
                           PRIMARY KEY (`id`),
                           UNIQUE KEY `conf_key` (`conf_key`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;