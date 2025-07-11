package com.example.library.models;

import lombok.*;

@Builder
@Getter
public class User {
    String name;
    String memberId;
    String contactNumber;
    String address;
}
