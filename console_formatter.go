package sgloggerconsole

import (
	"fmt"
	"runtime"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

// ConsoleFormatter определяет интерфейс для форматирования записей лога.
// Реализует контракт logrus.Formatter для интеграции с logrus.
type ConsoleFormatter interface {
	Format(entry *logrus.Entry) ([]byte, error)
}

// consoleFormatter реализует форматирование логов для консольного вывода.
// Добавляет цветовое кодирование, временные метки и информацию о вызывающей стороне.
type consoleFormatter struct {
}

// NewConsoleFormatter создает и возвращает новый экземпляр consoleFormatter.
// Возвращает:
//   - ConsoleFormatter - форматировщик для использования в logrus
//
// Пример использования:
//
//	logger := logrus.New()
//	logger.SetFormatter(NewConsoleFormatter())
func NewConsoleFormatter() ConsoleFormatter {
	return &consoleFormatter{}
}

// Format преобразует запись лога в форматированную строку с цветовой подсветкой.
// Формат вывода: "TIMESTAMP [COLORED_LEVEL] FILE(LINE) - MESSAGE [fields]"
//
// Параметры:
//   - entry - запись лога для форматирования
//
// Возвращает:
//   - []byte - отформатированную строку в виде байтового среза
//   - error - ошибка форматирования (всегда nil в текущей реализации)
//
// Цветовая схема:
//   - Debug:    Cyan (36)
//   - Info:     Green (32)
//   - Warn:     Yellow (33)
//   - Error:    Red (31)
//   - Default:  White (37)
//
// Пример вывода:
//   "2024-01-15 10:30:45 [INFO] main.go(25) - Application started user_id=123"
func (f *consoleFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var levelColor string
	var resetColor = "\033[0m"

	switch entry.Level {
	case logrus.DebugLevel:
		levelColor = "\033[36m" // Cyan
	case logrus.InfoLevel:
		levelColor = "\033[32m" // Green
	case logrus.WarnLevel:
		levelColor = "\033[33m" // Yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = "\033[31m" // Red
	default:
		levelColor = "\033[37m" // White
	}

	timestamp := time.Now().Format("2006-01-02 15:04:05")
	level := strings.ToUpper(entry.Level.String())

	file, line := getCaller()

	var fields string
	if len(entry.Data) > 0 {
		var fieldParts []string
		for k, v := range entry.Data {
			fieldParts = append(fieldParts, fmt.Sprintf("%s=%v", k, v))
		}
		fields = " " + strings.Join(fieldParts, " ")
	}

	message := fmt.Sprintf("%s [%s%s%s] %s(%d) - %s%s\n",
		timestamp,
		levelColor, level, resetColor,
		file, line,
		entry.Message, fields)

	return []byte(message), nil
}

// getCaller находит информацию о вызывающей стороне в стеке вызовов.
// Пропускает фреймы, связанные с logrus и пакетами логирования.
//
// Возвращает:
//   - string - имя файла (только последняя часть пути)
//   - int    - номер строки в файле
//
// Алгоритм:
//   - Просматривает стек вызовов начиная с 3-го фрейма (пропускает logrus)
//   - Игнорирует фреймы, содержащие "logrus" или "logger"
//   - Возвращает первый подходящий фрейм
//   - При неудаче возвращает "???" и 0
func getCaller() (string, int) {
	for i := 3; i <= 10; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			continue
		}

		fn := runtime.FuncForPC(pc)
		if fn == nil {
			continue
		}

		funcName := fn.Name()

		if strings.Contains(funcName, "logrus") || strings.Contains(file, "logrus") {
			continue
		}

		if strings.Contains(funcName, "logger") || strings.Contains(file, "logger") {
			continue
		}

		parts := strings.Split(file, "/")
		if len(parts) > 0 {
			file = parts[len(parts)-1]
		}

		return file, line
	}

	return "???", 0
}