package handlers;

import com.example.config.LoggerConfig;
import com.example.service.LogService;
import lombok.Getter;
import lombok.Setter;

public abstract class LogHandler {
    protected LogService logService;
    @Getter @Setter protected LogHandler nextLogHandler;
    public LogHandler(LogService svc) {
        logService = svc;
    }

    public abstract boolean canHandle(LoggerConfig config);
    public void handle(LoggerConfig config, String msg) {
        if (canHandle(config)) {
            logService.log(msg);
            return;
        }
        LogHandler logHandler = getNextLogHandler();
        if (logHandler != null) {
            logHandler.handle(config, msg);
        }
    }
}
