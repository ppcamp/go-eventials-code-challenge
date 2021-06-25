package pkg

// Para ler arquivos existem 2 abordagens
// - Ler o arquivo inteiro: Problema: arquivos grandes
// - Ler pedaços do arquivo: Problema: relação inversa de I/O por pedaço

// Dessa forma, fazendo uma modificação sobre a abordagem proposta
// por @ohm.patel1997, tenho uma função que será usada sobre cada linha
// dos consumidores e cada consumidor irá processar uma quantidade máxima
// de `ProcLines` de linhas, podendo ser menos, de acordo com o limite
// do `ChunkSize`

import (
	"bufio"
	"io"
	"log"
	"math"
	"os"
	"strings"
	"sync"
)

// Tamanho de bytes lidos de uma única vez
const ChunkSize = 250 * 1024

// Quantidade de linhas processadas por cada produtor
const ProcLines = 300

// Usando esta função, somente será necessário o envio do arquivo
// aberto. E a stream irá começar a partir da primeira linha
func SetUp(f *os.File, cb func(*string)) {
	// Abrindo por pedaços
	fileStream := bufio.NewReader(f)

	ProcessFile(*fileStream, cb)
}

// Abre pedaços de um arquivo e processa estes pedaços
// retornando para o listener o resultado processado.
//
// Funciona como um produtor-consumidor
// Retorna para o listener cada linha processada por cada um
// dos consumidores. Onde cada consumidor irá trabalhar sobre um
// chunk de dados.
//
// Esta função faz as chamadas para a criação dos consumidores
// Based on
func ProcessFile(fileStream bufio.Reader, cb func(*string)) error {

	buffPool := sync.Pool{New: func() interface{} {
		lines := make([]byte, ChunkSize)
		return lines
	}}

	stringPool := sync.Pool{New: func() interface{} {
		lines := ""
		return lines
	}}

	// Abrindo por pedaços
	// fileStream := bufio.NewReader(f)

	// Produtor-consumidor handler
	var wg sync.WaitGroup

	for {
		buf := buffPool.Get().([]byte)

		nbytes, err := fileStream.Read(buf)
		buf = buf[:nbytes] // corta o buffer para o total de chars lidos

		//region Faz algumas verificações sobre a leitura
		if nbytes == 0 {
			if err == io.EOF {
				// log.Println("Reached the end of the file")
				break
			}
			if err != nil {
				log.Panic(err)
			}
			return err
		}
		//endregion

		// Modifica o buffer, obrigando a ler até encontrar uma quebra de linha
		// isto garante que não tenha uma linha "incompleta" na leitura
		nextUntillNewline, err := fileStream.ReadBytes('\n')
		if err != io.EOF {
			buf = append(buf, nextUntillNewline...)
		}
		wg.Add(1)

		go func() {
			processFileChunk(buf, &buffPool, &stringPool, cb)
			// consume this chunk
			wg.Done()
		}()

	}

	// Só libera a função dps de fazer a chamada de todas as
	// goroutines. Este processamento é garantido pela variável wg
	// Pode ser entendido como uma fila de callbacks a serem processados
	wg.Wait()
	return nil
}

// Realiza o processamento dos produtores obtidos através do *ProcessFile*
//
// Ele cria vários produtores que irão processar o arquivo
// Cada consumidor passa a linha no momento para a callback function
// Cada consumidor trabalha processa uma relação de quantidade de linhas x chunk
// NOTE: que nessa implementação,posso ter mais threads do que a máquina possui
// Sendo baseada apenas no número de linhas máximo que cada produtor vai operar
func processFileChunk(
	buff []byte,
	buffPool *sync.Pool,
	stringPool *sync.Pool,
	cb func(*string)) {

	//region Buffer[]
	chunk := stringPool.Get().(string) // apenas para ter o tamanho máximo
	chunk = string(buff)               // cria a string
	buffPool.Put(&buff)
	//region

	//region String[]
	lines := strings.Split(chunk, "\n")
	stringPool.Put(&chunk)
	//endregion

	//region Obtendo configurações
	n := len(lines)             // num de linhas neste chunk
	noOfThread := n / ProcLines // num de produtores é baseado no num de linhas
	if n%ProcLines != 0 {
		noOfThread++
	}
	//endregion

	//region Executando tasks assíncronas de processamento
	// Flag de produtor-consumidor
	var wg2 sync.WaitGroup
	for i := 0; i < (noOfThread); i++ {
		wg2.Add(1)

		slen := i * ProcLines
		elen := int(math.Min(float64((i+1)*ProcLines), float64(len(lines))))

		go func(s int, e int) {
			defer wg2.Done() //to avoid deadlocks

			for i := s; i < e; i++ {
				text := lines[i]
				if len(text) == 0 {
					continue
				}
				cb(&text)
			}
		}(slen, elen)
	}
	wg2.Wait()
	//endregion

	lines = nil
}
