#Arquivo de configuração do NGINX
server {
    server_name imc.com www.imc.com;

    location / {
        proxy_pass http://localhost:9990;
    }
}