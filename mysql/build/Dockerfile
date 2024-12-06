#使うDockerイメージ
FROM mysql:8.1

#MySQL設定ファイルをイメージ内にコピー
COPY /mysql/db/my.cnf ./etc/mysql/conf.d

#docker run時に実行される
CMD ["mysqld"]
