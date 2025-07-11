package com.example.library;

import com.example.library.adapters.BookMemoryDataStore;
import com.example.library.adapters.BorrowTxnMemoryDataStore;
import com.example.library.adapters.UserMemoryDataStore;
import com.example.library.dependency.DI;
import com.example.library.models.User;
import com.example.library.usecase.BooksManager;
import com.example.library.usecase.TransactionManager;
import com.example.library.usecase.UserManager;

public class Main {
    public static void main(String[] args) {
        DI di = DI.builder()
            .bookRepository(new BookMemoryDataStore())
            .txnRepository(new BorrowTxnMemoryDataStore())
            .userRepository(new UserMemoryDataStore())
            .build();

        BooksManager booksManager = new BooksManager(di);
        TransactionManager transactionManager = new TransactionManager(di);
        UserManager userManager = new UserManager(di);
        booksManager.addNewBook("IFSN1", "Book 1", 2);
        booksManager.addNewBook("IFSN2", "Book 2", 1);
        booksManager.addNewBook("IFSN1", "Book 1", 3);
        booksManager.addNewBook("IFSN1", "Book 1", 1);
        
        userManager.addNewMember("Kapil", "9999955555", "Sarjapur Road, Bangalore");
        userManager.addNewMember("Pahul", "9999955556", "Sarjapur Road, Bangalore");

        User user = userManager.fetchMemberDetailsByContact("9999955555");
        transactionManager.borrowBook(user.getMemberId(), "IFSN1");
        transactionManager.borrowBook(user.getMemberId(), "IFSN2");
     
        transactionManager.getOverdueBooks(0L).forEach(ob -> {
            System.out.printf("IFSN %s, BORROW %d, MemberID %s, Txn ID %s\n", ob.getBookIfsn(), ob.getBorrowDate(), ob.getMemberId(), ob.getTransactionId());
        });
        System.out.println(userManager.fetchMemberDetailsByContact("9999955556"));
    }
}

