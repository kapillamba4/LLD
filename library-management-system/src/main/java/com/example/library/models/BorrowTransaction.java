package com.example.library.models;

import lombok.*;

@Builder
@Getter
public class BorrowTransaction {
    String transactionId;
    String memberId;
    String bookIfsn;
    long borrowDate;
}
