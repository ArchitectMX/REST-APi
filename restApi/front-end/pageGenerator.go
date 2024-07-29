package front_end

import (
	"advanced-tasks/berendeev/restApi/handler"
)

func PageGenerator() {
	go handler.HomePage()

	go handler.DeleteHandler()

	go handler.PostHandler()

	go handler.EditHandler()

	go handler.EditStatusHandler()
}
