package core;

import com.example.config.LogDestination;

import java.io.FileWriter;
import java.io.IOException;

public class FileLogger extends AbstractLogger {
    public static final LogDestination logDestination = LogDestination.FILE;
    private final String filePath;

    public FileLogger(String path) {
        this.filePath = path;
    }

    @Override
    public void write(String message) {
        try (FileWriter writer = new FileWriter(filePath, true)) {
            writer.write(message);
        } catch (IOException e) {
            System.err.println("Failed to write log to file: " + e.getMessage());
        }
    }
}
