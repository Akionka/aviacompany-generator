package main

var (
	femaleMidNames = []string{"Александровна",
		"Андреевна",
		"Архиповна",
		"Алексеевна",
		"Антоновна",
		"Аскольдовна",
		"Альбертовна",
		"Аркадьевна",
		"Афанасьевна",
		"Анатольевна",
		"Артемовна",
		"Богдановна",
		"Болеславовна",
		"Борисовна",
		"Вадимовна",
		"Васильевна",
		"Владимировна",
		"Валентиновна",
		"Вениаминовна",
		"Владиславовна",
		"Валериевна",
		"Викторовна",
		"Вячеславовна",
		"Геннадиевна",
		"Георгиевна",
		"Геннадьевна",
		"Григорьевна",
		"Даниловна",
		"Дмитриевна",
		"Евгеньевна",
		"Егоровна",
		"Ефимовна",
		"Ждановна",
		"Захаровна",
		"Ивановна",
		"Игоревна",
		"Ильинична",
		"Кирилловна",
		"Кузьминична",
		"Константиновна",
		"Леонидовна",
		"Леоновна",
		"Львовна",
		"Макаровна",
		"Матвеевна",
		"Михайловна",
		"Максимовна",
		"Мироновна",
		"Натановна",
		"Никифоровна",
		"Ниловна",
		"Наумовна",
		"Николаевна",
		"Олеговна",
		"Оскаровна",
		"Павловна",
		"Петровна",
		"Робертовна",
		"Рубеновна",
		"Руслановна",
		"Романовна",
		"Рудольфовна",
		"Святославовна",
		"Сергеевна",
		"Степановна",
		"Семеновна",
		"Станиславовна",
		"Тарасовна",
		"Тимофеевна",
		"Тимуровна",
		"Федоровна",
		"Феликсовна",
		"Филипповна",
		"Харитоновна",
		"Эдуардовна",
		"Эльдаровна",
		"Юльевна",
		"Юрьевна",
		"Яковлевна"}
	femaleLastNames = []string{"Смирнова",
		"Иванова",
		"Кузнецова",
		"Попова",
		"Соколова",
		"Лебедева",
		"Козлова",
		"Новикова",
		"Морозова",
		"Петрова",
		"Волкова",
		"Соловьева",
		"Васильева",
		"Зайцева",
		"Павлова",
		"Семенова",
		"Голубева",
		"Виноградова",
		"Богданова",
		"Воробьева",
		"Федорова",
		"Михайлова",
		"Беляева",
		"Тарасова",
		"Белова",
		"Комарова",
		"Орлова",
		"Киселева",
		"Макарова",
		"Андреева",
		"Ковалева",
		"Ильина",
		"Гусева",
		"Титова",
		"Кузьмина",
		"Кудрявцева",
		"Баранова",
		"Куликова",
		"Алексеева",
		"Степанова",
		"Яковлева",
		"Сорокина",
		"Сергеева",
		"Романова",
		"Захарова",
		"Борисова",
		"Королева",
		"Герасимова",
		"Пономарева",
		"Григорьева",
		"Лазарева",
		"Медведева",
		"Ершова",
		"Никитина",
		"Соболева",
		"Рябова",
		"Полякова",
		"Цветкова",
		"Данилова",
		"Жукова",
		"Фролова",
		"Журавлева",
		"Николаева",
		"Крылова",
		"Максимова",
		"Сидорова",
		"Осипова",
		"Белоусова",
		"Федотова",
		"Дорофеева",
		"Егорова",
		"Матвеева",
		"Боброва",
		"Дмитриева",
		"Калинина",
		"Анисимова",
		"Петухова",
		"Антонова",
		"Тимофеева",
		"Никифорова",
		"Веселова",
		"Филиппова",
		"Маркова",
		"Большакова",
		"Суханова",
		"Миронова",
		"Ширяева",
		"Александрова",
		"Коновалова",
		"Шестакова",
		"Казакова",
		"Ефимова",
		"Денисова",
		"Громова",
		"Фомина",
		"Давыдова",
		"Мельникова",
		"Щербакова",
		"Блинова",
		"Колесникова",
		"Карпова",
		"Афанасьева",
		"Власова",
		"Маслова",
		"Исакова",
		"Тихонова",
		"Аксенова",
		"Гаврилова",
		"Родионова",
		"Котова",
		"Горбунова",
		"Кудряшова",
		"Быкова",
		"Зуева",
		"Третьякова",
		"Савельева",
		"Панова",
		"Рыбакова",
		"Суворова",
		"Абрамова",
		"Воронова",
		"Мухина",
		"Архипова",
		"Трофимова",
		"Мартынова",
		"Емельянова",
		"Горшкова",
		"Чернова",
		"Овчинникова",
		"Селезнева",
		"Панфилова",
		"Копылова",
		"Михеева",
		"Галкина",
		"Назарова",
		"Лобанова",
		"Лукина",
		"Белякова",
		"Потапова",
		"Некрасова",
		"Хохлова",
		"Жданова",
		"Наумова",
		"Шилова",
		"Воронцова",
		"Ермакова",
		"Дроздова",
		"Игнатьева",
		"Савина",
		"Логинова",
		"Сафонова",
		"Капустина",
		"Кириллова",
		"Моисеева",
		"Елисеева",
		"Кошелева",
		"Костина",
		"Горбачева",
		"Орехова",
		"Ефремова",
		"Исаева",
		"Евдокимова",
		"Калашникова",
		"Кабанова",
		"Носкова",
		"Юдина",
		"Кулагина",
		"Лапина",
		"Прохорова",
		"Нестерова",
		"Харитонова",
		"Агафонова",
		"Муравьева",
		"Ларионова",
		"Федосеева",
		"Зимина",
		"Пахомова",
		"Шубина",
		"Игнатова",
		"Филатова",
		"Крюкова",
		"Рогова",
		"Кулакова",
		"Терентьева",
		"Молчанова",
		"Владимирова",
		"Артемьева",
		"Гурьева",
		"Зиновьева",
		"Гришина",
		"Кононова",
		"Дементьева",
		"Ситникова",
		"Симонова",
		"Мишина",
		"Фадеева",
		"Комиссарова",
		"Мамонтова",
		"Носова",
		"Гуляева",
		"Шарова",
		"Устинова",
		"Вишнякова",
		"Евсеева",
		"Лаврентьева",
		"Брагина",
		"Константинова",
		"Корнилова",
		"Авдеева",
		"Зыкова",
		"Бирюкова",
		"Шарапова",
		"Никонова",
		"Щукина",
		"Дьячкова",
		"Одинцова",
		"Сазонова",
		"Якушева",
		"Красильникова",
		"Гордеева",
		"Самойлова",
		"Князева",
		"Беспалова",
		"Уварова",
		"Шашкова",
		"Бобылева",
		"Доронина",
		"Белозерова",
		"Рожкова",
		"Самсонова",
		"Мясникова",
		"Лихачева",
		"Бурова",
		"Сысоева",
		"Фомичева",
		"Русакова",
		"Стрелкова",
		"Гущина",
		"Тетерина",
		"Колобова",
		"Субботина",
		"Фокина",
		"Блохина",
		"Селиверстова",
		"Пестова",
		"Кондратьева",
		"Силина",
		"Меркушева",
		"Лыткина",
		"Турова",
	}
	femaleNames = []string{"Агата",
		"Агафья",
		"Акулина",
		"Алевтина",
		"Александра",
		"Алина",
		"Алла",
		"Анастасия",
		"Ангелина",
		"Анжела",
		"Анжелика",
		"Анна",
		"Антонина",
		"Валентина",
		"Валерия",
		"Варвара",
		"Василиса",
		"Вера",
		"Вероника",
		"Виктория",
		"Галина",
		"Глафира",
		"Дарья",
		"Евгения",
		"Евдокия",
		"Евпраксия",
		"Евфросиния",
		"Екатерина",
		"Елена",
		"Елизавета",
		"Жанна",
		"Зинаида",
		"Зоя",
		"Иванна",
		"Ираида",
		"Ирина",
		"Ия",
		"Кира",
		"Клавдия",
		"Ксения",
		"Лариса",
		"Лидия",
		"Лора",
		"Лукия",
		"Любовь",
		"Людмила",
		"Майя",
		"Маргарита",
		"Марина",
		"Мария",
		"Марфа",
		"Милица",
		"Надежда",
		"Наина",
		"Наталья",
		"Нина",
		"Нинель",
		"Нонна",
		"Оксана",
		"Октябрина",
		"Олимпиада",
		"Ольга",
		"Пелагея",
		"Полина",
		"Прасковья",
		"Раиса",
		"Регина",
		"Светлана",
		"Синклитикия",
		"София",
		"Таисия",
		"Тамара",
		"Татьяна",
		"Ульяна",
		"Фаина",
		"Феврония",
		"Фёкла",
		"Элеонора",
		"Эмилия",
		"Юлия",
	}
)
