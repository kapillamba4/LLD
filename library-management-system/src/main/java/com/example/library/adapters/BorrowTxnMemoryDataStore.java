package com.example.library.adapters;

import java.util.*;

import com.example.library.models.BorrowTransaction;
import com.example.library.persistence.InMemoryDatabase;
import com.example.library.repository.BorrowTxnRepository;

public class BorrowTxnMemoryDataStore extends InMemoryDatabase<BorrowTransaction> implements BorrowTxnRepository {

    @Override
    public void issueBookToMember(BorrowTransaction transaction) {
        save(transaction.getTransactionId(), transaction);
    }

    @Override
    public void returnBookToShelf(String transactionId) {
        deleteByKey(transactionId);
    }

    @Override
    public List<BorrowTransaction> getBooksWithBorrowDateBefore(long borrowDate) {
        Map<String, BorrowTransaction> map = findAll();
        if (map == null) {
            return null;
        }
        return map.entrySet().stream().map(c -> c.getValue()).filter(bt -> bt.getBorrowDate() <= borrowDate).toList();
    }

    @Override
    public BorrowTransaction getTransactionById(String transactionId) {
        return findByKey(transactionId);
    }
}
