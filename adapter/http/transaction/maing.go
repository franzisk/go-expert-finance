package transaction

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/franzisk/go-expert-finance/model/transaction"
	"github.com/franzisk/go-expert-finance/util"
)

// CraeteTransaction para criar uma nova transação financeira
func CraeteTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("Corpo recebido:", string(body)) // Verifique o conteúdo recebido

	var res = transaction.Transactions{}

	// Faz a leitura do array de bytes, converte e joga toda a informação na variável de referência
	if err := json.Unmarshal(body, &res); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println("Erro no Unmarshal:", err) // Mostre o erro no console
		return
	}

	fmt.Print(res)
}

// GetTransactions vai retornar uma lista de transações
func GetTransactions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	encoder := json.NewEncoder(w)

	var transactions = transaction.Transactions{
		transaction.Transaction{
			Title:     "Pagamento Boleto",
			Amount:    1564.55,
			Type:      1,
			CreatedAt: util.StringToDateTime("2024-09-25T11:45:21"),
		},
	}

	// Para retornar um JSON precisa fazer um Encode com o pacote json
	if err := encoder.Encode(transactions); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
