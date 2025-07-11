package handlers;

import com.example.config.LogLevel;
import com.example.config.LoggerConfig;
import com.example.service.LogService;

import java.util.Objects;

public class ErrorLogHandler extends LogHandler {
    public ErrorLogHandler(LogService svc) {
        super(svc);
    }

    @Override
    public boolean canHandle(LoggerConfig config) {
        return Objects.equals(LogLevel.ERROR, config.getLogLevel());
    }
}
