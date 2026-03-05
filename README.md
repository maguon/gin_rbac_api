## server run

```shell
swag init --parseDependency --parseInternal
```

```shell
go get github.com/cosmtrek/air
air init
air
```


## install zhparser
```
wget http://www.xunsearch.com/scws/down/scws-1.2.3.tar.bz2
tar -xvf scws-1.2.3.tar.bz2 
./configure
make && make install

git clone https://github.com/amutu/zhparser.git
cd zhparser
make PG_CONFIG=/usr/pgsql-14/bin/pg_config
make install PG_CONFIG=/usr/pgsql-14/bin/pg_config

psql -U postgres -W
\c dbname
CREATE EXTENSION zhparser;
CREATE TEXT SEARCH CONFIGURATION zhcfg (PARSER = zhparser);
ALTER TEXT SEARCH CONFIGURATION zhcfg ADD MAPPING FOR n,v,a,i,e,l WITH simple;
SELECT to_tsvector('zhcfg','“今年保障房新开工数量虽然有所下调，但实际的年度在建规模以及竣工规模会超以往年份，相对应的对资金的需求也会创历>史纪录。”陈国强说。在他看来，与2011年相比，2012年的保障房建设在资金配套上的压力将更为严峻。');
SELECT to_tsvector('“今年保障房新开工数量虽然有所下调，但实际的年度在建规模以及竣工规模会超以往年份，相对应的对资金的需求也会创历>史纪录。”陈国强说。在他看来，与2011年相比，2012年的保障房建设在资金配套上的压力将更为严峻。');
SELECT to_tsquery('zhcfg','保障房资金压力');
```
