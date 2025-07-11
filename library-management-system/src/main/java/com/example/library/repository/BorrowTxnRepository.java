package com.example.library.repository;

import com.example.library.models.BorrowTransaction;
import java.util.*;

public interface BorrowTxnRepository {
    void issueBookToMember(BorrowTransaction transaction);
    void returnBookToShelf(String transactionId);
    List<BorrowTransaction> getBooksWithBorrowDateBefore(long borrowDate);
    BorrowTransaction getTransactionById(String transactionId);
}
