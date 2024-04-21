FROM debian:buster

# 不要なパッケージのインストールと特権昇格（パスワードなし）の許可設定
RUN apt update && \
    apt install -y apache2 libtinfo5 netcat sudo wget && \
    a2enmod cgid && \
    groupadd wheel && \
    usermod -aG wheel www-data && \
    echo "%wheel ALL=NOPASSWD: ALL" >> /etc/sudoers && \
    apt clean && \
    rm -rf /var/lib/apt/lists/*

# 脆弱性（CVE-2014-6271）を含むbashのインストール
RUN wget http://snapshot.debian.org/archive/debian/20130101T091755Z/pool/main/b/bash/bash_4.2%2Bdfsg-0.1_amd64.deb -O /tmp/bash_4.2_amd64.deb && \
    dpkg -i /tmp/bash_4.2_amd64.deb

# CGIプログラムの配置
COPY vulnerable /usr/lib/cgi-bin/

# ダミーファイルの配置
COPY danger.html /var/www/html/

RUN chown www-data:www-data /var/www/html/* && \
    chown www-data:www-data /usr/lib/cgi-bin/vulnerable && \
    chmod +x /usr/lib/cgi-bin/vulnerable

EXPOSE 80

CMD ["/usr/sbin/apache2ctl", "-DFOREGROUND"]