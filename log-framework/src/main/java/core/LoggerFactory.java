package core;

import com.example.config.LogDestination;
import com.example.config.LoggerConfig;

public class LoggerFactory {
    public static AbstractLogger getLogger(LogDestination logDestination, LoggerConfig loggerConfig) {
        switch (logDestination) {
            case FILE -> {
                return new FileLogger(loggerConfig.getDestinations().get(logDestination));
            }
            default ->  {
                return new ConsoleLogger();
            }
        }
    }
}
