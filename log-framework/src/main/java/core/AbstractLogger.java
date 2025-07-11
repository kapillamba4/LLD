package core;

import java.util.Date;

public abstract class AbstractLogger {
    public String getCurrentFormattedDate() {
        return new Date().toString();
    }

    public abstract void write(String message);
}
