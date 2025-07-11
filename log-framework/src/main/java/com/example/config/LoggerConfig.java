package com.example.config;

import lombok.Getter;
import lombok.Setter;

import java.util.HashMap;
import java.util.Map;

@Getter
@Setter
public class LoggerConfig {
    LogLevel logLevel;
    private Map<LogDestination, String> destinations;
    private String logFormat;
    
    public LoggerConfig() {
        this.logLevel = LogLevel.INFO;
        destinations = new HashMap<>();
        destinations.put(LogDestination.CONSOLE, "enabled");
        logFormat = "%{{date}} %{{logLevel}} %{{message}}";
    }
    
    public LoggerConfig(LogLevel logLevel, Map<LogDestination, String> destinations) {
        this.logLevel = logLevel;
        this.destinations = destinations;
        this.logFormat = "%{{date}} %{{logLevel}} %{{message}}";
    }
} 