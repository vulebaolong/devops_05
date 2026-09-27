package com.example.springbe;

public class NoDataException extends RuntimeException {
    public NoDataException(String message) {
        super(message);
    }
}
