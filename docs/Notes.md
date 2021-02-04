
# Anotações

## gRPC

- É um framework gratuito e open-source desenvolvido pelo Google;
- É parte da Cloud Native Computation Foudation (CNCF) como Docker e Kubernetes, por exemplo;
- Em alto nível permite que você defina Request e Response para um RPC e manipula todo o resto para você
- Moderno, rápido e eficiente. Construído sob o protocolo HTTP/2, com baixa latência
- Suporte para streaming
- Fácil de plugar autenticação, load balancing, log e monitoração
- Resolve uma série de problemas encontrados nas antigas implementações de RPC (ex. CORBA, JAVA RMI)

## Protocol Buffers

- É agnóstico da linguagem de programação
- Dados são binários serializados (small payloads)
- Conveniente para grande troca de dados
- Permite uma fácil evolução da API
- São auto gerados stubs e skeletons à partir do arquivo .proto
- Protocol Buffers, utilizado para definir mensagens e serviços
- Protocal Buffer, define regras para evoluir a API sem quebrar os clientes já existentes


## A eficiência do Protocol Buffers em relação ao JSON / XML

- Menor payload, as mensagens são menores, economia de banda de rede
- Parsing JSON/XML exige bastante processamento, são formatos para leitura dos humanos para para processamento do computador
- Protocol Buffers é um formato binários, menor processsamento
- gRPC por usar Protocol Buffers é mais rápido e com comunicação mais eficiente, bom para dispositivos com baixo poder de CPU


## Material de apoio

[Material do Curso gRPC Golang Master Class](https://github.com/simplesteph/grpc-go-course)

[gRPC Web Site](https://grpc.io)
