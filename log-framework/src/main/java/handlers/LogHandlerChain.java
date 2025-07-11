package handlers;

import com.example.service.LogService;

public class LogHandlerChain {
    public static LogHandler getHandlerChain(LogService service) {
        InfoLogHandler infoLogHandler = new InfoLogHandler(service);
        DebugLogHandler debugLogHandler = new DebugLogHandler(service);
        ErrorLogHandler errorLogHandler = new ErrorLogHandler(service);
        infoLogHandler.setNextLogHandler(debugLogHandler);
        debugLogHandler.setNextLogHandler(errorLogHandler);
        return infoLogHandler;
    }
}
