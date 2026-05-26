# 2. Авторизация
```JS
use lab12
db.auth("mikhaylo_ra", "g%e&j5&%#o%")
```

# 2.2. Проектирование базы данных неряционного типа
## Описание структуру NoSQL БД
```plaintext
db.users (
	"_id": ObjectId("..."),
	"username": <>,
	"uid": <>,
	"gid": <>,
	"host": {
        "ip": <>,
    },
    "email": <>
);

db.policies {
    "_id": ObjectId("..."),
    "policy_name": <>,
    "user_id": ObjectId("<>"),
    "status": <>,

    "services": [
		{
			"service_name": <>,
			"rule": {
				"rule_name": <>,
                "conditions": [
					{
						<>
					}
				],
                "effect": <>
            }
        }
	]
}


db.requests {
	"_id": ObjectId("..."),
	"user": {
		"username": <>,
		"host": {
			"ip": <>
		}
	},
	"context": {
		<>
	},
	"timestamp": <>
}

db.decisions {
	"_id": ObjectId("..."),
	"request": {
		"request_id": <>,
		"service"
	},
	"policy": {
		"policy_name": <>,
		"status": <>
	},
	"decision": {
		"result": <>,
		"rule_name": <>
	},
	"timestamp": ISODate(<>)
}
```

# 2.3. Ввод / редактирование/удаление данных в базе данных
## 2.3.1. Напишите запрос на добавление одной записи в коллекцию
```JS
db.users.insertOne(
    {
    username: "ivan.petrov",
    uid: 1001,
    gid: 1001,
    host: 
    {
        ip: "192.168.1.10"
    }
})
```

## 2.3.2. Напишите запрос для добавления нескольких записей с применением одной команды
```JS
db.users.insertMany([
     {
        _id: ObjectId('6a1456fd1c51cfb44fe165cf'),
        username: 'maria.sidorova',
        uid: 1002,
        gid: 1002,
        host: { ip: '192.168.1.11' },
        email: 'maria.sidorova@novsu.ru'
    },
    {
        _id: ObjectId('6a1456fd1c51cfb44fe165d0'),
        username: 'alexey.kozlov',
        uid: 1003,
        gid: 1003,
        host: { ip: '192.168.1.12' },
        email: 'alexey.kozlov@novsu.ru'
    },
    {
        _id: ObjectId('6a1544c85bfad964c4e165ce'),
        username: 'roman.mikhaylov',
        uid: 1000,
        gid: 1000,
        host: { ip: '127.0.0.1' },
        email: 'roman.mikhaylov@novsu.ru'
    }
])
```

## 2.3.3 Добавьте в коллекции несколько документов
```JS
db.policies.insertMany([
{
    _id: ObjectId('6a1457161c51cfb44fe165d1'),
    policy_name: 'Developer Policy',
    user_id: ObjectId('6a1456fd1c51cfb44fe165cf'),
    status: false,
    services: [
      {
        service_name: 'docker',
        rules: [
          {
            rule_name: 'Разрешить администраторам',
            conditions: [ { operator: 'contains', type: 'groups', value: 'sudo' } ],
            effect: true
          }
        ]
      }
    ]
  },
  {
    _id: ObjectId('6a14594e1c51cfb44fe165d7'),
    policy_name: 'Default SSH Policy',
    user_id: ObjectId('6a1456fd1c51cfb44fe165d0'),
    status: true,
    services: [
      {
        service_name: 'ssh',
        rules: [
          {
            rule_name: 'Разрешить утром',
            conditions: [ { operator: 'between', type: 'timestamp', value: '08:00-18:00' } ],
            effect: true
          }
        ]
      }
    ]
  },
  {
    _id: ObjectId('6a14594e1c51cfb44fe165d8'),
    policy_name: 'VPN Only Policy',
    user_id: ObjectId('6a1544c85bfad964c4e165ce'),
    status: true,
    services: [
      {
        service_name: 'ssh',
        rules: [
          {
            rule_name: 'Разрешить через VPN',
            conditions: [ { operator: 'regex', type: 'hostname', value: 'vpn-.*' } ],
            effect: true
          }
        ]
      }
    ]
  }
])
```

```JS
db.requests.insertMany([
{
    context: 
    {
        host: 
        {
            ip: "10.0.0.30",
            hostname: "load-balancer"
        },
        time: 
        {
            weekday: "wednesday",
            timestamp: ISODate("2026-05-15T20:30:00Z")
        },
        user: 
        {
            uid: 1030,
            gid: 1030,
            username: "vera.belova",
            groups: ["network"]
        },

        service: "nginx"
    },

    timestamp: ISODate("2026-05-15T20:30:00Z")
},

{
    context: 
    {
        host: 
        {
            ip: "172.16.0.35",
            hostname: "vpn-europe"
        },

        time: 
        {
            weekday: "friday",
            timestamp: ISODate("2026-05-17T12:30:00Z")
        },

        user: 
        {
            uid: 1032,
            gid: 1032,
            username: "nadezhda.komarova",
            groups: ["remote"]
        },

        service: "ssh"
    },

    timestamp: ISODate("2026-05-17T12:30:00Z")
},

{
    context: 
    {
        host: 
        {
            ip: "192.168.2.28",
            hostname: "devops-2"
        },

        time: 
        {
            weekday: "tuesday",
            timestamp: ISODate("2026-05-28T14:00:00Z")
        },

        user: 
        {
            uid: 1043,
            gid: 1043,
            username: "gennady.rodionov",
            groups: [
                "devops",
                "sudo"
            ]
        },

        service: "kubernetes"
    },

    timestamp: ISODate("2026-05-28T14:00:00Z")
},

{
    context: 
    {
        host: 
        {
            ip: "192.168.1.25",
            hostname: "workstation-12"
        },

        time: {
            weekday: "sunday",
            timestamp: ISODate("2026-05-19T15:30:00Z")
        },

        user: 
        {
            uid: 1034,
            gid: 1034,
            username: "ludmila.borisova",
            groups: [
                "developers",
                "docker"
            ]
        },

        service: "docker"
    },

    timestamp: ISODate("2026-05-19T15:30:00Z")
},

{
    context: 
    {
        host: 
        {
            ip: "172.16.0.36",
            hostname: "guest-network"
        },

        time: {
            weekday: "wednesday",
            timestamp: ISODate("2026-05-22T18:00:00Z")
        },

        user: 
        {
            uid: 1037,
            gid: 1037,
            username: "roman.belousov",
            groups: ["guests"]
        },

        service: "ftp"
    },

    timestamp: ISODate("2026-05-22T18:00:00Z")
}
])
```

```JS
db.decisions.insertOne(
    {
    request: 
    {
        request_id: 1,
        service: "ssh"
    },

    policy: 
    {
        policy_name: "Default SSH Policy",
        status: true
    },

    decision: 
    {
        result: true,
        rule_name: "Разрешить утром"
    },

    timestamp: new Date()
})
```

```JS
db.decisions.insertMany([
{
    request: 
    {
        request_id: 1,
        service: "ssh"
    },

    policy: 
    {
        policy_name: "Default SSH Policy",
        status: true
    },

    decision: 
    {
        result: true,
        rule_name: "Разрешить утром"
    },

    timestamp: new Date("2026-04-15T09:30:01Z")
},
{
    request: 
    {
        request_id: 2,
        service: "postgresql"
    },

    policy: 
    {
        policy_name: "VPN Only Policy",
        status: true
    },

    decision: 
    {
        result: true,
        rule_name: "Разрешить через VPN"
    },

    timestamp: new Date("2026-04-16T14:20:01Z")
},
{
    request: 
    {
        request_id: 3,
        service: "nginx"
    },

    policy: 
    {
        policy_name: "Restricted Access",
        status: false
    },

    decision: 
    {
        result: false,
        rule_name: "Запретить из DMZ"
    },

    timestamp: new Date("2026-04-17T22:15:01Z")
}
])
```

## 2.3.4. Напишите запрос на изменение данных по своему усмотрению
```JS
db.policies.updateOne(
    { policy_name: "Developer Policy" },
    {
        $set: {
            status: false
        }
    }
)
```

```JS
db.users.updateOne(
    { username: "ivan.petrov" },
    {
        $set: {
            "host.ip": "192.168.1.100"
        }
    }
)
```

## 2.3.5. Напишите запрос на удаление данных по своему усмотрению
```JS
db.users.deleteOne({

    username: "ivan.petrov"
})
```

## 2.4.1
```JS
db.users.find()
```

```JS
db.users.find().pretty()
```

```JS
db.users.find(
    {},
    {
        username: 1,
        uid: 1,
        _id: 0
    }
)
```

```JS
db.users.find(
    {
    uid: 
    {
        $gt: 1050
    }
})
```

```JS
db.users.find(
    {
    uid: 
    {
        $lt: 1010
    }
})
```

```JS
db.users.find().limit(1)

```JS
db.users.find().skip(1)

```JS
db.users.find().sort({
    uid: 1
})
```

```JS
db.users.find().sort({
    uid: -1
})
```

```JS
db.policies.find(
    {
        "services.rules.effect": true
    },
    {
        _id: 0,
        policy_name: 1
    }
)
```

```JS
db.decisions.find(
    {
        "decision.result": true
    },
    {
        _id: 0,
        "request.request_id": 1,
        "request.service": 1
    }
)
```

```JS
db.policies.countDocuments()
```

## 2.4.2 Напишите ряд запросов с применением регулярных выражений
```JS
db.users.createIndex(
{
    username: 1,
    uid: 1
})
```

```JS
db.users.createIndex(
{
    username: "text"
})
```

```JS
db.policies.createIndex(
{
    policy_name: "text"
})
```

```JS
db.users.find(
{
    username: "ivan.petrov"
}).explain("executionStats")
```

```JS
db.users.find(
{
    username: /^maria/
})
```

```JS
db.users.find(
{
    username: /ov/
})
```

```JS
db.users.find(
{
    username: /^\w+\.\w+$/
})
```

```JS
db.requests.find(
{
    "context.service": /^p/
})
```

## 2.4.4 Запросы: модификаторы массивов
```JS
db.requests.updateOne({
context: 
{
      host: { ip: '192.168.2.28', hostname: 'devops-2' },
      time: { weekday: 'tuesday', timestamp: ISODate('2026-05-28T14:00:00.000Z') },
      user: 
      {
        uid: 1043,
        gid: 1043,
        username: 'gennady.rodionov',
        groups: 
        [ 'devops', 'sudo' ]
      },
      service: 'kubernetes'
    }}, 
    {
        $push: 
        {
            "context.user.groups": "testik"
        }
    })
```

```JS
db.policies.updateOne(
{
    _id: ObjectId("6a14594e1c51cfb44fe165d7")
},
{
    $push: 
    {
        "services.0.rules.0.conditions": 
        {
            operator: "equals",
            type: "weekday",
            value: "sunday"
        }
    }
})
```

```JS
db.policies.updateOne(
{
    _id: ObjectId("6a14594e1c51cfb44fe165d7")
},
{
    $set: 
    {
        "services.0.rules.0.conditions.0.value": "08:00-18:00"
    }
})
```

```JS
db.policies.updateOne(
{
    _id: ObjectId("6a14594e1c51cfb44fe165d7")
},
{
    $pull: 
    {
        "services.0.rules.0.conditions": 
        {
            type: "weekday",
            value: "sunday"
        }
    }
})
```

## 2.4.5 Агрегирование данных
```JS
db.city.insertMany([
{
    "city": "Москва",
    "region": "Москва",
    "pop": 13104177,
    "loc": [
        55.7558,
        37.6176
    ]
},
{
    "city": "Санкт-Петербург",
    "region": "Санкт-Петербург",
    "pop": 5601911,
    "loc": [
        59.9386,
        30.3141
    ]
},
{
    "city": "Екатеринбург",
    "region": "Свердловская область",
    "pop": 1544376,
    "loc": [
        56.8389,
        60.6057
    ]
},
{
    "city": "Новосибирск",
    "region": "Новосибирская область",
    "pop": 1633595,
    "loc": [
        55.0084,
        82.9357
    ]
},
{
    "city": "Казань",
    "region": "Республика Татарстан",
    "pop": 1318604,
    "loc": [
        55.7963,
        49.1088
    ]
},
{
    "city": "Самара",
    "region": "Самарская область",
    "pop": 1171734,
    "loc": [
        53.1959,
        50.1008
    ]
},
{
    "city": "Нижний Новгород",
    "region": "Нижегородская область",
    "pop": 1225140,
    "loc": [
        56.3269,
        44.0059
    ]
},
{
    "city": "Краснодар",
    "region": "Краснодарский край",
    "pop": 1121215,
    "loc": [
        45.0355,
        38.9753
    ]
},
{
    "city": "Ростов-на-Дону",
    "region": "Ростовская область",
    "pop": 1142162,
    "loc": [
        47.2357,
        39.7015
    ]
},
{
    "city": "Уфа",
    "region": "Республика Башкортостан",
    "pop": 1144809,
    "loc": [
        54.7388,
        55.9721
    ]
},
{
    "city": "Владивосток",
    "region": "Приморский край",
    "pop": 603519,
    "loc": [
        43.1155,
        131.8855
    ]
},
{
    "city": "Архангельск",
    "region": "Архангельская область",
    "pop": 298186,
    "loc": [
        64.5399,
        40.5158
    ]
},
{
    "city": "Мурманск",
    "region": "Мурманская область",
    "pop": 267422,
    "loc": [
        68.9707,
        33.0749
    ]
},
{
    "city": "Улан-Удэ",
    "region": "Республика Бурятия",
    "pop": 436406,
    "loc": [
        51.8345,
        107.5840
    ]
},
{
    "city": "Тверь",
    "region": "Тверская область",
    "pop": 416219,
    "loc": [
        56.8587,
        35.9176
    ]
},
{
    "city": "Смоленск",
    "region": "Смоленская область",
    "pop": 316570,
    "loc": [
        54.7826,
        32.0453
    ]
},
{
    "city": "Киров",
    "region": "Кировская область",
    "pop": 468212,
    "loc": [
        58.6036,
        49.6679
    ]
},
{
    "city": "Кинешма",
    "region": "Ивановская область",
    "pop": 77000,
    "loc": [
        57.4425,
        42.1686
    ]
},
{
    "city": "Кондопога",
    "region": "Республика Карелия",
    "pop": 28500,
    "loc": [
        62.2057,
        34.2614
    ]
},
{
    "city": "Сафоново",
    "region": "Смоленская область",
    "pop": 41000,
    "loc": [
        55.1144,
        33.2333
    ]
}
])
```

```JS
db.city.aggregate(
[
    {
        $group: 
        {
            _id: "$region",
            totalPop: { $sum: "$pop" }
        }
    },
    {
        $match: 
        {
            totalPop: { $gte: 10000 }
        }
    }
])
```

```sql
SELECT region,
       SUM(pop) AS totalPop
FROM city
GROUP BY region
HAVING SUM(pop) >= 1000;
```

```JS
db.city.aggregate([
    {
        $group: 
        {
            _id: "$region",
            avgPop: { $avg: "$pop" }
        }
    }
])
```

```JS
db.city.aggregate([
    {
        $sort: 
        {
            pop: 1
        }
    },
    {
        $group: 
        {
            _id: "$region",
            minCity: { $first: "$city" },
            minPop: { $first: "$pop" },
            maxCity: { $last: "$city" },
            maxPop: { $last: "$pop" }
        }
    }
])
```

```JS
db.city.aggregate([
    {
        $group: 
        {
            _id: "$region",
            topCities: 
            {
                $topN: 
                {
                    n: 2,
                    sortBy: { pop: -1 },
                    output: 
                    {
                        city: "$city",
                        pop: "$pop"
                    }
                }
            }
        }
    }
])
```

## 2.4.6 Работа с курсорами
```JS
function countBigCities() {

    let cursor = db.city.find();

    let counter = 0;

    while (cursor.hasNext()) {

        let city = cursor.next();

        print(
            "Город: " + city.city +
            ", Регион: " + city.region +
            ", Население: " + city.pop
        );

        if (city.pop > 1000000) {
            counter++;
        }
    }

    print("Количество городов с населением более 1 млн: " + counter);
}
```

```JS
countBigCities();
```

## 2.4.7. Работа с функциями
```JS
var cursor = db.users.find(
    {
        username: /^[a-p]/i
    }
).limit(10);

cursor
```

```JS
db.city.find()
       .sort({ pop: -1 })
       .skip(5)
       .limit(5)
```

```JS
db.city.find(
    { pop: { $gt: 1000000 } }
).map(
    function(city) {
        return {
            city: city.city,
            population: city.pop
        };
    }
)
```

# 2.6 Связывание документов
## 2.6.1 Ручное связывание
```JS
db.tasks.insertMany([
    {
        user_id: ObjectId("6a1456fd1c51cfb44fe165cf"),
        name: "Разработка API авторизации",
        description: "Реализовать REST API для аутентификации и выдачи JWT токенов.",
        start_date: ISODate("2025-05-01T09:00:00Z"),
        end_date: ISODate("2025-05-10T18:00:00Z")
    },
    {
        user_id: ObjectId("6a1456fd1c51cfb44fe165d0"),
        name: "Настройка MongoDB",
        description: "Создать коллекции, индексы и пользователей базы данных.",
        start_date: ISODate("2025-05-03T10:00:00Z"),
        end_date: ISODate("2025-05-05T17:00:00Z")
    },
    {
        user_id: ObjectId("6a1544c85bfad964c4e165ce"),
        name: "Тестирование политик доступа",
        description: "Проверить корректность применения ABAC-политик для пользователей.",
        start_date: ISODate("2025-05-06T09:30:00Z"),
        end_date: ISODate("2025-05-12T16:30:00Z")
    },
    {
        user_id: ObjectId("6a1544c85bfad964c4e165ce"),
        name: "Подготовка документации",
        description: "Описать структуру базы данных и основные сценарии использования.",
        start_date: ISODate("2025-05-08T11:00:00Z"),
        end_date: ISODate("2025-05-15T18:00:00Z")
    },
    {
        user_id: ObjectId("6a1456fd1c51cfb44fe165d0"),
        name: "Мониторинг системы",
        description: "Настроить сбор метрик и контроль доступности сервисов.",
        start_date: ISODate("2025-05-11T08:00:00Z"),
        end_date: ISODate("2025-05-20T17:00:00Z")
    }
])
```

```JS
db.users.aggregate([
    {
        $match: {
            email: "roman.mikhaylov@novsu.ru"
        }
    },
    {
        $lookup: {
            from: "tasks",
            localField: "_id",
            foreignField: "user_id",
            as: "tasks"
        }
    }
])
```

## 2.6.2 Автоматическое связывание

```JS
db.authors.insertOne({
    author: "Лев Толстой",
    country: "Россия",
    birth_year: 1828,
    awards: [
        "Почётный академик Императорской Академии наук"
    ],
    books: []
})
```

```JS
db.books.insertMany([
{
    title: "Война и мир",
    year: 1869
},
{
    title: "Анна Каренина",
    year: 1877
},
{
    title: "Воскресение",
    year: 1899
}
])
```

```JS
db.authors.updateOne(
{
    _id: ObjectId("6a1548b35bfad964c4e165d4")
},
{
    $set: {
        books: [
            DBRef("books", ObjectId("6a1548b85bfad964c4e165d5")),
            DBRef("books", ObjectId("6a1548b85bfad964c4e165d6")),
            DBRef("books", ObjectId("6a1548b85bfad964c4e165d7"))
        ]
    }
})
```

## 2.6.3 Работа со вложенными документами
```JS
db.authors.insertMany([
{
    name: "Иван Петров",
    email: "ivan.petrov@novsu.ru",
    articles: [
        {
            title: "Основы MongoDB",
            content: "Введение в документоориентированные БД.",
            created_at: ISODate("2024-03-15")
        },
        {
            title: "Индексы в MongoDB",
            content: "Использование индексов для ускорения запросов.",
            created_at: ISODate("2023-06-20")
        },
        {
            title: "Агрегация данных",
            content: "Примеры использования aggregate.",
            created_at: ISODate("2022-11-10")
        }
    ]
},
{
    name: "Мария Сидорова",
    email: "maria.sidorova@novsu.ru",
    articles: [
        {
            title: "Язык запросов SQL",
            content: "Базовые конструкции SQL.",
            created_at: ISODate("2024-01-12")
        },
        {
            title: "Нормализация данных",
            content: "Нормальные формы БД.",
            created_at: ISODate("2023-05-08")
        },
        {
            title: "Транзакции",
            content: "ACID свойства транзакций.",
            created_at: ISODate("2021-09-14")
        }
    ]
},
{
    name: "Алексей Козлов",
    email: "alexey.kozlov@novsu.ru",
    articles: [
        {
            title: "Docker для начинающих",
            content: "Контейнеризация приложений.",
            created_at: ISODate("2024-02-01")
        },
        {
            title: "Kubernetes",
            content: "Оркестрация контейнеров.",
            created_at: ISODate("2023-04-17")
        },
        {
            title: "CI/CD",
            content: "Автоматизация разработки.",
            created_at: ISODate("2022-07-22")
        }
    ]
},
{
    name: "Елена Новикова",
    email: "elena.novikova@novsu.ru",
    articles: [
        {
            title: "Информационная безопасность",
            content: "Основные угрозы безопасности.",
            created_at: ISODate("2024-05-11")
        },
        {
            title: "Аутентификация",
            content: "Методы проверки пользователей.",
            created_at: ISODate("2023-08-19")
        },
        {
            title: "Авторизация",
            content: "Управление доступом.",
            created_at: ISODate("2022-12-05")
        }
    ]
},
{
    name: "Дмитрий Волков",
    email: "dmitry.volkov@novsu.ru",
    articles: [
        {
            title: "Linux Administration",
            content: "Работа с Linux-серверами.",
            created_at: ISODate("2024-04-18")
        },
        {
            title: "Сетевые службы",
            content: "SSH, FTP, DNS.",
            created_at: ISODate("2023-03-09")
        },
        {
            title: "Мониторинг систем",
            content: "Prometheus и Grafana.",
            created_at: ISODate("2021-10-30")
        }
    ]
}
])
```

```JS
db.authors.find(
    {
        email: "ivan.petrov@novsu.ru"
    },
    {
        _id: 0,
        name: 1,
        articles: 1
    }
)
```

```JS
db.authors.find(
{
    articles: {
        $elemMatch: {
            created_at: {
                $gte: new Date(
                    new Date().setFullYear(
                        new Date().getFullYear() - 3
                    )
                )
            }
        }
    }
},
{
    _id: 0,
    name: 1,
    email: 1
})
```

```JS
db.authors.aggregate([
{
    $match: {
        "articles.created_at": {
            $gte: new Date(
                new Date().setFullYear(
                    new Date().getFullYear() - 3
                )
            )
        }
    }
},
{
    $group: {
        _id: "$name",
        email: { $first: "$email" }
    }
}
])
```

## 2.6.4 Спроектируйте связанные коллекции документов, реализуйте заданные Вами связи. Произведите поиск по связанным документам

```JS
var users = [
    ObjectId("6a14594e1c51cfb44fe165d4"),
    ObjectId("6a14594e1c51cfb44fe165d5")
]

db.policies.find({
    user_id: {
        $in: users
    }
})
```

```JS
db.policies.aggregate([
{
    $lookup: {
        from: "users",
        localField: "user_id",
        foreignField: "_id",
        as: "user"
    }
},
{
    $unwind: "$user"
},
{
    $project: {
        _id: 0,
        username: "$user.username",
        policy_name: 1,
        services_count: {
            $size: "$services"
        }
    }
}
])
```

```JS
db.requests.find({
    "context.user.username": {
        $in: [
            "nadezhda.komarova",
            "gennady.rodionov",
            "roman.belousov"
        ]
    }
})
```

# 3. Работа с графическим клиентом для MongoDB
## 3.1.2
```JS
{
  "firstname": "Роман",
  "lastname": "Михайлов",
  "patronymic": "Александрович",
  "age": 20,
  "course": 3,
  "group": "3093до"
}
```

```JS
{
  "firstname": "Анна",
  "lastname": "Соколова",
  "patronymic": "Игоревна",
  "age": 19,
  "course": 2,
  "group": "2093до"
}
```

```JS
{
  "firstname": "Дмитрий",
  "lastname": "Кузнецов",
  "patronymic": "Сергеевич",
  "age": 21,
  "course": 4,
  "group": "4093до"
}
```

```JS
{
  "firstname": "Екатерина",
  "lastname": "Орлова",
  "patronymic": "Владимировна",
  "age": 20,
  "course": 3,
  "group": "3093до"
}
```

```JS
{
  "firstname": "Екатерина",
  "lastname": "Орлова",
  "patronymic": "Владимировна",
  "age": 20,
  "course": 3,
  "group": "3093до"
}
```

```JS
{
  "firstname": "Максим",
  "lastname": "Павлов",
  "patronymic": "Андреевич",
  "age": 22,
  "course": 4,
  "group": "4093до"
}
```

## 3.1.3. Создайте вложенные данные в документе
```JS
address: {
    city: "Великий Новгород",
    street: "Большая Санкт-Петербургская",
    building: 41
}
```

```JS
address: {
    city: "Псков",
    street: "Советская",
    building: 12
}
```

```JS
address: {
    city: "Санкт-Петербург",
    street: "Невский проспект",
    building: 85
}
```

```JS
address: {
    city: "Москва",
    street: "Тверская",
    building: 23
}
```

```JS
address: {
    city: "Тверь",
    street: "Вагжанова",
    building: 7
}
```

## 3.1.5.
```JS
{ rule_name: {$eq: "Разрешить в пятницу днем" } }
```

## 3.1.6. Разберитесь с поиском по нескольким условиям $ И
```JS
{
  "$and": [
    { "course": 3 },
    { "group": "3093до" }
  ]
}
```

## 3.1.7. Разберитесь с поиском по нескольким условиям $ ИЛИ
```JS
{
  "$or": [
    { "course": 3 },
    { "group": "4093до" }
  ]
}
```

## 3.1.8. Совпадение по исключению $ not
```JS
{
  "age": {
    "$not": {
      "$eq": 20
    }
  }
}
```

## 3.1.9. Совпадение с операторами сравнения ($lt, $gt)
```JS
{
  "age": {
    "$gt": 20
  }
}
```

```JS
{
  "age": {
    "$lt": 20
  }
}
```


## 3.1.10. Добавьте массив, напишите запрос для поиска всех документов, в которых в указанном поле хотя бы одно из значений удовлетворяет условию
```JS
{
  articles: {
    $elemMatch: {
      created_at: {
        $gte: { "$date": "2023-01-01T00:00:00.000Z" },
        $lte: { "$date": "2024-01-01T23:59:59.999Z" }
      }
    }
  }
}
```
