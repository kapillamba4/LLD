package handlers;

import com.example.config.LogLevel;
import com.example.config.LoggerConfig;
import com.example.service.LogService;

import java.util.List;

public class DebugLogHandler extends LogHandler {
    public DebugLogHandler(LogService svc) {
        super(svc);
    }

    @Override
    public boolean canHandle(LoggerConfig config) {
        return List.of(LogLevel.ERROR, LogLevel.DEBUG, LogLevel.INFO).contains(config.getLogLevel());
    }

}
