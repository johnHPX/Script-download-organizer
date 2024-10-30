# Script-download-organizer
<p>Feito em Go, este script é capaz de organizar sua pasta "Downloads" em diretorios rotulados com base nas extensões dos arquivos baixados. É funcional nos sistemas operacionais Linux e Windows!</p>

## Como Funciona: 
Ao executar pela primeira vez, o script imediatamente cria pastas rotuladas (nomes) de acordo com as principais extensões presentes em arquivos baixados da internet (como imagens, videos, arquivos compactados, códigos, músicas, etc). Após criar cada pasta o script criar uma subpasta, com o dia-mês-ano e move cada arquivo para seu devido diretório. Então, depois o script pergunta se o usuário deseja organizar novamente e caso o usuário digite `yes` o script assim faz. Caso contrário, para encerrar o script, basta digitar `no` ou fechar o terminal.

## Como usar:

### Windows:

<p>Basta baixar o arquivo `windows-script.exe` na pasta builds e executar com dois cliques no botão direito do mouse.</p>

### Linux:

<p>Basta baixar o arquivo `linux-script` na pasta builds e no terminal executar com o comando:</p>

```./linux-script```

<p><strong>Atenção!</strong> Caso deseje especificar um diretorio (dentro da pasta Downloads do seu sistema operacional) basta adicionar o paramêntro `-d` e depois o caminho.</p>

```./linux-script -d DownladsChromium```

### Ou gere seu próprio build!

<p>Se preferir, você pode instalar em seu sistema operacional, o Go na versão 1.18 ou superior e gerar seu próprio build ou executar direto.</p>

<p>Gerando o executavel e depois executando:</p>

```
go build
./main

```

<p>Executando direto:</p>

```
go run cmd/main.go

```
