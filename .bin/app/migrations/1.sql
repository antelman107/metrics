-- +migrate Up
CREATE TABLE metric (
  id         BIGSERIAL    NOT NULL,
  site_id    BIGINT       NOT NULL,
  error      TEXT         NOT NULL DEFAULT '',
  is_success BOOL         NOT NULL DEFAULT FALSE,
  duration   BIGINT       NOT NULL,
  created_at TIMESTAMP    NOT NULL DEFAULT NOW(),

  CONSTRAINT pk_metric_id PRIMARY KEY (id)
);

CREATE INDEX idx_metric_site_id ON metric (site_id);

CREATE TABLE site (
  id      BIGSERIAL    NOT NULL,
  url     TEXT       NOT NULL,

  CONSTRAINT pk_site_id PRIMARY KEY (id)
);

INSERT INTO site (url) VALUES
('google.com'),
('youtube.com'),
('facebook.com'),
('baidu.com'),
('wikipedia.org'),
('qq.com'),
('taobao.com'),
('yahoo.com'),
('tmall.com'),
('amazon.com'),
('google.co.in'),
('twitter.com'),
('sohu.com'),
('jd.com'),
('live.com'),
('instagram.com'),
('sina.com.cn'),
('weibo.com'),
('google.co.jp'),
('reddit.com'),
('vk.com'),
('360.cn'),
('login.tmall.com'),
('blogspot.com'),
('yandex.ru'),
('google.com.hk'),
('netflix.com'),
('linkedin.com'),
('pornhub.com'),
('google.com.br'),
('twitch.tv'),
('pages.tmall.com'),
('csdn.net'),
('yahoo.com.jp'),
('mail.ru'),
('aliexpress.com'),
('alipay.com'),
('office.com'),
('google.fr'),
('google.ru'),
('google.co.uk'),
('microsoftonline.com'),
('google.de'),
('ebay.com'),
('microsoft.com'),
('livejasmin.com'),
('t.co'),
('bing.com'),
('xvideos.com'),
('google.ca')
;

-- +migrate Down
DROP TABLE metric;
