Este projeto Go é uma ferramenta para monitorar a disponibilidade de links de sites. Ele verifica periodicamente (a cada 10 minutos) se os links estão funcionando e registra os resultados em um arquivo de log.

Como funciona:

Lista de URLs: O projeto lê uma lista de URLs de um arquivo de configuração (url.txt).
Verificação de Status: A cada 10 minutos, o projeto envia uma solicitação HTTP para cada URL na lista e verifica o código de status da resposta.
Registro de Log: Os resultados da verificação (URL, código de status, timestamp) são registrados em um arquivo de log.
Agendamento: O projeto usa um agendador para repetir a verificação a cada 10 minutos.
