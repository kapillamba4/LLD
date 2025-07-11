package com.example.library.models;

import lombok.*;

@Builder
@Getter
@Setter
public class Book {
    String ifsn;
    String bookName;
    int copies;
}
