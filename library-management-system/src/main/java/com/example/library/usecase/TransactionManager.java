package com.example.library.usecase;

import com.example.library.config.Constants;
import com.example.library.dependency.DI;
import com.example.library.models.BorrowTransaction;
import java.util.*;

public class TransactionManager {
    private DI diContainer;

    public TransactionManager(DI di) {
        this.diContainer = di;
    }

    public boolean borrowBook(String memberId, String ifsn) {
        if (diContainer.getBookRepository().borrowBook(ifsn)) {
            diContainer.getTxnRepository()
                    .issueBookToMember(BorrowTransaction.builder().bookIfsn(ifsn).borrowDate(System.currentTimeMillis())
                            .transactionId(UUID.randomUUID().toString()).memberId(memberId).build());
            return true;
        }
        return false;
    }

    public void returnBook(String transactionId) {
        BorrowTransaction txn = diContainer.getTxnRepository().getTransactionById(transactionId);
        diContainer.getBookRepository().returnBook(txn.getBookIfsn());
        diContainer.getTxnRepository()
                .returnBookToShelf(transactionId);
    }

    // books with borrow date more than 1 week ago
    public List<BorrowTransaction> getOverdueBooks(Long timestamp) {
        if (timestamp == null) {
            timestamp = Constants.sevenDaysInMillis;
        }
        return diContainer.getTxnRepository()
                .getBooksWithBorrowDateBefore(System.currentTimeMillis() - timestamp);
    }
}
