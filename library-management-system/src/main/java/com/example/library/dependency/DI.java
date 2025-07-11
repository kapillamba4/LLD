package com.example.library.dependency;

import com.example.library.repository.BookRepository;
import com.example.library.repository.BorrowTxnRepository;
import com.example.library.repository.UserRepository;

import lombok.Builder;
import lombok.Getter;

@Builder
@Getter
public class DI {
    BookRepository bookRepository;
    BorrowTxnRepository txnRepository;
    UserRepository userRepository;
}
