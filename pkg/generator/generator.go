package generator

import (
	"log"
	"os"

	"github.com/google/uuid"
)

var ids []uuid.UUID

func GenUUID(n int) ([]uuid.UUID, error) {
	os.Remove("result.txt")
	f, err := os.Create("result.txt")
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; n > i; i++ {
		id, err := uuid.NewRandom()
		if err != nil {
			return ids, err
		}
		ids = append(ids, id)
		f.Write([]byte(id.String() + "\n"))
	}
	return ids, nil
}
