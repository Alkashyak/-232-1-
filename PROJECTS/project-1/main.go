package main

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/golang-jwt/jwt/v5"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const (
	AsanIP   = "26.251.53.172:8081"
	BogdanIP = "26.127.253.74:8082"
)

func main() {
	sessii := make(map[int64]string)
	//  токен, который получили от BotFather.
	// Создаем экземпляр бота, используя токен
	bot, err := tgbotapi.NewBotAPI("6431776225:AAGEwnFM-XIY3ykX5kT7_3G-6lG8Dnw-waA")
	if err != nil {
		log.Fatal(err) // В случае ошибки при создании бота, выводим сообщение об ошибке и завершаем программу
	}

	bot.Debug = true // Включаем отладочный режим для бота

	// Настраиваем обработчик сообщений
	updateConfig := tgbotapi.NewUpdate(0)            // Создаем объект настроек для получения обновлений с указанного идентификатора
	updateConfig.Timeout = 60                        // Устанавливаем таймаут на получение обновлений в секундах
	updates, err := bot.GetUpdatesChan(updateConfig) // Получаем канал для получения обновлений от бота
	//Callback для авторизации для получения github_id, если он получин записываем его в map и ваводим сообщение об успешной авторизации
	http.HandleFunc("/gitid", func(w http.ResponseWriter, r *http.Request) { // Обработчик отвечающий на запроса к /gitid
		log.Printf("github_id:.")
		github_id := r.URL.Query().Get("githubid")
		chat_id, _ := strconv.ParseInt(r.URL.Query().Get("chatid"), 10, 64)
		log.Printf("github_id: %s", github_id)
		if github_id != "" {
			sessii[chat_id] = github_id
			bot.Send(tgbotapi.NewMessage(chat_id, "Вы успешно авторизировались!"))
		}
	})
	go func() { http.ListenAndServe(":6969", nil) }() // Запуск сервера на порту 6969

	log.Printf("Authorized on account %s", bot.Self.UserName)
	//Создаем новый объект обновления

	//программа настраивает канал обновлений для получения обновлений от бота, используя "bot.GetUpdatesChan(u)"
	// Запускаем бесконечный цикл для обработки обновлений, получаемых из канала updates
	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		//создание нового сообщения для отправки пользователю в указанный чат
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		//новое сообщение будет ответом на существующее сообщение с ID update.Message.MessageID
		msg.ReplyToMessageID = update.Message.MessageID
		//создание  клавиатуры с кнопками для ответов
		msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton("Где следующая пара"),
				tgbotapi.NewKeyboardButton("Расписание на сегодня"),
				tgbotapi.NewKeyboardButton("Расписание на завтра"),
				tgbotapi.NewKeyboardButton("Расписание на дни недели"),
				tgbotapi.NewKeyboardButton("Где преподаватель"),
			),
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton("Изменить данные(ФИО, группа)"),
				tgbotapi.NewKeyboardButton("Оставить комментарий к паре"),
				tgbotapi.NewKeyboardButton("Где группа"),
				tgbotapi.NewKeyboardButton("toadmin"),
				tgbotapi.NewKeyboardButton("Выйти"),
			),
		)
		if !checkS(sessii, update.Message.Chat.ID) { //проверяем есть ли для пользоателя открытая сессия
			msg.Text = "Привет! Я телеграмм бот c расписанием. \nЧтобы продолжить пользоваться вам нужно авторизироваться. \n "
			bot.Send(msg)
			msg.Text = link(strconv.FormatInt(update.Message.Chat.ID, 10))
			bot.Send(msg)
		} else {
			if check_data(update.Message.Chat.ID) == "true" {
				switch update.Message.Text {
				case "/start":
					msg.Text = "Привет! Я телеграмм бот c расписанием. \nНажми /help чтобы увидеть все команды."
				case "/help":
					msg.Text = "Список всех команд: \n- Где следующая пара\n- Расписание на день недели\n- Расписание на сегодня\n- Расписание на завтра\n- Оставить комментарий к паре /n- Где группа \n- Где преподаватель\n- toadmin"
				case "toadmin":
					if checkPrava(update.Message.Chat.ID, "Admin") { //Проверка роли для перехода в админ панель
						// Создаем HTTP-клиент
						client := http.Client{}
						// Формируем URL для HTTP-запроса, используя AsanIP
						requestURL := fmt.Sprintf("http://" + AsanIP + "/to_admin")
						// Создаем объект url.Values для передачи параметров в теле запроса
						form := url.Values{}
						form.Add("jwt", request_jwt_admin(sessii[update.Message.Chat.ID]))
						// Создаем новый HTTP-запрос методом POST по указанному URL
						request, _ := http.NewRequest("POST", requestURL, strings.NewReader(form.Encode()))
						// Устанавливаем заголовок Content-Type для указания типа передаваемых данных
						request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
						// Выполняем HTTP-запрос
						response, _ := client.Do(request)
						// Читаем тело ответа
						resBody, _ := io.ReadAll(response.Body)
						// Устанавливаем текст сообщения для отправки в Telegram как содержимое тела ответа
						msg.Text = string(resBody)
						// Отправляем сообщение в Telegram
						bot.Send(msg)
						// Закрываем тело ответа после использования
						defer response.Body.Close()
					} else {
						msg.Text = "Недостаточно прав"
						bot.Send(msg)
					}
				case "Где следующая пара":
					msg.Text = getSchedule(sessii, update.Message.Chat.ID, "get_schedule_next_lesson")
					bot.Send(msg)
				case "Расписание на сегодня":
					msg.Text = getSchedule(sessii, update.Message.Chat.ID, "get_schedule_that_day")
					bot.Send(msg)
				case "Расписание на завтра":
					msg.Text = getSchedule(sessii, update.Message.Chat.ID, "get_schedule_next_day")
					bot.Send(msg)
				case "Расписание на дни недели":
					newKeyboard := tgbotapi.NewMessage(update.Message.Chat.ID, "Выберите день недели:")
					newKeyboard.ReplyMarkup = tgbotapi.NewReplyKeyboard(
						tgbotapi.NewKeyboardButtonRow(
							tgbotapi.NewKeyboardButton("Расписание на понедельник"),
							tgbotapi.NewKeyboardButton("Расписание на вторник"),
							tgbotapi.NewKeyboardButton("Расписание на среду"),
							tgbotapi.NewKeyboardButton("Расписание на четверг"),
						),
						tgbotapi.NewKeyboardButtonRow(
							tgbotapi.NewKeyboardButton("Расписание на пятницу"),
							tgbotapi.NewKeyboardButton("Расписание на субботу"),
							tgbotapi.NewKeyboardButton("Расписание на воскресенье"),
							tgbotapi.NewKeyboardButton("Выйти"),
						),
					)
					bot.Send(newKeyboard)
				case "Расписание на понедельник":
					msg.Text = getSchedule(sessii, update.Message.Chat.ID, "0")
					bot.Send(msg)
				case "Расписание на вторник":
					msg.Text = getSchedule(sessii, update.Message.Chat.ID, "1")
					bot.Send(msg)
				case "Расписание на среду":
					msg.Text = getSchedule(sessii, update.Message.Chat.ID, "2")
					bot.Send(msg)
				case "Расписание на четверг":
					msg.Text = getSchedule(sessii, update.Message.Chat.ID, "3")
					bot.Send(msg)
				case "Расписание на пятницу":
					msg.Text = getSchedule(sessii, update.Message.Chat.ID, "4")
					bot.Send(msg)
				case "Расписание на субботу":
					msg.Text = "Выходной. В данный день пар нет."
					bot.Send(msg)
				case "Расписание на воскресенье":
					msg.Text = "Выходной. В данный день пар нет."
					bot.Send(msg)
				case "Выйти":
					delete(sessii, update.Message.Chat.ID)
					msg.Text = "Вы успешно вышли!"
					bot.Send(msg)
				case "Оставить комментарий к паре":
					if checkPrava(update.Message.Chat.ID, "Admin") || checkPrava(update.Message.Chat.ID, "Write") { //проверка роли для добавления коментария к паре
						var num_of_lesson, group, comment string
						msg.Text = "Введите номер пары"
						bot.Send(msg)
						for update := range updates {
							if update.Message == nil { // ignore any non-Message Updates
								continue
							}
							num_of_lesson = update.Message.Text
							break
						}
						msg.Text = "Введите номер группы (Например ПИ-б-о-232)"
						bot.Send(msg)
						for update := range updates {
							if update.Message == nil { // ignore any non-Message Updates
								continue
							}
							group = "группа " + update.Message.Text
							break
						}
						subgroup := "0"
						msg.Text = "Введите комментарий"
						bot.Send(msg)
						for update := range updates {
							if update.Message == nil { // ignore any non-Message Updates
								continue
							}
							comment = update.Message.Text
							break
						}
						chetW := ""
						msg.Text = "Введите четность недели (0 - чет; 1 - нечет)"
						bot.Send(msg)
						for update := range updates {
							if update.Message == nil { // ignore any non-Message Updates
								continue
							}
							chetW = update.Message.Text
							break
						}
						msg.Text = "Введите день недели (от 1 до 5)"
						bot.Send(msg)
						var Wkday = ""
						for update := range updates {
							if update.Message == nil { // ignore any non-Message Updates
								continue
							}
							Wkday = update.Message.Text
							break
						}
						// Создаем экземпляр HTTP-клиента
						client := http.Client{}
						// Формируем URL для HTTP-GET запроса, используя BogdanIP и функцию request_jwt_comment для создания JWT-токена
						requestURL := fmt.Sprintf("http://" + BogdanIP + "/get_change?jwtok=" + request_jwt_comment(sessii[update.Message.Chat.ID], "getting_prepod_comment", num_of_lesson, group, subgroup, chetW, comment, Wkday))
						// Создаем новый HTTP-GET запрос по указанному URL
						request, _ := http.NewRequest("GET", requestURL, nil)
						// Выполняем HTTP-GET запрос
						response, _ := client.Do(request)
						// Читаем тело ответа
						resBody, _ := io.ReadAll(response.Body)
						// Выводим тело ответа в лог
						log.Println(string(resBody))
						// Устанавливаем текст сообщения для отправки в Telegram как содержимое тела ответа
						msg.Text = string(resBody)
						// Отправляем сообщение в Telegram
						bot.Send(msg)
						// Закрываем тело ответа после использования
						defer response.Body.Close()
					} else {
						msg.Text = "Недостаточно прав"
						bot.Send(msg)
					}
				case "Где группа":
					if checkPrava(update.Message.Chat.ID, "Admin") || checkPrava(update.Message.Chat.ID, "Write") { //проверка роли
						group := ""
						msg.Text = "Введите номер группы (Например ПИ-б-о-232)"
						bot.Send(msg)
						for update := range updates {
							if update.Message == nil { // ignore any non-Message Updates
								continue
							}
							group = "группа " + update.Message.Text
							break
						}
						subgroup := "0"
						msg.Text = "Введите номер подгруппы (1 или 2)"
						bot.Send(msg)
						for update := range updates {
							if update.Message == nil { // ignore any non-Message Updates
								continue
							}
							subgroup = update.Message.Text
							break
						}
						client := http.Client{}
						requestURL := fmt.Sprintf("http://" + BogdanIP + "/get_change?jwtok=" + request_jwt_stud_loc("getting_students_location", group, subgroup))
						request, _ := http.NewRequest("GET", requestURL, nil)
						response, _ := client.Do(request)
						resBody, _ := io.ReadAll(response.Body)
						msg.Text = string(resBody)
						bot.Send(msg)
						defer response.Body.Close()
					} else {
						msg.Text = "Недостаточно прав"
						bot.Send(msg)
					}
				case "Где преподаватель":
					teacher := ""
					msg.Text = "Введите ФИО преподавателя (прим. Иванов И.И.)"
					bot.Send(msg)
					for update := range updates {
						if update.Message == nil { // ignore any non-Message Updates
							continue
						}
						teacher = update.Message.Text
						break
					}
					client := http.Client{}
					requestURL := fmt.Sprintf("http://" + BogdanIP + "/get_change?jwtok=" + request_jwt_prepod_loc("getting_prepod_location", teacher))
					request, _ := http.NewRequest("GET", requestURL, nil)
					response, _ := client.Do(request)
					resBody, _ := io.ReadAll(response.Body)
					msg.Text = string(resBody)
					bot.Send(msg)
					defer response.Body.Close()
				case "Изменить данные(ФИО, группа)":
					msg.Text = "Отправте ваше ФИО (прим. Иванов И.И.)"
					bot.Send(msg)
					for update := range updates {
						if update.Message == nil { // ignore any non-Message Updates
							continue
						}
						if send_data(update.Message.Chat.ID, update.Message.Text, "username") == "true" {
							msg.Text = "Данные успешно записаны."
							bot.Send(msg)
							msg.Text = ""
							break
						} else {
							msg.Text = "Ошибка! Не удалось записать данные."
							bot.Send(msg)
							msg.Text = ""
						}
					}
					msg.Text = "Отправте вашу группу (прим. ПИ-б-о-232)"
					bot.Send(msg)
					for update := range updates {
						if send_data(update.Message.Chat.ID, "группа "+update.Message.Text, "group") == "true" {
							msg.Text = "Данные успешно записаны."
							bot.Send(msg)
							msg.Text = ""
							break
						} else {
							msg.Text = "Ошибка! Не удалось записать данные."
							bot.Send(msg)
							msg.Text = ""
						}
					}
					msg.Text = "Отправте вашу подгруппу (прим. 1 или 2)"
					bot.Send(msg)
					for update := range updates {
						if send_data(update.Message.Chat.ID, update.Message.Text, "subgroup") == "true" {
							msg.Text = "Данные успешно записаны."
							bot.Send(msg)
							msg.Text = ""
							break
						} else {
							msg.Text = "Ошибка! Не удалось записать данные."
							bot.Send(msg)
							msg.Text = ""
						}
					}
				case "Назад":
					newKeyboard1 := tgbotapi.NewMessage(update.Message.Chat.ID, "Выберите действие:")
					newKeyboard1.ReplyMarkup = tgbotapi.NewReplyKeyboard(
						tgbotapi.NewKeyboardButtonRow(
							tgbotapi.NewKeyboardButton("Где следующая пара"),
							tgbotapi.NewKeyboardButton("Расписание на сегодня"),
							tgbotapi.NewKeyboardButton("Расписание на завтра"),
							tgbotapi.NewKeyboardButton("Расписание на дни недели"),
							tgbotapi.NewKeyboardButton("Где преподаватель"),
						),
						tgbotapi.NewKeyboardButtonRow(
							tgbotapi.NewKeyboardButton("Изменить данные(ФИО, группа)"),
							tgbotapi.NewKeyboardButton("Оставить комментарий к паре"),
							tgbotapi.NewKeyboardButton("Где группа"),
							tgbotapi.NewKeyboardButton("toadmin"),
							tgbotapi.NewKeyboardButton("Выйти"),
						),
					)
					bot.Send(newKeyboard1)
				default:
					msg.Text = "Я не понимаю, что вы хотите сказать."
					bot.Send(msg)
				}
			} else {
				msg.Text = "Необходимо отправить данные!"
				bot.Send(msg)
				msg.Text = "Отправте ваше ФИО (прим. Иванов И.И.)"
				bot.Send(msg)
				for update := range updates {
					if update.Message == nil { // ignore any non-Message Updates
						continue
					}
					if send_data(update.Message.Chat.ID, update.Message.Text, "username") == "true" {
						msg.Text = "Данные успешно записаны."
						bot.Send(msg)
						msg.Text = ""
						break
					} else {
						msg.Text = "Ошибка! Не удалось записать данные."
						bot.Send(msg)
						msg.Text = ""
					}
				}
				msg.Text = "Отправте вашу группу (прим. ПИ-б-о-232)"
				bot.Send(msg)
				for update := range updates {
					if send_data(update.Message.Chat.ID, "группа "+update.Message.Text, "group") == "true" {
						msg.Text = "Данные успешно записаны."
						bot.Send(msg)
						msg.Text = ""
						break
					} else {
						msg.Text = "Ошибка! Не удалось записать данные."
						bot.Send(msg)
						msg.Text = ""
					}
				}
				msg.Text = "Отправте вашу подгруппу (прим. 1 или 2)"
				bot.Send(msg)
				for update := range updates {
					if send_data(update.Message.Chat.ID, update.Message.Text, "subgroup") == "true" {
						msg.Text = "Данные успешно записаны."
						bot.Send(msg)
						msg.Text = ""
						break
					} else {
						msg.Text = "Ошибка! Не удалось записать данные."
						bot.Send(msg)
						msg.Text = ""
					}
				}
			}
		}
	}
}

func link(id string) string {
	// Создаём http-клиент с дефолтными настройками
	client := http.Client{}

	// Формируем строку запроса вместе с query string
	requestURL := fmt.Sprintf("http://"+AsanIP+"/reg?Id=%s", id)

	// Выполняем запрос на сервер. Ответ попадёт в переменную response
	request, _ := http.NewRequest("GET", requestURL, nil)
	response, _ := client.Do(request)

	resBody, _ := io.ReadAll(response.Body) // Получаем тело ответа
	defer response.Body.Close()             // Закрывает соединение с сервером

	return string(resBody)
}

func find(id string) map[string]interface{} {
	// Создаём http-клиент с дефолтными настройками
	client := http.Client{}

	// Формируем строку запроса вместе с query string
	requestURL := fmt.Sprintf("http://"+AsanIP+"/find?Tg_id=%s", id)

	// Выполняем запрос на сервер. Ответ попадёт в переменную response
	request, _ := http.NewRequest("GET", requestURL, nil)
	response, _ := client.Do(request)

	resBody, _ := io.ReadAll(response.Body) // Получаем тело ответа
	defer response.Body.Close()             // Закрывает соединение с сервером
	ans, _ := decode(string(resBody))
	return ans
}

func decode(tokenString string) (map[string]interface{}, error) {
	// Парсим токен, используя секретный ключ "7777777"
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("7777777"), nil
	})
	// Преобразуем токен в map с данными пользователя
	user, ok := token.Claims.(jwt.MapClaims)
	// Проверяем, что преобразование прошло успешно и токен валиден
	if ok && token.Valid {
		// Извлекаем права пользователя из токена и возвращаем их
		r := user["rights"]
		return r.(map[string]interface{}), nil
	} else {
		// В случае ошибки возвращаем пустую map и ошибку
		var empt map[string]interface{}
		return empt, err
	}
}

func checkS(sessii map[int64]string, Chat_ID int64) bool {
	for user, _ := range sessii {
		if user == Chat_ID {
			return true
		}
	}
	return false
}

// GET запрос к авторизации для проверки наличия всех данных(ФИО, группа)
func check_data(Chat_ID int64) string {
	client := http.Client{}
	requestURL := fmt.Sprintf("http://"+AsanIP+"/checkAbout?chatid=%d", Chat_ID)
	request, _ := http.NewRequest("GET", requestURL, nil)
	response, _ := client.Do(request)
	resBody, _ := io.ReadAll(response.Body)
	defer response.Body.Close()
	return string(resBody)
}

// GET запрос к авторизации для добавленния данных
func send_data(Chat_ID int64, message string, datatype string) string {
	client := http.Client{}
	requestURL := "http://" + AsanIP + "/updateData"

	form := url.Values{}
	form.Add("chatid", strconv.FormatInt(Chat_ID, 10)) //chat id пользователя которому нужно записать данные
	form.Add("data", message)                          //данные которые нужно записать
	form.Add("datatype", datatype)                     //тип данных (ФИО или группа)

	request, _ := http.NewRequest("POST", requestURL, strings.NewReader(form.Encode()))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, _ := client.Do(request)
	resBody, _ := io.ReadAll(response.Body) // Получаем тело ответ
	defer response.Body.Close()
	return string(resBody)
}

// POST запрос к авторизации для ПОЛУЧЕНИЯ JWT TOKEN
func request_jwt(GIT_ID string, action string) string {
	client := http.Client{}
	requestURL := "http://" + AsanIP + "/getJWT/schedule"

	form := url.Values{}
	form.Add("gitid", GIT_ID)
	form.Add("action", action)

	request, _ := http.NewRequest("POST", requestURL, strings.NewReader(form.Encode()))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, _ := client.Do(request)
	resBody, _ := io.ReadAll(response.Body) // Получаем тело ответ
	defer response.Body.Close()
	return string(resBody)
}

func request_jwt_admin(GIT_ID string) string {
	client := http.Client{}
	requestURL := "http://" + AsanIP + "/getJWT/admin"

	form := url.Values{}
	form.Add("gitid", GIT_ID)

	request, _ := http.NewRequest("POST", requestURL, strings.NewReader(form.Encode()))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, _ := client.Do(request)
	resBody, _ := io.ReadAll(response.Body) // Получаем тело ответ
	defer response.Body.Close()
	return string(resBody)
}
func request_jwt_comment(GIT_ID, action, lesson_number, main_group, sub_group, oddevenweek, comment, weekday string) string {
	client := http.Client{}
	requestURL := "http://" + AsanIP + "/getJWT/comment"

	form := url.Values{}
	form.Add("gitid", GIT_ID)
	form.Add("action", action)
	form.Add("lesson_number", lesson_number)
	form.Add("main_group", main_group)
	form.Add("sub_group", sub_group)
	form.Add("oddevenweek", oddevenweek)
	form.Add("comment", comment)
	form.Add("weekday", weekday)
	request, _ := http.NewRequest("POST", requestURL, strings.NewReader(form.Encode()))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, _ := client.Do(request)
	resBody, _ := io.ReadAll(response.Body) // Получаем тело ответ
	defer response.Body.Close()
	return string(resBody)
}
func request_jwt_stud_loc(action, main_group, sub_group string) string {
	client := http.Client{}
	requestURL := "http://" + AsanIP + "/getJWT/studloc"

	form := url.Values{}
	form.Add("action", action)
	form.Add("main_group", main_group)
	form.Add("sub_group", sub_group)
	request, _ := http.NewRequest("POST", requestURL, strings.NewReader(form.Encode()))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, _ := client.Do(request)
	resBody, _ := io.ReadAll(response.Body) // Получаем тело ответ
	defer response.Body.Close()
	return string(resBody)
}
func request_jwt_prepod_loc(action, prepod string) string {
	client := http.Client{}
	requestURL := "http://" + AsanIP + "/getJWT/prepodloc"

	form := url.Values{}
	form.Add("action", action)
	form.Add("prepod", prepod)

	request, _ := http.NewRequest("POST", requestURL, strings.NewReader(form.Encode()))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, _ := client.Do(request)
	resBody, _ := io.ReadAll(response.Body) // Получаем тело ответ
	defer response.Body.Close()
	return string(resBody)
}
func getSchedule(sessii map[int64]string, Chat_ID int64, action string) string {
	client := http.Client{}
	requestURL := fmt.Sprintf("http://" + BogdanIP + "/get_change?jwtok=" + request_jwt(sessii[Chat_ID], action))
	log.Println("token ok")

	request, _ := http.NewRequest("GET", requestURL, nil)

	response, _ := client.Do(request)
	resBody, _ := io.ReadAll(response.Body)
	log.Println(resBody)
	defer response.Body.Close()
	return string(resBody)
}
func checkPrava(Chatid int64, role string) bool {
	roles := find(strconv.FormatInt(Chatid, 10))

	// Проверяем, существует ли роль в словаре
	if val, ok := roles[role]; ok {
		// Если роль существует, возвращаем её статус
		return val.(bool)
	}

	// Если роль не существует, можно вернуть значение по умолчанию или обработать исключение
	return false
}
