# Seri Go Logger Console

Провайдер для консольного вывода логов для библиотеки seri-go-logger.

## Установка

```bash
go get github.com/SergeiKhanlarov/seri-go-logger-console
```

## Импорт

```go
import (
    sglogger "github.com/SergeiKhanlarov/seri-go-logger"
    console "github.com/SergeiKhanlarov/seri-go-logger-console"
)
```

## Быстрый старт

```go
package main

import (
    "context"
    
    sglogger "github.com/SergeiKhanlarov/seri-go-logger"
    "github.com/SergeiKhanlarov/seri-go-logger-console"
)

func main() {
    ctx := context.Background()
    
    // Создание конфигурации провайдера
    config := sgloggerconsole.ProviderConfig{
        LoggerConfig: sglogger.LoggerConfig{},
        level:        sglogger.LevelInfo,
    }
    
    // Создание форматировщика
    formatter := sgloggerconsole.NewConsoleFormatter()
    
    // Создание провайдера
    provider := sgloggerconsole.NewConsoleProvider(config, formatter)
    
    // Создание логгера
    logger := sglogger.NewLogger(
		sglogger.LoggerConfig{}, 
		sglogger.NewFieldsHandler(),
		provider)
    
    // Использование логгера
    logger.Info(ctx, "Приложение запущено", sglogger.Fields{
        "version": "1.0.0",
        "port":    8080,
    })
}
```

## Конфигурация

### ProviderConfig

```go
type ProviderConfig struct {
    sglogger.LoggerConfig        // Базовая конфигурация
    level       sglogger.Level   // Уровень логирования
}
```

### Уровни логирования

LevelDebug - отладочные сообщения
LevelInfo - информационные сообщения
LevelWarn - предупреждения
LevelError - ошибки
LevelFatal - критические ошибки

### Форматирование

Пакет предоставляет ConsoleFormatter с цветовым выводом:

```text
2024-01-15 10:30:45 [INFO] main.go(25) - Application started version=1.0.0 port=8080
2024-01-15 10:30:46 [ERROR] handler.go(78) - Database connection failed db_host=localhost
```
### Цветовая схема:

🔵 Debug - Cyan<br>
🟢 Info - Green<br>
🟡 Warn - Yellow<br>
🔴 Error/Fatal - Red<br>

### Особенности

✅ Цветовой вывод - автоматическое окрашивание по уровням логирования<br>
✅ Структурированные логи - поддержка полей в формате key-value<br>
✅ Информация о месте вызова - автоматическое определение файла и строки<br>
✅ Фильтрация по уровням - эффективная фильтрация ненужных сообщений<br>
✅ Производительность - минимальные накладные расходы<br>
✅ Совместимость - полная интеграция с seri-go-logger<br>

## 📄 Лицензия

MIT License - смотрите файл [LICENSE](LICENSE) для деталей.

Copyright (c) 2025 Ханларов Сергей