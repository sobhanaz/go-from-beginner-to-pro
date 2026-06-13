<div dir="rtl">

> 🌐 **زبان / Language:** فارسی (همین فایل) · [English](session-08.md)

# جلسه‌ی ۰۸ — استراکت‌ها و متدها 🏗️

**هدف (۱ ساعت):** نوع‌های خودت را تعریف کنی. یک **استراکت** فیلدهای مرتبط را در یک
نوع سفارشی گروه می‌کند (مثل `User` یا `Task`)، و **متدها** رفتار را به آن نوع‌ها
وصل می‌کنند. این روشی است که دنیای واقعی را در Go مدل می‌کنی — و ستون فقرات پروژه‌ی
نهایی است.

> **مرور جلسه‌ی ۰۷:** مپ‌ها و رشته‌ها انواع داخلی را کامل کردند. حالا نوع‌های *خودت*
> را می‌سازی.

---

## ۱. استراکت‌ها — گروه‌بندی داده‌ی مرتبط (۱۵ دقیقه)

یک **استراکت** مجموعه‌ای نوع‌دار از **فیلدهای** نام‌دار است. با `type` تعریفش کن:

</div>

```go
type User struct {
    ID     int
    Name   string
    Email  string
    Active bool
}
```

<div dir="rtl">

### ساخت مقادیر استراکت

</div>

```go
// ۱. فیلدهای نام‌دار (بهترین — واضح و مستقل از ترتیب).
u1 := User{
    ID:     1,
    Name:   "Sobhan",
    Email:  "sobhan@example.com",
    Active: true,
}

// ۲. ترتیبی (باید با ترتیب فیلد بخواند — برای استراکت بزرگ نکن).
u2 := User{2, "Ali", "ali@example.com", false}

// ۳. مقدار صفر — هر فیلد مقدار صفر نوعش را می‌گیرد.
var u3 User   // {0 "" "" false}
```

<div dir="rtl">

### دسترسی و تغییر فیلدها

</div>

```go
fmt.Println(u1.Name)   // Sobhan
u1.Name = "Sobhan A."  // فیلدها تغییرپذیرند
u1.Active = false
```

<div dir="rtl">

> 💡 **نام‌گذاری فیلد = دیده‌شدن.** فیلدی که با حرف **بزرگ** شروع شود (`Name`)
> **export‌شده** است (از پکیج‌های دیگر دیده می‌شود). حرف کوچک (`name`)
> **export‌نشده** (خصوصیِ پکیج) است. این قاعده‌ی بزرگی/کوچکی برای *همه‌چیز* در Go
> صدق می‌کند — نوع‌ها، توابع، فیلدها. یادت باشد!

فایل [`examples/session08/structs/structs.go`](../examples/session08/structs/structs.go) را اجرا کن.

---

## ۲. متدها — رفتار وصل‌شده به یک نوع (۲۰ دقیقه)

یک **متد** تابعی با یک **گیرنده‌ی** ویژه است که آن را به یک نوع گره می‌زند. گیرنده
در پرانتز *قبل* از نام متد می‌آید:

</div>

```go
type Rectangle struct {
    Width, Height float64
}

// (r Rectangle) گیرنده است. به‌صورت rect.Area() صدا بزن.
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

rect := Rectangle{Width: 3, Height: 4}
fmt.Println(rect.Area())   // 12
```

<div dir="rtl">

### گیرنده‌ی مقداری در مقابل گیرنده‌ی اشاره‌گری — تمایز حیاتی

</div>

```go
// گیرنده‌ی مقداری: یک کپی می‌گیرد. تغییرات روی اصلی اثر ندارد.
func (r Rectangle) Scale(factor float64) {
    r.Width *= factor   // فقط کپی را تغییر می‌دهد — اینجا بی‌فایده!
}

// گیرنده‌ی اشاره‌گری: خودِ استراکت را می‌گیرد. تغییرات می‌مانند.
func (r *Rectangle) ScaleInPlace(factor float64) {
    r.Width *= factor
    r.Height *= factor
}
```

<div dir="rtl">

> 🔑 **قاعده‌ی سرانگشتی:**
> - وقتی متد باید استراکت را **تغییر** دهد یا استراکت بزرگ است، **گیرنده‌ی
>   اشاره‌گری** `(r *T)`.
> - برای استراکت‌های کوچک که فقط می‌خوانی، **گیرنده‌ی مقداری** `(r T)`.
> - **یکدست باش:** اگر یک متد گیرنده‌ی اشاره‌گری لازم دارد، برای *همه‌ی* متدهای آن
>   نوع از گیرنده‌ی اشاره‌گری استفاده کن.

### متدها روی هر نوع نام‌دار، نه فقط استراکت

</div>

```go
type Celsius float64

func (c Celsius) ToFahrenheit() float64 {
    return float64(c)*9/5 + 32
}

temp := Celsius(100)
fmt.Println(temp.ToFahrenheit())   // 212
```

<div dir="rtl">

فایل [`examples/session08/methods/methods.go`](../examples/session08/methods/methods.go) را اجرا کن.

---

## ۳. سازنده‌ها — قرارداد `New...` (۱۰ دقیقه)

Go نه `class` دارد نه سازنده‌ی داخلی. اصطلاح، یک تابع ساده به نام `NewX` است که یک
مقدار آماده‌ی استفاده برمی‌گرداند:

</div>

```go
func NewUser(name, email string) User {
    return User{
        Name:   name,
        Email:  email,
        Active: true,   // پیش‌فرض منطقی
    }
}

u := NewUser("Sara", "sara@example.com")
```

<div dir="rtl">

این فقط یک قرارداد است، اما همه‌جای Go واقعی هست. از آن برای اعمال پیش‌فرض و اعتبارسنجی هنگام ساخت نوع‌هایت استفاده کن.

---

## ۴. embedding استراکت — ترکیب به‌جای وراثت (۱۵ دقیقه)

Go **وراثت ندارد**. به‌جایش **embedding** دارد: یک استراکت را *بدون نام فیلد* درون
دیگری بگذار، و فیلدها/متدهایش به استراکت بیرونی «ترفیع» می‌یابند.

</div>

```go
type Animal struct {
    Name string
}

func (a Animal) Speak() string {
    return a.Name + " makes a sound"
}

type Dog struct {
    Animal       // embedded — بدون نام فیلد
    Breed string
}

d := Dog{
    Animal: Animal{Name: "Rex"},
    Breed:  "Labrador",
}

fmt.Println(d.Name)      // "Rex"  — از Animal ترفیع یافت!
fmt.Println(d.Speak())   // "Rex makes a sound" — متد هم ترفیع یافت
fmt.Println(d.Breed)     // "Labrador"
```

<div dir="rtl">

این فلسفه‌ی «ترکیب به‌جای وراثت» Go است: نوع‌های بزرگ‌تر را با *ترکیب* نوع‌های
کوچک‌تر می‌سازی، نه با گسترش یک کلاس پایه.

فایل [`examples/session08/embedding/embedding.go`](../examples/session08/embedding/embedding.go) را اجرا کن.

---

## 🎯 تمرین‌ها (قبل از جلسه‌ی ۰۹ انجام بده!)

فایل `examples/session08/practice/practice.go` را بساز:

۱. **تعریف `Task`:** استراکت `Task` با فیلدهای `ID int`، `Title string`، `Done bool`، `Priority int` بساز. یکی با فیلدهای نام‌دار بساز و با `%+v` چاپ کن.
۲. **متد `Complete`:** متد گیرنده-اشاره‌گری `func (t *Task) Complete()` اضافه کن که `Done = true` کند. یک وظیفه بساز، صدایش بزن و چاپ کن تا ماندگاری تغییر را ببینی.
۳. **متد `String`:** `func (t Task) String() string` اضافه کن که خلاصه‌ی تک‌خطی خوبی برگرداند. وظیفه را چاپ کن و ببین Go خودکار از `String()` استفاده می‌کند.
۴. **سازنده:** `func NewTask(title string, priority int) Task` بنویس که وظیفه‌ای با `Done: false` و ID خودافزای (از یک متغیر پکیج به‌عنوان شمارنده) برگرداند.
۵. **embedding:** یک `BaseEntity` با `ID int` و `CreatedAt string` بساز، بعد آن را در `User` و `Product` embed کن. نشان بده هر دو فیلد `ID` و `CreatedAt` را مفت می‌گیرند.

---

## ✅ چک‌لیست جلسه‌ی ۰۸

- [ ] می‌توانم استراکت تعریف و مقادیر را ۳ جور بسازم (نام‌دار، ترتیبی، صفر)
- [ ] می‌توانم فیلدها را بخوانم و تغییر دهم
- [ ] می‌دانم بزرگ = export‌شده، کوچک = export‌نشده
- [ ] می‌توانم متد با گیرنده‌ی مقداری و اشاره‌گری بنویسم
- [ ] می‌دانم وقتی باید استراکت را تغییر دهم از گیرنده‌ی اشاره‌گری استفاده کنم
- [ ] از قرارداد سازنده‌ی `NewX` پیروی می‌کنم
- [ ] embedding استراکت (ترکیب به‌جای وراثت) را می‌فهمم
- [ ] هر ۵ تمرین را انجام دادم

**قبلی:** [→ جلسه‌ی ۰۷](session-07.fa.md) · **بعدی:** [جلسه‌ی ۰۹ — اشاره‌گرها ←](session-09.fa.md)

</div>
