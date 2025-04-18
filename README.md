# rp - Ruslan Password CLI менеджер паролей
**Описание**:  
Консольное приложение для безопасного хранения паролей. 

---

## 📜 **Команды**  

### **1. `init` — Инициализация хранилища**  
Создает новое зашифрованное хранилище в папке `~/.rp/data.enc`.  

**Пример**:  
```bash
./rp init
```  
**Действие**:  
- Запрашивает мастер-пароль (вводится скрыто).  
- Создает файл `~/.rp/data.enc` с зашифрованным пустым хранилищем.  

---

### **2. `set` — Сохранение пароля**  
Добавляет или обновляет пароль по ключу.  

**Флаги**:  
| Флаг | Описание                | Обязательный |  
|------|-------------------------|--------------|  
| `-k` | Ключ (например, `email`) | Да           |  
| `-v` | Пароль для сохранения    | Да           |  

**Пример**:  
```bash
./rp set -k github -v "балашиха"
```  

---

### **3. `get` — Получение пароля**  
Возвращает пароль по ключу.  

**Флаги**:  
| Флаг | Описание                | Обязательный |  
|------|-------------------------|--------------|  
| `-k` | Ключ (например, `email`) | Да           |  

**Пример**:  
```bash
./rp get -k github
```  
**Вывод**:  
```plaintext
Пароль для github: балашиха
```

---

### **4. `list` — Список всех ключей**  
Показывает все сохраненные ключи (без паролей).  

**Пример**:  
```bash
./rp list
```  
**Вывод**:  
```plaintext
Доступные ключи:
- github
- email
- bank
```

---

### **5. `delete` — Удаление пароля**  
Удаляет пароль по ключу.  

**Флаги**:  
| Флаг | Описание                | Обязательный |  
|------|-------------------------|--------------|  
| `-k` | Ключ (например, `email`) | Да           |  

**Пример**:  
```bash
./rp delete -k github
```  

---

## 🔒 **Безопасность**  
- **Шифрование**: AES-256-GCM.  
- **Хранение**: Локальный файл `~/.rp/data.enc`.  
- **Доступ**: Только после ввода мастер-пароля.  

**Важно**:  
- Мастер-пароль нельзя восстановить!  
- Файл `data.enc` не редактировать вручную.  

---

## 💾 **Расположение файлов**  
| Файл                  | Описание                     |  
|-----------------------|------------------------------|  
| `~/.rp/data.enc` | Зашифрованное хранилище      |  

---

## ❓ **Примеры workflows**  

### **Сценарий 1: Добавление и получение пароля**  
```bash
./rp init
./rp set -k email -v "secure123"
./rp get -k email
```

### **Сценарий 2: Резервное копирование**  
```bash
cp ~/.rp/data.enc ~/backup_passman.enc
```

---

## 🛠 **Дополнительно**  
- **Сброс хранилища**: Удалите файл `~/.rp/data.enc` и запустите `init` заново.  
- **Изменение мастер-пароля**: Не поддерживается (нужно создать новое хранилище).  

--- 

> ⚠️ **Предупреждение**: Не используйте простые мастер-пароли! Храните их в надежном месте (например, в менеджере паролей).
