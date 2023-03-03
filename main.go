package bitrix_go

import (
	"bitrixgo/src/client"
	"bitrixgo/src/types"
	"log"
)

func main() {
	c, err := client.NewClientWithWebhookAuth("https://bitrix.domain.itgrn", 1, "e0uqe8cvwso2mlbh")
	if err != nil {
		log.Println(err)
	}

	res2, _ := c.UpdateTask(12, types.Task{Title: "Go Editor", Responsible_id: 1})
	log.Println(res2)
}
