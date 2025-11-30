package sgloggerconsole

import (
	"context"
	"maps"
	"os"

	sglogger "github.com/SergeiKhanlarov/seri-go-logger"
	"github.com/sirupsen/logrus"
)

// consoleProvider реализует провайдер логирования для вывода в консоль.
// Использует logrus для форматирования и вывода логов с поддержкой структурированного логирования.
type consoleProvider struct {
	logrus *logrus.Logger  // Экземпляр logrus для логирования
	config ProviderConfig  // Конфигурация провайдера
}

// NewConsoleProvider создает и возвращает новый экземпляр провайдера консольного логирования.
//
// Параметры:
//   - config - конфигурация провайдера, включающая уровень логирования и базовые настройки
//   - formatter - форматировщик для преобразования записей лога в читаемый вид
//
// Возвращает:
//   - sglogger.LoggerProvider - интерфейс провайдера логирования
//
// Пример использования:
//
//	config := ProviderConfig{
//	    LoggerConfig: sglogger.LoggerConfig{...},
//	    level:        sglogger.LevelInfo,
//	}
//	formatter := NewConsoleFormatter()
//	provider := NewConsoleProvider(config, formatter)
//	logger := sglogger.NewLogger(provider)
func NewConsoleProvider(config ProviderConfig, formatter ConsoleFormatter) sglogger.LoggerProvider {
	log := logrus.New()
	
	log.SetOutput(os.Stdout)
	
	level, err := logrus.ParseLevel(config.GetLevel())
	if err != nil {
		level = logrus.InfoLevel 
	}
	log.SetLevel(level)
	
	log.SetFormatter(formatter)

	return &consoleProvider{
		logrus: log,
		config: config,
	}
}

// Write записывает сообщение лога в консоль с указанным уровнем и полями.
// Реализует интерфейс sglogger.LoggerProvider.
//
// Параметры:
//   - ctx - контекст выполнения (может использоваться для отмены или передачи метаданных)
//   - level - уровень важности сообщения
//   - message - текстовое сообщение лога
//   - fields - дополнительные структурированные данные в формате ключ-значение
//
// Возвращает:
//   - error - всегда возвращает nil, так как консольный вывод не порождает ошибок записи
//
// Логика работы:
//   - Проверяет, нужно ли логировать сообщение через ShouldLog
//   - Конвертирует поля в формат logrus
//   - Выбирает соответствующий метод логирования based on level
//   - Для LevelFatal вызывает os.Exit(1) после логирования
func (p *consoleProvider) Write(ctx context.Context, level sglogger.Level, message string, fields sglogger.Fields) error {
	if !p.ShouldLog(ctx, level) {
		return nil
	}

	logrusFields := logrus.Fields{}
	maps.Copy(logrusFields, fields)

	entry := p.logrus.WithFields(logrusFields)

	switch level {
	case sglogger.LevelDebug:
		entry.Debug(message)
	case sglogger.LevelInfo:
		entry.Info(message)
	case sglogger.LevelWarn:
		entry.Warn(message)
	case sglogger.LevelError:
		entry.Error(message)
	case sglogger.LevelFatal:
		entry.Fatal(message) // Вызывает os.Exit(1) после логирования
	}

	return nil
}

// ShouldLog определяет, должно ли сообщение с указанным уровнем быть залогировано.
// Используется для фильтрации сообщений по уровню важности.
//
// Параметры:
//   - ctx - контекст выполнения (в текущей реализации не используется)
//   - level - уровень проверяемого сообщения
//
// Возвращает:
//   - bool - true если сообщение должно быть залогировано, false если отфильтровано
//
// Принцип работы:
//   - Сравнивает уровень сообщения с минимальным уровнем из конфигурации
//   - Сообщения с уровнем ВЫШЕ или РАВНЫМ минимальному проходят фильтр
//   - Например, при LevelInfo: Debug отфильтруется, Info и выше пройдут
func (p *consoleProvider) ShouldLog(ctx context.Context, level sglogger.Level) bool {
	return level >= p.config.Level
}

// Close освобождает ресурсы провайдера логирования.
// В случае консольного провайдера не требует никаких действий.
//
// Параметры:
//   - ctx - контекст выполнения для контроля времени выполнения операции
//
// Возвращает:
//   - error - всегда возвращает nil, так как не требует очистки ресурсов
//
// Реализация интерфейса io.Closer для совместимости с паттернами управления ресурсами.
func (p *consoleProvider) Close(ctx context.Context) error {
	return nil
}