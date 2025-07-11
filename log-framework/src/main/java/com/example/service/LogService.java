package com.example.service;

import com.example.config.LoggerConfig;
import core.AbstractLogger;
import core.LoggerFactory;
import handlers.LogHandler;
import handlers.LogHandlerChain;

import java.util.List;

public class LogService {
    private final LoggerConfig loggerConfig;
    private LogHandler logHandler;
    private final List<AbstractLogger> loggerList;
    
    private LogService(LoggerConfig config) {
        loggerConfig = config;
        loggerList = loggerConfig.getDestinations().keySet().stream().map(logDestination -> LoggerFactory.getLogger(logDestination, loggerConfig)).toList();
    }
    
    public static LogService getLogger(LoggerConfig config) {
        LogService service = new LogService(config);
        service.logHandler = LogHandlerChain.getHandlerChain(service);
        return service;
    }

    public void log(String message) {
        loggerList.forEach(l -> l.write(message));
    }
    
    public void handleLog(String message) {
        if (logHandler != null) {
            logHandler.handle(this.loggerConfig, message);
        }
    }
}
