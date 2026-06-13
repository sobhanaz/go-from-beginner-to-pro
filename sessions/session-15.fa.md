<div dir="rtl">

> 🌐 **زبان / Language:** فارسی (همین فایل) · [English](session-15.md)

# جلسه‌ی ۱۵ — فایل‌ها، JSON و کدگذاری 📄

**هدف (۱ ساعت):** یاد بگیری داده را به برنامه‌ات بیاوری و از آن بیرون ببری.
**فایل‌ها** را می‌خوانی و می‌نویسی، و **JSON** را با `encoding/json` و **تگ‌های
استراکت** مسلط می‌شوی — مهم‌ترین مهارت سریال‌سازی برای کار بک‌اند. REST API تو با
JSON حرف می‌زند، پس این جلسه پل به پروژه‌ی نهایی است.

> **مرور جلسه‌ی ۱۴:** کتابخانه‌ی استاندارد را گشتی. JSON و ورودی/خروجی فایل دو
> پکیج دیگرند که تقریباً در هر برنامه‌ی واقعی استفاده می‌کنی.

---

## ۱. JSON — زبان APIهای وب (۲۵ دقیقه)

**JSON** (JavaScript Object Notation) قالب جهانی فرستادن داده بین سرویس‌هاست.
`encoding/json` در Go بین مقادیر Go و JSON تبدیل می‌کند.

دو عملیات، در جهت‌های مخالف:

- **Marshal**: مقدار Go ← JSON (برای فرستادن در یک پاسخ، JSON تولید می‌کنی).
- **Unmarshal**: JSON ← مقدار Go (JSON یک درخواست را تجزیه می‌کنی).

</div>

```go
type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

// Marshal: استراکت -> بایت‌های JSON
u := User{ID: 1, Name: "Sobhan", Email: "sobhan@example.com"}
data, err := json.Marshal(u)
// data = {"id":1,"name":"Sobhan","email":"sobhan@example.com"}

// Unmarshal: بایت‌های JSON -> استراکت (یک اشاره‌گر با & پاس بده)
var parsed User
err = json.Unmarshal([]byte(`{"id":2,"name":"Ali"}`), &parsed)
```

<div dir="rtl">

> 🔑 **دو چیز برای یادآوری:**
> ۱. `Unmarshal` به یک **اشاره‌گر** (`&parsed`) نیاز دارد تا بتواند متغیرت را پر کند.
> ۲. JSON فقط فیلدهای **export‌شده** (حرف‌بزرگ) را می‌بیند. فیلد حرف‌کوچک برای
>    `encoding/json` نامرئی است. (دوباره بزرگی/کوچکی مهم است — جلسه‌ی ۰۸!)

برای خروجی خوانا، از `json.MarshalIndent(v, "", "  ")` استفاده کن.

فایل [`examples/session15/jsonbasics/jsonbasics.go`](../examples/session15/jsonbasics/jsonbasics.go) را اجرا کن.

---

## ۲. تگ‌های استراکت — کنترل JSON (۲۰ دقیقه)

یک **تگ استراکت** رشته‌ی داخل بک‌تیک بعد از فیلد است. برای JSON نام کلید و رفتار را کنترل می‌کند:

</div>

```go
type Product struct {
    ID       int     `json:"id"`                 // کلید JSON «id» است نه «ID»
    Title    string  `json:"title"`
    Price    float64 `json:"price"`
    Discount float64 `json:"discount,omitempty"` // وقتی صفر است از JSON حذف شود
    Internal string  `json:"-"`                   // هرگز در JSON ظاهر نمی‌شود
}
```

<div dir="rtl">

| تگ | اثر |
|-----|-----|
| `json:"id"` | از `id` به‌عنوان کلید JSON استفاده کن (حرف‌کوچک، مناسب API) |
| `json:"discount,omitempty"` | وقتی مقدار صفر است فیلد را کلاً بیرون بگذار |
| `json:"-"` | هرگز این فیلد را سریال نکن (عالی برای رمز، داده‌ی داخلی) |

> 🔑 **`omitempty`** با حذف فیلدهای خالی/صفر JSON را تمیز نگه می‌دارد.
> **`json:"-"`** روشی است که رازها (مثل هش رمز) را از پاسخ‌های API بیرون نگه می‌داری
> — در مدل `User` پروژه‌ی نهایی استفاده‌اش می‌کنی.

هنگام unmarshal، کلیدهای ناشناخته‌ی JSON **نادیده** گرفته می‌شوند و کلیدهای غایب
**مقدار صفر** می‌مانند — پس کدت در برابر فیلدهای اضافی یا غایب مقاوم است.

فایل [`examples/session15/tags/tags.go`](../examples/session15/tags/tags.go) را اجرا کن.

---

## ۳. فایل‌ها — خواندن و نوشتن (۱۵ دقیقه)

برای بیشتر موارد، کمک‌کننده‌های کل-فایل همه‌ی چیزی است که نیاز داری:

</div>

```go
// نوشتن (می‌سازد یا بازنویسی می‌کند). 0644 = خواندن/نوشتن مالک، خواندن دیگران.
err := os.WriteFile("data.txt", []byte("hello\n"), 0644)

// خواندن کل فایل در حافظه.
content, err := os.ReadFile("data.txt")   // content یک []byte است
fmt.Println(string(content))
```

<div dir="rtl">

با JSON ترکیب کن تا داده‌ی ساختاریافته را ماندگار کنی (یک دیتابیس فقیرانه، و دقیقاً
همان‌طور که فایل‌های تنظیمات کار می‌کنند):

</div>

```go
cfg := Config{AppName: "TaskFlow", Port: 8080}
bytes, _ := json.MarshalIndent(cfg, "", "  ")
os.WriteFile("config.json", bytes, 0644)        // ذخیره

raw, _ := os.ReadFile("config.json")
var loaded Config
json.Unmarshal(raw, &loaded)                    // بارگذاری
```

<div dir="rtl">

> 💡 **مجوزهای فایل** مثل `0644` مجوزهای هشت‌هشتی یونیکس‌اند: `6`=خواندن+نوشتن برای
> مالک، `4`=خواندن برای گروه و دیگران. `0644` استاندارد یک فایل داده‌ی معمولی است؛
> `0755` (اجرا را اضافه می‌کند) استاندارد پوشه‌ها/فایل‌های اجرایی است.

> 📦 برای فایل‌های **بزرگ** به‌جای بارگذاری همه‌چیز در حافظه با `os.Open` +
> `bufio.Scanner` جریان‌سازی می‌کنی. اینجا ساده نگهش می‌داریم؛
> `ReadFile`/`WriteFile` اکثریت قریب‌به‌اتفاق نیازها را پوشش می‌دهند.

فایل [`examples/session15/files/files.go`](../examples/session15/files/files.go) را اجرا کن.

---

## 🎯 تمرین‌ها (قبل از جلسه‌ی ۱۶ انجام بده!)

فایل `examples/session15/practice/practice.go` را بساز:

۱. **رفت‌وبرگشت:** استراکت `Book` (`Title`، `Author`، `Year`، `Pages`) با تگ JSON تعریف کن. یکی را به JSON خوانا Marshal و چاپ کن، بعد یک رشته‌ی JSON را به `Book` بازگردانی (Unmarshal) و استراکت را چاپ کن.
۲. **پنهان کردن راز:** فیلد `Password string` با تگ `json:"-"` به یک استراکت `User` اضافه کن. Marshal کن و تأیید کن رمز هرگز در خروجی ظاهر نمی‌شود.
۳. **omitempty:** استراکتی با `Nickname string` اختیاری و `omitempty` بساز. یک‌بار خالی و یک‌بار پرشده Marshal کن؛ JSON را مقایسه کن.
۴. **ذخیره و بارگذاری فهرست:** یک `[]Task` بساز، Marshal کن، در فایل موقت بنویس، بازخوان، Unmarshal و اسلایس بازیابی‌شده را چاپ کن. فایل را پاک کن.
۵. **بارگذار تنظیمات:** `func loadConfig(path string) (Config, error)` بنویس که یک فایل JSON بخواند و یک `Config` پرشده (یا خطا) برگرداند. تست کن.

---

## ✅ چک‌لیست جلسه‌ی ۱۵

- [ ] می‌توانم مقدار Go را به JSON و JSON را به استراکت Unmarshal کنم
- [ ] یادم هست Unmarshal به اشاره‌گر (`&v`) نیاز دارد
- [ ] می‌دانم فقط فیلدهای export‌شده (حرف‌بزرگ) را `encoding/json` می‌بیند
- [ ] می‌توانم کلیدهای JSON را با تگ استراکت تغییر نام دهم
- [ ] می‌توانم از `omitempty` و `json:"-"` استفاده کنم
- [ ] می‌توانم کل فایل را با `os.ReadFile`/`os.WriteFile` بخوانم و بنویسم
- [ ] می‌توانم یک استراکت را به‌صورت JSON در فایل ماندگار و دوباره بارگذاری کنم
- [ ] هر ۵ تمرین را انجام دادم

**قبلی:** [→ جلسه‌ی ۱۴](session-14.fa.md) · **بعدی:** [جلسه‌ی ۱۶ — تست‌نویسی ←](session-16.fa.md)

</div>
