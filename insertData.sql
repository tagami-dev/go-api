insert into articles (title, contents, username, nice, created_at) values
('1st', 'The first article will be published', 'tgm', 2, now());

insert into articles (title, contents, username, nice) values
('2nd', 'The second article will be about computer science', 'tgm', 4);

insert into comments (article_id, message, created_at) values
(1, '1st hello world', now());

insert into comments (article_id, message) values
(1, 'awesome');