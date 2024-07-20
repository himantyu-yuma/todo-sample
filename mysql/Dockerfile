FROM mysql:8.0.28

# MySQLのaptパッケージの公開鍵を明示的に追加
# https://askubuntu.com/questions/1497140/how-to-fix-apt-update-after-to-upgrade-the-mysql-list-file
RUN apt-key --keyring /etc/apt/keyrings/mysql.gpg adv --keyserver keyserver.ubuntu.com --recv-keys B7B3B788A8D3785C

RUN apt-get update && apt-get -y install locales-all

ENV LANG ja_JP.UTF-8
ENV LANGUAGE ja_JP:ja
ENV LC_ALL ja_JP.UTF-8

COPY mysqld_charset.cnf /etc/mysql/conf.d/mysqld_charset.cnf
