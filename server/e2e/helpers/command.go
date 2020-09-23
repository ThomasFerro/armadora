package helpers
import (
	"bytes"
	"fmt"
	"net/http"
	"encoding/json"

	"github.com/ThomasFerro/armadora/infra"
)

func PostACommand(partyId string, command infra.Command, step string) error {
	url := fmt.Sprintf("http://localhost/parties/%v", partyId)
	marshalled, err := json.Marshal(command)
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewReader(marshalled))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return fmt.Errorf("%v - Wrong response code: %v", step, resp.StatusCode)
	}
	return nil
}