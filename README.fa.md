<div dir="rtl">

> 🌐 **زبان / Language:** فارسی (همین فایل) · [English](README.md)

# 🐹 زبان Go — از مبتدی تا حرفه‌ای

</div>

![Go](https://img.shields.io/badge/Go-1.25-00ADD8?logo=go&logoColor=white)
![Tests](https://img.shields.io/badge/tests-passing-brightgreen)
![Race](https://img.shields.io/badge/race--detector-clean-brightgreen)
![License](https://img.shields.io/badge/license-MIT-blue)

<div dir="rtl">

یک **دوره‌ی ۲۰ جلسه‌ای** کامل و خودآموز Go (انگلیسی + فارسی) **به‌علاوه‌ی دو
پروژه‌ی نمونه‌کار با سبک محصول واقعی** که می‌توانی اجرا کنی، در مصاحبه نشان دهی و
در رزومه‌ات بگذاری. همه‌چیز از صفر نوشته شده، کاملاً تست‌شده، و با Go 1.25.

---

## 📑 فهرست مطالب

- [داخلش چیست](#-داخلش-چیست)
- [ساختار مخزن](#-ساختار-مخزن)
- [شروع سریع](#-شروع-سریع)
- [دوره‌ی ۲۰ جلسه‌ای](#-دورهی-۲۰-جلسهای)
- [پروژه‌ی نمونه‌کار ۱ — TaskFlow (REST API)](#-پروژهی-نمونهکار-۱--taskflow-rest-api)
- [پروژه‌ی نمونه‌کار ۲ — gosearch (CLI)](#-پروژهی-نمونهکار-۲--gosearch-cli)
- [مفاهیم پوشش‌داده‌شده](#-مفاهیم-پوششدادهشده)
- [ردیاب پیشرفت](#-ردیاب-پیشرفت)

---

## 📦 داخلش چیست

| پوشه | چیست |
|------|------|
| [`sessions/`](sessions/) | ۲۰ فایل درس (هرکدام ~۱ ساعت)، به **انگلیسی و فارسی** |
| [`examples/`](examples/) | کد نمونه‌ی قابل‌اجرا برای جلسات ۱ تا ۱۷ — یک پوشه برای هر مفهوم |
| [`taskflow/`](taskflow/) | **پروژه‌ی نمونه‌کار ۱:** یک REST API با احراز هویت JWT (سرور) |
| [`gosearch/`](gosearch/) | **پروژه‌ی نمونه‌کار ۲:** یک ابزار جست‌وجوی متن همروند در خط فرمان |

سه ماژول Go اینجا هست: `golearn` (نمونه‌ها)، `taskflow` و `gosearch`.

---

## 🗂️ ساختار مخزن

</div>

```
.
├── README.md / README.fa.md      این فایل (انگلیسی / فارسی)
├── sessions/                     session-01..20 .md  (+ .fa.md فارسی)
├── examples/                     کد قابل‌اجرای جلسات ۱ تا ۱۷
│   ├── session01 … session17
│   └── (هر مفهوم در پوشه‌ی قابل‌اجرای خودش)
│
├── taskflow/                     ── پروژه‌ی نمونه‌کار REST API ──
│   ├── main.go                   نقطه‌ی ورود، خاموشی graceful
│   ├── internal/
│   │   ├── api/                  سرور، مسیرها، میان‌افزار، هندلرها، تست‌ها
│   │   ├── auth/                 JWT + هش رمز bcrypt
│   │   ├── config/              تنظیمات مبتنی بر محیط
│   │   ├── models/              انواع دامنه‌ی Task، User
│   │   └── store/               مخزن‌های SQLite (محدود به کاربر)
│   ├── Dockerfile, Makefile, demo.sh
│   ├── GETTING_STARTED.md        راهنمای گام‌به‌گام اجرا (با اسکرین‌شات)
│   └── RESUME.md                 بولت‌های رزومه + نکات مصاحبه
│
└── gosearch/                     ── پروژه‌ی نمونه‌کار CLI ──
    ├── main.go                   تجزیه‌ی فلگ، خروجی رنگی
    ├── internal/search/          اینترفیس Matcher + استخر کارگر همروند
    ├── Makefile, demo.sh
    └── README.md
```

<div dir="rtl">

---

## 🚀 شروع سریع

**یاد بگیر:** [`sessions/session-01.fa.md`](sessions/session-01.fa.md) را باز کن و نمونه‌ها را اجرا کن:

</div>

```bash
go run examples/session01/hello/hello.go
```

<div dir="rtl">

**REST API را اجرا کن:**

</div>

```bash
cd taskflow && go run .
# سپس: curl localhost:8080/health
```

<div dir="rtl">

**ابزار CLI را اجرا کن:**

</div>

```bash
cd gosearch && go run . -i -ext .go "func" .
```

<div dir="rtl">

دستورهایی که مدام استفاده می‌کنی:

</div>

```bash
go run .            # کامپایل + اجرا
go build ./...      # کامپایل همه‌چیز
go test ./...       # اجرای همه‌ی تست‌ها
go test -race ./... # بررسی کد همروند برای data race
go fmt ./...        # قالب‌بندی خودکار (همیشه)
```

<div dir="rtl">

---

## 📚 دوره‌ی ۲۰ جلسه‌ای

هر جلسه یک درس کامل با توضیح، نمونه‌ی قابل‌اجرا و تمرین است. به **انگلیسی**
(`session-NN.md`) و **فارسی** (`session-NN.fa.md`) موجود است.

### بخش ۱ — مبانی
| # | جلسه | چه چیزی یاد می‌گیری |
|---|------|---------------------|
| 01 | [سلام Go](sessions/session-01.fa.md) | اولین برنامه، `go run`/`go build`، پکیج‌ها، `gofmt` |
| 02 | [متغیرها و انواع](sessions/session-02.fa.md) | `var`، `:=`، ثابت‌ها، مقدار صفر، تبدیل نوع |
| 03 | [کنترل جریان](sessions/session-03.fa.md) | `if`، `switch`، `iota`، عملگرها |
| 04 | [حلقه‌ها و توابع](sessions/session-04.fa.md) | حلقه‌ی `for`، تعریف/فراخوانی توابع |
| 05 | [توابع پیشرفته](sessions/session-05.fa.md) | چند خروجی، variadic، closure، `defer` |

### بخش ۲ — ساختمان داده
| # | جلسه | چه چیزی یاد می‌گیری |
|---|------|---------------------|
| 06 | [آرایه‌ها و اسلایس‌ها](sessions/session-06.fa.md) | `append`، `make`، برش، `copy`، تله‌ی حافظه‌ی مشترک |
| 07 | [مپ‌ها، رشته‌ها و runeها](sessions/session-07.fa.md) | مپ، UTF-8، `rune` در برابر `byte` |
| 08 | [استراکت‌ها و متدها](sessions/session-08.fa.md) | نوع سفارشی، گیرنده‌ی مقداری/اشاره‌گری، embedding |
| 09 | [اشاره‌گرها](sessions/session-09.fa.md) | `&` و `*`، چه زمان و چرا |

### بخش ۳ — سبک Go
| # | جلسه | چه چیزی یاد می‌گیری |
|---|------|---------------------|
| 10 | [اینترفیس‌ها](sessions/session-10.fa.md) | اینترفیس ضمنی، چندریختی، `any` |
| 11 | [خطاها](sessions/session-11.fa.md) | خطای اصولی، `errors.Is/As`، `panic`/`recover` |
| 12 | [همروندی ۱](sessions/session-12.fa.md) | گوروتین، کانال، کلیدواژه‌ی `go` |
| 13 | [همروندی ۲](sessions/session-13.fa.md) | `select`، `WaitGroup`، `Mutex`، `context`، استخر کارگر |

### بخش ۴ — Go در دنیای واقعی
| # | جلسه | چه چیزی یاد می‌گیری |
|---|------|---------------------|
| 14 | [کتابخانه‌ی استاندارد](sessions/session-14.fa.md) | `strconv`، `time`، `sort`، `os` |
| 15 | [فایل‌ها، JSON و کدگذاری](sessions/session-15.fa.md) | ورودی/خروجی فایل، `encoding/json`، تگ استراکت |
| 16 | [تست‌نویسی](sessions/session-16.fa.md) | تست جدولی، بنچمارک، پوشش |
| 17 | [سرورهای HTTP](sessions/session-17.fa.md) | `net/http`، هندلر، مسیریابی، API به‌صورت JSON |

### بخش ۵ — پروژه‌ی نمونه‌کار
| # | جلسه | چه چیزی یاد می‌گیری |
|---|------|---------------------|
| 18 | [REST API + پایگاه‌داده](sessions/session-18.fa.md) | چیدمان پروژه، SQLite، endpointهای CRUD |
| 19 | [احراز هویت، میان‌افزار و تنظیمات](sessions/session-19.fa.md) | احراز هویت JWT، میان‌افزار، تنظیمات محیطی |
| 20 | [جلا و انتشار](sessions/session-20.fa.md) | Docker، خاموشی graceful، بولت‌های رزومه |

---

## 🏆 پروژه‌ی نمونه‌کار ۱ — TaskFlow (REST API)

یک **REST API مدیریت وظایف** چندکاربره با سبک محصول واقعی: هر کاربر ثبت‌نام می‌کند،
وارد می‌شود و فهرست وظایف خصوصی و اولویت‌دار خودش را مدیریت می‌کند. روی کتابخانه‌ی
استاندارد Go ساخته شده — بدون فریم‌ورک وب.

</div>

![TaskFlow demo](taskflow/docs/demo.png)

<div dir="rtl">

**ویژگی‌ها**
- 🔐 **احراز هویت JWT** با رمزهای هش‌شده با bcrypt
- 👥 **ایزولاسیون داده به‌ازای کاربر** که در لایه‌ی SQL اعمال می‌شود (`WHERE user_id = ?`)
- 🧩 **میان‌افزار**: لاگ ساختاریافته، بازیابی از panic، احراز هویت
- 🏷️ **اولویت وظیفه** (low/medium/high) + فیلترهای ترکیب‌پذیر (`?done=`، `?priority=`)
- 🗄️ **SQLite** با درایور خالص‌Go (بدون cgo) — لایه‌های تمیز هندلر ← مخزن
- 🧪 **تست‌های یکپارچگی** با `httptest`؛ ✅ 🐳 **Docker‌ای‌شده** به ایمیج distroless حدود ۲۱ مگابایت
- 🛑 **خاموشی graceful** روی SIGINT/SIGTERM

**اجرایش کن**

</div>

```bash
cd taskflow
go run .            # http://localhost:8080
make demo           # هر endpoint را اجرا و وضعیت + زمان را نشان می‌دهد
go test ./...       # همه‌ی تست‌ها
```

<div dir="rtl">

**API**

| متد | مسیر | احراز هویت | هدف |
|-----|------|------------|-----|
| GET | `/` · `/health` | – | فهرست API / liveness |
| POST | `/auth/register` · `/auth/login` | – | گرفتن JWT |
| GET | `/tasks` (`?done=`، `?priority=`) | ✅ | فهرست/فیلتر وظایف |
| POST | `/tasks` | ✅ | ساخت وظیفه |
| GET·PUT·DELETE | `/tasks/{id}` | ✅ | خواندن / به‌روزرسانی / حذف |

📖 راهنمای کامل: [taskflow/GETTING_STARTED.md](taskflow/GETTING_STARTED.md) ·
🧑‍💼 بولت‌های رزومه: [taskflow/RESUME.md](taskflow/RESUME.md)

---

## 🔎 پروژه‌ی نمونه‌کار ۲ — gosearch (CLI)

یک جست‌وجوی متن سریع و **همروند** در خط فرمان — یک mini-`grep`/`ripgrep`. به‌صورت
بازگشتی یک پوشه را با **استخر کارگر** اسکن می‌کند و متن ساده یا عبارت منظم را تطبیق
می‌دهد. فقط کتابخانه‌ی استاندارد، بدون وابستگی.

</div>

![gosearch demo](gosearch/docs/demo.png)

<div dir="rtl">

**ویژگی‌ها**
- 🔁 پیمایش پوشه + اسکن فایل به‌صورت همروند (استخر کارگر گوروتین، یکی به‌ازای هر CPU)
- 🔤 تطبیق ساده (مسیر سریع) یا **regex** کامل (`-r`)، پشت یک اینترفیس `Matcher`
- 🔡 بدون‌حساسیت به حروف (`-i`)، فیلتر پسوند (`-ext`)، کنترل فایل مخفی، حالت شمارش (`-c`)
- 🎨 خروجی رنگی با تطبیق‌های هایلایت‌شده (هنگام pipe خودکار خاموش)
- 🧪 تست جدولی + پوشه‌ی موقت، **بدون data race، پوشش ~۹۳٪**

**اجرایش کن**

</div>

```bash
cd gosearch
go run . func .                     # «func» را زیر پوشه‌ی فعلی پیدا کن
go run . -i -ext .go "todo" .       # بدون‌حساسیت به حروف، فقط فایل‌های .go
go run . -r "^func \w+\(" .         # regex
go test -race -cover ./...          # تست‌ها
```

<div dir="rtl">

---

## 🧠 مفاهیم پوشش‌داده‌شده

در طول دوره و پروژه‌ها، این مخزن این‌ها را نشان می‌دهد:

**زبان** — متغیرها، انواع، کنترل جریان، توابع، closure، `defer` · اسلایس، مپ،
رشته/rune · استراکت، متد، اشاره‌گر، embedding · اینترفیس و چندریختی · مدیریت خطای اصولی.

**همروندی** — گوروتین، کانال، `select`، `WaitGroup`، `Mutex`، `context`، استخر
کارگر، و آشکارساز ریس.

**دنیای واقعی** — کتابخانه‌ی استاندارد، ورودی/خروجی فایل، JSON، تست (واحد، جدولی،
یکپارچگی، بنچمارک، پوشش) و سرورهای HTTP.

**محصول** — احراز هویت JWT، هش رمز، میان‌افزار، معماری لایه‌ای، تنظیمات، SQLite،
Docker (چندمرحله‌ای، distroless) و خاموشی graceful.

---

## ✅ ردیاب پیشرفت

- [x] **جلسات ۰۱ تا ۲۰** — همه‌ی درس‌ها کامل (انگلیسی + فارسی)
- [x] **TaskFlow** — REST API ساخته، تست و Docker‌ای‌شده
- [x] **gosearch** — CLI همروند ساخته، تست‌شده، بدون data race

---

*با Go 1.25 ساخته شده است. از [جلسه‌ی ۰۱](sessions/session-01.fa.md) شروع کن، یا
مستقیم به [TaskFlow](taskflow/) و [gosearch](gosearch/) برو. کدنویسی خوبی داشته باشی! 🚀*

</div>
