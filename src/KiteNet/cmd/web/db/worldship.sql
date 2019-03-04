create table baidu_openid
(
  newOpenid varchar(50) not null
    primary key,
  oldOpenid varchar(50) not null
);


create table game_region
(
  id       int         not null
    primary key,
  name     varchar(20) not null,
  hot      int         not null,
  state    int         not null,
  initTime datetime    null
);

INSERT INTO worldship.game_region (id, name, hot, state, initTime) VALUES (1, 'test001', 1, 1, '2018-12-14 17:36:50');
INSERT INTO worldship.game_region (id, name, hot, state, initTime) VALUES (2, 'test002', 1, 1, '2018-12-14 17:40:16');
create table lucky
(
  diamond bigint not null
    primary key
);


create table pay
(
  billno   varchar(100)                          not null
    primary key,
  openid   varchar(50)                           not null,
  region   int                                   not null,
  uid      bigint                                not null,
  channel  int                                   not null,
  device   varchar(50)                           not null,
  account  varchar(64)                           null,
  money    int                                   null,
  orderid  varchar(100)                          null,
  level    int                                   not null,
  viplevel int                                   not null,
  status   int                                   not null,
  time     timestamp   default CURRENT_TIMESTAMP not null,
  payway   varchar(50)                           not null,
  origin   varchar(20) default '0'               not null
);


create table pay_billno
(
  billno  varchar(100)                          not null
    primary key,
  orderno varchar(100)                          not null,
  openid  varchar(50)                           not null,
  region  int                                   not null,
  uid     bigint                                not null,
  itemId  int                                   not null,
  channel int                                   not null,
  origin  varchar(20) default '0'               not null,
  device  varchar(50)                           not null,
  time    timestamp   default CURRENT_TIMESTAMP not null
);


create table pay_token
(
  openid    varchar(50)                           not null,
  region    int                                   not null,
  uid       bigint                                not null,
  qq_openid varchar(64)                           not null,
  token     varchar(100)                          not null,
  device    varchar(50)                           not null,
  time      timestamp   default CURRENT_TIMESTAMP not null,
  origin    varchar(20) default '0'               not null
);

create index i_token
  on pay_token (token);


create table review
(
  openid varchar(50)                            not null
    primary key,
  uid    bigint                                 not null,
  time   datetime default '2014-12-04 00:00:00' not null
);


create table share
(
  uid     bigint                              not null
    primary key,
  channel int                                 not null,
  level   int                                 not null,
  time    timestamp default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP
);


create table user
(
  id      bigint unsigned  not null
    primary key,
  code    varchar(50)      not null,
  type    varchar(50)      not null,
  channel int(11) unsigned not null,
  constraint i_code
    unique (code)
);


create table user_account
(
  account  varchar(50) not null
    primary key,
  password varchar(50) not null,
  openid   varchar(50) not null,
  mail     varchar(50) null
);

INSERT INTO worldship.user_account (account, password, openid, mail) VALUES ('10001-gqhuY', '', '1544683261109017462', null);
create table version
(
  `release` varchar(20) not null,
  test      varchar(20) not null
);

INSERT INTO worldship.version (`release`, test) VALUES ('1', '2');
create table whitelist
(
  type  int         not null,
  id    varchar(50) not null,
  state int         not null,
  notes varchar(50) null
);

