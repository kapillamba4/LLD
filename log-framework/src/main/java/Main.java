import com.example.config.LogDestination;
import com.example.config.LogLevel;
import com.example.config.LoggerConfig;
import com.example.service.LogService;

import java.util.HashMap;
import java.util.Map;

public class Main {
    public static void main(String[] args) {
        // Basic console logging
        System.out.println("\n1. Console Logging:");
        Map<LogDestination, String> consoleMap = new HashMap<>();
        consoleMap.put(LogDestination.CONSOLE, "enabled");
        
        LoggerConfig consoleConfig = new LoggerConfig(LogLevel.INFO, consoleMap);
        LogService consoleLogger = LogService.getLogger(consoleConfig);
        
        consoleLogger.handleLog("Hello from console logger!\n");
        
        // File logging
        System.out.println("\n2. File Logging:");
        Map<LogDestination, String> fileMap = new HashMap<>();
        fileMap.put(LogDestination.FILE, "test.log");
        
        LoggerConfig debugFileLoggerCfg = new LoggerConfig(LogLevel.DEBUG, fileMap);
        LogService debugFileLogger = LogService.getLogger(debugFileLoggerCfg);
        
        debugFileLogger.handleLog("This message goes to test.log file\n");
        System.out.println("Message written to test.log");
        
        // Multiple destinations
        System.out.println("\n3. Multi-destination Logging:");
        Map<LogDestination, String> logDest = new HashMap<>();
        logDest.put(LogDestination.CONSOLE, "enabled");
        logDest.put(LogDestination.FILE, "both.log");
        
        LogService errorLogger = LogService.getLogger(new LoggerConfig(LogLevel.ERROR, logDest));

        errorLogger.handleLog( "Error message - goes to both console and file!\n");
        System.out.println("Message also written to both.log");
    }
}