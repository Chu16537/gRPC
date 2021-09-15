rm *.pem

# 測試用可以不用密碼 -nodes 
# 1. 產生ca 私鑰 ＆ 簽名 ＆ csr
# openssl req -x509 -newkey rsa:4096 -days 365 -keyout ca-key.pem -out ca-cert.pem -subj "/C=TW/ST=Taiwan/L=Tianan/O=DK/OU=DK/CN=*.Chuxin/emailAddress=aa@chuxin.com"
openssl req -x509 -newkey rsa:4096 -days 365 -nodes -keyout ca-key.pem -out ca-cert.pem -subj "/C=TW/ST=Taiwan/L=Tianan/O=DK/OU=DK/CN=*.Chuxin/emailAddress=aa@chuxin.com"

# openssl x509 -in ca-cert.pem -noout -text

# 2. 產生 server 的 私鑰 ＆ csr 
# openssl req -newkey rsa:4096 -keyout server-key.pem -out server-req.pem -subj "/C=TW/ST=Taiwan/L=Tianan/O=s_DK/OU=s_DK/CN=*.s_Chuxin/emailAddress=aa@s_chuxin.com"
openssl req -newkey rsa:4096 -nodes -keyout server-key.pem -out server-req.pem -subj "/C=TW/ST=Taiwan/L=Tianan/O=s_DK/OU=s_DK/CN=*.s_Chuxin/emailAddress=aa@s_chuxin.com"


# 3. 用 ca 私鑰 對 server 的 csr 簽名 並取證書
openssl x509 -req -in server-req.pem -days 60 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out server-cert.pem -extfile server-ext.cnf

# openssl x509 -in server-cert.pem -noout -text

# 4. 創建 client 的 私鑰 ＆ csr 
openssl req -newkey rsa:4096 -nodes -keyout client-key.pem -out client-req.pem -subj "/C=TW/ST=Taiwan/L=Tianan/O=DK/OU=DK/CN=*.client.com/emailAddress=client@chuxin.com"

# 5.用 ca 私鑰 對 ㄏclient 的 csr 簽名 並取證書
openssl x509 -req -in client-req.pem -days 60 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out client-cert.pem -extfile client-ext.cnf



#  確認 證書 是否有效
# openssl verify -CAfile ca-cert.pem server-cert.pem
