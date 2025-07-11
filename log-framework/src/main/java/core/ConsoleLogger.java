package core;

import com.example.config.LogDestination;

public class ConsoleLogger extends AbstractLogger {
    public static final LogDestination logDestination = LogDestination.CONSOLE;

    @Override
    public void write(String message) {
        System.out.println(message);
    }
}
